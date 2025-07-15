package struc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func WriteStructIntoJson(v any) error {
	jFile, err := os.Create("struct.json")
	if err != nil {
		return nil
	}
	defer jFile.Close()

	err = json.NewEncoder(jFile).Encode(v)
	return err
}

func WritePrettyJson(v any, seperateFile string) error {
	if seperateFile == "" {
		seperateFile = "struct"
	}
	indentedJosn, err := json.MarshalIndent(v, " ", "   ")
	if err != nil {
		return err
	}
	jFile, err := os.Create(seperateFile + ".json")
	if err != nil {
		return err
	}
	defer jFile.Close()
	writer := bufio.NewWriter(jFile)
	writtenJson, err := writer.WriteString(string(indentedJosn))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("Json has been written pretty of size: ", writtenJson)
	return nil
}
