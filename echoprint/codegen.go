package echoprint

import (
	"encoding/json"
	"io/ioutil"
)

// ParseCodegenFile is a helper method for testing, parses a json file generated by codegen
// and returns and array of CodegenFp stucts
func ParseCodegenFile(path string) ([]CodegenFp, error) {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ParseCodegen(jsonData)
}

// ParseCodegen parses the json generated by codegen and returns and array of CodegenFp stucts
func ParseCodegen(jsonData []byte) ([]CodegenFp, error) {
	t := trackTime("ParseCodegen")
	defer t.finish()

	var fpList []CodegenFp
	err := json.Unmarshal(jsonData, &fpList)
	return fpList, err
}
