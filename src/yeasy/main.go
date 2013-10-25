package yeasy

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("[error]", err)
	}
}
