package main

import (
	"log"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/csv"
	"encoding/json"
)

// var tempFile *os.File

func uploadFile( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Uploading File")

	file,handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n",handler.Filename)
	fmt.Printf("Uploaded File: %+v\n",handler.Size)

	// Create a temporary file within our temp-file directory that follows
    // a particular naming pattern
		tempFile, err := ioutil.TempFile("temp-file", "upload-*.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
		// byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
		fmt.Fprintf(w, "Successfully Uploaded File\n")
		
		convert(tempFile.Name())
}

func setupRoutes(){
	http.HandleFunc("/upload",uploadFile)
	http.ListenAndServe(":8080",nil)
}

type Data struct{
	Id string `json:"id"`
	Medical_plan string `json:"medical_plan"`
	Dental_plan string `json:"dental_plan"`
	Employee_name string `json:"employee_name"`
}

//parse de csv para json
func convert(path string){
	fmt.Printf("File: %v\n", path)
	pwd, _ := os.Getwd()
	// fmt.Println(pwd)
	csvFile, err := os.Open(pwd+"/"+path)

	// fmt.Printf("csvFile: %T\n", csvFile)
	if err != nil{
		fmt.Println("Deu ruim")
		fmt.Println(err.Error())
	}
	reader := csv.NewReader(csvFile)
	var dado []Data
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err !=nil{
			log.Fatal(err)
		}  
		dado = append(dado, Data{
			Id: line[0],
			Medical_plan: line[1],
			Dental_plan: line[2],
			Employee_name: line[3],
		})
	}

	dadoJson, _ := json.Marshal(dado)
	fmt.Println(string(dadoJson))
}

func main() {
	fmt.Println("Go File Upload")
	setupRoutes()
}
