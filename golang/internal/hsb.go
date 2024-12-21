package hsb

import "math"

func HSBToRGB(h, s, v float64) (uint8, uint8, uint8) {
	if s == 0 {
		rgb := uint8(v * 255)
		return rgb, rgb, rgb
	}

	h = h * 6
	if h == 6 {
		h = 0
	}
	i := math.Floor(h)
	v1 := v * (1 - s)
	v2 := v * (1 - s*(h-i))
	v3 := v * (1 - s*(1-(h-i)))

	var r, g, b float64
	switch int(i) {
	case 0:
		r, g, b = v, v3, v1
	case 1:
		r, g, b = v2, v, v1
	case 2:
		r, g, b = v1, v, v3
	case 3:
		r, g, b = v1, v2, v
	case 4:
		r, g, b = v3, v1, v
	default:
		r, g, b = v, v1, v2
	}

	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}
