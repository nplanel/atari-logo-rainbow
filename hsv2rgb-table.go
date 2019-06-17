package main

import (
	"fmt"
	"math"
)

type HSV struct {
	H, S, V float64
}

type RGB struct {
	R, G, B float64
}

func (c *HSV) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSV from 0 to 1
		r = c.V * 255
		g = c.V * 255
		b = c.V * 255
	} else {
		h := c.H * 6
		if h == 6 {
			h = 0
		} //H must be < 1
		i := math.Floor(h) //Or ... var_i = floor( var_h )
		v1 := c.V * (1 - c.S)
		v2 := c.V * (1 - c.S*(h-i))
		v3 := c.V * (1 - c.S*(1-(h-i)))

		if i == 0 {
			r = c.V
			g = v3
			b = v1
		} else if i == 1 {
			r = v2
			g = c.V
			b = v1
		} else if i == 2 {
			r = v1
			g = c.V
			b = v3
		} else if i == 3 {
			r = v1
			g = v2
			b = c.V
		} else if i == 4 {
			r = v3
			g = v1
			b = c.V
		} else {
			r = c.V
			g = v1
			b = v2
		}

		r = r * 255 //RGB results from 0 to 255
		g = g * 255
		b = b * 255
	}
	rgb := &RGB{r, g, b}
	return rgb
}

func main() {
	steps := float64(41 * 2)
	for i := float64(0); i < 360; i += (360 / steps) {
		hsv := &HSV{i / 360.0, 1.0, 1.0}
		rgb := hsv.RGB()
		brightness := float64(0.25)
		fmt.Printf("0x%02x, 0x%02x, 0x%02x,\n", int(rgb.G*brightness), int(rgb.R*brightness), int(rgb.B*brightness))
	}
}
