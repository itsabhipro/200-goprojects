package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileJson struct {
	Name         string    `json:"name"`
	ModifiedDate string    `json:"modified_date"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	TrashPath    string    `json:"trash_path"`
	DeletedAt    time.Time `json:"delete_at"`
}

func LoadData() ([]FileJson, error) {

	file, err := os.Open("data.json")
	if os.IsNotExist(err) {
		file, err = os.Create("data.json")
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()
	var jsons []FileJson
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(file).Decode(&jsons)

	if err != nil {
		return jsons, nil
	}
	return jsons, nil
}

func WriteJsonIntoFile(path string) error {
	sf, err := os.Open(path)
	if err != nil {
		return err
	}
	defer sf.Close()
	file, err := os.OpenFile("data.json", os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	jsons, err := LoadData()
	if err != nil {
		return err
	}
	defer file.Close()
	var fileJson = FileJson{}
	info, err := sf.Stat()
	if err != nil {
		return err
	}
	fileJson.Name = info.Name()
	fileJson.Path = sf.Name()
	fileJson.TrashPath = "trash/" + info.Name()
	fileJson.ModifiedDate = info.ModTime().String()
	fileJson.Size = info.Size()
	fileJson.DeletedAt = time.Now().Local()
	jsons = append(jsons, fileJson)

	err = json.NewEncoder(file).Encode(jsons)
	if err != nil {
		return err
	}
	fmt.Println("File has been saved...")

	return nil
}

func GetTrashFileByName(name string) (*FileJson, error) {
	jsons, err := LoadData()
	if err != nil {
		return nil, err
	}
	for _, file := range jsons {
		if file.Name == name {
			return &file, nil
		}
	}
	return nil, fmt.Errorf("file not found in trash")
}
func saveJson(files []FileJson) error {
	file, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(files)
	if err != nil {
		return err
	}
	return nil
}
func UpadateTrash(filename string) error {
	files, err := LoadData()
	if err != nil {
		return err
	}
	var jsonFiles []FileJson
	trashFile, err := GetTrashFileByName(filename)
	if err != nil {
		return err
	}
	err = os.Remove(trashFile.TrashPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name != filename {
			jsonFiles = append(jsonFiles, file)
		}
	}
	err = saveJson(jsonFiles)
	if err != nil {
		return err
	}

	return nil
}
