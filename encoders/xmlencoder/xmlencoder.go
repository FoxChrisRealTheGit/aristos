package xmlencoder

import(
	"encoding/xml"
	"log"
	"os"
)


// struct types to write to

// basic of the msot basic implementation, needs work
// this will create a file with filename in xml format
// with the info in a provided interface
// needs better interface implementation
func EncodeSomeXML (filename string, i interface{}){

	//should probably check if file exists???
	f, err := os.Create(filename)
	defer f.Close()
	PrintFatalError(err)

	//needs error handling.. probably
	enc := xml.NewEncoder(f)
	enc.Indent(" ", "	")
	enc.Encode(&i)
}

func PrintFatalError(err error){
	if err != nil{
		log.Fatal("Error happened while processing file", err)
	}
}