// Copyright (c) 2013 David Lehmann <http://github.com/dtl>
// Use of this source code is governed by the MIT Licence
// that can be found in the LICENSE file.

package leapmotion

// Leap Motion data structures for JSON decoding.

import (
	"encoding/json"
)

// Leap Motion data structures for tracking data.
type Vector [3]float64
type Matrix [3]Vector

// A Frame represents one single tracking point. It contains general
// information about tracked motion in the Leap field of view and lists of
// specific values for hands, pointables like fingers and gestures.
type Frame struct {
	Id               int     `json:"id"`
	Timestamp        int     `json:"timestamp"`
	CurrentFrameRate float64 `json:"currentFrameRate"`
	R                Matrix  `json:"r"`
	S                float64 `json:"s"`
	T                Vector  `json:"t"`
	Hands            []struct {
		Id                     int     `json:"id"`
		Direction              Vector  `json:"direction"`
		TimeVisible            float64 `json:"timeVisible"`
		PalmNormal             Vector  `json:"palmNormal"`
		PalmPosition           Vector  `json:"palmPosition"`
		PalmVelocity           Vector  `json:"palmVelocity"`
		StabelizedPalmPosition Vector  `json:"stabelizedPalmPosition"`
		SphereCenter           Vector  `json:"sphereCenter"`
		SphereRadius           float64 `json:"sphereRadius"`
		R                      Matrix  `json:"r"`
		S                      float64 `json:"s"`
		T                      Vector  `json:"t"`
	} `json:"hands"`
	InteractionBox struct {
		Center Vector `json:"center"`
		Size   Vector `json:"size"`
	} `json:"interactionBox"`
	Pointables []struct {
		Id                    int     `json:"id"`
		HandId                int     `json:"handId"`
		Length                float64 `json:"length"`
		Direction             Vector  `json:"direction"`
		TimeVisible           float64 `json:"timeVisible"`
		TipPosition           Vector  `json:"tipPosition"`
		TipVelocity           Vector  `json:"tipVelocity"`
		StabilizedTipPosition Vector  `json:"stabilizedTipPosition"`
		TouchZone             string  `json:"touchZone"`
		TouchDistance         float64 `json:"touchDistance"`
		Tool                  bool    `json:"tool"`
	} `json:"pointables"`
	Gestures []struct {
		Id           int       `json:"id"`
		Center       Vector    `json:"center"`
		Duration     int       `json:"duration"`
		HandIds      []int     `json:"handIds"`
		Normal       []float64 `json:"normal"`
		PointableIds []int     `json:"pointableIds"`
		Progress     float64   `json:"progress"`
		Radius       float64   `json:"radius"`
		State        string    `json:"state"`
		Type         string    `json:"type"`
	} `json:"gestures"`
}

// Stringer interface implementation.
func (f *Frame) String() string {
	b, err := json.Marshal(f)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// Returns the palm position for a specific hand.
func (f *Frame) Position(hand, axis int) float64 {
	if (len(f.Hands) < 1) || (len(f.Hands) < hand) || (axis < 0) || (axis > 2) {
		return 0.0
	}

	return f.Hands[hand].PalmPosition[axis] / f.InteractionBox.Size[axis]
}

// Simplifyed access to the x-position of the first hand.
func (f *Frame) X() float64 {
	return f.Position(0, 0)
}

// Simplifyed access to the y-position of the first hand.
func (f *Frame) Y() float64 {
	return f.Position(0, 1)
}

// Simplifyed access to the z-position of the first hand.
func (f *Frame) Z() float64 {
	return f.Position(0, 2)
}
