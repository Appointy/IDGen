package idgen

import (
	"errors"
	"io"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/oklog/ulid"
)

const (
	// SeparatorByte is the byte representation of the separator '_' character.
	SeparatorByte = byte('_')

	// SeparatorString is the string representation of the separator '_' character.
	SeparatorString = string(SeparatorByte)
)

var (
	// ErrMalformed represents and error that is returned in case of malformed ulids
	ErrMalformed = errors.New("idgen: malformed ulid")
)

// New generates a lexically sorted, url safe Id with a prefix.
// Eg: cus_JSfjkdjf333j46, i.e. {prefix}_{ulid}
func New(prefix string) string {
	ent := getEntropy()
	res := prefix + SeparatorString + ulid.MustNew(ulid.Now(), ent).String()
	putEntropy(ent)
	return res
}

var entropyPool = sync.Pool{
	New: func() interface{} {
		ns := time.Now().UnixNano()
		return ulid.Monotonic(rand.New(rand.NewSource(ns)), uint64(ns)%math.MaxUint32)
	},
}

func getEntropy() io.Reader {
	return entropyPool.Get().(io.Reader)
}

func putEntropy(r io.Reader) {
	entropyPool.Put(r)
}

// Generator generates ids with the provided prefix
type Generator struct {
	Prefix string
}

// New generates a lexically sorted, url safe Id with available prefix
func (p Generator) New() string {
	return New(p.Prefix)
}

// Prefix returns the prefix of the id. Id should be of type {prefix}{Separator}*
// If id is malformed it returns an empty string.
func Prefix(id string) string {
	i := strings.IndexByte(id, SeparatorByte)
	if i < 0 {
		return ""
	}

	return id[:i]
}

// Time returns the Unix time in milliseconds encoded in the id.
// id should be of the form {prefix}{Separator}{ulid}. The prefix and separator are optional.
// ErrMalformed is returned in case of a malformed id.
func Time(id string) (uint64, error) {
	uid, err := ulid.Parse(id[strings.IndexByte(id, SeparatorByte)+1:])
	if err != nil {
		return 0, ErrMalformed
	}

	return uid.Time(), nil
}
