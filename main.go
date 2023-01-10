package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var notes []string

func readFromStdin() string {
	fmt.Print("Enter a command and data: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\r')
	//line = line[:len(line)-2]
	line = strings.TrimSpace(line)

	return line
}

func updateNoteBookAtIndex(index int, note []string) []string {
	//notes[index] =""
	var joinedNote string
	var newNote []string
	for i := 2; i < len(note); i++ {
		newNote = append(newNote, note[i])
	}
	joinedNote = strings.Join(newNote, " ")
	notes[index-1] = joinedNote
	fmt.Printf("[OK] The note at position %d  was successfully updated\n", index)

	return notes
}


func deleteNoteBookAtIndex(index int) []string {
	//length := index + 1
	notes = append(notes[:index-1], notes[index+1:]...)
	fmt.Printf("[OK] The note at position %d  was successfully deleted\n", index)

	return notes
}

func main() {
	var maximumNotes int

	if maximumNotes <= 0 {
		fmt.Print("Enter the maximum number of notes: ")
		_, _ = fmt.Scan(&maximumNotes)
	}

	for {
		line := readFromStdin()
		var joinedNote string
		var note []string

		splittedString := strings.Split(line, " ")

		if splittedString[0] != "create" && splittedString[0] != "list" && splittedString[0] != "update" && splittedString[0] != "delete" && splittedString[0] != "clear" && splittedString[0] != "exit" {
			fmt.Println("[Error] Unknown command")
		}

		if splittedString[0] == "create" && len(notes) >= maximumNotes {
			fmt.Println("[Error] Notepad is full")
		}
		if splittedString[1] == "create" && len(splittedString) <= 1 {
			fmt.Println("[Error] Missing note argument")
		}
		if splittedString[0] == "create" && len(splittedString) > 1 && len(notes) < maximumNotes {
			i := 1
			for ; i < len(splittedString); i++ {
				note = append(note, splittedString[i])
			}
			joinedNote = strings.Join(note, " ")
			notes = append(notes, joinedNote)
			note = nil
			fmt.Println("[OK] The note was successfully created")
		}
		if splittedString[0] == "list" && len(notes) <= 0 {
			fmt.Println("[Info] Notepad is empty")
		}
		if splittedString[0] == "list" {
			for i, noteList := range notes {
				//newNote := strings.TrimSpace(noteList)
				fmt.Printf("[Info] %d: %s\n", i+1, noteList)
			}
		}
		if splittedString[0] == "update" {
			if len(notes) <= 0 {
				fmt.Println("[Error] There is nothing to update")
			}
			if len(notes) > 0 {
				if len(splittedString) <= 1 {
					fmt.Println("[Error] Missing position argument")
				}
				if len(splittedString) > 1 && len(splittedString) <= 2 {
					fmt.Println("[Error] Missing note argument")
				}

				if len(splittedString) >= 3 {
					index, err := strconv.Atoi(splittedString[1])

					if err != nil {
						fmt.Printf("[Error] Invalid Position: %s\n", splittedString[1])
					}
					if err == nil {
						if index > len(notes) {
							fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", index, len(notes))
						}

						if index <= len(notes) {
							updateNoteBookAtIndex(index, splittedString)
						}
					}
				}
			}
		}
		if splittedString[0] == "delete" {
			if len(notes) <= 0 {
				fmt.Println("[Error] There is nothing to delete")
			}
			if len(notes) > 0 {
				if len(splittedString) <= 1 {
					fmt.Println("[Error] Missing position argument")
				}
				if len(splittedString) == 2 {
					index, err := strconv.Atoi(splittedString[1])

					if err != nil {
						fmt.Printf("[Error] Invalid Position:%s\n", splittedString[1])
					}
					if index > len(notes) {
						fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", index, len(notes))
					}

					if index <= len(notes) {
						deleteNoteBookAtIndex(index)
					}
				}
			}
		}
		if splittedString[0] == "clear" {
			notes = nil
			fmt.Println("[OK] All notes were successfully deleted")
		}

		if splittedString[0] == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		}
	}
}
