package convert

import (
	"encoding/json"
	"fmt"
)

func Encode2String(object interface{}) (string, error) {
	data, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return "", nil
	}
	return string(data), nil
}
func Decode2Object(str string, object interface{}) error {
	data := []byte(str)
	err := json.Unmarshal(data, &object)
	if err != nil {
		return err
	}
	return nil
}
func Object2Dict(object interface{}) (map[string]interface{},error) {
	var dict map[string]interface{}
	jsonData, err := json.Marshal(object)
	if err != nil{
		return dict,err
	}
	err = json.Unmarshal(jsonData, &dict)
	return dict,err
}