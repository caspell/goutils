package patterns

import (
	"fmt"
)

type GlobalObject struct {
	Name  string
	Value int
}

func (g GlobalObject) GetName() string {
	return g.Name
}

var Value1 string

func PatternSingleTone() {
	fmt.Println("single tone")
}
