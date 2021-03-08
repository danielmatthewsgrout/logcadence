package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"std/github.com/danielmatthewsgrout/logcadence/parser"
	"strings"
	"sync"
	"time"
)

/*  @danielmatthewsgrout */

//buffer for channel - to prevent the output from blocking the log parsing too much
const (
	//console colour codes
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {

	filename := flag.String("f", "", "filename or filename pattern to read")
	tsf := flag.String("t", "", "timestamp format name - see README.md for a list")
	start := flag.String("s", "", "starting timestamp - using format selected with \"-t\"")
	stop := flag.String("e", "", "ending timestamp (optional) - using format selected with \"-t\"")
	substring := flag.String("ss", "", "substring required to return line (optional)")
	maxLines := flag.Uint64("m", 0, "maximum lines to return (optional)")
	cols := flag.Bool("c", false, "use colours in output (not on Windows)")

	flag.Parse()

	colours := *cols && runtime.GOOS != "windows" //Windows doesn't support console colours, how dull!

	if *filename == "" {
		println("-f paramter required with valid filename or pattern.  Use -h to see all options.")
		return
	}

	if *tsf == "" {
		println("-t paramter required with valid timestamp format.  Use -h to see all options.")
		return
	}

	if *start == "" {
		println("-s paramter required with valid start time.  Use -h to see all options.")
		return
	}

	timestampFormat, err := parser.GetTimestampFormat(*tsf)

	if err != nil {
		println(err)
		println("See README.md for timestamp formats")
		return
	}

	startTime, err := time.Parse(timestampFormat, *start)

	if err != nil {
		fmt.Printf("Unable to parse start time %s - error: %s\n", *start, err)
		println("See README.md for timestamp formats")
		return
	}

	stopTime := time.Unix(0, 0)

	if *stop != "" {

		if stopTime, err = time.Parse(timestampFormat, *stop); err != nil {
			fmt.Printf("Unable to parse end time %s - error: %s\n", *stop, err)
			println("See README.md for timestamp formats")
			return
		}
	}

	files, err := filepath.Glob(*filename)

	if err != nil {
		fmt.Printf("invalid file path: %s\n", err)
	}

	out := make(chan string) //receiver for output from the parser(s)

	parsers := make([]parser.LogParser, 0)
	for _, s := range files {
		parser, err := parser.GetFileLogParser(s, startTime, stopTime, timestampFormat, *substring)
		if err != nil {
			fmt.Printf("error creating parser for file %s: %s\n", s, err)
			return
		}
		parsers = append(parsers, parser)
	}

	chanWait := sync.WaitGroup{}
	chanWait.Add(1)
	//start output writer
	go func() {
		defer chanWait.Done() //signal we have done everything we need to
		for s := range out {  //read channel in a loop until signalled to stop by channel closing
			if colours && *substring != "" && strings.Contains(s, *substring) {
				s = strings.ReplaceAll(s, *substring, green+*substring+reset) //make the substring search nice and green :)
			}
			println(s)
		}
	}()

	wg := sync.WaitGroup{}
	//start the parsers

	for _, p := range parsers {
		wg.Add(1) //wait group will be decremented by the parser
		go p.Parse(out, *maxLines, &wg)
	}

	//wait for all parsers to complete
	wg.Wait()
	close(out)
	chanWait.Wait() //wait for channel to finish printing everything in the buffer
}
