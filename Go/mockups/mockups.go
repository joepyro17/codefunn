package mockups

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"pyro.com/codefunn/models"
	. "pyro.com/codefunn/models"
)

func GetAllTutorials() Tutorials {
	jsonFile, err := os.Open("./mockups/tutorials.json")
	if err != nil {
		log.Fatal("Open Json failed")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tutorials models.Tutorials

	err = json.Unmarshal(byteValue, &tutorials)
	if err != nil {
		log.Fatal("Unmarshal Failed")
	}

	//for i:= 0; i< len(tutorials.Tutorials); i++ {
	//	fmt.Println(tutorials.Tutorials[i].Title)
	//}
	return tutorials
}