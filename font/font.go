package font

import "fmt"

type Font struct {
	family string
	size int
}

func New(family string, size int) (result *Font) {
	result = &Font{ family: "sans-serif", size: 14 }
	result.SetFamily(family).SetSize(size)
	return
}

func (f *Font) Family() string {
	return f.family
}

func (f *Font) SetFamily(family string) *Font {
	if len(family) > 0 {
		f.family = family
	}
	return f
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) SetSize(size int) *Font {
	if size > 0 {
		f.size = size
	}
	return f
}

func (f *Font) String() string {
	return fmt.Sprintf("{font-family: %q; font-size: %dpt;}", f.Family(), f.Size())
}