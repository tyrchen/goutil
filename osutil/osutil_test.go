package osutil

import (
	"os/exec"
	"testing"
)

const (
	FILENAME = "/tmp/a12x99_this_file_should_not_exists.tmp"
)

func TestFileExists(t *testing.T) {
	if FileExists(FILENAME) == true {
		t.Errorf("file %s exists (which should not)", FILENAME)
	}

	if exec.Command("touch", FILENAME).Run() == nil {
		if FileExists(FILENAME) == false {
			t.Errorf("file %s does not exists (which should exist)", FILENAME)
		}
		exec.Command("rm", "-f", FILENAME).Run()
	}
}
