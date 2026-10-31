package p

func (d *decoder) decodeValue(off int) (any, error) {
	// ruleid: recursive-decoder-without-depth
	return d.decodeValue(off + 1)
}

func (p *parser) parseNode() error {
	// ruleid: recursive-decoder-without-depth
	return p.parseNode()
}

func (d *decoder) decodeSafe(off int, depth int) (any, error) {
	// ok: recursive-decoder-without-depth
	return d.decodeSafe(off+1, depth+1)
}
