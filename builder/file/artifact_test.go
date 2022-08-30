package file

import (
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func TestNullArtifact(_ *testing.T) {
	var _ packersdk.Artifact = new(FileArtifact)
}
