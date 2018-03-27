package decoders

import(
	"encoding/xml"
	"fmt"
	"os"
)

func DecodeXMLConfig(v interface{}, filename string) error{
	fmt.Println("Decoding XML")
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil{
		return err
	}

	err = xml.NewDecoder(file).Decode(v)
	return err
}