package rule

import (
	"go/ast"
	"strings"

	"github.com/fatih/structtag"
	"github.com/mgechev/revive/lint"
)

// StructFormTagRule lints struct tags.
type StructFormTagRule struct{}

// Apply applies the rule to given file.
func (r *StructFormTagRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := lintStructFormTagRule{onFailure: onFailure}

	ast.Walk(w, file.AST)

	return failures
}

// Name returns the rule name.
func (r *StructFormTagRule) Name() string {
	return "struct-form-tag"
}

type lintStructFormTagRule struct {
	onFailure  func(lint.Failure)
	usedTagNbr map[string]bool // list of used tag numbers
}

func (w lintStructFormTagRule) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.StructType:
		if n.Fields == nil || n.Fields.NumFields() < 1 {
			return nil // skip empty structs
		}
		w.usedTagNbr = map[string]bool{} // init
		for _, f := range n.Fields.List {
			if f.Tag != nil {
				w.checkTaggedField(f)
			}
		}
	}

	return w

}

// checkTaggedField checks the tag of the given field.
// precondition: the field has a tag
func (w lintStructFormTagRule) checkTaggedField(f *ast.Field) {
	if len(f.Names) > 0 && !f.Names[0].IsExported() {
		w.addFailure(f, "tag on not-exported field "+f.Names[0].Name)
	}

	tags, err := structtag.Parse(strings.Trim(f.Tag.Value, "`"))
	if err != nil || tags == nil {
		w.addFailure(f.Tag, "malformed tag")
		return
	}
	tagForm, err := tags.Get("form")
	if err != nil {
		return
	}
	tagJson, _ := tags.Get("json")
	if tagForm != nil && tagJson == nil {
		w.addFailure(f.Tag, "tag form and json should exist simultaneously")
		return
	}
	if tagForm.Name != tagJson.Name {
		w.addFailure(f.Tag, "tag form and json should equal")
		return
	}
	return
}

func (w lintStructFormTagRule) addFailure(n ast.Node, msg string) {
	w.onFailure(lint.Failure{
		Node:       n,
		Failure:    msg,
		Confidence: 1,
	})
}
