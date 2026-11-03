package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

type Flags struct {
	Lines      bool
	Words      bool
	Chars      bool
	Bytes      bool
	LineLength bool
}

type FileCounters struct {
	countLines    int
	countWords    int
	countChars    int
	countBytes    int
	maxLineLength int
}

func getFlags(cmd *cobra.Command) Flags {
	var fls, tmp Flags
	fls.Lines, _ = cmd.Flags().GetBool("lines")
	fls.Words, _ = cmd.Flags().GetBool("words")
	fls.Chars, _ = cmd.Flags().GetBool("chars")
	fls.Bytes, _ = cmd.Flags().GetBool("bytes")
	fls.LineLength, _ = cmd.Flags().GetBool("max-line-length")
	if fls == tmp {
		fls.Lines = true
		fls.Words = true
		fls.Bytes = true
	}
	return fls
}

func isOneFlag(fls Flags) bool {
	cnt := 0
	iterFlags := reflect.ValueOf(fls)
	for i := range iterFlags.NumField() {
		if iterFlags.Field(i).Bool() {
			cnt += 1
		}
		if cnt > 1 {
			return false
		}
	}
	if cnt != 1 {
		return false
	}
	return true
}

func getWidth(fls Flags, countBytes int) int {
	if isOneFlag(fls) {
		return 0
	}
	return len(fmt.Sprint(countBytes))
}

func getOutput(fls Flags, fileCounters FileCounters, fileName string) string {
	width := getWidth(fls, fileCounters.countBytes)
	sliceCounters := make([]string, 0)
	iterFlags := reflect.ValueOf(fls)
	iterFileCounters := reflect.ValueOf(fileCounters)
	for i := range iterFlags.NumField() {
		if iterFlags.Field(i).Bool() {
			sliceCounters = append(sliceCounters, fmt.Sprintf("%*d", width, iterFileCounters.Field(i).Int()))
		}
	}
	sliceCounters = append(sliceCounters, fileName)
	return strings.Join(sliceCounters, " ")
}

func wc(cmd *cobra.Command, args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer f.Close()

	var fileCounters FileCounters
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if len(line) != 0 {
			fileCounters.countBytes += len(line)
			tmpLineLength := utf8.RuneCountInString(line)
			fileCounters.countChars += tmpLineLength
			if strings.Contains(line, "\n") {
				tmpLineLength -= 1
			}
			if fileCounters.maxLineLength < tmpLineLength {
				fileCounters.maxLineLength = tmpLineLength
			}
			fileCounters.countLines += strings.Count(line, "\n")
			fileCounters.countWords += len(strings.Fields(line))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	fls := getFlags(cmd)
	output := getOutput(fls, fileCounters, args[0])
	fmt.Printf("%s\n", output)
	return nil
}

func main() {
	rootCmd := &cobra.Command{Use: "mywc", Short: "My wc", Args: cobra.ExactArgs(1), RunE: wc}
	rootCmd.Flags().BoolP("bytes", "c", false, "Command print bytes")
	rootCmd.Flags().BoolP("chars", "m", false, "Command print chars")
	rootCmd.Flags().BoolP("lines", "l", false, "Command print lines")
	rootCmd.Flags().BoolP("max-line-length", "L", false, "Command print line length")
	rootCmd.Flags().BoolP("words", "w", false, "Command print words")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
