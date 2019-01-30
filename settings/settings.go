package settings

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	APP struct {
		Path string
	}
)

// execPath returns the executable path
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

// WorkDir returns absolute path of work directory.
func workDir() (string, error) {
	execPath, err := execPath()

	if err != nil {
		log.Fatal(2, "Fail to get app path: %v\n", err)
	}

	execPath = strings.Replace(execPath, "\\", "/", -1)

	i := strings.LastIndex(execPath, "/")
	if i == -1 {
		return execPath, nil
	}
	return execPath[:i], nil
}

func NewContext() {
	workDir, err := workDir()

	if err != nil {
		log.Fatal(2, "Fail to get work directory: %v\n", err)
	}

	APP.Path = workDir
}
