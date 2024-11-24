package main

import (
	"fmt"
	"os"
)

type inputBuffer struct {
	buffer       string
	bufferLength int
	inputLength  int
}

func createInputBuffer() inputBuffer {
	newInputBuffer := inputBuffer{"", 0, 0}
	return newInputBuffer
}

func printConsole() {
	fmt.Print("godb > ")
}

func readConsole(inputBuffer inputBuffer) inputBuffer {
	fmt.Scanf("%s", &inputBuffer.buffer)
	bytesRead := len(inputBuffer.buffer)
	if bytesRead <= 0 {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	inputBuffer.inputLength = bytesRead - 1

	return inputBuffer
}

func main() {
	inputBuffer := createInputBuffer()
	condition := true
	for condition {
		printConsole()
		inputBuffer = readConsole(inputBuffer)

		if inputBuffer.buffer == ".exit" {
			fmt.Println("bye!")
			os.Exit(0)
		} else {
			fmt.Printf("unrecognised command %v\n", inputBuffer.buffer)
		}

	}
}
