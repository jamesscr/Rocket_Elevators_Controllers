//The  Column function contructor has three methods from my column are :- findElevator, requestElevator and requestFloor
//
class Column {
	constructor(floors, elevators) {
		this.floors = floors;
		this.elevators = elevators;
		this.arrayExtButton = [];
		this.arrayElevators = [];
		for (let i of elevators) {
			this.arrayElevators.push(new Elevator(i, floors));
		}
		for (let i of elevators) {
			if (i == 0) {
				this.arrayExtButton.push(new ExtButton(i, "up", false));
			} else {
				this.arrayExtButton.push(new ExtButton(i, "up", false));
				this.arrayExtButton.push(new ExtButton(i, "down", false));
			}
		}


		findElevator(requestedFloor, direction) {
			let selectedElevator = null;
			let bestDiff = this.floors;
			for (let i of this.arrayElevators) {
				if (this.arrayElevators[i].direction === "up" && direction === "up" && requestedFloor > this.arrayElevators[i].currentFloor) {
					selectedElevator = this.arrayElevators[i];
				} else if (this.arrayElevators[i].direction === "down" && direction === "down" && requestedFloor < this.arrayElevators[i].currentFloor) {
					selectedElevator = this.arrayElevators[i];
				} else if (this.arrayElevators[i].status === "idle") {
					selectedElevator = this.arrayElevators[i];
				} else {
					for (let i of this.arrayElevators) {
						let diff = Math.abs(this.arrayElevators[i].currentFloor - requestedFloor);
						if (diff < bestDiff) {
							selectedElevator = this.arrayElevators[i];
							bestDiff = diff;
						}
					}
				}
			}
			return selectedElevator;
		}

		requestElevator(requestedFloor, direction) {
			if (requestedFloor > this.floors) return console.log("Floor " + requestedFloor + " doesn't exist!");

			let elevator = this.findElevator(requestedFloor, direction);

			elevator.addToQueue(requestedFloor);
			elevator.move();
			return elevator;
		}


		requestFloor(elevator, requestedFloor) {
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
		
		for(let i of this.floors){
			this.InButtonsList.push(new InButton(i));
		}
	}
	addToQueue() {
		this.queue.push(requestedFloor)

		if (this.direction == "up") {
			this.queue.sort((a, b) => a - b)
		}
		if (this.direction == "down") {
			this.queue.sort((a, b) => b - a)
		}
        //console.log(this.queue.join(", "));
	}
	move() {
		while (this.queue.length > 0) {

			let operate = this.queue[0];

			if (this.door === "open") {
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
		this.closeDoors();
		this.status = "is idle";
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
	// class ExtButton {

	// }

	// class InButton {

	// }




	// Test
	// console.log("\n\n_____________________senario_______________________\n\n\n")

	// function test1_requestElevator() {

	//code


// test1_requestElevator();

// console.log("\n\n______________________senario_______________________\n\n\n")

// function test2_requestElevator() {

	//code
// }

// test2_requestElevator();