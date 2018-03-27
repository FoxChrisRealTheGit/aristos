package decoders

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

//dont think this will work as id
// needs to accepta file, then output file as a base64
// should also be able to ouutput variable to pass to
// another service?

func Decodebase64Config( filename string) error {
	fmt.Println("Decoding base64")
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil{
		return err
	}
	buffer := make([]byte, 4096)
	dst := make([]byte, base64.StdEncoding.DecodedLen(5))
	_, err = base64.StdEncoding.Decode(dst, buffer[:5])
	if err != nil{
		log.Fatal(err)
	}
	filedst, _ := os.Create(filename)
	filedst.Write(dst)
	return file.Close()
}