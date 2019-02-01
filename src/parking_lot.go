
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
	"sort"
	"unsafe"
)

type Car struct {
	//SlotNumber int
	RegNumber string
	Colour string
}

var slotsCapacity int
var availableSlots []int
//var parkingInfo = make(map[int]interface{})
var mapSlotCar map[int]Car
var mapRegNumSlotNum map[string]int
var mapColourRegNums map[string][]string

func main() {
	
	if len(os.Args) >= 2 {
		fileData := readFromFile(os.Args[1])
	
		if fileData != "" {
			input := strings.Split(fileData, "\n")

			for _, line := range input {
				//args := strings.Split(line, " ")
				args := strings.Fields(line)
				switch args[0] {
				case "create_parking_lot":
					createParkingLot(args[1])

				case "park":
					parkCar(args[1], args[2])
				
				case "leave":
					leaveParking(args[1])
				
				case "status":
					getStatus()
				
				case "registration_numbers_for_cars_with_colour":
					getRegistration_numbers_for_cars_with_colour(args[1])

				case "slot_numbers_for_cars_with_colour":
					getSlot_numbers_for_cars_with_colour(args[1])
					
				case "slot_number_for_registration_number":
						getSlot_number_for_registration_number(args[1])

				case "default":
					fmt.Println("Incorrect input format")
				}
			}
		}
	} else {
		fmt.Print("Please provide input:\n")

		for {
			//fmt.Scan(&input)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			//input = strings.Replace(input, "\n", "", -1)
			
			if strings.TrimSpace(input) == "" {
				continue
			}

			args := strings.Fields(input)
			
			switch args[0] {
			case "create_parking_lot":
				createParkingLot(args[1])

			case "park":
				parkCar(args[1], args[2])
			
			case "leave":
				leaveParking(args[1])
			
			case "status":
				getStatus()
			
			case "registration_numbers_for_cars_with_colour":
				getRegistration_numbers_for_cars_with_colour(args[1])

			case "slot_numbers_for_cars_with_colour":
				getSlot_numbers_for_cars_with_colour(args[1])
				
			case "slot_number_for_registration_number":
				getSlot_number_for_registration_number(args[1])

			default:
				fmt.Println("Incorrect input format")
			}
		}
	}
}

func readFromFile(filepath string) string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File read error: ", err)
		return ""
	}
	strData := string(data)
	return strData
}

func createParkingLot(slots string) {
	nSlots, err := strconv.Atoi(slots)

	if err != nil {
		fmt.Println("Bad input format, not able to create parking lot")
	} else {
		fmt.Println("Created a parking lot with", nSlots, "slots")
		slotsCapacity = nSlots
		mapSlotCar = make(map[int]Car)
		mapRegNumSlotNum = make(map[string]int)
		mapColourRegNums = make(map[string][]string)
	}
	//parkingInfo = //create size
}

func parkCar(CarRegNumber, CarColour string) {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	} 

	if len(mapSlotCar) >= slotsCapacity {
		fmt.Println("Sorry, parking lot is full")
	} else {
		car := Car {
			RegNumber:  CarRegNumber,
			Colour: CarColour,
		}
	
		var slot int

		if len(availableSlots) > 0 {
			sort.Ints(availableSlots)
			slot = availableSlots[0]
		} else {
			slot = len(mapSlotCar) + 1
		}	

		mapSlotCar[slot] = car
		mapRegNumSlotNum[CarRegNumber] = slot
		
		regNumList := mapColourRegNums[CarColour]
		regNumList = append(regNumList, CarRegNumber)
		mapColourRegNums[CarColour] = regNumList

		fmt.Println("Allocated slot number:", slot)
		//availableSlots = append(availableSlots[:0], availableSlots[1:]...)
	}	
}

func leaveParking(strSlot string) {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	} 

	if len(mapSlotCar) > 0 {
		slot, _ := strconv.Atoi(strSlot)
		car := mapSlotCar[slot]

		if unsafe.Sizeof(car) != 0 {
			delete(mapSlotCar, slot)
			delete(mapRegNumSlotNum, car.RegNumber)

			availableSlots = append(availableSlots, slot)
			fmt.Println("Slot number", slot, "is free")
		} else {
			fmt.Println("Car in slot number", slot, "already left")
		}
	} else {
		fmt.Println("Parking lot is empty")
	}
}

func getStatus() {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	} 
	
	if len(mapSlotCar) > 0 {
		fmt.Println("Slot No.\t Registration No \t Colour")
		
		var keys []int
		for k := range mapSlotCar {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		//for slotNumber, car := range mapSlotCar {
		for _, k := range keys {
			car := mapSlotCar[k]
			fmt.Println(k,"\t\t", car.RegNumber, "\t\t", car.Colour)
		}
	} else {
		fmt.Println("Parking lot is empty")
	}
}

func getRegistration_numbers_for_cars_with_colour(colour string) {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	} 

	regNumList := mapColourRegNums[colour]
	var regNumbers string

	if len(regNumList) > 0 {
		for _, regNumber := range regNumList {
			regNumbers += regNumber + ", "
		}

		fmt.Println(regNumbers[:len(regNumbers)-2])		//to remove last ", "
	} else {
		fmt.Println("Not found")
	}
}

func getSlot_numbers_for_cars_with_colour(colour string) {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	} 

	regNumList := mapColourRegNums[colour]
	var strSlotNums string
	
	if len(regNumList) > 0 {
		for _, regNumber := range regNumList {
			strSlotNums += strconv.Itoa(mapRegNumSlotNum[regNumber]) + ", "
		}

		fmt.Println(strSlotNums[:len(strSlotNums)-2])		//to remove last ", "
	} else {
		fmt.Println("Not found")
	}
}

func getSlot_number_for_registration_number(regNumber string) {
	if slotsCapacity == 0 {
		fmt.Println("Sorry, parking lot is not created")
		return
	}

	slotNumber := mapRegNumSlotNum[regNumber]

	if slotNumber != 0 {
		fmt.Println(slotNumber)
	} else {
		fmt.Println("Not found")
	}
}
