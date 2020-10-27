package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff000",
		"black": "#ff83783",
		"green": "#f98098",
	}

	colors["white"] = "#fffff"
	delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for c, h := range m {
		fmt.Println("Colors is: ", c, " and hex is: ", h)
	}
}
