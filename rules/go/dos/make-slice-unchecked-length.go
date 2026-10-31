package p
func f(br []int64) {
	// ruleid: make-slice-unchecked-length
	buf := make([]byte, br[3]-br[2])
	// ruleid: make-slice-unchecked-length
	b2 := make([]byte, int(hdr.Len))
	n := int(hdr.Len)
	if n < 0 || n > max {
		return
	}
	// ok: make-slice-unchecked-length
	b3 := make([]byte, n)
	_, _, _ = buf, b2, b3
}
