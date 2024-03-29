package gtb

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// generates a video thumbnail and returns the path of the jpg file
func GetThumbnailVideo(videoPath string) (string, error) {
	basePath := strings.TrimSuffix(filepath.Base(videoPath), filepath.Ext(videoPath))
	imagePath := basePath + ".jpg"
	//command := fmt.Sprintf("ffmpeg -i %s -vframes 1 -an -s 320x320 -ss 0 %s", videoPath, imagePath)
	command := fmt.Sprintf("ffmpeg -i %s -vframes 1 -an -ss 0 %s", videoPath, imagePath)
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()

	if err != nil {
		return "", err
	}
	return imagePath, nil
}

// return video duration in seconds
func GetVideoDuration(path string) (int, error) {
	command := fmt.Sprintf("ffprobe -i %s -show_entries format=duration -v quiet -of csv=\"p=0\"", path)
	out, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		return 0, nil
	}
	durationStr := strings.ReplaceAll(string(out), "\n", "")

	dFloat, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, nil
	}
	return int(dFloat), nil
}
