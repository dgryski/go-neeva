package neeva

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestNeeva(t *testing.T) {

	var tests = []struct {
		input []byte
		want  string
	}{
		{
			[]byte{0xab},
			`0a163ca802692371b2d1a3035da3bb8f5e9b08ee82e2d5f41e532c1a`,
		},
		{
			[]byte{0xab, 0xcd},
			`1ed012c7c8b70abf4c4a79cb39b143aa29e0d72fd2a5048608c25f7c`,
		},

		{
			[]byte{0xab, 0xcd, 0xab, 0xcd},
			`2303d8e3278bec79aa71e55e018fc10ddc46636aff8fec4c51337167`,
		},
		{
			[]byte{0xab, 0xcd, 0xab, 0xcd, 0xab},
			`545fb189b261315f2f6c8e58be968f0412906f0ed49a5cb272d4b944`,
		},
	}

	for _, tt := range tests {
		want, _ := hex.DecodeString(tt.want)

		if got := Hash(tt.input); !bytes.Equal(got, want) {
			t.Errorf("Hash(%q)=%x, want %x", tt.input, got, want)
		}
	}
}
