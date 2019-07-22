package Design_Pattern

import (
	"errors"
	"log"
)

type Window interface {
	// The interface defines the general window functionality
	paint() // client can make use of this functionality
	getInfo()
}

func (Win HorizontalScrollDecorator) paint() {
	log.Println(Win.Name + " has drawn")
}

func (Win HorizontalScrollDecorator) getInfo() {
	log.Println(Win.Name + " has description")
}

func (Win VerticalScrollDecorator) paint() {
	log.Println(Win.Name + " has drawn")
}

func (Win VerticalScrollDecorator) getInfo() {
	log.Println(Win.Name + " has description")
}

type HorizontalScrollDecorator struct {
	Name      string
	Scrolling bool
}

type VerticalScrollDecorator struct {
	Name      string
	Scrolling bool
}

type SimpleWindow struct {
	Name string
}

func (Win SimpleWindow) paint() {
	log.Println(Win.Name + " has drawn")
}

func (Win SimpleWindow) getInfo() {

	log.Println(Win.Name + " has description")
}

func (Vz VerticalScrollDecorator) Scroll() string {

	Vz.Scrolling = true
	return Vz.Name + "  has vertical Scrolling"
}

func (Hz HorizontalScrollDecorator) Scroll() string {
	Hz.Scrolling = true
	return Hz.Name + "  has horizontal Scrolling"
}

func CreateWindow(name string, scroll string) (Window Window, err error) { // returns an interface type window
	switch scroll {
	case "Horizontal":
		var Hz HorizontalScrollDecorator
		Hz.Name = name
		log.Println(Hz.Scroll()) // Dynamically adds functionality to the Horizontal window
		return Hz, err
	case "Vertical":
		var Vz VerticalScrollDecorator
		Vz.Name = name
		log.Println(Vz.Scroll()) // Dynamically adds functionality to Vertical window
		return Vz, err

	case "Simple":
		var simpleWindow SimpleWindow
		simpleWindow.Name = name
		return simpleWindow, err
	default:
		err = errors.New("Failed to Create Window :" + name)
	}
	return
}

func main() {
	window1, _ := CreateWindow("Window1", "Simple")
	window2, _ := CreateWindow("Horizontal", "Horizontal")
	window3, _ := CreateWindow("Vertical", "Vertical")

	window1.paint()
	window1.getInfo()

	window2.paint()
	window2.getInfo()

	window3.paint()
	window3.getInfo()

}
