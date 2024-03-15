package helper

import "log"

func PanicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
