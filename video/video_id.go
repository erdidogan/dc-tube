package video

import "strings"

var urlList = [7]string{
	"://www.youtube.com/watch?v=",
	"://www.youtube.com/v/",
	"://youtu.be/",
	"://www.youtube.com/embed/",
	"://www.youtube.com/watch?app=desktop&v=",
	"://m.youtube.com/watch?app=desktop&v=",
	"://m.youtube.com/watch?v="}

func GetId(inputUrl string) string {
	path := findIdPath(inputUrl)
	id := findId(path)
	return id
}

func findIdPath(videoUrl string) string {
	idPath := ""
	for _, u := range urlList {
		if strings.Contains(videoUrl, u) {
			if strings.HasPrefix(videoUrl, "https") {
				idPath = videoUrl[len(u)+5:]
			} else {
				idPath = videoUrl[len(u)+4:]
			}
		}
	}
	return idPath
}

func findId(path string) string {
	if strings.Contains(path, "?") {
		return path[:strings.Index(path, "?")]
	} else if strings.Contains(path, "&") {
		return path[:strings.Index(path, "&")]
	} else if strings.Contains(path, "#") {
		return path[:strings.Index(path, "#")]
	} else {
		return path
	}
}
