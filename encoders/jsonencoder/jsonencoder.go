package jsonencoder

import(
	"encoding/json"
	"log"
	"os"
	
)


// struct types to write to

// basic of the msot basic implementation, needs work
// this will create a file with filename in json format
// with the info in a provided interface
// needs better interface handling
func EncodeSomeJSON (filename string, i interface{}){
	f, err := os.Create(filename)
	defer f.Close()
	PrintFatalError(err)
	err = json.NewEncoder(f).Encode(&i)
	PrintFatalError(err)
}

func PrintFatalError(err error){
	if err != nil{
		log.Fatal("Error happened while processing file", err)
	}
}