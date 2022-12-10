package dotmusgo

import (
	"github.com/ymz-ncnk/musgen"
	"github.com/ymz-ncnk/musgo"
)

// MiddleDot is used as filename prefix.
const MiddleDot = "·"

// New returns a new MusGo.
func New() (musgo.MusGo, error) {
	musGo, err := musgo.New()
	if err != nil {
		return musgo.MusGo{}, err
	}
	musGo.FilenameBuilder = FilenameBuilder
	return musGo, nil
}

// FilenameBuilder creates a filename for the generated file, uses MiddleDot as
// prefix.
func FilenameBuilder(td musgen.TypeDesc) string {
	return MiddleDot + musgo.DefaultFilenameBuilder(td)
}
