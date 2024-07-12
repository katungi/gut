package bundler

import (
	"fmt"
	"os/exec"
)

func Bundle(entry string, out string) error {
	fmt.Println("Bundling the JavaScript files")
	cmd := exec.Command("esbuild", entry, "--bundle", "--outfile="+out)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("could not bundle the JavaScript files: %v\n%s", err, output)
	}

	return nil
}
