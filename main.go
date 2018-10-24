package main

//Yikes.
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Gets rid of a few lines of code.
//Panics are better than completely destroying everything I hold dear.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//Gets working directory/Current Directory (needs to be wrapped in fmt.Print*())
func pwd() string {
	out, err := os.Getwd()
	check(err)
	return out
}

//Moves around the filesystem.
//I can probably get away without returning anything.
func cd(path string) int {
	if path == "" {
		path = "."
	}
	err := os.Chdir(path)
	check(err)
	return 0
}

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
}

//Enumerates Files/Folders in a directory.
//Can also probably get away with returning nothing.
//Need to figure out proper tab alignment.
func ls(path string) int {
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
		}
	}
	return 0
}

func cat(path string) string {
	file, err := ioutil.ReadFile(path)
	check(err)
	return string(file)
}

//Congrats! Here's how it works.
func main() {

	//Making some Declarations...
	var input commandline
	in := bufio.NewReader(os.Stdin)

	//Constantly Waiting to do something...
	for true {
		fmt.Print(pwd(), " $ ")          //Shells need to look pretty.
		line, err := in.ReadString('\n') //Get input
		check(err)
		args := strings.Split(strings.TrimSuffix(line, "\n"), " ") //Take off the newline, split everything by spaces.
		input.command = args[0]                                    //Set the actual command.

		//Performing pop/left shift in the input slice
		copy(args, args[1:])
		args = args[:len(args)-1]

		//Just to stop slice index out-of-bounds...
		if len(args) == 0 {
			args = args[:1]
			args[0] = ""
		}

		//Set the actual arguments to the remaining slice.
		input.arguments = args

		//Comand Switch
		switch input.command {
		case "cd":
			cd(input.arguments[0])
		case "ls":
			ls(input.arguments[0])
		case "cat":
			fmt.Print(cat(input.arguments[0]))
		case "pwd":
			fmt.Println(pwd())
		case "help":
			help("")
		default:
			help("gangweed") //rise the heck up. its gamer time
		}
	}
}

//Self explanatory. Stores the command for the switch and stores an array or arguments.
//Will likely change struct name at a later date.
type commandline struct {
	command   string
	arguments []string
}
