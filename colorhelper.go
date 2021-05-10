package colorhelper

import (
	"fmt"
	"image/color"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	hexTripletRegexp    *regexp.Regexp = regexp.MustCompile(`^#([a-fA-F0-9]{2})([a-fA-F0-9]{2})([a-fA-F0-9]{2})$`)
	hexQuadrupletRegexp *regexp.Regexp = regexp.MustCompile(`^#([a-fA-F0-9]{2})([a-fA-F0-9]{2})([a-fA-F0-9]{2})([a-fA-F0-9]{2})$`)
	rgbFunctionRegexp   *regexp.Regexp = regexp.MustCompile(`^rgb\(\s*(|\d+(?:\.\d+)?%)\s*,\s*(\d+|\d+(?:\.\d+)?%)\s*,\s*(?:\d+|\d+(?:\.\d+)?%)\s*\)$`)
	rgbaFunctionRegexp  *regexp.Regexp = regexp.MustCompile(`^rgba\(\s*(\d+|\d+(?:\.\d+)?%)\s*,\s*(\d+|\d+(?:\.\d+)?%)\s*,\s*(?:\d+|\d+(?:\.\d+)?%)\s*,\s*(\d+(?:\.\d+)?)\s*\)$`)
	hslFunctionRegexp   *regexp.Regexp = regexp.MustCompile(`^hsl\(\s*(\d+)\s*,\s*(\d+(?:\.\d+)?%)\s*,\s*\(\d+(?:\.\d+)?%\)\s*\)$`)
	hslaFunctionRegexp  *regexp.Regexp = regexp.MustCompile(`^hsl\(\s*(\d+)\s*,\s*(\d+(?:\.\d+)?%)\s*,\s*\(\d+(?:\.\d+)?%,\s*(\d+(?:\.\d+)?)\s*\)\s*\)$`)

	colorNamesToColorValues map[string]color.Color = map[string]color.Color{
		"aliceblue":            color.RGBA{0xF0, 0xF8, 0xFF, 0xFF},
		"antiquewhite":         color.RGBA{0xFA, 0xEB, 0xD7, 0xFF},
		"aqua":                 color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
		"aquamarine":           color.RGBA{0x7F, 0xFF, 0xD4, 0xFF},
		"azure":                color.RGBA{0xF0, 0xFF, 0xFF, 0xFF},
		"beige":                color.RGBA{0xF5, 0xF5, 0xDC, 0xFF},
		"bisque":               color.RGBA{0xFF, 0xE4, 0xC4, 0xFF},
		"black":                color.RGBA{0x00, 0x00, 0x00, 0xFF},
		"blanchedalmond":       color.RGBA{0xFF, 0xEB, 0xCD, 0xFF},
		"blue":                 color.RGBA{0x00, 0x00, 0xFF, 0xFF},
		"blueviolet":           color.RGBA{0x8A, 0x2B, 0xE2, 0xFF},
		"brown":                color.RGBA{0xA5, 0x2A, 0x2A, 0xFF},
		"burlywood":            color.RGBA{0xDE, 0xB8, 0x87, 0xFF},
		"cadetblue":            color.RGBA{0x5F, 0x9E, 0xA0, 0xFF},
		"chartreuse":           color.RGBA{0x7F, 0xFF, 0x00, 0xFF},
		"chocolate":            color.RGBA{0xD2, 0x69, 0x1E, 0xFF},
		"coral":                color.RGBA{0xFF, 0x7F, 0x50, 0xFF},
		"cornflowerblue":       color.RGBA{0x64, 0x95, 0xED, 0xFF},
		"cornsilk":             color.RGBA{0xFF, 0xF8, 0xDC, 0xFF},
		"crimson":              color.RGBA{0xDC, 0x14, 0x3C, 0xFF},
		"cyan":                 color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
		"darkblue":             color.RGBA{0x00, 0x00, 0x8B, 0xFF},
		"darkcyan":             color.RGBA{0x00, 0x8B, 0x8B, 0xFF},
		"darkgoldenrod":        color.RGBA{0xB8, 0x86, 0x0B, 0xFF},
		"darkgray":             color.RGBA{0xA9, 0xA9, 0xA9, 0xFF},
		"darkgrey":             color.RGBA{0xA9, 0xA9, 0xA9, 0xFF},
		"darkgreen":            color.RGBA{0x00, 0x64, 0x00, 0xFF},
		"darkkhaki":            color.RGBA{0xBD, 0xB7, 0x6B, 0xFF},
		"darkmagenta":          color.RGBA{0x8B, 0x00, 0x8B, 0xFF},
		"darkolivegreen":       color.RGBA{0x55, 0x6B, 0x2F, 0xFF},
		"darkorange":           color.RGBA{0xFF, 0x8C, 0x00, 0xFF},
		"darkorchid":           color.RGBA{0x99, 0x32, 0xCC, 0xFF},
		"darkred":              color.RGBA{0x8B, 0x00, 0x00, 0xFF},
		"darksalmon":           color.RGBA{0xE9, 0x96, 0x7A, 0xFF},
		"darkseagreen":         color.RGBA{0x8F, 0xBC, 0x8F, 0xFF},
		"darkslateblue":        color.RGBA{0x48, 0x3D, 0x8B, 0xFF},
		"darkslategray":        color.RGBA{0x2F, 0x4F, 0x4F, 0xFF},
		"darkslategrey":        color.RGBA{0x2F, 0x4F, 0x4F, 0xFF},
		"darkturquoise":        color.RGBA{0x00, 0xCE, 0xD1, 0xFF},
		"darkviolet":           color.RGBA{0x94, 0x00, 0xD3, 0xFF},
		"deeppink":             color.RGBA{0xFF, 0x14, 0x93, 0xFF},
		"deepskyblue":          color.RGBA{0x00, 0xBF, 0xFF, 0xFF},
		"dimgray":              color.RGBA{0x69, 0x69, 0x69, 0xFF},
		"dimgrey":              color.RGBA{0x69, 0x69, 0x69, 0xFF},
		"dodgerblue":           color.RGBA{0x1E, 0x90, 0xFF, 0xFF},
		"firebrick":            color.RGBA{0xB2, 0x22, 0x22, 0xFF},
		"floralwhite":          color.RGBA{0xFF, 0xFA, 0xF0, 0xFF},
		"forestgreen":          color.RGBA{0x22, 0x8B, 0x22, 0xFF},
		"fuchsia":              color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
		"gainsboro":            color.RGBA{0xDC, 0xDC, 0xDC, 0xFF},
		"ghostwhite":           color.RGBA{0xF8, 0xF8, 0xFF, 0xFF},
		"gold":                 color.RGBA{0xFF, 0xD7, 0x00, 0xFF},
		"goldenrod":            color.RGBA{0xDA, 0xA5, 0x20, 0xFF},
		"gray":                 color.RGBA{0x80, 0x80, 0x80, 0xFF},
		"grey":                 color.RGBA{0x80, 0x80, 0x80, 0xFF},
		"green":                color.RGBA{0x00, 0x80, 0x00, 0xFF},
		"greenyellow":          color.RGBA{0xAD, 0xFF, 0x2F, 0xFF},
		"honeydew":             color.RGBA{0xF0, 0xFF, 0xF0, 0xFF},
		"hotpink":              color.RGBA{0xFF, 0x69, 0xB4, 0xFF},
		"indianred":            color.RGBA{0xCD, 0x5C, 0x5C, 0xFF},
		"indigo":               color.RGBA{0x4B, 0x00, 0x82, 0xFF},
		"ivory":                color.RGBA{0xFF, 0xFF, 0xF0, 0xFF},
		"khaki":                color.RGBA{0xF0, 0xE6, 0x8C, 0xFF},
		"lavender":             color.RGBA{0xE6, 0xE6, 0xFA, 0xFF},
		"lavenderblush":        color.RGBA{0xFF, 0xF0, 0xF5, 0xFF},
		"lawngreen":            color.RGBA{0x7C, 0xFC, 0x00, 0xFF},
		"lemonchiffon":         color.RGBA{0xFF, 0xFA, 0xCD, 0xFF},
		"lightblue":            color.RGBA{0xAD, 0xD8, 0xE6, 0xFF},
		"lightcoral":           color.RGBA{0xF0, 0x80, 0x80, 0xFF},
		"lightcyan":            color.RGBA{0xE0, 0xFF, 0xFF, 0xFF},
		"lightgoldenrodyellow": color.RGBA{0xFA, 0xFA, 0xD2, 0xFF},
		"lightgray":            color.RGBA{0xD3, 0xD3, 0xD3, 0xFF},
		"lightgrey":            color.RGBA{0xD3, 0xD3, 0xD3, 0xFF},
		"lightgreen":           color.RGBA{0x90, 0xEE, 0x90, 0xFF},
		"lightpink":            color.RGBA{0xFF, 0xB6, 0xC1, 0xFF},
		"lightsalmon":          color.RGBA{0xFF, 0xA0, 0x7A, 0xFF},
		"lightseagreen":        color.RGBA{0x20, 0xB2, 0xAA, 0xFF},
		"lightskyblue":         color.RGBA{0x87, 0xCE, 0xFA, 0xFF},
		"lightslategray":       color.RGBA{0x77, 0x88, 0x99, 0xFF},
		"lightslategrey":       color.RGBA{0x77, 0x88, 0x99, 0xFF},
		"lightsteelblue":       color.RGBA{0xB0, 0xC4, 0xDE, 0xFF},
		"lightyellow":          color.RGBA{0xFF, 0xFF, 0xE0, 0xFF},
		"lime":                 color.RGBA{0x00, 0xFF, 0x00, 0xFF},
		"limegreen":            color.RGBA{0x32, 0xCD, 0x32, 0xFF},
		"linen":                color.RGBA{0xFA, 0xF0, 0xE6, 0xFF},
		"magenta":              color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
		"maroon":               color.RGBA{0x80, 0x00, 0x00, 0xFF},
		"mediumaquamarine":     color.RGBA{0x66, 0xCD, 0xAA, 0xFF},
		"mediumblue":           color.RGBA{0x00, 0x00, 0xCD, 0xFF},
		"mediumorchid":         color.RGBA{0xBA, 0x55, 0xD3, 0xFF},
		"mediumpurple":         color.RGBA{0x93, 0x70, 0xD8, 0xFF},
		"mediumseagreen":       color.RGBA{0x3C, 0xB3, 0x71, 0xFF},
		"mediumslateblue":      color.RGBA{0x7B, 0x68, 0xEE, 0xFF},
		"mediumspringgreen":    color.RGBA{0x00, 0xFA, 0x9A, 0xFF},
		"mediumturquoise":      color.RGBA{0x48, 0xD1, 0xCC, 0xFF},
		"mediumvioletred":      color.RGBA{0xC7, 0x15, 0x85, 0xFF},
		"midnightblue":         color.RGBA{0x19, 0x19, 0x70, 0xFF},
		"mintcream":            color.RGBA{0xF5, 0xFF, 0xFA, 0xFF},
		"mistyrose":            color.RGBA{0xFF, 0xE4, 0xE1, 0xFF},
		"moccasin":             color.RGBA{0xFF, 0xE4, 0xB5, 0xFF},
		"navajowhite":          color.RGBA{0xFF, 0xDE, 0xAD, 0xFF},
		"navy":                 color.RGBA{0x00, 0x00, 0x80, 0xFF},
		"oldlace":              color.RGBA{0xFD, 0xF5, 0xE6, 0xFF},
		"olive":                color.RGBA{0x80, 0x80, 0x00, 0xFF},
		"olivedrab":            color.RGBA{0x6B, 0x8E, 0x23, 0xFF},
		"orange":               color.RGBA{0xFF, 0xA5, 0x00, 0xFF},
		"orangered":            color.RGBA{0xFF, 0x45, 0x00, 0xFF},
		"orchid":               color.RGBA{0xDA, 0x70, 0xD6, 0xFF},
		"palegoldenrod":        color.RGBA{0xEE, 0xE8, 0xAA, 0xFF},
		"palegreen":            color.RGBA{0x98, 0xFB, 0x98, 0xFF},
		"paleturquoise":        color.RGBA{0xAF, 0xEE, 0xEE, 0xFF},
		"palevioletred":        color.RGBA{0xD8, 0x70, 0x93, 0xFF},
		"papayawhip":           color.RGBA{0xFF, 0xEF, 0xD5, 0xFF},
		"peachpuff":            color.RGBA{0xFF, 0xDA, 0xB9, 0xFF},
		"peru":                 color.RGBA{0xCD, 0x85, 0x3F, 0xFF},
		"pink":                 color.RGBA{0xFF, 0xC0, 0xCB, 0xFF},
		"plum":                 color.RGBA{0xDD, 0xA0, 0xDD, 0xFF},
		"powderblue":           color.RGBA{0xB0, 0xE0, 0xE6, 0xFF},
		"purple":               color.RGBA{0x80, 0x00, 0x80, 0xFF},
		"red":                  color.RGBA{0xFF, 0x00, 0x00, 0xFF},
		"rosybrown":            color.RGBA{0xBC, 0x8F, 0x8F, 0xFF},
		"royalblue":            color.RGBA{0x41, 0x69, 0xE1, 0xFF},
		"saddlebrown":          color.RGBA{0x8B, 0x45, 0x13, 0xFF},
		"salmon":               color.RGBA{0xFA, 0x80, 0x72, 0xFF},
		"sandybrown":           color.RGBA{0xF4, 0xA4, 0x60, 0xFF},
		"seagreen":             color.RGBA{0x2E, 0x8B, 0x57, 0xFF},
		"seashell":             color.RGBA{0xFF, 0xF5, 0xEE, 0xFF},
		"sienna":               color.RGBA{0xA0, 0x52, 0x2D, 0xFF},
		"silver":               color.RGBA{0xC0, 0xC0, 0xC0, 0xFF},
		"skyblue":              color.RGBA{0x87, 0xCE, 0xEB, 0xFF},
		"slateblue":            color.RGBA{0x6A, 0x5A, 0xCD, 0xFF},
		"slategray":            color.RGBA{0x70, 0x80, 0x90, 0xFF},
		"slategrey":            color.RGBA{0x70, 0x80, 0x90, 0xFF},
		"snow":                 color.RGBA{0xFF, 0xFA, 0xFA, 0xFF},
		"springgreen":          color.RGBA{0x00, 0xFF, 0x7F, 0xFF},
		"steelblue":            color.RGBA{0x46, 0x82, 0xB4, 0xFF},
		"tan":                  color.RGBA{0xD2, 0xB4, 0x8C, 0xFF},
		"teal":                 color.RGBA{0x00, 0x80, 0x80, 0xFF},
		"thistle":              color.RGBA{0xD8, 0xBF, 0xD8, 0xFF},
		"tomato":               color.RGBA{0xFF, 0x63, 0x47, 0xFF},
		"turquoise":            color.RGBA{0x40, 0xE0, 0xD0, 0xFF},
		"violet":               color.RGBA{0xEE, 0x82, 0xEE, 0xFF},
		"wheat":                color.RGBA{0xF5, 0xDE, 0xB3, 0xFF},
		"white":                color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		"whitesmoke":           color.RGBA{0xF5, 0xF5, 0xF5, 0xFF},
		"yellow":               color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
		"yellowgreen":          color.RGBA{0x9A, 0xCD, 0x32, 0xFF},
	}
)

// relativeLuminance calculates the relative luminance of the given color
func relativeLuminance(c color.Color) float64 {
	srgba := SRGBAModel.Convert(c).(SRGBA)

	var r, g, b float64
	if srgba.R <= 0.3928 {
		r = srgba.R / 12.92
	} else {
		r = math.Pow((srgba.R+0.055)/1.055, 2.4)
	}
	if srgba.G <= 0.3928 {
		g = srgba.G / 12.92
	} else {
		g = math.Pow((srgba.G+0.055)/1.055, 2.4)
	}
	if srgba.R <= 0.3928 {
		b = srgba.B / 12.92
	} else {
		b = math.Pow((srgba.B+0.055)/1.055, 2.4)
	}

	return 0.2126*r + 0.7152*g + 0.0722*b
}

// EqualColors returns true if both colors converted to the RGBA colorspace are equal.
func EqualColors(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

// PickBestTextColor compares each of the given text colors and chooses the best one for the given background color.
// If no text colors are given, this function chooses between black and white.
func PickBestTextColor(backgroundColor color.Color, textColors ...color.Color) color.Color {
	backgroundColorRelativeLuminance := relativeLuminance(backgroundColor)
	if len(textColors) == 0 {
		if backgroundColorRelativeLuminance > math.Sqrt(1.05*0.05)-0.05 {
			return color.Black
		} else {
			return color.White
		}
	} else {
		var bestTextColor color.Color
		var bestContrastRatio float64
		for _, textColor := range textColors {
			textColorRelativeLuminance := relativeLuminance(textColor)

			var contrastRatio float64
			if textColorRelativeLuminance > backgroundColorRelativeLuminance {
				contrastRatio = (textColorRelativeLuminance + 0.05) / (backgroundColorRelativeLuminance + 0.05)
			} else {
				contrastRatio = (backgroundColorRelativeLuminance + 0.05) / (textColorRelativeLuminance + 0.05)
			}

			if contrastRatio > bestContrastRatio {
				bestContrastRatio = contrastRatio
				bestTextColor = textColor
			}
		}

		return bestTextColor
	}
}

// ParseCSSColorRepresentation parses a legal CSS color value into a color.Color
// The model of the returned color will be the same as the value passed in
// For example, if the CSS value uses the hsl() function, its color model will be colorhelper.HSLAModel
func ParseCSSColorRepresentation(colorRepresentation string) (color.Color, error) {
	if hexTripletRegexp.MatchString(colorRepresentation) {
		hexTripletValue := hexTripletRegexp.FindStringSubmatch(colorRepresentation)
		var r, g, b uint8
		if ri, err := strconv.ParseInt(hexTripletValue[1], 16, 8); err == nil {
			r = uint8(ri)
		} else {
			return nil, fmt.Errorf("cannot parse red value %v of %v as hex triplet: %v", hexTripletValue[1], colorRepresentation, err)
		}
		if gi, err := strconv.ParseInt(hexTripletValue[2], 16, 8); err == nil {
			g = uint8(gi)
		} else {
			return nil, fmt.Errorf("cannot parse green value %v of %v as hex triplet: %v", hexTripletValue[2], colorRepresentation, err)
		}
		if bi, err := strconv.ParseInt(hexTripletValue[3], 16, 8); err == nil {
			b = uint8(bi)
		} else {
			return nil, fmt.Errorf("cannot parse blue value %v of %v as hex triplet: %v", hexTripletValue[3], colorRepresentation, err)
		}
		return color.RGBA{r, g, b, 0xFF}, nil
	} else if hexQuadrupletRegexp.MatchString(colorRepresentation) {
		hexQuadrupletValue := hexQuadrupletRegexp.FindStringSubmatch(colorRepresentation)
		var r, g, b, a uint8
		if ri, err := strconv.ParseUint(hexQuadrupletValue[1], 16, 8); err == nil {
			r = uint8(ri)
		} else {
			return nil, fmt.Errorf("cannot parse red value %v of %v as hex quadruplet: %v", hexQuadrupletValue[1], colorRepresentation, err)
		}
		if gi, err := strconv.ParseUint(hexQuadrupletValue[2], 16, 8); err == nil {
			g = uint8(gi)
		} else {
			return nil, fmt.Errorf("cannot parse green value %v of %v as hex quadruplet: %v", hexQuadrupletValue[2], colorRepresentation, err)
		}
		if bi, err := strconv.ParseUint(hexQuadrupletValue[3], 16, 8); err == nil {
			b = uint8(bi)
		} else {
			return nil, fmt.Errorf("cannot parse blue value %v of %v as hex quadruplet: %v", hexQuadrupletValue[3], colorRepresentation, err)
		}
		if ai, err := strconv.ParseUint(hexQuadrupletValue[4], 16, 8); err == nil {
			a = uint8(ai)
		} else {
			return nil, fmt.Errorf("cannot parse alpha value %v of %v as hex quadruplet: %v", hexQuadrupletValue[4], colorRepresentation, err)
		}
		return color.RGBA{r, g, b, a}, nil
	} else if rgbFunctionRegexp.MatchString(colorRepresentation) {
		rgbFunctionValue := rgbFunctionRegexp.FindStringSubmatch(colorRepresentation)
		var r, g, b uint8
		if strings.HasSuffix(rgbFunctionValue[1], "%") {
			if rf, err := strconv.ParseFloat(strings.TrimSuffix(rgbFunctionValue[1], "%"), 64); err == nil {
				if rf > 100.0 {
					return nil, fmt.Errorf("invalid percentage red value %v in %v", rgbFunctionValue[1], colorRepresentation)
				}
				r = uint8(rf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage red value %v of %v as an rgb() function call: %v", rgbFunctionValue[1], colorRepresentation, err)
			}
		} else {
			if ri, err := strconv.ParseUint(rgbFunctionValue[1], 10, 8); err == nil {
				r = uint8(ri)
			} else {
				return nil, fmt.Errorf("cannot parse integer red value %v of %v as an rgb() function call: %v", rgbFunctionValue[2], colorRepresentation, err)
			}
		}
		if strings.HasSuffix(rgbFunctionValue[2], "%") {
			if gf, err := strconv.ParseFloat(strings.TrimSuffix(rgbFunctionValue[2], "%"), 64); err == nil {
				if gf > 100.0 {
					return nil, fmt.Errorf("invalid percentage green value %v in %v", rgbFunctionValue[2], colorRepresentation)
				}
				g = uint8(gf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage green value %v of %v as an rgb() function call: %v", rgbFunctionValue[2], colorRepresentation, err)
			}
		} else {
			if gi, err := strconv.ParseUint(rgbFunctionValue[2], 10, 8); err == nil {
				g = uint8(gi)
			} else {
				return nil, fmt.Errorf("cannot parse integer green value %v of %v as an rgb() function call: %v", rgbFunctionValue[2], colorRepresentation, err)
			}
		}
		if strings.HasSuffix(rgbFunctionValue[3], "%") {
			if bf, err := strconv.ParseFloat(strings.TrimSuffix(rgbFunctionValue[3], "%"), 64); err == nil {
				if bf > 100.0 {
					return nil, fmt.Errorf("invalid percentage blue value %v in %v", rgbFunctionValue[3], colorRepresentation)
				}
				b = uint8(bf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage blue value %v of %v as an rgb() function call: %v", rgbFunctionValue[3], colorRepresentation, err)
			}
		} else {
			if bi, err := strconv.ParseUint(rgbFunctionValue[3], 10, 8); err == nil {
				b = uint8(bi)
			} else {
				return nil, fmt.Errorf("cannot parse integer blue value %v of %v as an rgb() function call: %v", rgbFunctionValue[3], colorRepresentation, err)
			}
		}
		return color.RGBA{r, g, b, 0xFF}, nil
	} else if rgbaFunctionRegexp.MatchString(colorRepresentation) {
		rgbaFunctionValue := rgbFunctionRegexp.FindStringSubmatch(colorRepresentation)
		var r, g, b, a uint8
		if strings.HasSuffix(rgbaFunctionValue[1], "%") {
			if rf, err := strconv.ParseFloat(strings.TrimSuffix(rgbaFunctionValue[1], "%"), 64); err == nil {
				if rf > 100.0 {
					return nil, fmt.Errorf("invalid percentage red value %v in %v", rgbaFunctionValue[1], colorRepresentation)
				}
				r = uint8(rf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[1], colorRepresentation, err)
			}
		} else {
			if ri, err := strconv.ParseUint(rgbaFunctionValue[1], 10, 8); err == nil {
				r = uint8(ri)
			} else {
				return nil, fmt.Errorf("cannot parse integer red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[2], colorRepresentation, err)
			}
		}
		if strings.HasSuffix(rgbaFunctionValue[2], "%") {
			if gf, err := strconv.ParseFloat(strings.TrimSuffix(rgbaFunctionValue[2], "%"), 64); err == nil {
				if gf > 100.0 {
					return nil, fmt.Errorf("invalid percentage red value %v in %v", rgbaFunctionValue[2], colorRepresentation)
				}
				g = uint8(gf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[2], colorRepresentation, err)
			}
		} else {
			if gi, err := strconv.ParseUint(rgbaFunctionValue[2], 10, 8); err == nil {
				g = uint8(gi)
			} else {
				return nil, fmt.Errorf("cannot parse integer red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[2], colorRepresentation, err)
			}
		}
		if strings.HasSuffix(rgbaFunctionValue[3], "%") {
			if bf, err := strconv.ParseFloat(strings.TrimSuffix(rgbaFunctionValue[3], "%"), 64); err == nil {
				if bf > 100.0 {
					return nil, fmt.Errorf("invalid percentage red value %v in %v", rgbaFunctionValue[3], colorRepresentation)
				}
				b = uint8(bf / 100.0 * 255.0)
			} else {
				return nil, fmt.Errorf("cannot parse percentage red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[3], colorRepresentation, err)
			}
		} else {
			if bi, err := strconv.ParseUint(rgbaFunctionValue[3], 10, 8); err == nil {
				b = uint8(bi)
			} else {
				return nil, fmt.Errorf("cannot parse integer red value %v of %v as an rgba() function call: %v", rgbaFunctionValue[3], colorRepresentation, err)
			}
		}
		if af, err := strconv.ParseFloat(rgbaFunctionValue[4], 64); err == nil {
			if af < 0.0 || af > 1.0 {
				return nil, fmt.Errorf("invalid alpha value %v in %v", rgbaFunctionValue[4], colorRepresentation)
			}
			a = uint8(af * 255.0)
		} else {
			return nil, fmt.Errorf("cannot parse float alpha value %v in %v", rgbaFunctionValue[4], colorRepresentation)
		}
		return color.RGBA{r, g, b, a}, nil
	} else if hslFunctionRegexp.MatchString(colorRepresentation) {
		hslFunctionValue := hslFunctionRegexp.FindStringSubmatch(colorRepresentation)
		var h, s, l float64
		if hi, err := strconv.ParseUint(hslFunctionValue[1], 10, 16); err == nil {
			if hi > 360 {
				return nil, fmt.Errorf("invalid hue value %v for %v", hslFunctionValue[1], colorRepresentation)
			}
			h = float64(hi)
		} else {
			return nil, fmt.Errorf("cannot parse hue value %v of %v as hsl() function call: %v", hslFunctionValue[1], colorRepresentation, err)
		}
		if sf, err := strconv.ParseFloat(strings.TrimSuffix(hslFunctionValue[2], "%"), 64); err == nil {
			if sf > 100.0 {
				return nil, fmt.Errorf("invalid saturation value %v for %v", hslFunctionValue[2], colorRepresentation)
			}
			s = sf / 100.0
		} else {
			return nil, fmt.Errorf("cannot parse saturation value %v of %v as hsl() function call: %v", hslFunctionValue[2], colorRepresentation, err)
		}
		if lf, err := strconv.ParseFloat(strings.TrimSuffix(hslFunctionValue[3], "%"), 64); err == nil {
			if lf > 100.0 {
				return nil, fmt.Errorf("invalid lightness value %v for %v", hslFunctionValue[3], colorRepresentation)
			}
			l = lf / 100.0
		} else {
			return nil, fmt.Errorf("cannot parse lightness value %v of %v as hsl() function call: %v", hslFunctionValue[2], colorRepresentation, err)
		}
		return HSLA{h, s, l, 1.0}, nil
	} else if hslaFunctionRegexp.MatchString(colorRepresentation) {
		hslaFunctionValue := hslFunctionRegexp.FindStringSubmatch(colorRepresentation)
		var h, s, l, a float64
		if hi, err := strconv.ParseUint(hslaFunctionValue[1], 10, 16); err == nil {
			if hi > 360 {
				return nil, fmt.Errorf("invalid hue value %v for %v", hslaFunctionValue[1], colorRepresentation)
			}
			h = float64(hi)
		} else {
			return nil, fmt.Errorf("cannot parse hue value %v of %v as hsla() function call: %v", hslaFunctionValue[1], colorRepresentation, err)
		}
		if sf, err := strconv.ParseFloat(strings.TrimSuffix(hslaFunctionValue[2], "%"), 64); err == nil {
			if sf > 100.0 {
				return nil, fmt.Errorf("invalid saturation value %v for %v", hslaFunctionValue[2], colorRepresentation)
			}
			s = sf / 100.0
		} else {
			return nil, fmt.Errorf("cannot parse saturation value %v of %v as hsla() function call: %v", hslaFunctionValue[2], colorRepresentation, err)
		}
		if lf, err := strconv.ParseFloat(strings.TrimSuffix(hslaFunctionValue[3], "%"), 64); err == nil {
			if lf > 100.0 {
				return nil, fmt.Errorf("invalid lightness value %v for %v", hslaFunctionValue[3], colorRepresentation)
			}
			l = lf / 100.0
		} else {
			return nil, fmt.Errorf("cannot parse lightness value %v of %v as hsla() function call: %v", hslaFunctionValue[2], colorRepresentation, err)
		}
		if af, err := strconv.ParseFloat(hslaFunctionValue[4], 64); err == nil {
			if af < 0.0 || af > 1.0 {
				return nil, fmt.Errorf("invalid alpha value %v in %v", hslaFunctionValue[4], colorRepresentation)
			}
			a = af
		} else {
			return nil, fmt.Errorf("cannot parse float alpha value %v in %v", hslaFunctionValue[4], colorRepresentation)
		}
		return HSLA{h, s, l, a}, nil
	} else if namedColorValue, ok := colorNamesToColorValues[strings.ToLower(colorRepresentation)]; ok {
		return namedColorValue, nil
	} else {
		return nil, fmt.Errorf("invalid color value %v", colorRepresentation)
	}
}

const (
	AnyRepresentation           int = iota // Use best representation based upon the color.Model of the color.Color
	HexTripletRepresentation               // #xxxxxx
	HexQuadrupletRepresentation            // #xxxxxxxx
	RGBFunctionRepresentation              // rgb(r,g,b)
	RGBAFunctionRepresentation             // rgba(r,g,b,a)
	HSLFunctionRepresentation              // hsl(h,s,l)
	HSLAFunctionRepresentation             // hsla(h,s,l,a)
)

// MakeCSSColorRepresentation generates a legal CSS color value from a color.Color
func MakeCSSColorRepresentation(colorValue color.Color, colorRepresentation int) string {
	switch colorRepresentation {
	case HexTripletRepresentation:
		colorRGBA := color.RGBAModel.Convert(colorValue).(color.RGBA)
		return fmt.Sprintf("#%02X%02X%02X", colorRGBA.R, colorRGBA.G, colorRGBA.B)
	case HexQuadrupletRepresentation:
		colorRGBA := color.RGBAModel.Convert(colorValue).(color.RGBA)
		return fmt.Sprintf("#%02X%02X%02X%02X", colorRGBA.R, colorRGBA.G, colorRGBA.B, colorRGBA.A)
	case RGBFunctionRepresentation:
		colorRGBA := color.RGBAModel.Convert(colorValue).(color.RGBA)
		return fmt.Sprintf("rgb(%d,%d,%d)", colorRGBA.R, colorRGBA.G, colorRGBA.B)
	case RGBAFunctionRepresentation:
		colorRGBA := color.RGBAModel.Convert(colorValue).(color.RGBA)
		return fmt.Sprintf("rgba(%d,%d,%d,%f)", colorRGBA.R, colorRGBA.G, colorRGBA.B, float64(colorRGBA.A)/255.0)
	case HSLFunctionRepresentation:
		colorHSLA := HSLAModel.Convert(colorValue).(HSLA)
		return fmt.Sprintf("hsl(%d,%f%%,%f%%)", uint16(colorHSLA.H), colorHSLA.S*100.0, colorHSLA.L*100.0)
	case HSLAFunctionRepresentation:
		colorHSLA := HSLAModel.Convert(colorValue).(HSLA)
		return fmt.Sprintf("hsla(%d,%f%%,%f%%,%f)", uint16(colorHSLA.H), colorHSLA.S*100.0, colorHSLA.L*100.0, colorHSLA.A)
	default:
		switch trueColorValue := colorValue.(type) {
		case SRGBA:
			if trueColorValue.A == 1.0 {
				return fmt.Sprintf("rgb(%f%%,%f%%,%f%%)", trueColorValue.R*100.0, trueColorValue.G*100.0, trueColorValue.B*100.0)
			} else {
				return fmt.Sprintf("rgba(%f%%,%f%%,%f%%,%f)", trueColorValue.R*100.0, trueColorValue.G*100.0, trueColorValue.B*100.0, trueColorValue.A)
			}
		case HSLA:
			if trueColorValue.A == 1.0 {
				return fmt.Sprintf("hsl(%d,%f%%,%f%%)", uint16(trueColorValue.H), trueColorValue.S*100.0, trueColorValue.L*100.0)
			} else {
				return fmt.Sprintf("hsla(%d,%f%%,%f%%,%f)", uint16(trueColorValue.H), trueColorValue.S*100.0, trueColorValue.L*100.0, trueColorValue.A)
			}
		default:
			colorRGBA := color.RGBAModel.Convert(colorValue).(color.RGBA)
			if colorRGBA.A == 255 {
				return fmt.Sprintf("rgb(%d,%d,%d)", colorRGBA.R, colorRGBA.G, colorRGBA.B)
			} else {
				return fmt.Sprintf("rgba(%d,%d,%d,%f)", colorRGBA.R, colorRGBA.G, colorRGBA.B, float64(colorRGBA.A)/255.0)
			}
		}
	}
}
