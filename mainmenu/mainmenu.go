package mainmenu

import "fmt"

func MainMenu()  {
	for {

		fmt.Println("\n ---Hotel Management Syestem ---")
		fmt.Println("1. Register Guest")
		fmt.Println("2. Reserve Room")
		fmt.Println("3. Check out for guest")
		fmt.Println("4. View Room status")
		fmt.Println("5. exit")
		fmt.Println("Choose an option :")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			var guestName string
			fmt.Println("Enter guest name:")
			fmt.Scanln(&guestName)
			fmt.Println("**** Guest created ****")

		case 2:
			var guestID int
			var roomNumber int

			fmt.Println("Enter guest id:")
			fmt.Scanln(&guestID)
			fmt.Println("Enter room number:")
			fmt.Scanln(&roomNumber)



			
		}




	


	}
}