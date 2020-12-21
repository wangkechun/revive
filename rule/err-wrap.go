package rule

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/mgechev/revive/lint"
)

// ErrWrapRule lints struct tags.
type ErrWrapRule struct{}

// Apply applies the rule to given file.
func (r *ErrWrapRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}
	if strings.HasSuffix(file.Name, "_test.go") {
		return failures
	}

	w := lintErrWrapRule{onFailure: onFailure}

	ast.Walk(w, file.AST)

	return failures
}

// Name returns the rule name.
func (r *ErrWrapRule) Name() string {
	return "err-wrap"
}

type lintErrWrapRule struct {
	onFailure  func(lint.Failure)
	usedTagNbr map[string]bool // list of used tag numbers
}

func (w lintErrWrapRule) Visit(node ast.Node) ast.Visitor {
	posNode := node
	switch n := node.(type) {
	case *ast.IfStmt:
		existsErrWrap := false
		ast.Inspect(n.Body, func(node ast.Node) bool {
			if _, ok := node.(*ast.IfStmt); ok {
				return false
			}
			if call, ok := node.(*ast.CallExpr); ok {
				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				if !strings.HasPrefix(sel.Sel.Name, "Wrap") {
					return true
				}
				ident, ok := sel.X.(*ast.Ident)
				if !ok {
					return true
				}
				if ident.Name != "errors" {
					return true
				}
				posNode = ident
				existsErrWrap = true
			}
			return true
		})

		existsErrCheck := false
		ast.Inspect(n.Cond, func(node ast.Node) bool {
			if _, ok := node.(*ast.IfStmt); ok {
				return false
			}
			if node, ok := node.(*ast.Ident); ok {
				if node.Name == "err" {
					existsErrCheck = true
				}
			}
			return true
		})
		existsOrErrCheck := false
		ast.Inspect(n.Cond, func(node ast.Node) bool {
			if _, ok := node.(*ast.IfStmt); ok {
				return false
			}
			if node, ok := node.(*ast.BinaryExpr); ok {
				if node.Op == token.LOR {
					existsOrErrCheck = true
				}
			}
			return true
		})
		if !existsErrWrap {
			return w
		}
		if existsErrCheck && !existsOrErrCheck {
			return w
		}
		w.addFailure(posNode, "errors.Wrap nil")
	}
	return w
}

func (w lintErrWrapRule) addFailure(n ast.Node, msg string) {
	w.onFailure(lint.Failure{
		Node:       n,
		Failure:    msg,
		Confidence: 1,
	})
}
