package main

import (
	"flag"
	"fmt"
	"github.com/TrueAbility/sudosh2asciinema"
	"os"
)

func main() {
	var directory string
	var help bool

	flag.StringVar(&directory, "directory", "/var/log/sudosh", "sudosh logfile directory")
	flag.StringVar(&directory, "d", "/var/log/sudosh", "sudosh logfile directory")
	flag.BoolVar(&help, "help", false, "Show usasge")
	flag.BoolVar(&help, "h", false, "Show usasge")

	flag.Parse()

	if directory == "" || help {
		fmt.Println("Converts a directory of sudosh files to asciinema files.")
		fmt.Println("USAGE: sudosh-dir -d <logfile-directory>")
		os.Exit(0)
	}

	fmt.Printf("Converting all files in %s\n", directory)
	sudosh2asciinema.ConvertDirectory(directory)
	fmt.Println("Done")
}
