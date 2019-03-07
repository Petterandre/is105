package main

import ("fmt"
		"io"
		"log"
		"os")

func main() {
	
	f1 := FileToByteslice("files/text1.txt")
	f2 := FileToByteslice("files/text2.txt")
	fmt.Print(f1)
	fmt.Print("\n")
	fmt.Print(f2)

}


func FileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	return byteSlice

}