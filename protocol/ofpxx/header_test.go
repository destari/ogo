package ofpxx

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestHelloMarshalBinary(t *testing.T) {
	s := "   01 00 00 08 00 00 00 02 " +
		"00 00 10 01"
	s = strings.Replace(s, " ", "", -1)
	bytes, _ := hex.DecodeString(s)

	h, _ := NewHello(1)
	data, _ := h.MarshalBinary()

	b := hex.EncodeToString(bytes)
	d := hex.EncodeToString(data)
	if (len(b) != len(d)) || (b != d) {
		t.Log("Exp:", b)
		t.Log("Rec:", d)
		t.Errorf("Received length of %d, expected %d", len(data), len(bytes))
	}
}