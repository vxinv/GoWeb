package main
 type Employee struct {
 	ID int
 	Name string
	Address string
	Sarary int
	Position string 
 }
 func  main(){
  var dilbert Employee
  dilbert.Sarary=5000;
  position:=&dilbert.Position;
  *position="Senior"+*position;
  var employOfTheMonth *Employee=&dilbert
  employOfTheMonth.Position+="lalal"
  
 }

