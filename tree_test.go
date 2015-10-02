package binarytree_test

import (
    "github.com/droxer/binarytree"
    "testing"
)

func TestInsertAndSearch(t *testing.T) {
    bt := binarytree.New()
    bt.Insert(10)
    bt.Insert(7)
    bt.Insert(1)
    bt.Insert(13)
    bt.Insert(16)

    if !bt.Lookup(7) {
        t.Fatalf("should found 10")
    }

    if bt.Lookup(11) {
        t.Fatalf("should not found 11")
    }
}
