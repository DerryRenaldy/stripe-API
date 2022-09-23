package helper

import (
	log "github.com/sirupsen/logrus"
)

func PrintError(err error) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if err != nil {
		log.Println(err)
		return
	}
}
