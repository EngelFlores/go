package main

import (
	"fmt"
	"io"
	"log"
	"encoding/csv"
	"os"
)

type Data struct {
	Id                                                                                                                             int
	Medical_plan, Dental_plan, Employee_name, Language, Claimant_name, Relationship_type, Gender, Effective_date, Termination_date string
}

//parse de csv para json
func convert(path string) {
	pwd, _ := os.Getwd()
	csvFile, err := os.Open(pwd + "/" + path)
	ourData, err := os.Open(pwd + "/temp-file/ourData.csv")

	if err != nil {
		fmt.Println(err.Error())
	}
	var clienteData []Data = createStruct(csvFile)
	var nossosDados []Data = createStruct(ourData)

	read(clienteData, nossosDados)
}

func createStruct(csvFile *os.File) []Data {
	reader := csv.NewReader(csvFile)

	var dados []Data
	var count = 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if line[0] != "id" {
			count++
		}
		if count > 0 {
			dados = append(dados, Data{
				Id:                count,
				Medical_plan:      line[1],
				Dental_plan:       line[2],
				Employee_name:     line[3],
				Language:          line[4],
				Claimant_name:     line[5],
				Relationship_type: line[6],
				Gender:            line[7],
				Effective_date:    line[8],
				Termination_date:  line[9],
			})
		}

	}

	return dados
}
