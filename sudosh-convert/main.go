package main

import (
	"flag"
	"fmt"
	"github.com/TrueAbility/sudosh2asciinema"
	"os"
)

func main() {
	var timing, script, output string

	flag.StringVar(&timing, "timing", "", "sudosh timing filename")
	flag.StringVar(&timing, "t", "", "sudosh timing filename")
	flag.StringVar(&script, "script", "", "sudosh script filename")
	flag.StringVar(&script, "s", "", "sudosh script filename")
	flag.StringVar(&output, "output", "", "output filename")
	flag.StringVar(&output, "o", "", "output filename")

	flag.Parse()

	if timing == "" || script == "" || output == "" {
		fmt.Println("USAGE: sudosh-convert -t <timing> -s <script> -o <output>")
		os.Exit(0)
	}

	timingTimestamp := sudosh2asciinema.FindTimeStampFromFilename(timing)
	scriptTimestamp := sudosh2asciinema.FindTimeStampFromFilename(script)
	if timingTimestamp != scriptTimestamp {
		fmt.Println("WARNING: Timestamp mismatch!")
	}

	su := sudosh2asciinema.SudoshHistory{
		TimeFilename:   timing,
		ScriptFilename: script,
	}

	su.Convert(output)
}
