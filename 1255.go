package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cases, higher int
	var letters [26]int
	fmt.Scanln(&cases)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < cases; i++ {
		for j := 0; j < 26; j++ {
			letters[j] = 0
		}
		if text, err := in.ReadString('\n'); err == nil {
			for c := 0; c < len(text); c++ {
				s := fmt.Sprintf("%d", text[c])
				if aux, err := strconv.Atoi(s); err == nil {
					if aux > 64 && aux < 91 {
						aux = aux - 65
						letters[aux]++
					} else if aux > 92 && aux < 123 {
						aux = aux - 97
						letters[aux]++
					}
				}
			}
		}
		higher = letters[0]
		for j := 1; j < 26; j++ {
			if letters[j] > higher {

				higher = letters[j]
			}
		}
		for j := 0; j < 26; j++ {
			if letters[j] == higher {
				fmt.Printf("%c", j+97)
			}
			letters[j] = 0
		}
		fmt.Printf("\n")
	}
}
