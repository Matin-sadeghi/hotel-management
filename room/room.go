package room

import ( 
	"fmt"
	"hotel-management/guest")

type Room struct {
	Number int
	Occupied bool
	Guest *guest.Guest
}

var rooms []Room

func InitRooms(count int) {
	for i := 0; i < count; i++ {
		rooms = append(rooms,Room{Number:i+1,Occupied:false,Guest:nil})
	}
}

func ReserveRoom(number int, guest *guest.Guest){
	for i, room := range rooms {

		if room.Number == number{
			if room.Occupied {
				fmt.Println("This room is already reserved")
			}else{
				rooms[i].Occupied = true
				rooms[i].Guest = guest
				fmt.Printf("Room %d is reserved for %s . \n",number,guest.Name)
			}
			return
		}

		
	}

	fmt.Println("Room Not Found")
}

func CheckoutRoom(number int){
	for _, room := range rooms {

		if room.Number == number{
			if room.Occupied  {
				room.Occupied = false
				room.Guest = nil

				fmt.Println("This room is already check out")
			}else{
				fmt.Println("This room is free")
			}
			return
		}

		
	}

	fmt.Println("Room Not Found")
}