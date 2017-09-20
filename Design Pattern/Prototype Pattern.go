package Design_Pattern

import (
	"log"
	"runtime"
)

type Struct1 struct {
	Name string
}

func main() {

	// In golang struct are immutable, if we

	a := Struct1{"Old things"}
	b := a
	b.Name = "New Value"
	runtime.GC()
	log.Println("A:", a.Name, "B:", b.Name)

	c := &Struct1{"Old things"}
	d := a
	d.Name = "New Value"

	log.Println("C:", c.Name, "D:", d.Name)

}
