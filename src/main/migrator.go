package main

import (
	"datamapper"
	"models"
)

func main() {

	// iterate through the attributes of a Data Model instance
	for name, mtype := range datamapper.attributes(&Note{}) {
		fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	}

	fmt.Printf(" TABLE DEF: %s", datamapper.table_def(&Note{}))
	fmt.Printf(" TABLE DEF: %s", datamapper.table_def(&NoteList{}))
	fmt.Printf(" TABLE DEF: %s", datamapper.table_def(&NoteMapping{}))
}
