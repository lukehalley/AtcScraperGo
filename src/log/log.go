package atchelpers

import "log"

func LogSeparator(NewLine bool) {

	Separator := "------------------------------"

	if NewLine {
		log.Print(Separator)
		log.Print("")
	} else {
		log.Print(Separator)
	}

}
