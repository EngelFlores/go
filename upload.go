package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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