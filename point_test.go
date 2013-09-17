package geo

import (
	"fmt"
	"testing"
)

// Tests that a call to NewPoint should return a pointer to a Point with the specified values assigned correctly.
func TestNewPoint(t *testing.T) {
	p := NewPoint(40.5, 120.5)

	if p == nil {
		t.Error("Expected to get a pointer to a new point, but got nil instead.")
	}

	if p.lat != 40.5 {
		t.Errorf("Expected to be able to specify 40.5 as the lat value of a new point, but got %f instead", p.lat)
	}

	if p.lng != 120.5 {
		t.Errorf("Expected to be able to specify 120.5 as the lng value of a new point, but got %f instead", p.lng)
	}
}

// Tests that calling GetLat() after creating a new point returns the expected lat value.
func TestLat(t *testing.T) {
	p := NewPoint(40.5, 120.5)

	lat := p.Lat()

	if lat != 40.5 {
		t.Error("Expected a call to GetLat() to return the same lat value as was set before, but got %f instead", lat)
	}
}

// Tests that calling GetLng() after creating a new point returns the expected lng value.
func TestLng(t *testing.T) {
	p := NewPoint(40.5, 120.5)

	lng := p.Lng()

	if lng != 120.5 {
		t.Error("Expected a call to GetLng() to return the same lat value as was set before, but got %f instead", lng)
	}
}

// Seems brittle :\
func TestGreatCircleDistance(t *testing.T) {
	// Test that SEA and SFO are ~ 1091km apart, accurate to 100 meters.
	sea := &Point{lat: 47.44745785, lng: -122.308065668024}
	sfo := &Point{lat: 37.6160933, lng: -122.3924223}
	sfoToSea := 1090.7

	dist := sea.GreatCircleDistance(sfo)

	if !(dist < (sfoToSea+0.1) && dist > (sfoToSea-0.1)) {
		t.Error("Unnacceptable result.", dist)
	}
}

func TestPointAtDistanceAndBearing(t *testing.T) {
	sea := &Point{lat: 47.44745785, lng: -122.308065668024}
	p := sea.PointAtDistanceAndBearing(1090.7, 180)

	// Expected results of transposing point
	// ~1091km at bearing of 180 degrees
	resultLat := 37.616572
	resultLng := -122.308066

	withinLatBounds := p.lat < resultLat+0.001 && p.lat > resultLat-0.001
	withinLngBounds := p.lng < resultLng+0.001 && p.lng > resultLng-0.001
	if !(withinLatBounds && withinLngBounds) {
		t.Error("Unnacceptable result.", fmt.Sprintf("[%f, %f]", p.lat, p.lng))
	}
}
