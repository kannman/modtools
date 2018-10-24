package cli

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/kannman/modtools/output"
)

// ReadListJSON reads dependency information and returns a Json string.
func ReadListJSON() string {
	cmd := exec.Command("go", "list", "-json", "-m", "all")

	cmd.Env = buildEnv()
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	output.OnError(err, "Error connecting to 'go list -json -m all' stdout")

	err = cmd.Start()
	output.OnError(err, "Error starting 'go list -json -m all'")

	scanner := bufio.NewScanner(stdout)
	json := ""
	for scanner.Scan() {
		json += scanner.Text()
	}

	err = cmd.Wait()
	output.OnError(err, "Error while waiting for 'go list -json -m all'")

	return json
}

// BuildEnv creates the environment in which to run the commands.
func buildEnv() []string {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")
	return env
}
