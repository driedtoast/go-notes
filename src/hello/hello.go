package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

type ThingOne struct {
	Name string
}

type ThingTwo struct {
	ThingOne
}

func (thing ThingOne) displayName() (string) {
	return thing.Name
}


var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m["Bell Labs"].test())

    thingTwo := ThingTwo{ThingOne{"hello"}}
    fmt.Println("THIS THING  %s",thingTwo.displayName())
}

// Adds a method to vertex
func (vertx Vertex) test() (float64, float64) {
  return vertx.Lat, vertx.Long
}
