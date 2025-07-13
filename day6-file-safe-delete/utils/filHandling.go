package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func RestoreFile(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	trashFile, err := GetTrashFileByName(info.Name())
	if err != nil {
		return err
	}
	dstFile, er := os.Create(trashFile.Path)
	srcFile, err := os.Open(trashFile.TrashPath)
	if er != nil {
		return er
	}
	if err != nil {
		return err
	}
	defer dstFile.Close()
	defer srcFile.Close()
	dstFile.Seek(0, 0)
	writer := bufio.NewWriter(dstFile)
	reader := bufio.NewReader(srcFile)
	dataCopied, err := io.Copy(writer, reader)
	if err != nil {
		return err
	}
	fmt.Println("Data transfered :", dataCopied)
	writer.Flush()
	return nil
}

func MoveToTrash(file *os.File) error {

	err := os.MkdirAll("trash", 0755)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	fmt.Println(info.Name())
	dstWriterFile, err := os.Create("trash/" + info.Name())
	if err != nil {
		return err
	}
	defer dstWriterFile.Close()

	dstWriter := bufio.NewWriter(dstWriterFile)
	srcReader := bufio.NewReader(file)
	defer dstWriter.Flush()
	file.Seek(0, 0)
	transferDone, err := io.Copy(dstWriter, srcReader)
	if err != nil {
		return err
	}
	fmt.Println("Data transfered : ", transferDone)
	choiceReader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to delete your file ? y: yes,any key:cancel: ")
	choice, err := choiceReader.ReadString('\n')
	if err != nil {
		return err
	}
	if len(choice) < 2 {
		return fmt.Errorf("only one character is allowed")
	}
	choice = strings.TrimSpace(choice)
	if strings.ToLower(choice) == "y" {
		err := WriteJsonIntoFile(file.Name())
		if err != nil {
			return err
		}
		file.Close()
		err = os.Remove(file.Name())
		if err != nil {
			return err
		}
		fmt.Println("Your file have been deleted")
	}
	return nil
}

func MoveFileToDestination(dst string, src string) error {
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	srcReaderFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcReaderFile.Close()
	dstWriter := bufio.NewWriter(dstFile)
	srcReader := bufio.NewReader(srcReaderFile)
	fileSize, err := io.Copy(dstWriter, srcReader)
	if err != nil {
		return err
	}
	fmt.Println("File has been transfer: ", fileSize)
	return nil
}
