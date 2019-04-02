package main

import "log"

type Window interface {
	// The interface defines the general window functionality
	draw() // client can make use of this functionality
	getDescription()
}

type HorizontalScrollDecorator struct {
	Name      string
	Scrolling bool
	Window    *SimpleWindow
}

type VerticalScrollDecorator struct {
	Name      string
	Scrolling bool
	Window    *SimpleWindow
}

type SimpleWindow struct {
	Name string
}

func (Win SimpleWindow) draw() {
	log.Println(Win.Name + " has drawn")
}

func (Win SimpleWindow) getDescription() {

	log.Println(Win.Name + " has description")
}

func (Win HorizontalScrollDecorator) Scroll() string {
	Win.Scrolling = true
	return Win.Name + "  has horizontal Scrolling"
}

func (Win VerticalScrollDecorator) Scroll() string {
	Win.Scrolling = true
	return Win.Name + "  has vertical Scrolling"
}

func (Win HorizontalScrollDecorator) draw() {
	Win.Window.draw()
	log.Println(Win.Name + "  has drawing")
}

func (Win VerticalScrollDecorator) draw() {
	Win.Window.draw()
	log.Println(Win.Name + "  has drawing")
}

func (Win HorizontalScrollDecorator) getDescription() {
	Win.Window.getDescription()
	log.Println(Win.Name + " has description")
}

func (Win VerticalScrollDecorator) getDescription() {
	Win.Window.getDescription()
	log.Println(Win.Name + " has description")
}

func HorizontalScrollWindow(name string, window *SimpleWindow) Window {
	var horizontalWindow HorizontalScrollDecorator
	horizontalWindow.Name = "Horizontal Window"
	horizontalWindow.Scrolling = true
	horizontalWindow.Window = window
	log.Println(horizontalWindow.Scroll())
	return horizontalWindow
}

func VerticalScrollWindow(name string, window *SimpleWindow) Window {
	var horizontalScrollWindow VerticalScrollDecorator
	horizontalScrollWindow.Name = "Vertical Window"
	horizontalScrollWindow.Scrolling = true
	horizontalScrollWindow.Window = window
	log.Println(horizontalScrollWindow.Scroll())
	return horizontalScrollWindow
}

func main() {
	var simplewindow SimpleWindow
	simplewindow.Name = "Simple window"
	window1 := HorizontalScrollWindow("Window1", &simplewindow)
	window1.draw()
	window1.getDescription()

	window2 := VerticalScrollWindow("Window2", &simplewindow)
	window2.draw()
	window2.getDescription()
}
