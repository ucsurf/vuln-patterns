int f(const uint8_t *p, uint32_t off, uint32_t len, uint32_t size) {
    // ruleid: length-check-addition-overflow
    if (off + len > size) return -1;
    // ruleid: length-check-addition-overflow
    if (size < off + len) return -1;
    // ok: length-check-addition-overflow
    if (off > size || len > size - off) return -1;
    return p[off];
}
int g(int slots, int elements, int *top, int size, int *limit) {
    // ok: length-check-addition-overflow
    if (slots < elements + 2) return 1;
    // ok: length-check-addition-overflow
    if (limit < top + size) return 1;
    return 0;
}
