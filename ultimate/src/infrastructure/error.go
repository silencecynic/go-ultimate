package infrastructure

import "log"

func Error(argc error)  {
	if argc != nil {
		log.Fatalln(argc)
	}
}