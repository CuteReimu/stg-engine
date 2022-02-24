package utils

import (
	"math"
	"strconv"
)

func RoundToPoint(rad, length float64) Point {
	sin, cos := math.Sincos(rad)
	return Point{length * cos, length * sin}
}

type Point struct {
	X, Y float64
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p Point) Sub(p2 Point) Point {
	return Point{p.X - p2.X, p.Y - p2.Y}
}

func (p Point) Diff(p2 Point) (float64, float64) {
	return p.X - p2.X, p.Y - p2.Y
}

// String returns a string representation of p like "(3,4)".
func (p Point) String() string {
	return "(" + strconv.Itoa(int(p.X)) + "," + strconv.Itoa(int(p.Y)) + ")"
}

type Rectangle struct {
	Min, Max Point
}

// String returns a string representation of r like "(3,4)-(6,5)".
func (r Rectangle) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Dx returns r's width.
func (r Rectangle) Dx() float64 {
	return r.Max.X - r.Min.X
}

// Dy returns r's height.
func (r Rectangle) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height.
func (r Rectangle) Size() Point {
	return Point{
		r.Max.X - r.Min.X,
		r.Max.Y - r.Min.Y,
	}
}

// Rect build a utils.Rectangle
func Rect(x, y, w, h float64) Rectangle {
	return Rectangle{Point{x, y}, Point{x + w, y + h}}
}
