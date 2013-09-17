package datamapper

import (
	"bytes"
	"fmt"
	"reflect"
)

func table_def(m interface{}) string {
	var buffer bytes.Buffer

	// TODO need to get name
	buffer.WriteString(" CREATE TABLE ")
	buffer.WriteString(get_interface(m).Name())
	buffer.WriteString(" (")
	start := true

	for name, mtype := range attributes(m) {
		if len(mtype.Name()) == 0 {
			continue
		}
		if start {
			start = false
		} else {
			buffer.WriteString(",")
		}
		buffer.WriteString(name)
		buffer.WriteString(" ")
		buffer.WriteString(db_type(mtype.Name()))
	}
	buffer.WriteString(" );")
	return buffer.String()
}

func db_type(field_type string) string {
	db_type := "varchar(255)"
	if field_type == "int" {
		db_type = "integer"
	}
	return db_type
}

func get_interface(m interface{}) reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := get_interface(m)

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}
