package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	tapeSize = 30000
	debug    = false
)

type Interpreter struct {
	tape    []byte
	pointer int
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		tape:    make([]byte, tapeSize),
		pointer: 0,
	}
}

func (i *Interpreter) Interpret(code string) error {
	reader := strings.NewReader(code)
	return i.interpret(reader, os.Stdout, os.Stdin)
}

func (i *Interpreter) interpret(r io.Reader, w io.Writer, stdin io.Reader) error {
	code, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("reading code: %w", err)
	}

	stdinReader := bufio.NewReader(stdin)
	codeLen := len(code)

	for pos := 0; pos < codeLen; pos++ {
		switch code[pos] {
		case '>':
			if i.pointer >= tapeSize-1 {
				return errors.New("pointer out of bounds")
			}
			i.pointer++
		case '<':
			if i.pointer <= 0 {
				return errors.New("pointer out of bounds")
			}
			i.pointer--
		case '+':
			i.tape[i.pointer]++
		case '-':
			i.tape[i.pointer]--
		case '.':
			if _, err := w.Write([]byte{i.tape[i.pointer]}); err != nil {
				return fmt.Errorf("writing output: %w", err)
			}
		case ',':
			b, err := stdinReader.ReadByte()
			if err != nil {
				return fmt.Errorf("reading input: %w", err)
			}
			i.tape[i.pointer] = b
		case '[':
			if i.tape[i.pointer] == 0 {
				bracketCount := 1
				for bracketCount > 0 {
					pos++
					if pos >= codeLen {
						return errors.New("unmatched opening bracket")
					}
					if code[pos] == '[' {
						bracketCount++
					} else if code[pos] == ']' {
						bracketCount--
					}
				}
			}
		case ']':
			if i.tape[i.pointer] != 0 {
				bracketCount := 1
				for bracketCount > 0 {
					pos--
					if pos < 0 {
						return errors.New("unmatched closing bracket")
					}
					if code[pos] == ']' {
						bracketCount++
					} else if code[pos] == '[' {
						bracketCount--
					}
				}
			}
		}

		if debug {
			fmt.Fprintf(os.Stderr, "Instruction: %c, Pointer: %d, Value: %d\n",
				code[pos], i.pointer, i.tape[i.pointer])
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename.bf>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	if filepath.Ext(filename) != ".bf" {
		fmt.Fprintln(os.Stderr, "Error: File must have .bf extension")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	code, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	interpreter := NewInterpreter()
	if err := interpreter.Interpret(string(code)); err != nil {
		fmt.Fprintf(os.Stderr, "Error during interpretation: %v\n", err)
		os.Exit(1)
	}
}

