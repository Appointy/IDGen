package idgen_test

import (
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/oklog/ulid"

	"github.com/srikrsna/idgen"
)

func TestPrefix(t *testing.T) {
	prefix := "cus"
	testPrefix(t, idgen.New(prefix), prefix)
}

func TestULID(t *testing.T) {
	id := idgen.New("cus")
	if _, err := ulid.Parse(id[len("cus_"):]); err != nil {
		t.Errorf("func New not generating ulids: %v", err)
	}
}

func TestURLSafe(t *testing.T) {
	for range [20]struct{}{} {
		id := idgen.New("cus")

		if url.PathEscape(id) != id {
			t.Errorf("value generated by New is not url path safe, expected: %s, got: %s", url.PathEscape(id), id)
		}

		if url.QueryEscape(id) != id {
			t.Errorf("value generated by New is not url query safe, expected: %s, got: %s", url.QueryEscape(id), id)
		}
	}
}

func TestUniqueness(t *testing.T) {
	set := map[string]bool{}
	for range [1000]struct{}{} {
		id := idgen.New("cus")
		if set[id] {
			t.Errorf("generating repeated strings")
		}

		set[id] = true
	}
}

func TestLexicalOrder(t *testing.T) {
	var ii [1000]string
	for k := range ii {
		ii[k] = idgen.New("cus")
	}

	if !sort.StringsAreSorted(ii[:]) {
		t.Errorf("should generate lexically sorted ids")
	}
}

func TestPrefixGenerator_New(t *testing.T) {
	const prefix = "cus"
	pg := idgen.PrefixGenerator{
		Prefix: prefix,
	}
	testPrefix(t, pg.New(), prefix)
}

func testPrefix(t *testing.T, id, prefix string) {
	t.Helper()
	if !strings.HasPrefix(id, prefix+"_") {
		t.Errorf("New should return a ulid with the given prefix, expected: cus, got: %s", id)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = idgen.New("cus")
	}
}

func BenchmarkNewParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = idgen.New("cus")
		}
	})
}
