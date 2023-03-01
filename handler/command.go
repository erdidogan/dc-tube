package handler

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const PATH = "yt-dlp"

func Execute(args []string) (string, error) {
	path, err := getPath()

	if err != nil {
		log.Println("Path not found!")
		return "", err
	}

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(path, args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Println("Execution Error: " + fmt.Sprint(err))
		log.Println("Details: " + stderr.String())
		return "", err
	}
	return out.String(), nil

}

func getPath() (string, error) {
	path, err := exec.LookPath(PATH)
	if err != nil {
		log.Println("Path not found ", err)
		return "", err
	}
	return path, nil
}
