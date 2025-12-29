package main

import (
	"fmt"
	"strings"
)

var display Display

func main() {
	// Print intro message
	fmt.Println("Homework 5: Geometry Using Go Interfaces")
	fmt.Println("CS 341, Fall 2025")
	fmt.Println("")
	fmt.Println("This application allows you to draw various shapes")
	fmt.Println("of different colors via interfaces in Go.")
	fmt.Println("")

	//
	// SOME PRINT STATEMENTS YOU WILL NEED CAN BE FOUND BELOW
	//
	// Ask user for dimensions for display
	var width, height int
	fmt.Print("Enter the number of rows (x) that you would like the display to have: ")
	fmt.Scan(&width)
	fmt.Print("Enter the number of columns (y) that you would like the display to have: ")
	fmt.Scan(&height)

	// Set up the display
	display.initialize(width, height)

	// Main loop for shape drawing
	for {
		//
		// Menu options
		fmt.Println("")
		fmt.Println("Select a shape to draw: ")
		fmt.Println("	 R for a rectangle")
		fmt.Println("	 T for a triangle")
		fmt.Println("	 C for a circle")
		fmt.Println(" or X to stop drawing shapes.")
		fmt.Print("Your choice --> ")

		var userInput string
		fmt.Scan(&userInput)
		userInput = strings.ToUpper(userInput)

		if userInput == "X" {
			break
		} else if userInput == "R" {
			//
			// Drawing a rectangle
			var lowerX, lowerY, upperX, upperY int
			var colorName string

			fmt.Print("Enter the X and Y values of the lower left corner of the rectangle: ")
			fmt.Scan(&lowerX, &lowerY)
			fmt.Print("Enter the X and Y values of the upper right corner of the rectangle: ")
			fmt.Scan(&upperX, &upperY)
			fmt.Print("Enter the color of the rectangle: ")
			fmt.Scan(&colorName)
			colorName = strings.ToLower(colorName)

			// Create the rectangle
			box := Rectangle{
				ll: Point{x: lowerX, y: lowerY},
				ur: Point{x: upperX, y: upperY},
			}
			fmt.Println(box.printShape())

			if outOfBounds(box.ll, &display) || outOfBounds(box.ur, &display) {
				fmt.Println(outOfBoundsErr)
				continue
			}

			selectedColor, colorExists := colorMap[colorName]
			if !colorExists {
				fmt.Println(colorUnknownErr)
				continue
			}
			box.c = selectedColor
			drawErr := box.draw(&display)
			if drawErr != nil {
				fmt.Println(drawErr)
				continue
			}
			fmt.Println("Rectangle drawn successfully.")

		} else if userInput == "T" {
			//
			// Drawing a triangle
			var x1, y1, x2, y2, x3, y3 int
			var colorName string

			fmt.Print("Enter the X and Y values of the first point of the triangle: ")
			fmt.Scan(&x1, &y1)

			fmt.Print("Enter the X and Y values of the second point of the triangle: ")
			fmt.Scan(&x2, &y2)

			fmt.Print("Enter the X and Y values of the third point of the triangle: ")
			fmt.Scan(&x3, &y3)

			fmt.Print("Enter the color of the triangle: ")
			fmt.Scan(&colorName)
			colorName = strings.ToLower(colorName)

			// Create the triangle
			shape := Triangle{
				pt0: Point{x: x1, y: y1},
				pt1: Point{x: x2, y: y2},
				pt2: Point{x: x3, y: y3},
			}

			fmt.Println(shape.printShape())
			if outOfBounds(shape.pt0, &display) || outOfBounds(shape.pt1, &display) || outOfBounds(shape.pt2, &display) {
				fmt.Println(outOfBoundsErr)
				continue
			}

			selectedColor, colorExists := colorMap[colorName]
			if !colorExists || colorUnknown(selectedColor) {
				fmt.Println(colorUnknownErr)
				continue
			}
			shape.c = selectedColor

			if drawErr := shape.draw(&display); drawErr != nil {
				fmt.Println(drawErr)
				continue
			}

			fmt.Println("Triangle drawn successfully.")
		} else if userInput == "C" {
			//
			// Drawing a circle
			var centerPosX, centerPosY, rad int
			var colorName string

			fmt.Print("Enter the X and Y values of the center of the circle: ")
			fmt.Scan(&centerPosX, &centerPosY)

			fmt.Print("Enter the value of the radius of the circle: ")
			fmt.Scan(&rad)

			fmt.Print("Enter the color of the circle: ")
			fmt.Scan(&colorName)
			colorName = strings.ToLower(colorName)

			// Create and draw the circle
			roundShape := Circle{
				center: Point{x: centerPosX, y: centerPosY},
				r:      rad,
			}
			fmt.Println(roundShape.printShape())
			if outOfBounds(Point{x: centerPosX - rad, y: centerPosY - rad}, &display) || outOfBounds(Point{x: centerPosX + rad, y: centerPosY + rad}, &display) {
				fmt.Println(outOfBoundsErr)
				continue
			}

			selectedColor, colorExists := colorMap[colorName]
			if !colorExists || colorUnknown(selectedColor) {
				fmt.Println(colorUnknownErr)
				continue
			}
			roundShape.c = selectedColor
			drawErr := roundShape.draw(&display)
			if drawErr != nil {
				fmt.Println(drawErr)
				continue
			}
			fmt.Println("Circle drawn successfully.")
		} else {
			fmt.Println("**Error, unknown command. Try again.")
		}
	}

	//
	// Saving the results in a file
	var fileName string
	fmt.Print("Enter the name of the .ppm file in which the results should be saved: ")
	fmt.Scan(&fileName)

	saveErr := display.screenShot(fileName)
	if saveErr != nil {
		fmt.Println(saveErr)
	}

	//
	// Exiting program
	fmt.Println("Done. Exiting program...")
}
