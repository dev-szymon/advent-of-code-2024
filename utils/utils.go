package utils

import (
	"log"
	"strconv"
)

func MustMakeInt(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("error converting string to int: %s\n", n)
	}
	return num
}
