package main

//Yikes.
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var input userInput //User input (command and arguments)
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
	fmt.Println("\texec\t--\tExecutes a file based on the path given.\n\t\t\tUse -n(oreturn) or -r(eturn) to get the output.\n\t\t\texec [-OPTION] [FILE]")
	fmt.Println()
	fmt.Println("\thelp\t--\tLists this dialogue.")
	fmt.Println()
	fmt.Println("\tos\t--\tReturns the OS.")
	fmt.Println()
	fmt.Println("\trm\t--\tDeletes the specified file")
	fmt.Println()
	fmt.Println("\trf\t--\tDeletes the specified folder")
	fmt.Println()
	fmt.Println("\tremote\t--\tDownloads a file from url to the desired path/name.\n\t\t\tUSAGE: remote [URL] [PATH/FILE]")
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

//DownloadFile The code behind remote()
func DownloadFile(filepath string, url string) {

	// Create the file
	out, err := os.Create(filepath)
	check(err)
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	check(err)
}

//Downloads a file to the current directory (or whatever directory specified by the path prefixing the desired filename)
func remote(url string, path string) {
	DownloadFile(path, url)
} //End remote()

//Absolute clusterfuck. Will comment better at a later date.
//Oh, by the way, it's the execute command. <3
func ex(commands []string) {
	var passedArgs userInput //Going to filter out the first option (-n/-r) using the same pop/leftshift alg
	passedArgs.cmd = commands[0]
	passedArgs.argv = sliceShorten(commands)

	if passedArgs.cmd == "-n" { //The option to execute without returning anything
		if runtime.GOOS == "windows" { //Windows doesn't even run this, but I included it anyway LOL
			err := exec.Command("cmd", passedArgs.argv...).Run() //Just checking if there's a problem with running the file. 'cmd -c' may also be a viable option for the command string.
			check(err)
			fmt.Println("Probably succeeded.") //There is no way to know if you don't return the output :)

		} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" { //MacOS or Linux
			err := exec.Command("/bin/bash", passedArgs.argv...).Run()
			check(err)
			fmt.Println("Probably succeeded.)")

		} else {
			fmt.Println("Weird OS detected (How did you even get this to compile???")

		}

	} else if passedArgs.cmd == "-r" {
		if runtime.GOOS == "windows" {
			cmdOut, err := exec.Command("cmd", passedArgs.argv...).Output()
			check(err)
			fmt.Print(string(cmdOut), '\n')

		} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
			cmdOut, err := exec.Command("/bin/bash", passedArgs.argv...).Output()
			check(err)
			fmt.Print(string(cmdOut), '\n')

		} else {
			fmt.Println("Weird OS detected (How did you even get this to compile???")
		}

	} else {
		fmt.Println("Learn to RTFM and try again.") //No options specified
	}
} //End ex()

//Deletes a file.
//rf should be used for folders (safety)
func rm(path string) {
	fileInfo, infoerr := os.Stat(path)
	check(infoerr)
	if fileInfo.IsDir() == true {
		fmt.Println("This is a directory.\nHint: use 'rf'")
	} else {
		err := os.Remove(path)
		check(err)
	} //End conditional
} //End rm()

//Removes a folder.
//rm should be used for files (safety)
func rf(path string) {
	folderInfo, infoerr := os.Stat(path)
	check(infoerr)
	if folderInfo.IsDir() == false {
		fmt.Println("This is a file.\nHint: use 'rm'")
	} else {
		err := os.Remove(path)
		check(err)
	} //End conditional
} //End rf()

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

		//Performing pop/left shift in the input slice, assigns to the input struct.
		input.argv = sliceShorten(args)

		//Comand Switch
		cmdSwitch(input.cmd)
	} //This never ends (for true loop)
} //End main()

//Self explanatory. Stores the command for the switch and stores an array or arguments.
type userInput struct {
	cmd  string
	argv []string
} //End Structure definition.

func sliceShorten(slice []string) []string { //Pop/left shift on a given slice
	copy(slice, slice[1:])
	slice = slice[:len(slice)-1]
	if len(slice) == 0 {
		slice = slice[:1]
		slice[0] = ""
	} //End conditional
	return slice
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
	case "os":
		fmt.Println(runtime.GOOS)
	case "rm":
		rm(input.argv[0])
	case "rf":
		rf(input.argv[0])
	case "remote":
		remote(input.argv[0], input.argv[1])
	default:
		help("gangweed") //rise tf up. its gamer time
	} //End switch
} //End cmdSwitch()
