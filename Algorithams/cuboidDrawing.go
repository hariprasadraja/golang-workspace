package main

import "fmt"

// DrawCuboid draws a Cuboid in the terminal with the given x,y,z co-oridinate values
// Example
// 		DrawCuboid(5,5,5)
func DrawCuboid(drawX, drawY, drawZ int) {

	drawCubeLine := func(n, drawX, drawY int, cubeDraw string) {
		fmt.Printf("%*s", n+1, cubeDraw[:1])
		for d := 9*drawX - 1; d > 0; d-- {
			fmt.Print(cubeDraw[1:2])
		}

		fmt.Print(cubeDraw[:1])
		fmt.Printf("%*s\n", drawY+1, cubeDraw[2:])
	}

	fmt.Printf("Cuboid %d %d %d:\n", drawX, drawY, drawZ)
	drawCubeLine(drawY+1, drawX, 0, "+-")
	for i := 1; i <= drawY; i++ {
		drawCubeLine(drawY-i+1, drawX, i-1, "/ |")
	}

	drawCubeLine(0, drawX, drawY, "+-|")
	for i := 4*drawZ - drawY - 2; i > 0; i-- {
		drawCubeLine(0, drawX, drawY, "| |")
	}

	drawCubeLine(0, drawX, drawY, "| +")
	for i := 1; i <= drawY; i++ {
		drawCubeLine(0, drawX, drawY-i, "| /")
	}

	drawCubeLine(0, drawX, 0, "+-\n")
}

func main() {
	fmt.Println("Enter 3 dimensions of Cuboid : ")
	var l, b, h int
	fmt.Scanln(&l)
	fmt.Scanln(&b)
	fmt.Scanln(&h)
	DrawCuboid(l, b, h)
}
