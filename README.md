# logcadence

Log snipping by timestamp

Source requires Golang >= v1.16

Binaries available on the releases page

## How to Use

```bash
$ logcandence -h

Usage of logcadence:
  -f string
        filename or filename pattern to read
  -t string
        timestamp format name - see README.md for a list
  -s string
        starting timestamp - using format selected with "-t"
  -e string
        ending timestamp (optional) - using format selected with "-t"
  -m uint
        maximum lines to return per file (optional)
  -ss string
        substring required to return line (optional)
  -c    use colours in output (*nix only) 
```

**Note:** when using wildcards in the filename it reads multiple files in parallel so the output from this will not be ordered.

### Example

```bash
logcadence -c -f "/var/log/kern.log*" -s "Mar  7 06:31:13" -t STAMP -e "Mar  7 23:59:01" -ss "auth"
```

Github doesn't support colours in Markdown so imagine the "auth" bits are green. :)

```log
Mar  7 06:32:43 hostname kernel: [1329582.856664] wlp0x00x0: deauthenticating from 00:00:00:00:00:00 by local choice (Reason: 3=DEAUTH_LEAVING)
Mar  7 06:32:49 hostname kernel: [1329589.080753] wlp0x00x0: authenticate with 00:00:00:00:00:00
Mar  7 06:32:49 hostname kernel: [1329589.087006] wlp0x00x0: send auth to 00:00:00:00:00:00 (try 1/3)
Mar  7 06:32:50 hostname kernel: [1329589.127305] wlp0x00x0: authenticated
Mar  7 15:24:23 hostname kernel: [1361481.976513] wlp0x00x0: authenticate with 00:00:00:00:00:00
Mar  7 15:24:23 hostname kernel: [1361481.989113] wlp0x00x0: send auth to 00:00:00:00:00:00 (try 1/3)
```

## Timestamp Formats

Use the variable name as the parameter on the command line

### from Golang "time" package

```golang
ANSIC = "Mon Jan _2 15:04:05 2006"
UnixDate= "Mon Jan _2 15:04:05 MST 2006"
RubyDate= "Mon Jan 02 15:04:05 -0700 2006"
RFC822  = "02 Jan 06 15:04 MST"
RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
RFC850  = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z= "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
RFC3339 = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen = "3:04PM"
Stamp  = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
```

### custom formats

```golang
DDMMYYYYHHmmss = "02012006150405"
DDMMYYYYHHmmsssss = "02012006150405.000"
STAMPYEARMICRO = "02 Jan 2006-15:04:05.000000"
JAVA1 = "2006-01-02 15:04:05,000"
JAVA2 = "2006-01-02 15:04:05.000"
ISO8601COMPLETE = "2006-01-02T15:04:05.00-07:00"
ISO8601SECONDS = "2006-01-02T15:04:05-07:00"
```
