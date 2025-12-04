package asciiart

import (
	"fmt"
	"strings"
)

func cheak(str string) bool {

	return strings.Trim(str, "\\n") == ""

}

func PrintSymbole(arr [][]string, woord string) {

	str := woord
	if cheak(woord) {
		str = woord[0 : len(woord)-2]
	}

	words := strings.Split(str, "\\n")
	for _, val := range words {
		if val == "" {
			fmt.Println()
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, sVal := range val {
					if sVal >= '~' && sVal <= ' ' {
						continue
					} else {
						if sVal >= 0 && int(rune(sVal)-32) < len(arr) {
							fmt.Print(arr[int(rune(sVal)-32)][i])
						}
					}
				}
				fmt.Println()
			}
		}
	}
}
