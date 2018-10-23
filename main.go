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
		//case "cat": 			//Not Yet implemented.
		//	cat()				//Thus commented out.
		case "pwd":
			fmt.Println(pwd())
		}
	}
}

//Self explanatory. Stores the command for the switch and stores an array or arguments.
//Will likely change struct name at a later date.
type commandline struct {
	command   string
	arguments []string
}
