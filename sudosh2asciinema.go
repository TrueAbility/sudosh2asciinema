// Convert sudosh timing and script files to asciinema json format v1
package sudosh2asciinema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// SudoSH has two files, script and time
type SudoshHistory struct {
	TimeFilename   string
	ScriptFilename string
}

type Command struct {
	Delay   float64
	Command []byte
}

func (command *Command) MarshalJSON() ([]byte, error) {
	str, err := json.Marshal(string(command.Command))
	if err != nil {
		return str, err
	}
	json := fmt.Sprintf(`[%.6f, %s]`, command.Delay, str)
	return []byte(json), nil
}

type AsciinemaFile struct {
	Filename string
	Version  int               `json:"version"`
	Width    int               `json:"width"`
	Height   int               `json:"height"`
	Duration float64           `json:"duration"`
	Command  string            `json:"command"`
	Title    string            `json:"title"`
	Env      map[string]string `json:"env"`
	StdOut   []Command         `json:"stdout"`
}

// Convert from SudoSH to Asciinema format
func (su SudoshHistory) Convert(output string) error {
	timing, err := os.Open(su.TimeFilename)
	if err != nil {
		fmt.Println("Error reading timing file", timing)
		os.Exit(-1)
	}
	defer timing.Close()

	script, err := os.Open(su.ScriptFilename)
	if err != nil {
		fmt.Println("Error reading script file", script)
		os.Exit(-1)
	}
	defer script.Close()

	lines, duration, err := mungeLines(timing, script)
	asciiFile := NewAsciiFile(output, lines, duration)
	err = asciiFile.Write()
	return err
}

// Convert a Directory of sudosh files to asciinema
func ConvertDirectory(directory string) error {
	sudoshfiles, err := findScriptTimeFilesInDir(directory)
	if err != nil {
		fmt.Printf("Error getting files from %s", directory)
		return err
	}
	for _, sudoshFile := range sudoshfiles {
		timestamp := FindTimeStampFromFilename(sudoshFile.TimeFilename)
		fmt.Println("Processing sudosh file", timestamp)

		filename := fmt.Sprintf("%s/%s.json", directory, timestamp)
		sudoshFile.Convert(filename)
	}
	return nil
}

func FindTimeStampFromFilename(filename string) string {
	tokens := strings.Split(filename, "-")
	timestamp := strings.Join(tokens[3:5], "-")
	return timestamp
}

func findScriptTimeFilesInDir(directory string) ([]SudoshHistory, error) {
	glob := fmt.Sprintf("%s/root-root-script-*", directory)
	var collection []SudoshHistory
	files, err := filepath.Glob(glob)
	if err != nil {
		fmt.Printf("Error globing files in %s error was %s", glob, err)
		return collection, err
	}

	for _, scriptFile := range files {
		timeFile := strings.Replace(scriptFile, "-script-", "-time-", 1)
		collection = append(collection,
			SudoshHistory{
				TimeFilename:   timeFile,
				ScriptFilename: scriptFile,
			})
	}
	return collection, err
}

// combines the delay and command into one struct
func mungeLines(timing *os.File, script *os.File) ([]Command, float64, error) {
	duration := 0.0
	var lines []Command
	var err error
	for {
		var i int
		var delay float64
		var n int
		n, err := fmt.Fscanln(timing, &delay, &i)
		if i < 0 {
			i = 0
		}
		if n == 0 || err != nil {
			break
		}
		b := make([]byte, i)
		n, err = script.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		lines = append(lines, Command{
			Delay:   delay,
			Command: b,
		})
		duration += delay
	}
	return lines, duration, err
}

// get a new asciifile struct
func NewAsciiFile(output string, lines []Command, duration float64) *AsciinemaFile {
	ascii := &AsciinemaFile{
		Filename: output,
		Version:  1,
		Width:    100,
		Height:   40,
		Duration: duration,
		Command:  "/opt/bin/sudosh",
		Title:    "SudoSH",
		StdOut:   lines,
	}
	return ascii
}

// write the json to a file
func (ascii *AsciinemaFile) Write() error {
	output, err := os.Create(ascii.Filename)
	if err != nil {
		fmt.Println("Error creating file", ascii.Filename, err)
		os.Exit(-1)
	}
	defer output.Close()

	b, err := json.Marshal(ascii)
	if err != nil {
		fmt.Println("Error creationg Asciinema json", err)
		return err
	}
	output.Write(b)
	return err
}
