package dotmusgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo"
)

type IntAlias int

func TestGenerate(t *testing.T) {
	musGo, err := New()
	if err != nil {
		t.Fatal(err)
	}
	unsafe := false
	err = musGo.Generate(reflect.TypeOf((*IntAlias)(nil)).Elem(), unsafe)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateWithFilename(t *testing.T) {
	musGo, err := New()
	if err != nil {
		t.Fatal(err)
	}
	conf := musgo.NewConf()
	conf.T = reflect.TypeOf((*IntAlias)(nil)).Elem()
	conf.Filename = "AnotherIntAlias.go"
	err = musGo.GenerateAs(conf)
	if err != nil {
		t.Fatal(err)
	}
}
