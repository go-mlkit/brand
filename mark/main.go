// Command mark writes the go-mlkit mark, svg/color/go-mlkit.svg:
//
//	cd mark && go run . > ../svg/color/go-mlkit.svg
//
// go-mlkit is machine learning, and the mark is a decision tree: a root that
// splits in two, each branch splitting again, with a leaf at the bottom of
// every path. Trees and forests are the classifier the library is named for,
// and a tree is also what clustering and cross-validation draw.
//
// Nothing is placed by hand. The program builds a complete binary tree of a
// given depth, puts level l at an even vertical step and its 2^l nodes at an
// even horizontal step, centres the whole on the square, and joins each node
// to its two children by a line that stops on the rim of each ring. Change the
// depth and the tree redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#F0D060" // the gradient's light stop
	dark  = "#7A5C00" // its dark stop

	stroke = 13.0 // the family's line width
	centre = 128.0

	levels = 3    // levels of the tree: 1 root, 2 children, 4 leaves
	radius = 13.0 // radius of a ring, between stroke centres
	hstep  = 54.0 // horizontal distance between two neighbours on the last level
	vstep  = 62.0 // vertical distance between two levels
)

type point struct{ x, y float64 }

// position is the centre of node i (0-based, left to right) on level l.
func position(l, i int) point {
	span := hstep * math.Pow(2, float64(levels-1-l)) // gap between nodes on this level
	n := 1 << l
	x := centre + (float64(i)-float64(n-1)/2)*span
	y := centre + (float64(l)-float64(levels-1)/2)*vstep
	return point{x, y}
}

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

// edge is the line from a parent to a child, shortened at both ends so that it
// meets the rings on their rim.
func edge(p, c point) string {
	dx, dy := c.x-p.x, c.y-p.y
	n := math.Hypot(dx, dy)
	ux, uy := dx/n, dy/n
	return fmt.Sprintf(`<path d="M%s %sL%s %s"/>`,
		num(p.x+ux*radius), num(p.y+uy*radius),
		num(c.x-ux*radius), num(c.y-uy*radius))
}

func main() {
	var items []string
	for l := 0; l < levels; l++ {
		for i := 0; i < 1<<l; i++ {
			p := position(l, i)
			if l+1 < levels {
				items = append(items, edge(p, position(l+1, 2*i)), edge(p, position(l+1, 2*i+1)))
			}
		}
	}
	for l := 0; l < levels; l++ {
		for i := 0; i < 1<<l; i++ {
			p := position(l, i)
			items = append(items, fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(p.x), num(p.y), num(radius)))
		}
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-mlkit">
  <defs><linearGradient id="mlgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#mlgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(items, "\n    "))
}
