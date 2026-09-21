package guest

type Guest struct {
	ID int
	Name string
}

var GuestIdCounter int = 0
func RegisterationGuest(name string) Guest{
	GuestIdCounter ++ 
	guest :=Guest{ID:GuestIdCounter,Name:name}
	return guest
}