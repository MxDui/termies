package menu

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: -m [menu name]")
	fmt.Println("\tThe menu name is the name of the menu you want to use. The menu name can be one of the following:")
	fmt.Println(" \t -default")
	fmt.Println(" \t -debug")
	fmt.Println(" \t -report")

	fmt.Println("Usage of flags for debug menu:")
	fmt.Println("termies -m [menu name] -[flag name] [flag value]")
	fmt.Println("\t -d [duration] \t The duration flag is used to set the duration of the debug session. The duration flag is only used in the debug menu.")
	fmt.Println("\t -n [program name] \t The program name flag is used to set the name of the program you are debugging. The program name flag is only used in the debug menu.")

	fmt.Println("Usage of flags for report menu:")
	fmt.Println("termies -m report ")

	fmt.Println("Usage of flags for help menu:")
	fmt.Println("termies -m help ")

}
