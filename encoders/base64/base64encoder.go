package base64

import(
	"encoding/base64"
	"log"
	"os"
	"io/ioutil"
)

// I dont think this will work as it is presently
//but it is close?
func EncodeSomeBase64(filename string) error{
	f, err := os.Create(filename)
	defer f.Close()
	PrintFatalError(err)

	filebytes, err := ioutil.ReadFile(filename)
	PrintFatalError(err)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(filebytes)))
	base64.StdEncoding.Encode(dst, filebytes)

	return err
}


func PrintFatalError(err error){
	if err != nil{
		log.Fatal("Error happened while processing file", err)
	}
}