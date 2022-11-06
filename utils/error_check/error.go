package error_check

import "fmt"

func PrintErrorMsg(err error, msg interface{}) {
	if err != nil {
		fmt.Printf("printError: msg: %s, err: %v", msg, err)
	}
}
