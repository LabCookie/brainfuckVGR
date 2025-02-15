package main

import (
	"fmt"
	"unicode"
)

func interpret(text string) {
	var cells [64]int
	var variable_cells [10]int
	var def_cells [10]int

	var pointer int = 0
	var str_pointer int = 0
	var output string = ""
	var loop_start_pos int
	var called_pos int
	var calling bool

	for str_pointer < len(text) {
		var char rune = rune(text[str_pointer])

		if char == '+' {
			cells[pointer]++
		} else if char == '-' {
			cells[pointer]--
		} else if char == '<' {
			pointer--
		} else if char == '>' {
			pointer++
		} else if char == '.' {
			output += string(rune(cells[pointer]))
		} else if char == 'u' {
			cells[pointer] = int(rune(output[len(output)-1]))
		} else if char == '[' {
			loop_start_pos = str_pointer
		} else if char == ']' {
			if cells[pointer] > 0 {
				str_pointer = loop_start_pos
			}
		} else if char == 'v' {
			variable_cells[int(text[str_pointer+1]-'0')] = cells[pointer]
		} else if char == 'g' {
			cells[pointer] = variable_cells[int(text[str_pointer+1]-'0')]
		} else if char == 'r' {
			cells[pointer] = 0
		} else if char == 'c' {
			for _, element := range cells {
				output += "|"
				output += fmt.Sprint(element)
			}
			output += "|\n"
		} else if char == 'n' {
			output += "\n"
		} else if char == 'V' {
			for _, element := range variable_cells {
				output += "|"
				output += fmt.Sprint(element)
			}
			output += "|\n"
		} else if char == 'f' {
			if unicode.IsDigit(rune(text[str_pointer+1])) && !calling {
				def_cells[int(text[str_pointer+1]-'0')] = str_pointer + 1
			} else {
				str_pointer = called_pos
				calling = false
			}
		} else if char == 'F' {
			str_pointer = def_cells[int(text[str_pointer+1]-'0')]
			calling = true
		}

		str_pointer++
	}
	fmt.Println(output)
}

func main() {
	interpret(`
		f0++++++[>++++++++++<-]>+++++f

		F0+.

		++++++++[>++++++++++<-]>.v0
		++++++[>++++++++++<-]>+++++v1


		g0++++++v2.
		>+++[<----->-]<.
		>g2----.

		ncV
	`)
}
