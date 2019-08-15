package code

import (
	"fmt"
	"strings"
)

func Wave(words string) {
	for i:=1;i<=len(words);i++{
		fmt.Println(strings.ToUpper(words))
		// fmt.Println(i,words)
	}

}
