package rule

import (
	"github.com/mgechev/revive/lint"
	"go/ast"
	"regexp"
)

// I18nRule lints struct tags.
type I18nRule struct{}

// Apply applies the rule to given file.
func (r *I18nRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := lintI18nRule{onFailure: onFailure}

	ast.Walk(w, file.AST)

	return failures
}

// Name returns the rule name.
func (r *I18nRule) Name() string {
	return "i18n"
}

type lintI18nRule struct {
	onFailure  func(lint.Failure)
	usedTagNbr map[string]bool // list of used tag numbers
}

var zhRegex = regexp.MustCompile("[\u4e00-\u9fa5]+")

func (w lintI18nRule) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.CallExpr:
		if sel, ok := n.Fun.(*ast.SelectorExpr); ok {
			if sel.Sel.Name == "Tr" {
				return nil
			}
			if sel.Sel.Name == "T" {
				return nil
			}
		}
	case *ast.BasicLit:
		if zhRegex.MatchString(n.Value) {
			w.addFailure(n, "existence of non-English characters")
		}
	}
	return w
}

func (w lintI18nRule) addFailure(n ast.Node, msg string) {
	w.onFailure(lint.Failure{
		Node:       n,
		Failure:    msg,
		Confidence: 1,
	})
}
