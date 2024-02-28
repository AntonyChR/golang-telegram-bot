package gtb

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)


func GetThumbnailVideo(path string, t string) (pathImage string) {
	command := fmt.Sprintf("ffmpeg -i %s -ss %s -vframes 1 thumbnail.jpg", path, t)
	return command
}

func GetVideoDuration(path string) (int, error){
	command := fmt.Sprintf("ffprobe -i %s -show_entries format=duration -v quiet -of csv=\"p=0\"", path)
	out, err:= exec.Command("bash", "-c", command).Output()

	if err != nil{
		return 0, nil
	}
	durationStr:= strings.ReplaceAll(string(out), "\n", "")

	dFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil{
		return 0, nil
	}


	return int(dFloat), nil 
	
}
