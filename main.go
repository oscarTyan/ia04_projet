package main

import (
	"fmt"
	"ia04/comsoc"
)

func main() {
	prefs := [][]comsoc.Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := comsoc.MajoritySWF(prefs)
	fmt.Println(res)
}
