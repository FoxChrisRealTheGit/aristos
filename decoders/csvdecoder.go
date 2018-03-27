package decoders

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//need better interface conv to [][]string before adding to decoding handler

func DecodeCSVConfig(v [][]string, filename string) error {
	fmt.Println("Decoding CSV")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			return err
		}
		fmt.Println("CSV Row:", record)
	}
	return err
}
