package utils

import (
	"fmt"
	"os"
)

func ErrHandle(errMsg string, e ...error) {
	for _, v := range e {
		if v != nil {
			fmt.Printf("%s: %v", errMsg, e)
			os.Exit(1)
		}
	}
}
