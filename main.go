package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type metaCommandResults int
type prepareResults int
type statementType int

const (
	metaCommandSuccess metaCommandResults = iota
	metaCommandUnrecognised
)

const (
	prepareSuccess prepareResults = iota
	prepareUnrecognised
)

const (
	statementSelect statementType = iota
	statementInsert
)

type inputBuffer struct {
	buffer       string
	bufferLength int
	inputLength  int
}

type statement struct {
	stype statementType
}

func doMetaCommands(inputBuffer inputBuffer) metaCommandResults {
	metacommand := metaCommandUnrecognised
	if inputBuffer.buffer == ".exit" {
		fmt.Println("bye!")
		os.Exit(0)
	} else {
		metacommand = metaCommandUnrecognised
	}
	return metacommand
}

func prepareStatements(inputBuffer inputBuffer, statement *statement) prepareResults {
	if strings.Split(inputBuffer.buffer, " ")[0] == "insert" {
		statement.stype = statementInsert
		return prepareSuccess
	}
	if strings.Split(inputBuffer.buffer, " ")[0] == "select" {
		statement.stype = statementSelect
		return prepareSuccess
	}
	return prepareUnrecognised
}

func executeStatement(statement *statement) {
	switch statement.stype {
	case statementInsert:
		fmt.Println("This is where the logic of insert statement should go")
		break
	case statementSelect:
		fmt.Println("This is where the logic of select statement should go")
		break
	}
}

func createInputBuffer() inputBuffer {
	newInputBuffer := inputBuffer{"", 0, 0}
	return newInputBuffer
}

func printConsole() {
	fmt.Print("godb > ")
}

func readConsole(inputBuffer inputBuffer) inputBuffer {
	var stdin *bufio.Reader
	var line []rune
	stdin = bufio.NewReader(os.Stdin)
	for {
		c, _, err := stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Println("Error reading input")
			os.Exit(1)
		}
		line = append(line, c)
	}
	inputBuffer.buffer = string(line)
	bytesRead := len(inputBuffer.buffer)
	if bytesRead <= 0 {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	inputBuffer.bufferLength = len(inputBuffer.buffer)
	inputBuffer.inputLength = bytesRead - 1

	return inputBuffer
}

func main() {
	inputBuffer := createInputBuffer()
	statement := statement{0}
	condition := true
	for condition {
		printConsole()
		inputBuffer = readConsole(inputBuffer)
		if strings.Split(inputBuffer.buffer, "")[0] == "." {
			switch doMetaCommands(inputBuffer) {
			case metaCommandSuccess:
				continue
			case metaCommandUnrecognised:
				fmt.Printf("unrecognised command %v\n", inputBuffer.buffer)
				continue
			}
		}

		switch prepareStatements(inputBuffer, &statement) {
		case prepareSuccess:
			break
		case prepareUnrecognised:
			fmt.Printf("Unrecognised keyword at the start of '%s'.\n", inputBuffer.buffer)
			continue
		}

		executeStatement(&statement)
		fmt.Println("Executed")

	}
}
