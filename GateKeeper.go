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

func licenseCheck(licensePlate string) {

	// Registered licence plates
	licensePlates := [4]string{"645A990", "DYP990", "L0D090", "SUV003"}

	if licensePlate == licensePlates[0] {
		timeCheck()
	} else if licensePlate == licensePlates[1] {
		timeCheck()
	} else if licensePlate == licensePlates[2] {
		timeCheck()
	} else if licensePlate == licensePlates[3] {
		timeCheck()
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein.")
	}

}

func main() {
	var license string

	fmt.Print("Uw kenteken: ")
	fmt.Scan(&license)

	licenseCheck(license)
}
