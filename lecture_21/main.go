package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	resp := []byte("\x00" + "romdevelopfile@gmail.com" + "\x00" + "Bangladesh71")
	fmt.Println(string(resp), resp)

	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))

	fmt.Println(sEnc)
}
