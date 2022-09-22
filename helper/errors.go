package helper

import (
	"fmt"
	"log"
)

func PrintError(err error) {
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return
	}
}
