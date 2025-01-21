package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	CREATE = "create_parking_lot"
	PARK   = "park"
	LEAVE  = "leave"
	STATUS = "status"
)

func main() {

	var parkingLot []string
	var freeParkingSlot []int
	var space int

	inputFile, err := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	csvRead := csv.NewReader(inputFile)

	for {

		record, err := csvRead.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		command := strings.Split(record[0], " ")

		// fmt.Println(record)
		switch {
		case strings.EqualFold(command[0], CREATE):

			n, err := strconv.Atoi(command[1])

			if err != nil {
				log.Fatal(err)
			}

			space = n

			fmt.Printf("Create %d Parking Lot\n", n)
		case strings.EqualFold(command[0], PARK):

			if len(parkingLot) == space && len(freeParkingSlot) == 0 {

				fmt.Println("Sorry, parking lot is full")

			} else {

				var slot int

				if len(freeParkingSlot) > 0 {

					slot = freeParkingSlot[0]
					freeParkingSlot = freeParkingSlot[1:]
					parkingLot[slot] = command[1]

				} else {

					parkingLot = append(parkingLot, command[1])
					slot = len(parkingLot) - 1

				}

				fmt.Printf("Allocated slot number: %d\n", slot+1)

			}

		case strings.EqualFold(command[0], LEAVE):

			hour, err := strconv.Atoi(command[2])
			if err != nil {
				log.Fatal(err)
			}

			found := false
			tempI := 0
			for i, v := range parkingLot {
				if strings.EqualFold(v, command[1]) {
					parkingLot[i] = ""
					tempI = i
					freeParkingSlot = append(freeParkingSlot, i)
					found = true
					continue
				}
			}

			if !found {
				fmt.Printf("Registration number %s not found\n", command[1])
			} else {

				var charge int

				if hour <= 2 {
					charge = 10
				} else {
					charge = 10 + (hour-2)*10
				}

				fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", command[1], tempI+1, charge)
			}

			for i := 0; i < len(freeParkingSlot); i++ {

				for j := i + 1; j < len(freeParkingSlot); j++ {

					if freeParkingSlot[i] > freeParkingSlot[j] {
						temp := freeParkingSlot[i]
						freeParkingSlot[i] = freeParkingSlot[j]
						freeParkingSlot[j] = temp

					}

				}

			}

		case strings.EqualFold(command[0], STATUS):
			fmt.Println("Slot No. Registration No.")
			for i, v := range parkingLot {
				if v == "" {
					continue

				}
				fmt.Println(i+1, " "+v)
			}
		default:
			fmt.Println("Invalid Command")

		}

	}

}
