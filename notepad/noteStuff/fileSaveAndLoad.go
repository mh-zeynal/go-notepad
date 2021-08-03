package noteStuff

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//this function converts note to json style and
//saves it on a file with .json format
//param: newFile -> a Note that must be saved on file
func SaveFile(newFile Note)  {
	newJson, _ := json.Marshal(newFile)
	address := "notes\\" + newFile.Title + ".json"
	os.Create(address)
	ioutil.WriteFile(address, newJson, 0777)
}

//this function loads a .json file and converts it to Note struct
//and returns a Note struct
func LoadFile(fileName string) Note {
	b, _ := ioutil.ReadFile("notes\\" + fileName + ".json")
	var pingJSON Note
	json.Unmarshal(b, &pingJSON)
	return pingJSON
}

//a modification to scanner that to scan a line with spaces
//and returns an string that has been scanned
func scanner() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.Replace(line, "\n", "", strings.LastIndex(line, "\n"))
	return line
}

//makes a new note and uses SaveFile to save it in "notes" directory
func MakeNewNote() {
	fmt.Print("\033[H\033[2J")
	note := Note{}
	fmt.Print("enter the title of your note:")
	note.Title = scanner()
	fmt.Println("enter your content:")
	note.Content = scanner()
	note.Date = time.Now().String()[:19]
	SaveFile(note)
	fmt.Println("your note has been saved\npress enter to continue...")
	scanner()
}

//loads files in "notes" directory and
//you can choose one of them to see its content
func LoadFiles()  {
	fmt.Print("\033[H\033[2J")
	files, _ := ioutil.ReadDir("notes")
	if len(files) != 0 {
		printFileContent()
	} else {
		fmt.Println("there is no file inside directory")
	}
	fmt.Println("press enter to continue...")
	scanner()
}

//returns an slice of all .json files in "notes" directory
func getFilesList() *[]string {
	files, _ := ioutil.ReadDir("notes")
	list := make([]string, 0)
	for _, file := range files{
		list = append(list, strings.Replace(file.Name(), ".json", "", 1))
	}
	return &list
}

//gets getFilesList and prints its elements
func printFileContent()  {
	list := *getFilesList()
	counter := 1
	for _, fileName := range list{
		fmt.Print(counter); fmt.Print("."); fmt.Println(fileName)
		counter++
	}
	var choice int
	fmt.Print("choose a file:")
	fmt.Scanln(&choice)
	note := LoadFile(list[choice - 1])
	fmt.Println("\033[31m", "\n title:", "\033[0m", note.GetTitle())
	fmt.Println("\033[32m", "create date:", "\033[0m", note.GetDate())
	fmt.Println(note.GetContent())
	fmt.Println()
}

//asks from user to choose a file and
//add some text to it via getNewContentFromUser
func EditFileContent()  {
	fmt.Print("\033[H\033[2J")
	list := *getFilesList()
	counter := 1
	for _, fileName := range list{
		fmt.Print(counter); fmt.Print("."); fmt.Println(fileName)
		counter++
	}
	var choice int
	fmt.Print("choose a file:")
	fmt.Scanln(&choice)
	note := LoadFile(list[choice - 1])
	note.Content += "\n" + getNewContentFromUser()
	SaveFile(note)
	fmt.Println("note has been edited\npress enter to continue...")
	scanner()
}

//gets a text from console and adds it to an intended note
//returns user's input text
func getNewContentFromUser() string {
	fmt.Println("write down your text:")
	newText := scanner()
	return newText
}
