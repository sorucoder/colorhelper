package colorhelper

import (
	"image/color"
	"math"
)

// SRGBA defines an sRGB color.
type SRGBA struct {
	R, G, B, A float64
}

var SRGBAModel color.Model = color.ModelFunc(srgbaModel)

func srgbaModel(c color.Color) color.Color {
	if _, ok := c.(SRGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return SRGBA{float64(r) / 65535.0, float64(g) / 65535.0, float64(b) / 65535.0, float64(a) / 65535.0}
}

func (srgba SRGBA) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(srgba.R * 65535.0), uint32(srgba.G * 65535.0), uint32(srgba.B * 65535.0), uint32(srgba.A * 65535.0)
}

// HSLA defines a HSLA color.
type HSLA struct {
	H, S, L, A float64
}

var HSLAModel color.Model = color.ModelFunc(hslaModel)

// HSLtoRGB converts an HSL triple to an sRGB triple.
func HSLtoRGB(h, s, l float64) (r, g, b float64) {
	C := (1.0 - math.Abs(2.0*l-1)) * s
	H := h / 60.0
	X := C * (1.0 - math.Abs(math.Mod(H, 2.0)-1.0))

	var R, G, B float64
	switch {
	case H >= 0.0 && H <= 1.0:
		R, G, B = C, X, 0.0
	case H > 1.0 && H <= 2.0:
		R, G, B = X, C, 0.0
	case H > 2.0 && H <= 3.0:
		R, G, B = 0.0, C, X
	case H > 3.0 && H <= 4.0:
		R, G, B = 0.0, X, C
	case H > 4.0 && H <= 5.0:
		R, G, B = X, 0.0, C
	case H > 5.0 && H <= 6.0:
		R, G, B = C, 0.0, X
	}
	m := l - C/2.0

	r = R + m
	g = G + m
	b = B + m

	return
}

// RGBtoHSL converts an sRGB triple to an HSL triple
func RGBtoHSL(r, g, b float64) (h, s, l float64) {
	M := math.Max(math.Max(r, g), math.Max(g, b))
	m := math.Min(math.Min(r, g), math.Min(g, b))
	C := M - m

	var H float64
	if C != 0 {
		switch M {
		case r:
			H = math.Mod((g-b)/C, 6.0)
		case g:
			H = (b-r)/C + 2.0
		case b:
			H = (r-g)/C + 4.0
		}
	}

	h = H * 60.0
	l = (M + m) / 2.0
	if l > 0.0 || l < 1.0 {
		s = C / (1.0 - math.Abs(2.0*l-1.0))
	}

	return
}

func hslaModel(c color.Color) color.Color {
	if _, ok := c.(HSLA); ok {
		return c
	}
	r, g, b, a := c.RGBA()

	var hsla HSLA
	hsla.H, hsla.S, hsla.L = RGBtoHSL(float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0)
	hsla.A = float64(a) / 65535.0

	return hsla
}

func (hsla HSLA) RGBA() (r, g, b, a uint32) {
	R, G, B := HSLtoRGB(hsla.H, hsla.S, hsla.L)
	A := hsla.A

	r = uint32(R * 65535.0)
	g = uint32(G * 65535.0)
	b = uint32(B * 65535.0)
	a = uint32(A * 65535.0)

	return
}
