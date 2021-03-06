//The  Column function contructor has three methods from my column are :- findElevator, requestElevator and requestFloor
class Column {
	constructor(floors, elevators) {
		this.floors = floors;
		this.elevators = elevators;
		this.arrayExtButton = [];
		this.arrayElevators = [];
		for (let i = 0; i < elevators; i++) {
			this.arrayElevators.push(new Elevator(0, floors));
		}
		for (let i = 0; i < floors; i++) {
			if (i == 0) {
				this.arrayExtButton.push(new ExtButton(i, "up", false));
			} else {
				this.arrayExtButton.push(new ExtButton(i, "up", false));
				this.arrayExtButton.push(new ExtButton(i, "down", false));
			}
		}
	}
    
	findElevator(requestedFloor, direction) {
		let selectedElevator = null;
		let bestDiff = this.floors;

		for (let i = 0; i < this.arrayElevators.length; i++) {
			if (this.arrayElevators[i].direction === "up" && direction === "up" && requestedFloor > this.arrayElevators[i].currentFloor) {
				selectedElevator = this.arrayElevators[i];
			} else if (this.arrayElevators[i].direction === "down" && direction === "down" && requestedFloor < this.arrayElevators[i].currentFloor) {
				selectedElevator = this.arrayElevators[i];
			} else if (this.arrayElevators[i].status === "idle") {
				selectedElevator = this.arrayElevators[i];
			} else {
				for (let i = 0; i < this.arrayElevators.length; i++) {
					let diff = Math.abs(this.arrayElevators[i].currentFloor - requestedFloor);
					if (diff < bestDiff) {
						selectedElevator = this.arrayElevators[i];
						bestDiff = diff;
					}
				}
			}
		}
		console.log("System has find the best elevator on floor " + selectedElevator.currentFloor);
		return selectedElevator;
	}

	requestElevator(requestedFloor, direction) {
		if (requestedFloor > this.floors) return console.log("Floor " + requestedFloor + " doesn't exist!");
		console.log("Call for an elevator to the floor " + requestedFloor);

		let elevator = this.findElevator(requestedFloor, direction);

		elevator.addToQueue(requestedFloor);
		elevator.move();
		//console.log(elevator);
		return elevator;
	}
	
	requestFloor(elevator, requestedFloor) {
		console.log("Moving elevator on floor " + elevator.currentFloor + " to the floor " + requestedFloor);
		elevator.addToQueue(requestedFloor);
		elevator.closeDoors();
		elevator.move();
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

		for (let i = 0; i < this.floors; i++) {
			this.InButtonsList.push(new InButton(i));
		}
	}
	
	addToQueue(requestedFloor) {
		this.queue.push(requestedFloor)

		if (this.direction == "up") {
			this.queue.sort((a, b) => a - b)
		}
		if (this.direction == "down") {
			this.queue.sort((a, b) => b - a)
		}
		//console.log(this.queue.join(", "));
		console.log("Added floor " + requestedFloor + " to the elevator's queue which is now:  " + this.queue.join(", "));
	}
	
	move() {
		//console.log("Move elevator");
		while (this.queue.length > 0) {

			let operate = this.queue[0];

			if (this.door === "open") {
				console.log("Within 5 seconds the door closes if passage is not obstructed");
				this.closeDoors();
			}
			if (operate === this.currentFloor) {
				this.queue.shift();
				this.openDoors();
			}
			if (operate > this.currentFloor) {
				this.status = "moving";
				this.direction = "up";
				this.moveUp();
			}
			if (operate < this.currentFloor) {
				this.status = "moving";
				this.direction = "down";
				this.moveDown()
			}
		}
		console.log("Within 5 seconds door the closes if passage is not obstructed");
		this.closeDoors();
		console.log("Elevator is now idle on floor " + this.currentFloor);
		this.status = "is idle";
	}
	
	moveUp() {
		this.currentFloor++;
		console.log("Moving up Elevator on floor " + this.currentFloor);
	}

	moveDown() {
		this.currentFloor--;
		console.log("Moving down Elevator on floor " + this.currentFloor);
	}

	openDoors() {
		this.door = "open"
		console.log("Open door on floor " + this.currentFloor);
	}

	closeDoors() {
		this.door = "closed"
		console.log("Close door on floor " + this.currentFloor);
	}

}

class ExtButton {
	constructor(requestFloor, direction) {
		this.requestFloor = requestFloor;
		this.direction = direction;
	}
}

class InButton {
	constructor(floor) {
		this.floor = floor;
	}
}

// Test

console.log("\nTest\n")

function test1_requestElevator() {
	column = new Column(10, 2);
	column.arrayElevators[1].currentFloor = 2
	column.arrayElevators[1].direction = null
	column.arrayElevators[1].status = "idle"
	column.arrayElevators[1].queue = []
	column.arrayElevators[0].currentFloor = 6
	column.arrayElevators[0].direction = null
	column.arrayElevators[0].status = "idle"
	column.arrayElevators[0].queue = []
	
	//Calling the function to request an elevator
	column.requestElevator(7, "up");
}
test1_requestElevator();
// //test 2
 
// function test2_requestFloor() {
// 	column2 = new Column(10, 2);
// 	column2.arrayElevators[1].currentFloor = 3
// 	column2.arrayElevators[1].direction = "up"
// 	column2.arrayElevators[1].status = "moving"
// 	column2.arrayElevators[1].queue = [3]
	
// 	column2.requestFloor(7)
// }

// test2_requestFloor()