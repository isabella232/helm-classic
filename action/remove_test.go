package action

import (
	"bytes"
	"os"
	"testing"

	"github.com/deis/helm/log"
	pretty "github.com/deis/pkg/prettyprint"
)

func TestRemove(t *testing.T) {
	tmpHome := createTmpHome()

	Fetch("kitchensink", "", tmpHome)

	var output bytes.Buffer
	log.Stderr = &output

	Remove("kitchensink", tmpHome)

	expected := pretty.Colorize("{{.Green}}--->{{.Default}} ") + "All clear! You have successfully removed kitchensink from your workspace.\n"

	actual := output.String()
	if actual != expected {
		t.Errorf("Expected %v - Got %v ", expected, actual)
	}

	// reset log
	log.Stdout = os.Stdout
}