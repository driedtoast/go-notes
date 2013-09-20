package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Flag expected from the commmand line
var createFlag = flag.Bool("c", false, `Create file \n`)
var tagFlag = flag.String("t", "default", `Tag the note with values / Search notes based on list option \n`)
var listFlag = flag.Bool("l", true, `Lists the notes in storage\n`)

// TODO integrate with this
var createMigration = flag.Bool("m", false, `Create migration / schema file \n`)
var lineNumbers = flag.Int("n", 100, `Number of lines to save to notes \n`)

// TODO::
// create a flag for type of note / line storage
// Store the blob / text into the db
// Process lines based on certain criteria?
// call commands based on flags
// clean up command line parsing

type CommandInput struct {
	tags string
	body string
}

func main() {
	flag.Parse()
	fmt.Println("Creating a note with", *lineNumbers, "line numbers.")
	fmt.Println("List flag ", *listFlag)

	// Load all the strings into a process
	// TODO need to get the strings via non params as well
	ch := make(chan CommandInput)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			ch <- CommandInput{*tagFlag, scanner.Text()}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		close(ch)
	}()

	// Process all the strings from the std input
	var max = *lineNumbers
	var count int = 0
	var tags string

loop:
	for {
		count++

		select {
		case commandInput, ok := <-ch:
			if !ok {
				fmt.Println("saved note. <id here>")
				break loop
			}
			tags = commandInput.tags
			fmt.Println(commandInput.body)
		}
		if count > max {
			count = 0
			fmt.Println("reached max lines ")
			break
		}

	}
	fmt.Print(tags)
	fmt.Println("saved note. <id here>")
}
