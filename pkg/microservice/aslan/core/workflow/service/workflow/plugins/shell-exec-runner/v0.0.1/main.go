package main

import (
	"os"
	"os/exec"

	"github.com/spf13/viper"

	"github.com/koderover/zadig/pkg/tool/log"
)

// envs
const (
	EXEC_CMD = "EXEC_CMD"
)

func main() {
	log.Init(&log.Config{
		Level:       "info",
		Development: false,
		MaxSize:     5,
	})
	viper.AutomaticEnv()

	execCmd := viper.GetString(EXEC_CMD)

	// write the sh command
	err := os.WriteFile("/script.sh", []byte(execCmd), 0777)
	if err != nil {
		log.Errorf("failed to write script into /script.sh, error: %s", err)
		os.Exit(1)
	}

	args := []string{"-c", "/script.sh"}
	cmd := exec.Command("sh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Errorf("failed to execute /script.sh, error: %s", err)
		os.Exit(1)
	}

	log.Infof("shell script execution complete")
}
