package main

//Yikes.
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var input userInput //User inpute (command and arguments)
var args []string   //Temporary argument storage

//Gets rid of a few lines of code.
//Panics are better than completely destroying everything I hold dear.
func check(err error) {
	if err != nil {
		panic(err)
	} //End conditional
} //End check()

//Gets working directory/Current Directory (needs to be wrapped in fmt.Print*())
func pwd() string {
	out, err := os.Getwd()
	check(err)
	return out
} //End pwd()

//Moves around the filesystem.
//I can probably get away without returning anything.
func cd(path string) {
	if path == "" {
		path = "."
	}
	err := os.Chdir(path)
	check(err)
} //End cd()

//The help command (or lack thereof.)
func help(command string) {
	if command != "" {
		fmt.Print("Invalid command.\n\n")
	}
	fmt.Println("Commands:\t\tDescription:")
	fmt.Println()
	fmt.Println("\tcd\t--\tChanges current/present working directory.")
	fmt.Println()
	fmt.Println("\tls\t--\tEnumerates all files & folders in a directory.\n\t\t\tLists file sizes in bytes.")
	fmt.Println()
	fmt.Println("\tpwd\t--\tReturns the current working directory.\n\t\t\tWhy would you use this, the shell does it already...")
	fmt.Println()
	fmt.Println("\tcat\t--\tPrints a file to the shell.")
	fmt.Println()
	fmt.Println("\texec\t--\tComing soon!")
	fmt.Println()
	fmt.Println("\thelp\t--\tLists this dialogue.")
} //End help()

//Enumerates Files/Folders in a directory.
//Need to figure out proper tab alignment.
func ls(path string) {
	if path == "" {
		path = "."
	}
	files, err := ioutil.ReadDir(path)
	check(err)
	for i, file := range files {
		file = files[i]
		if file.IsDir() == true {
			fmt.Println(file.Name(), "\t", "DIR")
		} else {
			fmt.Println(file.Name(), "\t", "FILE", "\t", file.Size())
		} //End conditional
	} //End for loop
} //End ls()

func cat(path string) string {
	file, err := ioutil.ReadFile(path)
	check(err)
	return string(file)
} //End cat()

//Congrats! Here's how it works.
func main() {
	in := bufio.NewReader(os.Stdin)

	//Constantly Waiting to do something...
	for true {
		fmt.Print(pwd(), " $ ")          //Shells need to look pretty.
		line, err := in.ReadString('\n') //Get input
		check(err)
		args = strings.Split(strings.TrimSuffix(line, "\n"), " ") //Take off the newline, split everything by spaces.
		input.cmd = args[0]                                       //Set the actual command.

		//Performing pop/left shift in the input slice
		sliceShorten(args)

		//Set the actual arguments to the remaining slice.
		input.argv = args

		//Comand Switch
		cmdSwitch(input.cmd)
	} //This never ends (for true loop)
} //End main()

//Self explanatory. Stores the command for the switch and stores an array or arguments.
type userInput struct {
	cmd  string
	argv []string
} //End Structure definition.

func sliceShorten(slice []string) { //Pop/left shift on the arguments from the raw input.
	copy(slice, slice[1:])
	args = slice[:len(slice)-1]
	if len(args) == 0 {
		args = args[:1]
		args[0] = ""
	} //End conditional
} //End sliceShorten()

func cmdSwitch(command string) {
	switch command {
	case "cd":
		cd(input.argv[0])
	case "ls":
		ls(input.argv[0])
	case "cat":
		fmt.Print(cat(input.argv[0]))
	case "pwd":
		fmt.Println(pwd())
	case "help":
		help("")
	default:
		help("gangweed") //rise tf up. its gamer time
	} //End switch
} //End cmdSwitch()
