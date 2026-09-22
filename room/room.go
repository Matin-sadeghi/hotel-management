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
	for _, i := range rooms {

		if rooms[i].Number == number{
			if rooms[i].Occupied {
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