package bundler

import (
	"fmt"
	"os/exec"
)

func Bundle(entry string, out string) error {
	cmd := exec.Command("esbuild", entry, "--bundle", "--outfile="+out, "--loader:.js=jsx", "--format=cjs")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("could not bundle the JavaScript files: %v\n%s", err, output)
	}

	return nil
}