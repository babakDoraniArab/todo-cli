package infrastructure

import (
	"encoding/json"
	
	"os"
	"path/filepath"


	"github.com/babakDoraniArab/todo-cli/pkg/helper"
)



const (
	AppDir       = ".todoApp"
	
)



func Initialization(storageFileName string ,storageData interface{}) (string,error) {

	homeDir, err := os.UserHomeDir()
	helper.CheckErr(err, "Error getting user home directory:")

	TodoAppDir := filepath.Join(homeDir, AppDir)

	_, err = os.Stat(TodoAppDir)
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir(TodoAppDir, 0755)
		helper.CheckErr(err, "error creating new folder for .TodoAppDir")

	}

	storageFileAddr := filepath.Join(TodoAppDir,storageFileName)
	err = createFileIfNotExist(storageFileAddr,storageData)

	helper.CheckErr(err,"could not create the file")
	return storageFileAddr,nil

}

func createFileIfNotExist(filepath string, data interface{}) error {
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {

		file, _ := json.MarshalIndent(data, "", " ")
		err = os.WriteFile(filepath, file, 0755)
		helper.CheckErr(err, "could not create the DB file ")
	}
	return nil

}