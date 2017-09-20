package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
	"io/ioutil"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "/home/parvathavarthinik/kpm/goworkspace/src/github.com/jung-kurt/gofpdf/font/")
	//pdf.AddFont("Calligrapher", "", "TSCu_SaiIndira.json")
	//pdf.AddFont("Calligrapher", "", "SaiEmbed-forPDF.json")

	json, _ := ioutil.ReadFile("/home/parvathavarthinik/kpm/goworkspace/src/github.com/jung-kurt/gofpdf/font/UniIla.Sundaram-01.json")

	zip, _ := ioutil.ReadFile("/home/parvathavarthinik/kpm/goworkspace/src/github.com/jung-kurt/gofpdf/font/UniIla.Sundaram-01.z")

	pdf.AddFontFromBytes("Calligrapher", "", json, zip)
	pdf.AddPage()
	//ln := pdf.UnicodeTranslatorFromDescriptor("")
	log.Print(`t`)
	pdf.SetCompression(true)
	pdf.SetFont("Calligrapher", "", 35)
	pdf.Cell(0, 10, `தயாதளமனகரைந`)
	fileStr := "gofpdf.pdf"

	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		log.Print(err)
	}

}
