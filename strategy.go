package main

import (
	"fmt"
	"os"
)

type Response struct {
	ID   string
	Type string
}

type SpecificResponse struct {
	Response
	Data string
}

type File interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

type ReadDirFile interface {
	File
	ReadDir(n int) ([]os.FileInfo, error)
}

func (resp Response) Read(stringa []byte) (int, error) {
	var err error
	resp.ID = string(stringa)
	if err != nil {
		return 1, err
	}
	fmt.Printf("RISPOSTA : %s ", resp.ID)
	return 1, err
}

/* func main() {
	var sr SpecificResponse
	sayHello := "Hello World!"
	sr.Response.Read([]byte(sayHello))
	fmt.Printf(sr.ID)
	fmt.Printf(sr.Response.ID)
} */
