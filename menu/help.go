package menu

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: -m [menu name]")
	fmt.Println("Available menus:")
	fmt.Println("Menu: default")
	fmt.Println("Menu: debug")
	fmt.Println("Menu: report")
	fmt.Println("Usage of default menu:")
	fmt.Println("termies -m default")
	fmt.Println("The default menu is the main menu of the program. It displays the current data and allows the user to start a debug session.")
	fmt.Println("Usage of debug menu:")
	fmt.Println("termies -m debug -d [duration] -n [program name]")
	fmt.Println("The debug menu is used to debug the program. It starts a debug session and saves the data to the database.")
	fmt.Println("Usage of report menu:")
	fmt.Println("termies -m report")
	fmt.Println("The report menu is used to generate a report of the data in the database.")
	fmt.Println("Usage of help menu:")
	fmt.Println("termies -m help")
	fmt.Println("The help menu is used to display the help menu.")

}
