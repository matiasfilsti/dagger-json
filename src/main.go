package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// type Valores struct {
// 	Source       int `json:"source"`
// 	Environments []struct {
// 		Name      string `json:"name"`
// 		Variables []struct {
// 			Valvalue map[string]string
// 		} `json:"variables"`
// 	} `json:"environments"`
// }

type Valores struct {
	Source       int `json:"source"`
	Environments []struct {
		Name      string            `json:"name"`
		Variables map[string]any `json:"variables"`
	} `json:"environments"`
}



func main() {
	file, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("error abriendo archivo")
		log.Panic(err)
	}

	var datafile Valores

	err = json.Unmarshal(file, &datafile)
	if err != nil {
		fmt.Println(err)
		return
	}


	for _, value := range datafile.Environments {
		fmt.Println(value.Name)
		CreateFile(value.Name)
		file, err := os.OpenFile(value.Name, os.O_APPEND|os.O_WRONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		for z, a := range value.Variables {		
			fmt.Printf("%s: ", z )
			//file.WriteString("%v",z)
			for _, j := range a.(map[string]interface{}){
				fmt.Printf("%v \n", j)
				d2 := z+": "+j.(string)+"\n"
				file.WriteString(d2)
				// result := []byte("%v %v", z, j)
			    //result := fmt.Printf("%v %v",z, j)

				
			}

		}

	}
}


func CreateFile(filename string){
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}

// func GetValues(data Valores) ([]string, []string) {
// 	for _, value := range data.Environments {
// 		fmt.Println(value.Name)
// 		for z, a := range value.Variables {
// 			fmt.Printf("%s: ", z )
// 			for _, j := range a.(map[string]interface{}){
// 				fmt.Printf("%v \n", j)
// 			}
// 		}
// 	}
// }



// func findAllValuesByName(data interface{}, name string) []string {
// 	names := []string{}

// 	switch data.(type) {
// 	case map[string]interface{}:
// 		for key, value := range data.(map[string]interface{}) {
// 			if key == name {
// 				names = append(names, value.(string))
// 			}

// 			// Recorrer recursivamente los valores del mapa
// 			subNames := findAllValuesByName(value, name)
// 			names = append(names, subNames...)
// 		}
// 	case []interface{}:
// 		for _, value := range data.([]interface{}) {
// 			// Recorrer recursivamente los valores del slice
// 			subNames := findAllValuesByName(value, name)
// 			names = append(names, subNames...)
// 		}
// 	}
// 	return names
// }
