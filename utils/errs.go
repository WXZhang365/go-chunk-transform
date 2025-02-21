package utils

import (
	"fmt"
	"os"
)

func ErrHandle(errMsg string, e ...error) {
	for _, v := range e {
		if v != nil {
			fmt.Printf("%s: %v \n", errMsg, e)
			os.Exit(1)
		}
	}
}

func ErrChan(err error, channel chan error) {
	if err != nil {
		channel <- err
	}
}
