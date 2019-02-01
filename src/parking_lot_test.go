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
	if slotsCapacity != 0 {
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	} else {
		createParkingLot("6")
		
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	}

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	}
}

func TestLeaveParking(t *testing.T) {
	if slotsCapacity != 0 {
		leaveParking("2")
	} else {
		createParkingLot("6")
		
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
		leaveParking("2")
	}

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	}
}

func TestGetStatus(t *testing.T) {
	if slotsCapacity == 0 {
		createParkingLot("6")

		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	}
	getStatus()

	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetRegistration_numbers_for_cars_with_colour(t *testing.T) {
	if slotsCapacity == 0 {
		createParkingLot("6")
	
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	}
	
	getRegistration_numbers_for_cars_with_colour("White")
	
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetSlot_numbers_for_cars_with_colour(t *testing.T) {
	if slotsCapacity == 0 {
		createParkingLot("6")
	
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	}

	getSlot_numbers_for_cars_with_colour("White")
	
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}

func TestGetSlot_number_for_registration_number(t *testing.T) {
	if slotsCapacity == 0 {
		createParkingLot("6")
	
		parkCar("KA-01-HH-1234", "White")
		parkCar("KA-01-HH-9999", "White")
	}

	getSlot_number_for_registration_number("KA-01-HH-1234")
	
	if slotsCapacity == 0 {
		t.Errorf("Sorry, parking lot is not created")
	} 
}