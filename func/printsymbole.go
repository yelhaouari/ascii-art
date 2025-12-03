package asciiart

import (
	"fmt"
	"strings"
)

func PrintSymbole(arr [][]string, woord string) {
	words := strings.Split(woord, "\n")
	for _, val := range words {
		if val == "" || val == "\\n" {
			fmt.Println()
			continue
		} else {
			for i := 0; i < 9-1; i++ {
				for _, sVal := range val {

					if sVal >= 0 && int(rune(sVal)-32) < len(arr) {
						fmt.Print(arr[int(rune(sVal)-32)][i])
					}
				}
				fmt.Println()
			}
		}
	}
}
