package main

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/golang/freetype"
	"image"
	"fmt"
	"bytes"
	"image/jpeg"
	"io/ioutil"
	"image/draw"
	"log"
)

func main() {

	dataFont, err := ioutil.ReadFile("/home/parvathavarthinik/Downloads/UnicodeFonts-IlaSundaram-Set1/Link to UniIla.Sundaram-01.ttf")
	f, err := freetype.ParseFont(dataFont)
	if err != nil {
		fmt.Printf("%v", err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, 800, 600))
	draw.Draw(dst, dst.Bounds(), image.White, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(image.Black)
	c.SetFont(f)
	c.DrawString("யளனகயளக", freetype.Pt(0, 16))
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, dst, nil)

	if err != nil {
		fmt.Printf("%v", err)
	}

	reader := bytes.NewReader(buf.Bytes())
	//textName := "text1"
	//pdf.RegisterImageReader(textName, "jpg", reader)
	pdf.AddFontFromReader("new", "", reader)
	pdf.SetFont("new", "", 35)
	pdf.Cell(0, 10, `தயாதளமனகரைந`)
	fileStr := "gofpdf.pdf"
	err = pdf.OutputFileAndClose(fileStr)
	if err != nil {
		log.Print(err)
	}

	//pdf.Image(textName, 15, 15, 0, 0, false, "jpg", 0, "")
	//pdf.OutputFileAndClose("test.pdf")
}
