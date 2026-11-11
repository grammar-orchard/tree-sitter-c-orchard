package tree_sitter_c_orchard_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c_orchard "codeberg.org/grammar-orchard/tree-sitter-c-orchard/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_c_orchard.Language())
	if language == nil {
		t.Errorf("Error loading C grammar")
	}
}
