package video

import (
	"discord-tube/handler"
)

func GetAudioUrl(videoId string) (string, error) {
	//140  m4a   audio only
	args := []string{"-v", "-f", "140", "-g", videoId}
	return handler.Execute(args)
}

func GetTitle(videoId string) (string, error) {
	args := []string{"--get-title", videoId}
	return handler.Execute(args)
}
