package main

import (
	"strconv"
	"strings"
)

// Triangle represents a triangle.
type Triangle struct {
	x1, y1 int
	x2, y2 int
	x3, y3 int
}

func (t *Triangle) init(s string) {
	var tmp []int
	for _, p := range strings.Split(s, ",") {
		t, _ := strconv.Atoi(p)
		tmp = append(tmp, t)
	}
	t.x1, t.y1 = tmp[0], tmp[1]
	t.x2, t.y2 = tmp[2], tmp[3]
	t.x3, t.y3 = tmp[4], tmp[5]
}

func (t *Triangle) det() int {
	return (t.y2-t.y3)*(t.x1-t.x3) + (t.x3-t.x2)*(t.y1-t.y3)
}

func (t *Triangle) barycentric(x, y int) (int, int, int) {
	l1 := (t.y2-t.y3)*(x-t.x3) + (t.x3-t.x2)*(y-t.y3)
	l2 := (t.y3-t.y1)*(x-t.x3) + (t.x1-t.x3)*(y-t.y3)
	l3 := t.det() - l1 - l2
	return l1, l2, l3
}

func (t *Triangle) isInside(x, y int) bool {
	d := t.det()
	a, b, c := t.barycentric(x, y)
	if d > 0 {
		return 0 < a && a < d && 0 < b && b < d && 0 < c && c < d
	}
	return 0 > a && a > d && 0 > b && b > d && 0 > c && c > d
}
