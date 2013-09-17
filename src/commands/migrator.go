package commands

import (
	"datamapper"
	"fmt"
	"models"
)

func run() {

	// iterate through the attributes of a Data Model instance
	for name, mtype := range datamapper.Attributes(&models.Note{}) {
		fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	}

	fmt.Printf(" TABLE DEF: %s", datamapper.Table_def(&models.Note{}))
	fmt.Printf(" TABLE DEF: %s", datamapper.Table_def(&models.NoteList{}))
	fmt.Printf(" TABLE DEF: %s", datamapper.Table_def(&models.NoteMapping{}))
}
