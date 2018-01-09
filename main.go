package main

import (
	"PartialFileProcessing/csfile"
	"fmt"
)

func main() {
	const PartSize = 1024 * 1024
	name := "E:/golang/GOPATH/src/filemaker/date.text"
	//name := "E:/golang/GOPATH/src/PartialFileProcessing/main.go"
	var Date []byte = make([]byte, PartSize)
	var file csfile.FilePartialProcessing
	file.InitProcessing(name, PartSize)
	fmt.Println(file)
	for i := 0; i < file.Fornum; i++ {
		file.ReadPartFile(Date, i)
		/*for j := 0; j < PartSize; j++ {
			fmt.Printf("%c", Date[j])
		}*/
		//fmt.Println(file)
	}
	file.FileCloss()
	//fmt.Println(file)
	fmt.Printf("size=%d\n", file.FileSize)
	fmt.Printf("nice code\n")

}
