package daos

import (
	"log"
	"os"
)

func WriteLog(errMsg string) {
	fd, err := os.OpenFile("./log/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open file failed:", err)
		return
	}
	defer fd.Close()

	_, err = fd.WriteString(errMsg)
	if err != nil {
		log.Println("write file failed:", err)
		return
	}
}
