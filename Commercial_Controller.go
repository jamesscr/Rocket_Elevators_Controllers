package main

import (
	"fmt"
	"sort"
	"time"
)

// Construction of Battery, Columns and Elevators
type AppController struct {
	BatteryNum int
	arraybattery     []Battery
	columnNumber      int
	userInput      string
}


// for battery \\
type Battery struct {
	columnNumber   int
	columnList []Column
}

// for columns \\
type Column struct {
	ColumnNumber  int
	elvPerColumn int
	ElevatorList  []Elevator
}

// for elevator \\
type Elevator struct {
	elevatorNumber         int
	currentFloor      int
	floorList         []int
	state    		  string
	Direction 		  string
	doorstate         bool
	Column            Column
}

// New Controller, Create Battery
func NewController(BatteryNum int) AppController {
	controller := new(AppController)
	controller.BatteryNum = 1
	for index := 0; index < BatteryNum; index++ {
		battery := NewBattery(index)
		controller.arraybattery = append(controller.arraybattery, *battery)
	}
	return *controller
}

// NewBattery to create column
func NewBattery(columnNumber int) *Battery {
	battery := new(Battery)
	battery.columnNumber = 4
	for index := 0; index < battery.columnNumber; index++ {
		Column := NewColumn(index)
		battery.columnList = append(battery.columnList, *Column)
	}
	return battery
}


// NewColumn to create Elevator 
func NewColumn(elvPerColumn int) *Column {
	Column := new(Column)
	Column.elvPerColumn = 5
	for index := 0; index < Column.elvPerColumn; index++ {
		Elevator := NewElevator()
		Column.ElevatorList = append(Column.ElevatorList, *Elevator)
	}
	return Column
}

// NewElevator to define elevator
func NewElevator() *Elevator {
		Elevator := new(Elevator)
		Elevator.currentFloor = 7
		Elevator.floorList = []int{}
		Elevator.state = "idle"
		Elevator.Direction = "stop"
		Elevator.doorstate = true
		return Elevator
	}


// from userInput for RequestElevator
func (controller *AppController) RequestElevator(FloorNumber, RequestedFloor int) Elevator { 
	fmt.Println("Request Elevator to floor : ", FloorNumber)
	time.Sleep(2000 * time.Millisecond)
	var Column = controller.arraybattery[0].FindBestColumn(FloorNumber)
	controller.userInput = "down"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userInput)
	Elevator.SendRequest(FloorNumber)
	Elevator.SendRequest(RequestedFloor)
	return Elevator
}

// from userInput for AssignElevator
func (controller *AppController) AssignElevator(RequestedFloor int) Elevator {
	fmt.Println("Request Elevator to floor : ", RequestedFloor)
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Button Light On")
	Column := controller.arraybattery[0].FindBestColumn(RequestedFloor) 
	controller.userInput = "up"
	var Elevator = Column.FindBestElevator(RequestedFloor, controller.userInput)
	var FloorNumber = 7
	Elevator.SendRequest(FloorNumber)
	Elevator.SendRequest(RequestedFloor)
	return Elevator
}

//find the best Column
func (b *Battery) FindBestColumn(RequestedFloor int) Column { 
	if RequestedFloor >= 1 && RequestedFloor <= 7 {
		return b.columnList[0]
	} else if RequestedFloor > 8 && RequestedFloor <= 27 || RequestedFloor == 7 {
		return b.columnList[1]
	} else if RequestedFloor > 28 && RequestedFloor <= 47 || RequestedFloor == 7 {
		return b.columnList[2]
	} else if RequestedFloor > 48 && RequestedFloor <= 66 || RequestedFloor == 7 {
		return b.columnList[3]
	}
	return b.columnList[3] //need to change 3 for ""
}

//find the beat elevator
func (c *Column) FindBestElevator(RequestedFloor int, userInput string) Elevator {
	var selectedElevator = c.ElevatorList[4] 
	for _, e := range c.ElevatorList {
		if RequestedFloor < e.currentFloor && e.Direction == "down" && userInput == "down" {
			selectedElevator = e
		} else if e.state == "idle" {
			selectedElevator = e
		} else if e.Direction != userInput && e.state == "moving" || e.state == "stopped" {
			selectedElevator = e
		} else if e.Direction == userInput && e.state == "moving" || e.state == "stopped" {
			selectedElevator = e
		}
	}
	return selectedElevator
}

//request Elevator
func (e *Elevator) SendRequest(RequestedFloor int) {
	e.floorList = append(e.floorList, RequestedFloor)
	if RequestedFloor > e.currentFloor {

		sort.Ints(e.floorList)
	} else if RequestedFloor < e.currentFloor {

		sort.Sort(sort.Reverse(sort.IntSlice(e.floorList)))
	}
	e.OperateElevator(RequestedFloor)
}

//OperateElevator movement
func (e *Elevator) OperateElevator(RequestedFloor int) {
	if RequestedFloor == e.currentFloor {
		e.OpenDoor()
	} else if RequestedFloor > e.currentFloor {
		e.state = "moving"
		e.MoveUp(RequestedFloor)
		e.state = "stopped"
		e.OpenDoor()
		e.state = "moving"
	} else if RequestedFloor < e.currentFloor {
		e.state = "moving"
		e.MoveDown(RequestedFloor)
		e.state = "stopped"
		e.OpenDoor()
		e.state = "moving"
	}
}

//function OpenDoor and CloseDoor
func (e *Elevator) OpenDoor() {
	fmt.Println("Door is Opening")
	time.Sleep(1 * time.Second)
	fmt.Println("Door is Open")
	time.Sleep(1 * time.Second)
	e.CloseDoor()
}
func (e *Elevator) CloseDoor() {
	if e.doorstate == true {
		fmt.Println("Door is Closing")
		time.Sleep(1 * time.Second)
		fmt.Println("Door is Close")
		time.Sleep(1 * time.Second)
	} else if e.doorstate {
		e.OpenDoor()
		fmt.Println("Door can not be close please make sur door is not obstruct")
	}
}

//function to mMoveUp
func (e *Elevator) MoveUp(RequestedFloor int) {
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Current Floor :", e.currentFloor)
	for RequestedFloor > e.currentFloor {
		e.currentFloor += 1
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Reach the destination floor : ", e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Floor : ", e.currentFloor)
	}
}

// function to MoveDown
func (e *Elevator) MoveDown(RequestedFloor int) {
	fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Current Floor :", e.currentFloor)
	for RequestedFloor < e.currentFloor {
		e.currentFloor -= 1
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Reach the destination floor : ", e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", e.Column.ColumnNumber, " Elevator : #", e.elevatorNumber, " Floor : ", e.currentFloor)
	}
}

// main function in which test is performed \\
func main()  {
	controller := NewController(1)
	
	/* Scenario 4:
    With first column (or Column A) serving the basements B1 to B6, with elevator A1 idle at B4, A2 idle at 1st floor, A3 at B3 and going to B5, A4 at B6 and going to 1st floor, and A5 at B1 going to B6, someone is at B3 and requests the 1st floor. Elevator A4 is expected to be sent.
    */
	//column A

	controller.arraybattery[0].columnList[0].ElevatorList[0].currentFloor = 3
	controller.arraybattery[0].columnList[0].ElevatorList[0].state = "idle"
	controller.arraybattery[0].columnList[0].ElevatorList[0].Direction = "stop"
	// controller.arraybattery[0].columnList[0].ElevatorList[0].SendRequest(3)

	// controller.arraybattery[0].columnList[0].ElevatorList[1].currentFloor = 7
	// controller.arraybattery[0].columnList[0].ElevatorList[1].state = "idle"
	// controller.arraybattery[0].columnList[0].ElevatorList[1].Direction = "stop"
	// // controller.arraybattery[0].columnList[0].ElevatorList[1].SendRequest(7)

	// controller.arraybattery[0].columnList[0].ElevatorList[2].currentFloor = 4
	// controller.arraybattery[0].columnList[0].ElevatorList[2].state = "moving"
	// controller.arraybattery[0].columnList[0].ElevatorList[2].Direction = "down"
	// //controller.arraybattery[0].columnList[0].ElevatorList[2].SendRequest(2)

	// controller.arraybattery[0].columnList[0].ElevatorList[3].currentFloor = 1
	// controller.arraybattery[0].columnList[0].ElevatorList[3].state = "moving"
	// controller.arraybattery[0].columnList[0].ElevatorList[3].Direction = "up"
	// //controller.arraybattery[0].columnList[0].ElevatorList[3].SendRequest(7)

	// controller.arraybattery[0].columnList[0].ElevatorList[4].currentFloor = 6
	// controller.arraybattery[0].columnList[0].ElevatorList[4].state = "moving"
	// controller.arraybattery[0].columnList[0].ElevatorList[4].Direction = "down"
	// //controller.arraybattery[0].columnList[0].ElevatorList[4].SendRequest(1)

	// // controller.AssignElevator(1)
	controller.RequestElevator(4, 7)
}