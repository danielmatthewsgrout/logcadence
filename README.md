# logcadence

Log snipping by timestamp.  Supports multiple input files.

Source requires Golang >= v1.16

Binaries available on the releases page

## How to Use

```text
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
  -c    use colours in output (not on Windows) 
```

**Note:** when using wildcards in the filename it reads multiple files in parallel so the output from this will not be ordered.

### Example

```bash
logcadence -c -f "/var/log/kern.log*" -s "Mar  7 06:31:13" -t STAMP -e "Mar  7 23:59:01" -ss "auth"
```

Github doesn't support colours in Markdown so imagine the "daniel" bits are green and the filename is yellow with red brackets. :)

```log
[/var/log/auth.log.1] Feb  2 20:23:51 daniel-XPS-13-9360 systemd-logind[911]: Power key pressed.
[/var/log/auth.log.1] Feb  2 20:23:51 daniel-XPS-13-9360 systemd-logind[911]: Powering Off...
[/var/log/auth.log] Feb 10 07:15:01 daniel-XPS-13-9360 CRON[29090]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:15:01 daniel-XPS-13-9360 CRON[29090]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:17:01 daniel-XPS-13-9360 CRON[29448]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:17:01 daniel-XPS-13-9360 CRON[29448]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:25:01 daniel-XPS-13-9360 CRON[30888]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:25:01 daniel-XPS-13-9360 CRON[30888]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:30:01 daniel-XPS-13-9360 CRON[31781]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:30:01 daniel-XPS-13-9360 CRON[31781]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:35:01 daniel-XPS-13-9360 CRON[32655]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:35:01 daniel-XPS-13-9360 CRON[32655]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:39:44 daniel-XPS-13-9360 gnome-keyring-daemon[2345]: asked to register item /org/freedesktop/secrets/collection/login/1, but it's already registered
[/var/log/auth.log] Feb 10 07:45:01 daniel-XPS-13-9360 CRON[2096]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:45:01 daniel-XPS-13-9360 CRON[2096]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 07:46:42 daniel-XPS-13-9360 gdm-password]: gkr-pam: unlocked login keyring
[/var/log/auth.log] Feb 10 07:55:01 daniel-XPS-13-9360 CRON[5040]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 07:55:01 daniel-XPS-13-9360 CRON[5040]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:05:01 daniel-XPS-13-9360 CRON[6856]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:05:01 daniel-XPS-13-9360 CRON[6856]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:15:01 daniel-XPS-13-9360 CRON[8792]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:15:01 daniel-XPS-13-9360 CRON[8792]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:17:01 daniel-XPS-13-9360 CRON[9146]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:17:01 daniel-XPS-13-9360 CRON[9146]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:25:01 daniel-XPS-13-9360 CRON[10572]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:25:01 daniel-XPS-13-9360 CRON[10572]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:30:01 daniel-XPS-13-9360 CRON[11445]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:30:01 daniel-XPS-13-9360 CRON[11445]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:35:01 daniel-XPS-13-9360 CRON[12332]: pam_unix(cron:session): session opened for user root by (uid=0)
[/var/log/auth.log] Feb 10 08:35:01 daniel-XPS-13-9360 CRON[12332]: pam_unix(cron:session): session closed for user root
[/var/log/auth.log] Feb 10 08:39:45 daniel-XPS-13-9360 gnome-keyring-daemon[2345]: asked to register item /org/freedesktop/secrets/collection/login/1, but it's already registered
[/var/log/auth.log.1] Feb  2 20:23:51 daniel-XPS-13-9360 systemd-logind[911]: System is powering down.
[/var/log/auth.log.1] Feb  2 20:23:51 daniel-XPS-13-9360 sshd[1070]: Received signal 15; terminating.
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 sshd[1081]: Server listening on 0.0.0.0 port 22.
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 sshd[1081]: Server listening on :: port 22.
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 systemd-logind[879]: New seat seat0.
[/var/log/auth.log] Feb 10 11:22:40 daniel-XPS-13-9360 systemd-logind[917]: Operation 'sleep' finished.
[/var/log/auth.log] Feb 10 11:22:44 daniel-XPS-13-9360 gdm-password]: gkr-pam: unlocked login keyring
[/var/log/auth.log] Feb 10 11:22:51 daniel-XPS-13-9360 pkexec: pam_unix(polkit-1:session): session opened for user root by (uid=1000)
[/var/log/auth.log] Feb 10 11:22:51 daniel-XPS-13-9360 pkexec[14650]: daniel: Executing command [USER=root] [TTY=unknown] [CWD=/home/daniel] [COMMAND=/usr/bin/cpufreqctl --no-turbo --set=0]
[/var/log/auth.log] Feb 10 11:22:51 daniel-XPS-13-9360 pkexec: pam_unix(polkit-1:session): session opened for user root by (uid=1000)
[/var/log/auth.log] Feb 10 11:22:51 daniel-XPS-13-9360 pkexec[14657]: daniel: Executing command [USER=root] [TTY=unknown] [CWD=/home/daniel] [COMMAND=/usr/bin/cpufreqctl --no-turbo --set=1]
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 systemd-logind[879]: Watching system buttons on /dev/input/event3 (Power Button)
[/var/log/auth.log] Feb 10 11:22:51 daniel-XPS-13-9360 pkexec: pam_unix(polkit-1:session): session opened for user root by (uid=1000)
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 systemd-logind[879]: Watching system buttons on /dev/input/event1 (Power Button)
[/var/log/auth.log.1] Feb  2 20:24:29 daniel-XPS-13-9360 systemd-logind[879]: Watching system buttons on /dev/input/event0 (Lid Switch)
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
