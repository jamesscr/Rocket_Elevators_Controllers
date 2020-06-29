package main

import "fmt"

//Constructures
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



// main function
func main()  {
	fmt.Println("test")
}