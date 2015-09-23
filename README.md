
# sudosh2asciinema

`sudosh` is a shell that captures a users input into two files per session.
These log files are stored in `/var/log/sudosh` by default and are named

    root-root-time-1442065357-i6pmm9Eq6T6wQCOx
    root-root-script-1442065357-i6pmm9Eq6T6wQCOx

 - root-root is the user
 - root-root-time is the delay between commands
 - root-root-script is the actual text typed and/or return value
 - 1442065357 is a timestamp
 - The last part appears to be a random bit to prevent conflicts

Asciinema.org is an online terminal recording and playback site. Their
javascript terminal player does a terrific job of playing back terminal
recordings.

Asciinema V1 file format is a JSON file that combines delay and output into
a `stdout` key

The JSON structure looks like:

    version: 1
    width: tty-width
    height: tty-height
    duration: length-of-recording
    command: the-command-that-was-run
    title: title-of-recording
    env: env-variables
    stdout: an-array-of-commands, eg [0.000000, "ls -l"]

sudosh2asciinema converts the sudosh files the asciinema formatn

## USAGE

### Command Line

sudosh-convert
    Convert a single file
    sudosh-convert -t <timingfile> -s <scriptfile> -o <outputfile>

sudosh-dir
    Convert all the files in a directory
    sudosh-dir -d <directory>


### Library

    import "github.com/TrueAbility/sudosh2asciinema"

Then look at Convert() or ConvertDirectory()

## License

Copyright (c) 2015 TrueAbility, Inc.
MIT License, see LICENSE.txt for details
