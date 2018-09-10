package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type parking struct {
	registId string
	color    string
}

func main() {
	var parkingData [1000000]parking
	var parkMax int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		firstSpace := strings.Index(input, " ")
		lastSpace := strings.LastIndex(input, " ")
		length := len(input)
		if strings.Index(input, "create_parking_lot") == 0 {
			parkLength, _ := strconv.Atoi(input[firstSpace+1 : length])
			parkMax = parkLength
			for i := 1; i <= parkMax; i++ {
				parkingData[i].registId = "-1"
				parkingData[i].color = "-1"
			}
			fmt.Printf("Created a parking lot with %d slots\n", parkMax)

		} else if strings.Index(input, "park") == 0 {
			registrationNumber := input[firstSpace+1 : lastSpace]
			color := input[lastSpace+1 : length]
			i := 1
			for parkingData[i].color != "-1" && parkingData[i].registId != "-1" && i <= parkMax {
				i++
			}
			if i > parkMax {
				fmt.Println("Sorry, parking lot is full")
			} else {
				parkingData[i].color = color
				parkingData[i].registId = registrationNumber
				fmt.Printf("Allocated slot number: %d\n", i)
			}
		} else if strings.Index(input, "leave") == 0 {
			indexDelete, _ := strconv.Atoi(input[firstSpace+1 : length])
			parkingData[indexDelete].color = "-1"
			parkingData[indexDelete].registId = "-1"
			fmt.Printf("Slot number %d is free\n", indexDelete)
		} else if strings.Index(input, "status") == 0 {
			fmt.Println("Slot No. Registration No. Colour")
			for i := 1; i <= parkMax; i++ {
				if parkingData[i].color != "-1" && parkingData[i].registId != "-1" {
					fmt.Printf("%d %s %s\n", i, parkingData[i].registId, parkingData[i].color)
				}
			}
		} else if strings.Index(input, "registration_numbers_for_cars_with_colour") == 0 {
			colorFind := input[firstSpace+1 : length]
			firstPrinted := false
			for i := 1; i <= parkMax; i++ {
				if parkingData[i].color == colorFind {
					if firstPrinted == false {
						fmt.Print(parkingData[i].registId)
						firstPrinted = true
					} else {
						fmt.Printf(", %s", parkingData[i].registId)
					}
				}
			}
			fmt.Print("\n")
		} else if strings.Index(input, "slot_numbers_for_cars_with_colour") == 0 {
			colorFind := input[firstSpace+1 : length]
			firstPrinted := false
			for i := 1; i <= parkMax; i++ {
				if parkingData[i].color == colorFind {
					if firstPrinted == false {
						fmt.Print(i)
						firstPrinted = true
					} else {
						fmt.Printf(",% d", i)
					}
				}
			}
			fmt.Print("\n")
		} else if strings.Index(input, "slot_number_for_registration_number") == 0 {
			registerIDFind := input[firstSpace+1 : length]
			i := 1
			for i <= parkMax && parkingData[i].registId != registerIDFind {
				i++
			}
			if i > parkMax {
				fmt.Println("Not found")
			} else {
				fmt.Println(i)
			}
		}
	}
}
