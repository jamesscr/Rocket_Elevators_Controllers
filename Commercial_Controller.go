package main

import (
	"fmt"
	"sort"
	"time"
)

// AppController Construct
type AppController struct {
	BatteryNum int
	arraybattery     []Battery
	columnNumber      int
	userInput      string
}


// Battery Construct
type Battery struct {
	columnNumber   int
	columnList []Column
}

// Column Construct
type Column struct {
	ColumnNumber  int
	elvPerColumn int
	ElevatorList  []Elevator
}

// Elevator Construct
type Elevator struct {
	elevatorNumber         int
	currentFloor      int
	floorList         []int
	state    		  string
	Direction 		  string
	doorstate         bool
	Column            Column
}

// NewController from Battery
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


// RequestElevator from userInput
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

// AssignElevator from userInput
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

// FindBestColumn with ResquestFloor
func (b *Battery) FindBestColumn(RequestedFloor int) Column {
	if RequestedFloor >= 1 && RequestedFloor <= 7 { // B6(1) => B1(6) + RC(7)
		return b.columnList[0]
	} else if RequestedFloor > 8 && RequestedFloor <= 26 || RequestedFloor == 7 { //RC(7) + 2=>20=8=>26
		return b.columnList[1]
	} else if RequestedFloor > 27 && RequestedFloor <= 46 || RequestedFloor == 7 { //RC(7) + 21=>40=27=>46
		return b.columnList[2]
	} else if RequestedFloor > 47 && RequestedFloor <= 66 || RequestedFloor == 7 { //RC(7) + 41=>60=47=>60
		return b.columnList[3]
	}
	return b.columnList[3] //need to change 3 for ""
}

// FindBestElevator with RequestFloor, Direction
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

// SendRequest to RequestFloor
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

// OpenDoor Elevator
func (e *Elevator) OpenDoor() {
	fmt.Println("Door is Opening")
	time.Sleep(1 * time.Second)
	fmt.Println("Door is Open")
	time.Sleep(1 * time.Second)
	e.CloseDoor()
}

// CloseDoor Elevator
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

// MoveUp Elevator
func (e *Elevator) MoveUp(RequestedFloor int) {
	fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Current Floor :", + e.currentFloor)
	for RequestedFloor > e.currentFloor {
		e.currentFloor ++
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Reach the destination floor : ", + e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Floor : ", + e.currentFloor)
	}
}

// MoveDown Elevator
func (e *Elevator) MoveDown(RequestedFloor int) {
	fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Current Floor :", + e.currentFloor)
	for RequestedFloor < e.currentFloor {
		e.currentFloor --
		if RequestedFloor == e.currentFloor {
			time.Sleep(1 * time.Second)
			fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Reach the destination floor : ", + e.currentFloor)
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Column : ", + e.Column.ColumnNumber, " Elevator : #", + e.elevatorNumber, " Floor : ", + e.currentFloor)
	}
}

// main function in which test is performed \\
func main()  {
	controller := NewController(1)

	/* Scenario 4:
    With first column (or Column A) serving the basements B1 to B6, with elevator A1 idle at B4, A2 idle at 1st floor, A3 at B3 and going to B5, A4 at B6 and going to 1st floor, and A5 at B1 going to B6, someone is at B3 and requests the 1st floor. Elevator A4 is expected to be sent.
    */
	//column A
	
	// controller.arraybattery[0].columnList[0].ElevatorList[0].currentFloor = 3
	// controller.arraybattery[0].columnList[0].ElevatorList[0].state = "idle"
	// controller.arraybattery[0].columnList[0].ElevatorList[0].Direction = "stop"
	// controller.arraybattery[0].columnList[0].ElevatorList[0].SendRequest(3)
	
	// controller.arraybattery[0].columnList[0].ElevatorList[1].currentFloor = 7
	// controller.arraybattery[0].columnList[0].ElevatorList[1].state = "idle"
	// controller.arraybattery[0].columnList[0].ElevatorList[1].Direction = "stop"
	// controller.arraybattery[0].columnList[0].ElevatorList[1].SendRequest(7)
	
	// controller.arraybattery[0].columnList[0].ElevatorList[2].currentFloor = 4
	// controller.arraybattery[0].columnList[0].ElevatorList[2].state = "moving"
	// controller.arraybattery[0].columnList[0].ElevatorList[2].Direction = "down"
	// controller.arraybattery[0].columnList[0].ElevatorList[2].SendRequest(2)

	controller.arraybattery[0].columnList[0].ElevatorList[3].currentFloor = 1
	controller.arraybattery[0].columnList[0].ElevatorList[3].state = "moving"
	controller.arraybattery[0].columnList[0].ElevatorList[3].Direction = "up"
	controller.arraybattery[0].columnList[0].ElevatorList[3].SendRequest(7)
	
	// controller.arraybattery[0].columnList[0].ElevatorList[4].currentFloor = 6
	// controller.arraybattery[0].columnList[0].ElevatorList[4].state = "moving"
	// controller.arraybattery[0].columnList[0].ElevatorList[4].Direction = "up"
	// controller.arraybattery[0].columnList[0].ElevatorList[4].SendRequest(1)

	// controller.AssignElevator(4)
	controller.RequestElevator(4, 7)

	/* Scenario 1:
    With second column (or column B) serving floors from 2 to 20, with elevator B1 at 20th floor going to 5th, B2 at 3rd floor going to 15th, B3 at 13th floor going to 1st, B4 at 15th floor going to 2nd, and B5 at 6th floor going to 1st, someone is at 1st floor and requests the 20th floor, elevator B5 is expected to be sent
	*/

	// column B

	// controller.arraybattery[0].columnList[1].ElevatorList[0].currentFloor = 26
	// controller.arraybattery[0].columnList[1].ElevatorList[0].state = "moving"
	// controller.arraybattery[0].columnList[1].ElevatorList[0].Direction = "down"
	// controller.arraybattery[0].columnList[1].ElevatorList[0].SendRequest(11)

	// controller.arraybattery[0].columnList[1].ElevatorList[1].currentFloor = 9
	// controller.arraybattery[0].columnList[1].ElevatorList[1].state = "moving"
	// controller.arraybattery[0].columnList[1].ElevatorList[1].Direction = "up"
	// controller.arraybattery[0].columnList[1].ElevatorList[1].SendRequest(21)

	// controller.arraybattery[0].columnList[1].ElevatorList[2].currentFloor = 19
	// controller.arraybattery[0].columnList[1].ElevatorList[2].state = "moving"
	// controller.arraybattery[0].columnList[1].ElevatorList[2].Direction = "down"
	// controller.arraybattery[0].columnList[1].ElevatorList[2].SendRequest(7)

	// controller.arraybattery[0].columnList[1].ElevatorList[3].currentFloor = 21
	// controller.arraybattery[0].columnList[1].ElevatorList[3].state = "moving"
	// controller.arraybattery[0].columnList[1].ElevatorList[3].Direction = "down"
	// controller.arraybattery[0].columnList[1].ElevatorList[3].SendRequest(8)

	// controller.arraybattery[0].columnList[1].ElevatorList[4].currentFloor = 12
	// controller.arraybattery[0].columnList[1].ElevatorList[4].state = "moving"
	// controller.arraybattery[0].columnList[1].ElevatorList[4].Direction = "down"
	// controller.arraybattery[0].columnList[1].ElevatorList[4].SendRequest(7)

	// // controller.AssignElevator(26)
	// controller.RequestElevator(7, 26)

	/* Scenario 2:
    With third column (or column C) serving floors from 21 to 40, with elevator C1 at 1st floor going to 21th, C2 at 23st floor going to 28th, C3 at 33rd floor going to 1st, C4 at 40th floor going to 24th, and C5 at 39nd floor going to 1st, someone is at 1st floor and requests the 36th floor, elevator C1 is expected to be sent
	*/

	//column C

	// controller.arraybattery[0].columnList[2].ElevatorList[0].currentFloor = 7
	// controller.arraybattery[0].columnList[2].ElevatorList[0].state = "moving"
	// controller.arraybattery[0].columnList[2].ElevatorList[0].Direction = "up"
	// controller.arraybattery[0].columnList[2].ElevatorList[0].SendRequest(27)

	// controller.arraybattery[0].columnList[2].ElevatorList[1].currentFloor = 29
	// controller.arraybattery[0].columnList[2].ElevatorList[1].state = "moving"
	// controller.arraybattery[0].columnList[2].ElevatorList[1].Direction = "up"
	// controller.arraybattery[0].columnList[2].ElevatorList[1].SendRequest(34)

	// controller.arraybattery[0].columnList[2].ElevatorList[2].currentFloor = 39
	// controller.arraybattery[0].columnList[2].ElevatorList[2].state = "moving"
	// controller.arraybattery[0].columnList[2].ElevatorList[2].Direction = "down"
	// controller.arraybattery[0].columnList[2].ElevatorList[2].SendRequest(7)

	// controller.arraybattery[0].columnList[2].ElevatorList[3].currentFloor = 46
	// controller.arraybattery[0].columnList[2].ElevatorList[3].state = "moving"
	// controller.arraybattery[0].columnList[2].ElevatorList[3].Direction = "down"
	// controller.arraybattery[0].columnList[2].ElevatorList[3].SendRequest(30)

	// controller.arraybattery[0].columnList[2].ElevatorList[4].currentFloor = 45
	// controller.arraybattery[0].columnList[2].ElevatorList[4].state = "moving"
	// controller.arraybattery[0].columnList[2].ElevatorList[4].Direction = "down"
	// controller.arraybattery[0].columnList[2].ElevatorList[4].SendRequest(7

	// //controller.AssignElevator(27)
	// controller.RequestElevator(27, 42)

	/* Scenario 3:
    With fourth column (or column D) serving floors from 41 to 60, with elevator D1 at 58th floor going to 1st, D2 at 50th floor going to 60th, D3 at 46th floor going to 58th, D4 at 1st floor going to 54th, and D5 at 60th floor going to 1st, someone is at 54th floor and requests the 1st floor, elevator D1 is expected to pick him up
	*/

	//column D

	// controller.arraybattery[0].columnList[3].ElevatorList[0].currentFloor = 64
	// controller.arraybattery[0].columnList[3].ElevatorList[0].state = "moving"
	// controller.arraybattery[0].columnList[3].ElevatorList[0].Direction = "down"
	// controller.arraybattery[0].columnList[3].ElevatorList[0].SendRequest(7)

	// controller.arraybattery[0].columnList[3].ElevatorList[1].currentFloor = 56
	// controller.arraybattery[0].columnList[3].ElevatorList[1].state = "moving"
	// controller.arraybattery[0].columnList[3].ElevatorList[1].Direction = "up"
	// controller.arraybattery[0].columnList[3].ElevatorList[1].SendRequest(66)

	// controller.arraybattery[0].columnList[3].ElevatorList[2].currentFloor = 52
	// controller.arraybattery[0].columnList[3].ElevatorList[2].state = "moving"
	// controller.arraybattery[0].columnList[3].ElevatorList[2].Direction = "down"
	// controller.arraybattery[0].columnList[3].ElevatorList[2].SendRequest(64)

	// controller.arraybattery[0].columnList[3].ElevatorList[3].currentFloor = 7
	// controller.arraybattery[0].columnList[3].ElevatorList[3].state = "moving"
	// controller.arraybattery[0].columnList[3].ElevatorList[3].Direction = "down"
	// controller.arraybattery[0].columnList[3].ElevatorList[3].SendRequest(60)

	// controller.arraybattery[0].columnList[3].ElevatorList[4].currentFloor = 66
	// controller.arraybattery[0].columnList[3].ElevatorList[4].state = "moving"
	// controller.arraybattery[0].columnList[3].ElevatorList[4].Direction = "down"
	// controller.arraybattery[0].columnList[3].ElevatorList[4].SendRequest(7)

    // controller.AssignElevator(60)
    // controller.RequestElevator(60, 7)
}