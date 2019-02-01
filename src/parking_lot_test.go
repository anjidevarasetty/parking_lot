package main

import(
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T)  {
	
	fileData := readFromFile("input.txt")

	if fileData == "" {
		t.Errorf("Input file data is mepty")
	}
	
	//assert.Equalf("File read error:",fileData,""Input file content is empty")
}

func TestCreateParkingLot(t *testing.T) {
	createParkingLot("6")
	
	if slotsCapacity == 0 {
		t.Errorf("Bad input format, not able to create parking lot")
	}
}

func TestParkCar(t *testing.T) {
	parkCar("KA-01-HH-1234", "White")
	parkCar("KA-01-HH-9999", "White")

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 

	createParkingLot("6")
	parkCar("KA-01-BB-0001", "Black")
	parkCar("KA-01-HH-2701", "Blue")

	if len(mapSlotCar) >= slotsCapacity {
		t.Errorf("Sorry, parking lot is full")
	}
}

func TestLeaveParking(t *testing.T) {
	leaveParking("3")

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	}
}

func TestGetStatus(t *testing.T) {
	getStatus()

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 

	createParkingLot("6")

	parkCar("KA-01-HH-1234", "White")
	parkCar("KA-01-HH-9999", "White")

	getStatus()

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetRegistration_numbers_for_cars_with_colour(t *testing.T) {
	getRegistration_numbers_for_cars_with_colour("White")

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 

	createParkingLot("6")
	parkCar("KA-01-HH-1234", "White")
	parkCar("KA-01-HH-9999", "White")

	getRegistration_numbers_for_cars_with_colour("White")
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetSlot_numbers_for_cars_with_colour(t *testing.T) {
	getSlot_numbers_for_cars_with_colour("White")

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 

	createParkingLot("6")
	parkCar("KA-01-HH-1234", "White")
	parkCar("KA-01-HH-9999", "White")

	getSlot_numbers_for_cars_with_colour("White")
	
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetSlot_number_for_registration_number(t *testing.T) {
	getSlot_number_for_registration_number("KA-01-HH-1234")

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 

	createParkingLot("6")
	parkCar("KA-01-HH-1234", "White")
	parkCar("KA-01-HH-9999", "White")

	getSlot_number_for_registration_number("KA-01-HH-1234")
	
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}
