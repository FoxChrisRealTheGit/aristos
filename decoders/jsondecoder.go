package decoders

import(
	"encoding/json"
	"fmt"
	"os"
)

func DecodeJSONConfig(v interface{}, filename string) error{
	fmt.Println("Decoding JSON")
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil{
		return err
	}
	return json.NewDecoder(file).Decode(v)
}