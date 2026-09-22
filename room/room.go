package room

import "hotel-management/guest"

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