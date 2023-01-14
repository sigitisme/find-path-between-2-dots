package models_test

import (
	"amartha/models"
	"testing"
)

func TestLocation(t *testing.T) {
	newLocation := new(models.Location)

	//move to down
	downOneStep := newLocation.Down()

	downOneStep2 := newLocation.Down()

	if !downOneStep.Compare(downOneStep2) {
		t.Error("downOneStep must be same with downOneStep2")
	}

	if downOneStep.Compare(newLocation) {
		t.Error("downOneStep can't be same with newLocation")
	}

	if downOneStep.String() != "1:0" {
		t.Error("Down() not working correctly")
	}

	//move to right
	rightOneStep := newLocation.Right()

	if rightOneStep.Compare(newLocation) {
		t.Error("rightOneStep can't be same with newLocation")
	}

	if rightOneStep.String() != "0:1" {
		t.Error("Right() not working correctly")
	}

	//move to up
	upOneStep := newLocation.Up()

	if upOneStep.Compare(newLocation) {
		t.Error("upOneStep can't be same with newLocation")
	}

	if upOneStep.String() != "-1:0" {
		t.Error("Up() not working correctly")
	}

	//move to up
	leftOneStep := newLocation.Left()

	if leftOneStep.Compare(newLocation) {
		t.Error("leftOneStep can't be same with newLocation")
	}

	if leftOneStep.String() != "0:-1" {
		t.Error("Left() not working correctly")
	}

	//checkBoundaries
	boundary := newLocation.CheckBoundaries(4, 4)

	if !(boundary.InboundDown && boundary.InboundRight) {
		t.Error("CheckBoundaries not working correctly")
	}

	if boundary.InboundLeft && boundary.InboundUp {
		t.Error("CheckBoundaries not working correctly")
	}

	centerLocation := &models.Location{2, 2}

	newBoundary := centerLocation.CheckBoundaries(4, 4)

	if !(newBoundary.InboundDown && newBoundary.InboundRight) {
		t.Error("CheckBoundaries not working correctly")
	}

	if !(newBoundary.InboundLeft && newBoundary.InboundUp) {
		t.Error("CheckBoundaries not working correctly")
	}
}
