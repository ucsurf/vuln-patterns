def decode(chunks):
    out = b''
    for c in chunks:
        # ruleid: bytes-accumulate-in-loop
        out += c
    bits = ()
    while more():
        # ruleid: bytes-accumulate-in-loop
        bits = bits + next_frag()
    parts = []
    for c in chunks:
        # ok: bytes-accumulate-in-loop
        parts.append(c)
    return b''.join(parts)
