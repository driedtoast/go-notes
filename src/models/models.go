package models

type Note struct {
	id   int64
	name string
}

type NoteList struct {
	id   int64
	name string
}

type NoteMapping struct {
	note_id int64
	list_id int64
}

// TODO stop wording and ngram analysis
