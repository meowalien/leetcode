package Number_of_1_Bits

var bitCountTable [256]int

// Initialize the lookup table
func init() {
	for i := 0; i < 256; i++ {
		bitCountTable[i] = (i & 1) + bitCountTable[i>>1]
	}
}

func hammingWeight1(n int) int {
	u := uint64(n)
	count := 0
	for i := 0; i < 8; i++ { // 8 個 byte (8 * 8 = 64 bits)
		count += bitCountTable[(u>>(i*8))&0xff]
	}
	return count
}
