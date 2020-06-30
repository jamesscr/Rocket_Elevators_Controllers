package main

import "fmt"

// Construction of Battery, Columns and Elevators
// for battery \\
type battery struct{
	columns  []column
	columnsNumber int
	totalFloors int
}

// for columns \\
type column struct{
	elevatorsQuant []elevator
	elevatorNumber int
	floorsQuant []int
	floorNumber int
	columnSelect []string
	columnID string
}

// for elevator \\
type elevator struct{
	elevatorID string
	currentFloor int
	currentDirection string
	quant int
}

func battery(columnsNumber int, elevatorNumber int)  *battery{
	battery.columnsNumber = columnsNumber
	columns := []column {column[0], column[1], column[2], column[3]}
	
	//loop through the list of columns (B0,C1,C2,C3)
	if i := 0; i <=columnsNumber; i++) {
		if c==0 {
			column := column{
				columnID: "columnB0",
				elevatorNumber: elevatorNumber,
				floorNumber: 66,
			}
			column.append(column)
		}else if c==1 {
			column := column{
				columnID: "columnC1",
				elevatorNumber: elevatorNumber,
				floorNumber: 66,
			}
			columns.append(column)
		}else if c==2 {
			column := column{
				columnID: "columnC2",
				elevatorNumber: elevatorNumber,
				floorNumber: 66,
			}
			columns.append(column)
		}else if c==2 {
			column := column{
				columnID: "columnC3",
				elevatorNumber: elevatorNumber,
				floorNumber: 66,
			}
			columns.append(column)
		}
	}
}

// constructor for columns \\
func column(columnID string, floorNumber int, elevatorNumber int)  {
	this.columnID=columnID
	this.elevatorNumber=elevatorNumber
	this.floorNumber=floorNumber
	    elevators:= [5]{elevator[0], elevator[1], elevator[2], elevator[3], elevator[4]}
    for e:=0; e<elevatorNumber; e++{
		if e==0 {
			elevator:= elevator{
				elevatorID: "elevator" + (e+1),
			}
			elevators.append(elevator)
		}
	}
}

// constructor for elevators \\
func elevator(elevatorID string)  {
	this.elevatorID=elevatorID
}

// request elevator function to chose which elevator \\
func requestElevator()  {
	
}

// function to find best elevator\\
func findBestElevator()  {
	
}

// main function in which test is performed \\
func main()  {
	fmt.Println("test")
}