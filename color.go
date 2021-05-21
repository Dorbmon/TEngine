package TEngine

import (
	"TEngine/thirdparty"
	"strconv"
)

type RColor struct {
	R, G, B, Alpha uint8
}

func (v *RColor) ToRGBA() (r, g, b, a uint8) {
	return v.R, v.G, v.B, v.Alpha
}
func ToRColor(r, g, b, a uint8) RColor {
	return RColor{R: r, G: g, B: b, Alpha: a}
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}
func (z *RColor) ToHex() int {
	r := t2x(int64(z.R))
	g := t2x(int64(z.G))
	b := t2x(int64(z.B))
	v, _ := strconv.Atoi(r + g + b)
	return v
}
func NewColorFromHex(Hex int, Alpha uint8) RColor {
	r, g, b := thirdparty.HexToRGB(Hex)
	return RColor{
		R:     r,
		G:     g,
		B:     b,
		Alpha: Alpha,
	}
}
