package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// WritePID write pid to the given file path.
func WritePID(filepath string, forceStart bool) error {
	if pid, err := ReadPID(filepath); err == nil {
		if !forceStart {
			return fmt.Errorf("Instance of Manager API maybe running with a pid %d. If not, please run Manager API with '-f' or '--force' flag\n", pid)
		}
		fmt.Printf("Force starting new instance. Another instance of Manager API maybe running with pid %d\n", pid)
	}
	pid := os.Getpid()
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(strconv.Itoa(pid)); err != nil {
		return err
	}
	return nil
}

// ReadPID reads the pid from the given file path.
func ReadPID(filepath string) (int, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return -1, err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, fmt.Errorf("invalid pid: %s", err)
	}
	return pid, nil
}
