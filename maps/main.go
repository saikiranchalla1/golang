package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#FFFFFF"

	fmt.Println(colors)
	fmt.Println(colors["white"])

	delete(colors, "white")

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " has Hexacode of " + hex)
	}
}
