package dotmusgo

import (
	"github.com/ymz-ncnk/musgen"
	"github.com/ymz-ncnk/musgo"
)

// Prefix for filename.
const Prefix = "a__"

// New returns a new MusGo.
func New() (musgo.MusGo, error) {
	musGo, err := musgo.New()
	if err != nil {
		return musgo.MusGo{}, err
	}
	musGo.FilenameBuilder = FilenameBuilder
	return musGo, nil
}

// FilenameBuilder creates a filename for the generated file.
func FilenameBuilder(td musgen.TypeDesc) string {
	return Prefix + musgo.DefaultFilenameBuilder(td)
}
