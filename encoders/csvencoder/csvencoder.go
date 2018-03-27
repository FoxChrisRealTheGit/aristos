package csvencoder

import(
	"encoding/csv"
	"log"
	"os"
)


// struct types to write to

// basic of the msot basic implementation, needs work
// this will create a file with filename in json format
// with the info in a provided interface
// needs better input handling
func EncodeSomeCSV (filename string, i [][]string){
	f, err := os.Create(filename)
	PrintFatalError(err)
	defer f.Close()

	w := csv.NewWriter(f)
	PrintFatalError(err)

	for _, q := range i {
		if err := w.Write(q); err != nil{
			log.Fatal(err)
		}
	}
	w.Flush()

	err = w.Error()
	PrintFatalError(err)
}

func PrintFatalError(err error){
	if err != nil{
		log.Fatal("Error happened while processing file", err)
	}
}