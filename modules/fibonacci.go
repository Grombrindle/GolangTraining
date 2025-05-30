package modules

// var kaka []uint8

func FibSeqRecursive(n int, seq []uint64) []uint64 {
	if n <= 0 {
		return seq
	}
	length := len(seq)
	if length == 0 {
		seq = append(seq, 0)
	} else if length == 1 {
		seq = append(seq, 1)
	} else {
		seq = append(seq, seq[length-1]+seq[length-2])
	}
	return FibSeqRecursive(n-1, seq)
}
