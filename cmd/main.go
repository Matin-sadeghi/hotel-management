package main

import (
	"hotel-management/mainmenu"
	"hotel-management/room"
	"fmt"
)


func main(){
	var numberOfRooms int
	fmt.Println("Enter number of rooms:")
	fmt.Scanln(&numberOfRooms)
	room.InitRooms(numberOfRooms)

	mainmenu.MainMenu()
}