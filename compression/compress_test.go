package compression

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCompressEmptyData(t *testing.T) {
	comp := LZWCompress(make([]byte, 0))

	if len(comp.Data) != 0 {
		t.Errorf("Should have empty data")
	}
}

func TestCompressComplexData(t *testing.T) {
	message := "A robot may not injure a human being or, through inaction, allow a human being to " +
		"come to harm. A robot must obey the orders given to it by human beings, except where such " +
		"orders would conflict with the First Law. A robot must protect its own existence."
	want, _ := os.ReadFile("../test_assets/complex-data.textlzw")
	got := LZWCompress([]byte(message))
	if diff := cmp.Diff(got.Data, want); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}
