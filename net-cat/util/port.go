package util

import (
	"os"
	"strconv"
	"errors"
)

func GetPort() (error, string) {
	cmd := os.Args[1:]

	if len(cmd) > 1 {
		return errors.New("[USAGE]: ./TCPChat $port"), ""
	} else if len(cmd) == 0 {
		return nil, ":8989"
	} else {
		_, err := strconv.Atoi(cmd[0])
		if err != nil || len(cmd[0]) < 4 {
			return errors.New("[USAGE]: ./TCPChat $port"), ""
		} else {
			return nil, ":" + cmd[0]
		}	
	}
}