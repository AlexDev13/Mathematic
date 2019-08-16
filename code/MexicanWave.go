package code

import (
	"fmt"
	"strings"
)

func Wave(words string) {
	for i:=1;i<=len(words);i++{
		for i:=0;i<len(words);i++{
			fmt.Println(strings.ToUpper(words[i]))
		}		
	}

}
