package room

import "hotel-management/guest"

type Room struct {
	Number int
	Occupied bool
	Guest *guest.Guest
}