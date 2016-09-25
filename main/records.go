package main

import (
	"GoPasswords/CryptoHelper"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**

Format of record file:

FileName: Site name
Line 1: User name
Line 2: Password

**/

const recordsFolder string = "records/"

type record struct {
	sitename string
	username string
	password string
}

var recordMap map[string]string

// loadRecords loads a list of records into memory (only the site names)
func loadRecords() {
	fileNames, err := ioutil.ReadDir(dataDirectory + recordsFolder)

	// defer our error function
	// if the folder doesn't exist, make it and try again
	defer func(err error) {
		if err != nil {
			if os.IsNotExist(err) {
				createRecordsFolder()
				loadRecords()
			} else {
				panic("Error reading records: " + err.Error())
			}
		}
	}(err)

	if err == nil {
		// initialie recordNames to len of fileNames
		// this might be more than required if there are stray directories in the folder
		recordMap = make(map[string]string)

		for _, fileInfo := range fileNames {
			if !fileInfo.IsDir() {
				encryptedFileName := fileInfo.Name()
				decryptedFileName := CryptoHelper.DecryptString(encryptedFileName, decryptionKey, initializationVector)

				recordMap[decryptedFileName] = encryptedFileName
			}
		}
	}
}

func createRecordsFolder() {
	fmt.Println("WARNING: records folder does not exist, I will create it")

	err := os.Mkdir(dataDirectory+recordsFolder, 0777)
	if err != nil {
		panic("Error creating records folder: " + err.Error())
	}
}

// searchRecords searches all records for the query string
// and returns a slice with all of the record names that contain the query string
func searchRecords(query string) []string {
	var resultSet []string

	for plainText := range recordMap {
		if strings.Contains(query, plainText) {
			resultSet = append(resultSet, plainText)
		}
	}

	return resultSet
}

func recordExists(recordName string) bool {
	return recordMap[recordName] != ""
}

// deleteRecord deletes the given record on disk
// it is the caller's responsibility to ensure the record actually exists
func deleteRecord(toDelete record) {
	fileName := recordMap[toDelete.sitename]

	err := os.Remove(fileName)

	if err != nil {
		panic("Error removing record (" + toDelete.sitename + " -> " + fileName + "): " + err.Error())
	}

	delete(recordMap, toDelete.sitename)
}

// createNewRecord will create and save a new record in the records folder
// it is the caller's responsibility to ensure there is no duplicate
func createNewRecord(newRecord record) {
	newRecord.encryptContents()

	file, err := os.Create(dataDirectory + recordsFolder + newRecord.sitename)
	if err != nil {
		panic("Error saving record: " + err.Error())
	}

	file.WriteString(newRecord.username + "\n")
	file.WriteString(newRecord.password)

	file.Close()
}

func (r *record) encryptContents() {
	sitenameEncrypted := CryptoHelper.EncryptString(r.sitename, decryptionKey, initializationVector)
	usernameEncrypted := CryptoHelper.EncryptString(r.username, decryptionKey, initializationVector)
	passwordEncrypted := CryptoHelper.EncryptString(r.password, decryptionKey, initializationVector)

	r.sitename = base64.URLEncoding.EncodeToString(sitenameEncrypted)
	r.username = base64.URLEncoding.EncodeToString(usernameEncrypted)
	r.password = base64.URLEncoding.EncodeToString(passwordEncrypted)
}
