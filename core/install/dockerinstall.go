package install

import (
	"fmt"
	"os"
	"os/exec"
)

func DockerInstall() {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", "../common/script/install-docker.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println("docker install failed,", err)
		os.Exit(-1)
	}
}
