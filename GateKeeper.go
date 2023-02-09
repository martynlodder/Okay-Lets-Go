package main

import (
	"fmt"
	"time"
)

func timeCheck() {
	datetime := time.Now()  // Fetch the current date and time
	time := datetime.Hour() // Fetch only the hour

	if time > 7 && time < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken!")
	} else if time >= 12 && time < 17 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken!")
	} else if time > 17 && time <= 23 {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten.")
	}
}

func main() {
	timeCheck()
}
