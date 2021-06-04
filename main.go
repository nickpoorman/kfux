// Copyright 2021 Nick Poorman
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func usage() {
	log.Println("Usage: kfux [-pp]")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var pretty = flag.Bool("p", false, "Pretty print by json decoding and splitting into new lines")
	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	if *pretty {
		if err := prettyPrint(os.Stdin, os.Stdout); err != nil {
			writeError("d", err)
		}
	}
}

func prettyPrint(in io.Reader, out io.Writer) error {
	// Read all the data in
	data, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	var line string
	if err := json.Unmarshal(data, &line); err != nil {
		return err
	}

	// Split on all the newlines
	lines := splitNewLines(line)

	// Write it out
	for _, line := range lines {
		if _, err := fmt.Fprintf(out, "%s\n", line); err != nil {
			return err
		}
	}

	return nil
}

func writeError(op string, err error) {
	fmt.Fprintln(os.Stderr, fmt.Errorf("error doing %s: %w", op, err).Error())
}

func splitNewLines(line string) []string {
	if line == "" {
		return nil
	}

	lines := splitNewLine(line)
	if len(lines) == 1 {
		return lines
	}

	var newLines []string
	for _, l := range lines {
		spls := splitNewLines(l)
		newLines = append(newLines, spls...)
	}
	return newLines
}

func splitNewLine(line string) []string {
	return strings.Split(line, "\n")
}

func removePrefix(line string, char string) string {
	for strings.HasPrefix(line, char) {
		line = strings.Replace(line, char, "", 1)
	}
	return line
}

func removeSuffix(line string, char string) string {
	for strings.HasSuffix(line, char) {
		line = line[:strings.LastIndex(line, char)]
	}
	return line
}
