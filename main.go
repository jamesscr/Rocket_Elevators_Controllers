package main

import "fmt"

//Construction of Battery, Columns and Elevators

// for battery
type battery struct{
	columns  []column
	columnsNumber int
	totalFloors int
}

// for columns
type column struct{
	elevatorsQuant []elevator
	elevatorNumber int
	floorsQuant []int
	floorNumber int
	columnSelect []string
	columnID string
}

// for elevator
type elevator struct{
	elevatorID string
	currentFloor int
	currentDirection string
	quant int
}



// main function in which test is performed
func main()  {
	fmt.Println("test")
}