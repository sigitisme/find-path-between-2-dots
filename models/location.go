package models

import "fmt"

// Location is to hold x, y position
type Location struct {
	X int
	Y int
}

// String it to print the location
func (l *Location) String() string {
	return fmt.Sprintf("%d:%d", l.X, l.Y)
}

// Compare returns `true` if `l` has same x and y values with `comp`
func (l *Location) Compare(comp *Location) bool {
	if l.X == comp.X && l.Y == comp.Y {
		return true
	}

	return false
}

// Down returns new location if `l` is going down
func (l *Location) Down() *Location {
	return &Location{
		X: l.X + 1,
		Y: l.Y,
	}
}

// Right returns new location if `l` is going right
func (l *Location) Right() *Location {
	return &Location{
		X: l.X,
		Y: l.Y + 1,
	}
}

// Up returns new location if `l` is going up
func (l *Location) Up() *Location {
	return &Location{
		X: l.X - 1,
		Y: l.Y,
	}

}

// Left returns new location if `l` is going left
func (l *Location) Left() *Location {
	return &Location{
		X: l.X,
		Y: l.Y - 1,
	}
}

// Boundary is used to check boundaries of the location with corresponding matrix
type Boundary struct {
	InboundDown  bool
	InboundRight bool
	InboundUp    bool
	InboundLeft  bool
}

// CheckBoundaries checks if `l` is still in the boundary when it moves
// returns Boundary
func (l *Location) CheckBoundaries(m, n int) Boundary {
	return Boundary{
		InboundDown:  l.X+1 < m,
		InboundRight: l.Y+1 < n,
		InboundUp:    l.X-1 >= 0,
		InboundLeft:  l.Y-1 >= 0,
	}
}
