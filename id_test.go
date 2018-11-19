package idgen_test

import (
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"

	"github.com/srikrsna/idgen"
)

func TestNew(t *testing.T) {
	prefix := "cus"
	testNew(t, idgen.New(prefix), prefix)
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
	for range [10000]struct{}{} {
		id := idgen.New("cus")
		if set[id] {
			t.Errorf("generating repeated strings")
		}

		set[id] = true
	}
}

func TestUniquenessParallel(t *testing.T) {
	var wg sync.WaitGroup

	ids := make(chan string, 1000)
	for i := 0; i < 100; i++ {
		go func(ids chan<- string) {
			wg.Add(1)
			defer wg.Done()

			for i := 0; i < 100; i++ {
				ids <- idgen.New("usr")
			}
		}(ids)
	}

	go func() {
		wg.Wait()
		close(ids)
	}()

	set := map[string]bool{}
	for id := range ids {
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
		for _, v := range ii {
			t.Log(v)
		}
	}
}

func TestGenerator_New(t *testing.T) {
	const prefix = "cus"
	pg := idgen.Generator{
		Prefix: prefix,
	}
	testNew(t, pg.New(), prefix)
}

func TestPrefix(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "Empty",
			id:   "",
			want: "",
		},
		{
			name: "Good",
			id:   idgen.New("cus"),
			want: "cus",
		},
		{
			name: "ulid",
			id:   idgen.New("c")[2:],
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idgen.Prefix(tt.id); got != tt.want {
				t.Errorf("Prefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    uint64
		wantErr error
	}{
		{
			name:    "Empty",
			id:      "",
			want:    0,
			wantErr: idgen.ErrMalformed,
		},
		{
			name:    "Valid",
			id:      "cus_" + ulid.MustNew(ulid.MaxTime(), rand.New(rand.NewSource(time.Now().UnixNano()))).String(),
			want:    ulid.MaxTime(),
			wantErr: nil,
		},
		{
			name:    "Valid Without Prefix",
			id:      ulid.MustNew(ulid.MaxTime(), rand.New(rand.NewSource(time.Now().UnixNano()))).String(),
			want:    ulid.MaxTime(),
			wantErr: nil,
		},
		{
			name:    "Malformed",
			id:      "cus_" + uuid.Must(uuid.NewUUID()).String(),
			wantErr: idgen.ErrMalformed,
			want:    0,
		},
		{
			name:    "Just Prefix",
			id:      "cus_",
			wantErr: idgen.ErrMalformed,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idgen.Time(tt.id)
			if err != tt.wantErr {
				t.Errorf("Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testNew(t *testing.T, id, prefix string) {
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
