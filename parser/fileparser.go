package parser

/*  @danielmatthewsgrout */

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type fileLogParser struct {
	file            *os.File
	fileName        string
	scanner         *bufio.Scanner
	start           time.Time
	stop            time.Time
	substring       string
	timestampRegex  *regexp.Regexp
	timestampFormat string
	useColour       bool
}

const (
	bufferSize = 1000

	//console colour codes
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

//GetFileLogParser gets a fileLogParser for a given path and sends output to the given channel for timestamps start and stop and substring using timestamp pattern
func GetFileLogParser(filename string, startTime, stopTime time.Time, timestampformat, substring string, useColour bool) (LogParser, error) {

	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("Unable to open file %s - error: %s", filename, err)
	}

	tsRe, err := getRegExpForTimestampPattern(timestampformat)

	if err != nil {
		return nil, fmt.Errorf("Unable to create regular expression for %s - error: %s", timestampformat, err)
	}

	return &fileLogParser{
		file:            f,
		scanner:         bufio.NewScanner(f),
		start:           startTime,
		stop:            stopTime,
		substring:       substring,
		timestampRegex:  tsRe,
		timestampFormat: timestampformat,
		fileName:        filename,
		useColour:       useColour,
	}, nil
}

//Parse parse the file and push to output to the channel defined in the builder for maxlines - needs wait group to signal completion
func (f *fileLogParser) Parse(out chan string, maxLinesToOutput uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer f.file.Close()

	foundStart := false

	var lineNumber uint64

	//read file line by line
	for f.scanner.Scan() {

		//read this line as a string
		s := f.scanner.Text()

		//do we have a match for the timestamp format?
		if f.timestampRegex.MatchString(s) {
			ts := f.timestampRegex.FindString(s)
			t, err := time.Parse(f.timestampFormat, ts)
			foundStart = err == nil && (t.Equal(f.start) || t.After(f.start)) //true if start timestamp found or after
		}

		if foundStart { //we know where the start is so output this line and look for the end timestamp

			//see if we are at the end - if needed, don't bother with regex unless we have a valid end time
			if f.stop.After(f.start) && f.timestampRegex.MatchString(s) {
				ts := f.timestampRegex.FindString(s)
				t, err := time.Parse(f.timestampFormat, ts)
				if err == nil && t.After(f.stop) { //if this passes we found the end timestamp
					return
				}
			}
			//if substring search is set then do that check
			if f.substring == "" || strings.Contains(s, f.substring) {
				lineNumber++

				//send this message to the receiver
				if f.useColour {
					if f.useColour && f.substring != "" && strings.Contains(s, f.substring) {
						s = strings.ReplaceAll(s, f.substring, green+f.substring+reset) //make the substring search nice and green :)
					}
					out <- fmt.Sprintf(red+"["+yellow+"%s"+red+"]"+reset+" %s", f.fileName, s)
				} else {
					out <- fmt.Sprintf("[%s] %s", f.fileName, s)
				}
			}
		}

		//don't read beyond the maximum number of lines - ignore test if maxLines is 0
		if maxLinesToOutput != 0 && lineNumber == maxLinesToOutput {
			return
		}
	}

	return
}
