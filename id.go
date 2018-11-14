package idgen

import (
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

// New generates a lexically sorted, url safe Id with a prefix.
// Eg: cus_JSfjkdjf333j46, i.e. {prefix}-{ulid}
func New(prefix string) string {
	ent := getEntropy()
	res := prefix + "_" + ulid.MustNew(ulid.Now(), ent).String()
	putEntropy(ent)
	return res
}

var entropyPool = sync.Pool{
	New: func() interface{} {
		return ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	},
}

func getEntropy() io.Reader {
	return entropyPool.Get().(io.Reader)
}

func putEntropy(r io.Reader) {
	entropyPool.Put(r)
}

// PrefixGenerator is id generator with the provided prefix
type PrefixGenerator struct {
	Prefix string
}

// New generates a lexically sorted, url safe Id with available prefix
func (p PrefixGenerator) New() string {
	return New(p.Prefix)
}
