package main

import "testing"

func TestGivenATriangleWhenGetAreaThenMustReturnArea(t *testing.T) {
	shape := triangle{base: 10, height: 10}

	area := shape.getArea()

	if area != 50 {
		t.Errorf("Expected 50, but got %v", area)
	}
}

func TestGivenASquareWhenGetAreaThenMustReturnArea(t *testing.T) {
	shape := square{sideLenght: 10}

	area := shape.getArea()

	if area != 100 {
		t.Errorf("Expected 100, but got %v", area)
	}
}
