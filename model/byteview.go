package model

type ByteView struct {
	B []byte
}

func (v *ByteView) Len() int {
	return len(v.B)
}

func (v *ByteView) ByteSlice() []byte {
	return CloneBytes(v.B)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.B)
}

func CloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
