package main

import (
	"os"
	"text/template"
)

type Header struct {
	Background string
	Color      string
}

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	change := &Header{"white", "black"}
	tmpl, err := template.New("template").ParseFiles("design.css")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(w, "design.css", change)
	if err != nil {
		panic(err)
	}
	//sweaters := Inventory{"wool", 17}
	//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	//if err != nil { panic(err) }
	//err = tmpl.Execute(os.Stdout, sweaters)
	//if err != nil { panic(err) }

	//filePath := "design.css"
	//
	//var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	//defer file.Close()
	//
	//// write some text to file
	//file.WriteString("halo\n")
	//
	//file.WriteString("mari belajar golang\n")
	//
	//
	//// save changes
	//err = file.Sync()
	//if err !=nil{
	//	panic(err)
	//}

}
