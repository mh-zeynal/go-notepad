package main

import (
	"fmt"
	"notepad/noteStuff"
)

func main() {
	for {
		fmt.Print("1.write a new note\n2.read a note\n3.edit a note\n4.exit\n->")
		var userChoice int
		fmt.Scanln(&userChoice)
		if userChoice == 1{
			noteStuff.MakeNewNote()
		} else if userChoice == 2{
			noteStuff.LoadFiles()
		}else if userChoice == 3{
			noteStuff.EditFileContent()
		} else if userChoice == 4{
			break
		}
	}

}

