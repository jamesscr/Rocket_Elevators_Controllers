//The  Column function contructor has three methods from my column are :- findElevator, requestElevator and requestFloor
//
class Column {
	constructor(floors, elevators) {
		this.floors = floors;
		this.elevators = elevators;
		this.arrayExtButton = [];
        this.arrayElevators = [];
        //console.log(this.arrayExtButton);
        //console.log(this.arrayElevators);
        
    

	findElevator(requestedFloor, direction) {

		
	}

	requestElevator(requestedFloor, direction) {
		
        
    }
    

	requestFloor(elevator, requestedFloor) {
		
    }
    
    
}

class Elevator {
	constructor(currentFloor, floors) {

		this.direction = null;
		this.floors = floors;
		this.currentFloor = currentFloor;
		this.status = "idle";
		this.queue = [];
		this.InButtonsList = [];
		this.door = "closed";

		
	}
	addToQueue() {
		
	}
	move() {
		
	}
	moveUp() {
		
	}

	moveDown() {
		
	}

	openDoors() {
		
	}

	closeDoors() {
		
	}

}

class ExtButton {
	
}

class InButton {
	
}




//Test
console.log("\n\n_____________________senario_______________________\n\n\n")

function testing1_requestElevator() {

	
}

testing1_requestElevator();

console.log("\n\n______________________senario_______________________\n\n\n")

function Test2_requestElevator() {

	
}

test2_requestElevator();