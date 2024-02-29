package gtb

import (
	"fmt"
	"os/exec"
)

func VerifyRequirements() error {
	//ferify ffmpeg install
	cmd := exec.Command("bash", "-c", "ffmpeg -version")

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("error checking ffmpeg availability: %w", err)
	}

	return nil
}
