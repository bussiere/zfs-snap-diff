package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

// lastElement splits a string by sep and returns the last element
func lastElement(str, sep string) string {
	fields := strings.Split(str, sep)
	return fields[len(fields)-1]
}

// zfs executes the 'zfs' command with the provided arguments.
// if the 'zfs' command return code is 0, it returns stdout
// else it returns stderr
func zfs(args string) (string, error) {
	log.Printf("execute: zfs %s\n", args)
	cmd := exec.Command("zfs", strings.Split(args, " ")...)

	var stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	if cmdErr := cmd.Run(); cmdErr != nil {
		return stderrBuf.String(), cmdErr
	}

	return strings.TrimRight(stdoutBuf.String(), "\n"), nil
}

// envHasSet returns true, if 'key' is in the environment
func envHasSet(key string) bool {
	return len(os.Getenv(key)) > 0
}
