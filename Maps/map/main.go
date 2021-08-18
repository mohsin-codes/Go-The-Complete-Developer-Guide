package main

import "fmt"

func main() {
	//Declaring and initializing maps
	// 1. First Method
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"purple": "#800080",
		"white": "#ffffff",
		"black": "#000000",
		"Yellow": "#ffff00",
		"Pink": "#ffc0cb",
		"Grey": "#808080",

	}

	// 2. Second Method
	// var color map[string]string
	// color["purple"] = "#800080"
	// color["white"] = "#ffffff"
	// color["black"] = "#000000"

	// 3. Third Method
	// var colors2 = make(map[string]string)
	// colors2["Yellow"] = "#ffff00"
	// colors2["Pink"] = "#ffc0cb"
	// colors2["Grey"] = "#808080"

	fmt.Println(colors)
	// fmt.Println(color)
	// fmt.Println(colors2)

	printMap(colors)
}

func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println(color, hex)
	}
}
