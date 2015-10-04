package binarytree_test

import (
    "github.com/droxer/binarytree"
    "math/rand"
    "testing"
    "time"
)

func init() {
    seed := time.Now().Unix()
    rand.Seed(seed)
}

func perm(n int) (out []int) {
    for _, v := range rand.Perm(n) {
        out = append(out, v)
    }
    return
}

func TestInsertAndSearch(t *testing.T) {
    bt := binarytree.New()
    bt.Insert(10)
    bt.Insert(7)
    bt.Insert(1)
    bt.Insert(13)
    bt.Insert(16)

    if !bt.Get(7) {
        t.Fatalf("should found 10")
    }

    if bt.Get(11) {
        t.Fatalf("should not found 11")
    }
}

func BenchmarkInsert10000(b *testing.B) {
    benchmarkInsert(10000, b)
}

func BenchmarkInsert1000(b *testing.B) {
    benchmarkInsert(1000, b)
}

func BenchmarkGet10000(b *testing.B) {
    benchmarkGet(10000, b)
}

func BenchmarkGet1000(b *testing.B) {
    benchmarkGet(1000, b)
}

func benchmarkInsert(size int, b *testing.B) {
    b.StopTimer()
    values := perm(size)
    b.StartTimer()

    i := 0
    for i < b.N {
        bt := binarytree.New()

        for _, v := range values {
            bt.Insert(v)
            i++

            if i >= b.N {
                return
            }
        }

    }
}

func benchmarkGet(size int, b *testing.B) {
    b.StopTimer()
    insert := perm(size)
    get := perm(size)
    b.StartTimer()

    i := 0
    for i < b.N {
        b.StopTimer()
        bt := binarytree.New()
        for _, v := range insert {
            bt.Insert(v)
            i++

            if i >= b.N {
                return
            }
        }
        b.StartTimer()

        for _, v := range get {
            bt.Get(v)
            i++

            if i >= b.N {
                return
            }
        }
    }
}
