package utility

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Setup os keys and value
//Reference from .env file the variables to be set os.Setenv
//Observe that the Function name starts with a Capitalized letter
//For us to refer to it outside of this package
//It's like scoped methods on other languages
func Setup() {
	envConfig, err := GetCWD("/.env")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(envConfig)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	configLoaded := string(data)
	contents := strings.Split(configLoaded, "\n")
	configMap := map[string]string{}
	for k, value := range contents {
		configMap[string(k)] = value
		keyValue := strings.Split(value, "=")
		if len(keyValue) > 1 {
			err := os.Setenv(keyValue[0], keyValue[1])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

//GetCWD : Get Current Working Directory + file path to be checked
func GetCWD(file string) (filePath string, err error) {
	path, errorNotFound := os.Getwd()
	if errorNotFound != nil {
		log.Fatal(errorNotFound)
		return "Not existing", fmt.Errorf("file [%s] does not exist", file)
	}
	fileName := path + file
	_, errorMessage := os.Stat(fileName)
	if os.IsNotExist(errorMessage) {
		return "Not existing", fmt.Errorf("file [%s] does not exist", fileName)
	}
	return fileName, nil
}
