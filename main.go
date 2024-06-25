package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

// name used in the repl prompts
var cliName string = "Pokedex"
 
// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
    fmt.Print(strings.ToLower(cliName), "> ")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
    fmt.Printf("\nWelcome to the %v! \n\n", cliName)
    fmt.Println("Usage:")
    fmt.Println("help: Displays a help message")
    fmt.Printf("exit: Exit the %v \n\n", cliName)
    return errors.New("something didn't work")
}

func commandExit() error {
    os.Exit(0)
    return errors.New("something didn't work")
}


func main() {
    /* 
    Starts up an interactive REPL when you run it.
    It should print a prompt, wait for you to type in a command,
    and then print the result of that command.
    */

    commands := map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
    }

    // initialize input reader
    reader := bufio.NewScanner(os.Stdin)
    // print REPL prompt
    printPrompt()
    for reader.Scan() {
        cmd := commands[strings.ToLower(reader.Text())]
        cmd.callback()
        // print REPL prompt
        printPrompt()
    }
}
