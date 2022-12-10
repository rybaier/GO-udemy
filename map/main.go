package main

import "fmt"

func main() {
	// ways of making maps

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "@4bf745",
	}

	//passing a new key:value pair into map
	colors["white"] = "#ffffff"
	// colors[10] = "#RRRRRR"

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(" Hex code for ", color, "is", hex)
	}
}
