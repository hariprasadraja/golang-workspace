package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "/home/parvathavarthinik/Downloads/UnicodeFonts-IlaSundaram-Set1/UniIla.Sundaram-01.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "")
	pdf.WritePdf("hello.pdf")

}
