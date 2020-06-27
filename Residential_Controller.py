# defining my kind of blueprint
class Column():
    def __init__(self, floors, elevators):
        self.floors = floors
        self.elevators = elevators
        self.arrayButtons = []
        self.arrayelevators = []
        for i in range(elevators):
            self.arrayelevators.append(Elevator(0, floors))
        for i in range(floors):
            if i == 0:
                self.arrayButtons.append(ExtButton(i, "up"))
            else:
                self.arrayButtons.append(ExtButton(i, "up"))
                self.arrayButtons.append(ExtButton(i, "down"))

    def findElevator(self, requestedFloor, direction):

        bestDiff = self.floors
        chosenElevator = None

        for elevator in self.arrayelevators:
            if elevator.direction == "up" and direction == "up" and requestedFloor < elevator.currentFloor:
                chosenElevator = elevator
            elif elevator.direction == "down" and direction == "down" and requestedFloor > elevator.currentFloor:
                chosenElevator = elevator
            elif elevator.status == "idle":
                chosenElevator = elevator
            else:
                for elevator in self.arrayelevators:
                    diff = abs(elevator.currentFloor - requestedFloor)
                    if diff < bestDiff:
                        chosenElevator = elevator
                        bestDiff = diff
        print("Best elevator on floor " + str(chosenElevator.currentFloor))
        return chosenElevator

    def requestElevator(self, requestedFloor, direction):
        print("Called an elevator to the floor " + str(requestedFloor))

        elevator = self.findElevator(requestedFloor, direction)

        elevator.addToQueue(requestedFloor)
        elevator.move()
        return elevator

    def requestFloor(self, elevator, requestedFloor):
        print("Moving elevator on floor " + str(elevator.currentFloor) + " to the floor " + str(requestedFloor))
        elevator.addToQueue(requestedFloor)
        elevator.closeDoors()
        elevator.move()

class Elevator():
    def __init__(self, currentFloor, floors):
        self.direction =  None
        self.floors = floors
        self.currentFloor = currentFloor
        self.status = "idle"
        self.queue = []
        self.IntButtonsList = []
        self.door = "closed"

        for i in range(self.floors):
            self.IntButtonsList.append(IntButton(i))

    def addToQueue(self, requestedFloor):
        self.queue.append(requestedFloor)

        if self.direction == "up":
            self.queue.sort(reverse=True)
        if self.direction == "down":
            self.queue.sort(reverse=True)

        print("Added floor " + str(requestedFloor) + " elevators driving passengers on " + ', '.join(str(x) for x in self.queue))

    def move(self):
        print ("Moving elevator")
        #while len function to returns the number 
        while len(self.queue) > 0:

            operate = self.queue[0]

            if self.door == "open":
                print("within 5 seconds if passage is not obstructed")
                self.closeDoors()
            if operate == self.currentFloor:
                del self.queue[0]
                self.openDoors()
            elif operate > self.currentFloor:
                self.status = "moving"
                self.direction = "up"
                self.moveUp()
            elif operate < self.currentFloor:
                self.status = "moving"
                self.direction = "down"
                self.moveDown()
        print("within 5 seconds if passage is not obstructed")
        self.closeDoors()
        print("Elevator is now idle")
        self.status = "idle"

    def moveUp(self):
        self.currentFloor += 1
        print("moving up Elevator on floor " + str(self.currentFloor))

    def moveDown(self):
        self.currentFloor -= 1
        print("moving down Elevator on floor " + str(self.currentFloor))

    def openDoors(self):
        self.door = "open"
        print("open door on floor ")

    def closeDoors(self):
        self.door="closed"
        print("close door on floor ")

class ExtButton():
    def __init__(self, requestFloor, direction):
        self.requestFloor = requestFloor
        self.direction = direction

class IntButton():
    def __init__(self, floor):
        self.floor = floor


# testing #


def test_requestElevator():
                column1 = Column(10, 2)
                column1.arrayelevators[0].currentFloor = 6
                column1.arrayelevators[0].direction  =  "up"
                column1.arrayelevators[0].status =  "idle"
                column1.arrayelevators[0].queue = [1, 5, 8]
                column1.arrayelevators[1].currentFloor = 3
                column1.arrayelevators[1].direction  =  "up"
                column1.arrayelevators[1].status =  "moving"
                column1.arrayelevators[1].queue = [2, 1]
                
                column1.requestElevator(4, "down")
                column1.requestElevator(9, "down")

test_requestElevator()
