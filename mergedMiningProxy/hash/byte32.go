package hash

// Byte32 is a type that provides a cute way of expressing this trivial
// fixed size array type, and is useful because the type is often used in
// slices, and it prevents a forest of square brackets when that is done.
type Byte32 [32]byte

// Assign 赋值
func (data *Byte32) Assign(value []byte) {
	copy(data[:], value)
}

// Reverse 反转，颠倒字节的前后顺序
func (data *Byte32) Reverse() {
	x := 0
	y := len(data) - 1

	for x < y {
		data[x], data[y] = data[y], data[x]
		x++
		y--
	}
}
