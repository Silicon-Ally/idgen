package idgen

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefaultGenerator(t *testing.T) {
	gen, err := New(rand.New(rand.NewSource(0)))
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	var got []string
	for i := 0; i < 10; i++ {
		got = append(got, gen.NewID())
	}

	want := []string{
		"mUNERA0rI3cvTK5UHomc",
		"jcEQvymkzADmxkuPMaHw",
		"xmE5tL31SrWWgqtwpMU9",
		"7R8wIBbUt0RwI0U48s6b",
		"aWsz1lmm7WgegogMR7sp",
		"JPZHaPT1w5aIuvnRPswX",
		"n1i7jt4PARUDb1oUtPos",
		"aT4LnXs2oH8gqe8wQHol",
		"gK5k9lETzUcTwCQmwFJP",
		"ivbFQtOWSXBcA59bmfKk",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected random IDs generated (-want +got)\n%s", diff)
	}
}

func TestGeneratorWithOptions(t *testing.T) {
	gen, err := New(rand.New(rand.NewSource(0)),
		WithCharSet([]rune("0123456789")),
		WithDefaultLength(20))
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	var got []string
	for i := 0; i < 10; i++ {
		got = append(got, gen.NewID())
	}

	want := []string{
		"44365677888798261004",
		"58601620127854830094",
		"74689780270842907464",
		"07784390913647839433",
		"48212326206208889001",
		"19910996422689399403",
		"78887357076970225560",
		"49791703471400780329",
		"08464103768524628793",
		"27998142091068852380",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected random IDs generated (-want +got)\n%s", diff)
	}
}

func TestNewIDWithLength(t *testing.T) {
	gen, err := New(rand.New(rand.NewSource(0)))
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	var got []string
	for i := 0; i < 5; i++ {
		id, err := gen.NewIDWithLength(30)
		if err != nil {
			t.Fatalf("gen.NewIDWithLength: %v", err)
		}
		got = append(got, id)
	}

	want := []string{
		"mUNERA0rI3cvTK5UHomcjcEQvymkzA",
		"DmxkuPMaHwxmE5tL31SrWWgqtwpMU9",
		"7R8wIBbUt0RwI0U48s6baWsz1lmm7W",
		"gegogMR7spJPZHaPT1w5aIuvnRPswX",
		"n1i7jt4PARUDb1oUtPosaT4LnXs2oH",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected random IDs generated (-want +got)\n%s", diff)
	}
}
