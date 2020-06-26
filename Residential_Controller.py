# defining my kind of blueprint
class Column():
    def __init__(self, floors, elevators):
        self.floors = floors
        self.elevators = elevators
        self.extButtonsList = []
        self.elevatorsList = []
        for i in range(elevators):
            self.elevatorsList.append(Elevator(0, floors))
        for i in range(floors):
            if i == 0:
                self.extButtonsList.append(ExtButton(i, "up"))
            else:
                self.extButtonsList.append(ExtButton(i, "up"))
                self.extButtonsList.append(ExtButton(i, "down"))

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
                    return chosenElevator

    def requestElevator(self, requestedFloor, direction):
        elevator = self.findElevator(requestedFloor, direction)

        elevator.addToQueue(requestedFloor)
        elevator.move()
        return elevator

    def requestFloor(self, elevator, requestedFloor):
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
                   
            def move(self):
                #while len function to returns the number 
                while len (self.queue) > 0:
                    
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
                
                self.closeDoors()
                
                self.status = "idle"
                
               

            def moveUp(self):

            def moveDown(self):

            def openDoors(self):

            def closeDoors(self):

        class ExtButton():

        class IntButton():

        # testing

        # def test1_requestElevator():

        # Test1_requestElevator()

        # def test2_requestFloor():

        # Test2_requestFloor()
