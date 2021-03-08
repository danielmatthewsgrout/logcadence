package parser

/*  @danielmatthewsgrout */

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"
)

//LogParser parses logs
type LogParser interface {
	Parse(out chan string, maxLinesToOutput uint64, wg *sync.WaitGroup)
}

func getRegExpForTimestampPattern(p string) (*regexp.Regexp, error) {

	pattern := ""

	for _, c := range p {

		if unicode.IsNumber(c) { //this is a number so we need numbers here
			pattern += "\\d"
		} else if unicode.IsLetter(c) { //this is a letter
			pattern += "[a-zA-Z]"
		} else if c == '_' { //optional number
			pattern += "[\\d\\s]"
		} else if c == ' ' { //le space
			pattern += "\\s"
		} else { //'tis neither lord, escape!
			pattern += "\\" + string(c)
		}

	}
	return regexp.Compile(pattern)
}

//GetTimestampFormat get a Golang compatible timestamp format
func GetTimestampFormat(f string) (string, error) {

	switch strings.ToUpper(f) {
	case "ANSIC":
		return time.ANSIC, nil
	case "UNIXDATE":
		return time.UnixDate, nil
	case "RUBYDATE":
		return time.RubyDate, nil
	case "RFC822":
		return time.RFC822, nil
	case "RFC822Z":
		return time.RFC822Z, nil
	case "RFC850":
		return time.RFC850, nil
	case "RFC1123":
		return time.RFC1123, nil
	case "RFC1123Z":
		return time.RFC1123Z, nil
	case "RFC3339":
		return time.RFC3339, nil
	case "RFC3339NANO":
		return time.RFC3339Nano, nil
	case "KITCHEN":
		return time.Kitchen, nil
	case "STAMP":
		return time.Stamp, nil
	case "STAMPMILLI":
		return time.StampMilli, nil
	case "STAMPMICRO":
		return time.StampMicro, nil
	case "STAMPNANO":
		return time.StampNano, nil
	case "DDMMYYYYHHMMSS":
		return "02012006150405", nil
	case "DDMMYYYYHHMMSSSSS":
		return "02012006150405.000", nil
	case "STAMPYEARMICRO":
		return "02 Jan 2006-15:04:05.000000", nil
	case "JAVA1":
		return "2006-01-02 15:04:05,000", nil
	case "JAVA2":
		return "2006-01-02 15:04:05.000", nil
	case "ISO8601COMPLETE":
		return "2006-01-02T15:04:05.00-07:00", nil
	case "ISO8601SECONDS":
		return "2006-01-02T15:04:05-07:00", nil
	default:
		return time.ANSIC, fmt.Errorf("Unrecognised timestamp format string %s", f)
	}

}

//01 Mar 2021-00:00:00.245713
//12 Dec 2020-09:54:00.002542
//2020-11-16 01:41:17,125
