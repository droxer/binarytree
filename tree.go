package binarytree

type BinaryTree struct {
    root *node
}

func New() *BinaryTree {
    return &BinaryTree{}
}

type node struct {
    data  int
    left  *node
    right *node
}

func (b *BinaryTree) Insert(data int) {
    b.root = b.insert(b.root, data)
}

func (b *BinaryTree) insert(n *node, data int) *node {
    if n == nil {
        return &node{
            data: data,
        }
    }

    if data <= n.data {
        n.left = b.insert(n.left, data)
    } else {
        n.right = b.insert(n.right, data)
    }

    return n
}

func (b *BinaryTree) Lookup(data int) bool {
    return b.lookup(b.root, data)
}

func (b *BinaryTree) lookup(n *node, data int) bool {
    if n == nil {
        return false
    }

    if data == n.data {
        return true
    } else if data < n.data {
        return b.lookup(n.left, data)
    } else {
        return b.lookup(n.right, data)
    }
}
