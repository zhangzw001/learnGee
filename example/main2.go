package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "test123.log"
	file,_ := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE,0644)

	//写点啥
	_,_ = file.Write([]byte("123456789"))
	defer file.Close()

	file,_ = os.Open(filename)
	defer file.Close()

	buf := make([]byte, 6)
	n, _ := file.Read(buf)
	fmt.Printf("%v 读取的字节: %v\n", n,string(buf))	//读取的字节: 123456


	buf2:= make([]byte, 6)
	n, _ = file.Read(buf2)
	fmt.Printf("%v 读取的字节: %v\n", n, string(buf2))	//读取的字节: 789

	//s := strings.NewReader("abcdefghij")

}
