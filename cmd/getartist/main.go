package getartist

import "fmt"

type Artist struct {
	Name string
}

func (artist Artist) Printer() { fmt.Printf("\nThis is a test %s", artist.Name) }
