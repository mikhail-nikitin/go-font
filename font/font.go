package font

import "fmt"

type Font interface {
	fmt.Stringer
	Family() string
	SetFamily(string)
	Size() int
	SetSize(int)
}

type font struct {
	family string
	size int
}

func New(family string, size int) Font {
	f := &font{ family: "sans-serif", size: 14 }
	f.SetFamily(family)
	f.SetSize(size)
	return f
}

func (f *font) Family() string {
	return f.family
}

func (f *font) SetFamily(family string) {
	if len(family) > 0 {
		f.family = family
	}
}

func (f *font) Size() int {
	return f.size
}

func (f *font) SetSize(size int) {
	if size >= 5 && size <= 144 {
		f.size = size
	}
}

func (f *font) String() string {
	return fmt.Sprintf("{font-family: %q; font-size: %dpt;}", f.Family(), f.Size())
}