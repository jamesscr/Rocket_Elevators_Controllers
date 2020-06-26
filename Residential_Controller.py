# defining my kind of blueprint
class Column():
    def __init__(self, floors, elevators):
        self.floors = floors
        self.elevators = elevators
        self.extButtonsList = []
        self.elevatorsList = []
        print("This building has " + elevators + " elevators driving passengers on " + floors + " floors")
        for i in range(elevators):
            self.elevatorsList.append(Elevator(0, floors))
        for i in range(floors):
            if i == 0:
                self.extButtonsList.append(extButton(i, "up"))
            else:
                self.extButtonsList.append(extButton(i, "up"))
                self.extButtonsList.append(extButton(i, "down"))

    # method findElevator with parameters self, requestedFloor and direction

    def findElevator(self, requestedFloor, direction):
        chosenElevator = None
        bestDiff = self.floors

        for elevator in self.elevatorsList:
            if elevator.direction == "up" and direction == "up" and requestedFloor < elevator.currentFloor:
                chosenElevator = elevator
            elif elevator.direction == "down" and direction == "down" and requestedFloor > elevator.currentFloor:
                chosenElevator = elevator
            elif elevator.status == "idle":
                chosenElevator = elevator
            else:
                for elevator in self.elevatorsList:
                    diff = abs(elevator.currentFloor - requestedFloor)
                    if diff < bestDiff:
                        chosenElevator = elevator
                        bestDiff = diff
                    print("Best elevator on floor " + str(chosenElevator.currentFloor))    
                    return chosenElevator

    def requestElevator(self, requestedFloor, direction):
        print("Call for an elevator to the floor " + str(requestedFloor))
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
                self.direction = None
                self.floors = floors
                self.currentFloor = currentFloor
                self.status = "idle"
                self.queue = []
                self.intButtonsList = []
                self.door = "close"
                for i in range(self.floors):
                    self.intButtonsList.append(IntButton(i))

            def addToQueue(self, requestedFloor):
                self.queue.append(requestedFloor)

                if self.direction == "up":
                    self.queue.sort(reverse=True)
                if self.direction == "down":
                    self.queue.sort(reverse=True)
                print("Added floor " + str(requestedFloor) + " to the elevator's queue. Current queue: " + ', '.join(str(x) for x in self.queue))    

            def move(self):
                # while len function to returns the number
                while len(self.queue) > 0:

                    operate = self.queue[0]

            if self.door == "open":
                print("Waiting 7 seconds for the doorway to be cleared")
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
                self.door = "closed"
                print("close door on floor ")

        class ExtButton():
            def __init__(self, requestFloor, direction):
                self.requestFloor = requestFloor
                self.direction = direction

        class IntButton():
            def __init__(self, floor):
                self.floor = floor

            # testing

            # def test1_requestElevator():
                column1 = Column(10, 2)
                column1.elevatorsList[0].currentFloor = 1
                column1.elevatorsList[0].direction  =  "up"
                column1.elevatorsList[0].status =  "moving"
                column1.elevatorsList[0].queue = [4, 6, 7]
                column1.elevatorsList[1].currentFloor = 6
                column1.elevatorsList[1].direction  =  "down"
                column1.elevatorsList[1].status =  "moving"
                column1.elevatorsList[1].queue = [4, 3]
                column1.requestElevator(1, "down")

            # Test1_requestElevator()

            # def test2_requestFloor():

            # Test2_requestFloor()
