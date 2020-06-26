# defining my kind of blueprint
class Column():
    def __init__(self, floors, elevators):
        self.floors = floors
        self.elevators = elevators
        self.externalButtonsList = []
        self.elevatorsList = []
        
                
    #method findElevator with parameters self, requestedFloor and direction
    def findElevator(self, requestedFloor, direction):

    def requestElevator(self, requestedFloor, direction):
        

    def requestFloor(self, elevator, requestedFloor):
       
class Elevator():
    def __init__(self, currentFloor, floors):
        self.direction =  None
        self.floors = floors
        self.currentFloor = currentFloor
        self.status = "idle"
        self.queue = []
        self.internalButtonsList = []
        self.door = "close"


    def addToQueue(self, requestedFloor):
        self.queue.append(requestedFloor)

    def move(self):
        

    def moveUp(self):
      

    def moveDown(self):
        

    def openDoors(self):
      

    def closeDoors(self):
       
class ExternalButton():
  

class InternalButton():
    
#testing

def test1_requestElevator():
    

#Test1_requestElevator()



def test2_requestFloor():
    

#Test2_requestFloor()