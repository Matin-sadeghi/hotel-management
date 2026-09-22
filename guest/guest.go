package guest

type Guest struct {
	ID int
	Name string
}

var GuestIdCounter int = 0
var guests []Guest
func RegisterationGuest(name string) Guest{
	GuestIdCounter ++ 
	guest :=Guest{ID:GuestIdCounter,Name:name}
	guests = append(guests,guest)
	return guest
}

func FindGuestById(ID int) *Guest {
	for _,  guest:= range guests {
		if guest.ID == ID{
			return &guest
		}
	}
	return nil
}