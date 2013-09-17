package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Flag expected from the commmand line
var createFlag = flag.Bool("c", false, `Create file \n`)

// TODO integrate with this
var createMigration = flag.Bool("m", false, `Create migration / schema file \n`)
var lineNumbers = flag.Int("l", 100, `Number of lines to save to notes \n`)

// TODO::
// create a flag for type of note / line storage
// Store the blob / text into the db
// Process lines based on certain criteria?
// call commands based on flags
// clean up command line parsing

func main() {
	flag.Parse()
	fmt.Println("Creating a note with", *lineNumbers, "line numbers.")

	// Load all the strings into a process
	ch := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		close(ch)
	}()

	// Process all the strings from the std input
	var max = *lineNumbers
	var count int = 0
loop:
	for {
		count++

		select {
		case s, ok := <-ch:
			if !ok {
				break loop
			}
			fmt.Println(s)
		}
		if count > max {
			count = 0
			fmt.Println("reached max lines ")
		}

	}
	fmt.Println("saved note. <id here>")
}
