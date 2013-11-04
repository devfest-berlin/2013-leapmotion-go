package leapmotion

import (
	"encoding/json"
	"testing"
)

func TestFrame(t *testing.T) {
	frame := Frame{}
	if err := json.Unmarshal(testFrame, &frame); err != nil {
		t.Fatalf("Couln't unmarshal JSON to Frame: %v", err)
	}
}

var testFrame = []byte(`
{
  "id": 1,
  "timestamp": 1,
  "currentFrameRate": 100.0,
  "r": [
    [-1.0, 0, 1.0],
    [-1.0, 0, 1.0],
    [-1.0, 0, 1.0]
  ],
  "s": -75.0,
  "t": [-1.0, 0, 1.0],
  "hands": [
    {
      "id": 100,
      "direction": [-1.0, 0, 1.0],
      "timeVisible": 1.0,
      "palmNormal": [-1.0, 0, 1.0],
      "palmPosition": [-1.0, 0, 1.0],
      "palmVelocity": [-1.0, 0, 1.0],
      "stabilizedPalmPosition": [-1.0, 0, 1.0],
      "sphereCenter": [-1.0, 0, 1.0],
      "sphereRadius": 50.0,
      "r": [
        [-1.0, 0, 1.0],
        [-1.0, 0, 1.0],
        [-1.0, 0, 1.0]
      ],
      "s": 1.0,
      "t": [-1.0, 0, 1.0]
    },
    {
      "id": 200,
      "direction": [-1.0, 0, 1.0],
      "timeVisible": 1.0,
      "palmNormal": [-1.0, 0, 1.0],
      "palmPosition": [-1.0, 0, 1.0],
      "palmVelocity": [-1.0, 0, 1.0],
      "stabilizedPalmPosition": [-1.0, 0, 1.0],
      "sphereCenter": [-1.0, 0, 1.0],
      "sphereRadius": 50.0,
      "r": [
        [-1.0, 0, 1.0],
        [-1.0, 0, 1.0],
        [-1.0, 0, 1.0]
      ],
      "s": 1.0,
      "t": [-1.0, 0, 1.0]
    }
  ],
  "interactionBox": {
    "center": [0, 200, 0],
    "size": [221.418, 221.418, 154.742]
  },
  "pointables": [
    {
      "id": 300,
      "handId": 100,
      "length": 70.0,
      "direction": [-1.0, 0, 1.0],
      "timeVisible": 1.0,
      "tipPosition": [-1.0, 0, 1.0],
      "tipVelocity": [-1.0, 0, 1.0],
      "stabilizedTipPosition": [-1.0, 0, 1.0],
      "touchZone": "hovering",
      "touchDistance": 0.5,
      "tool": false
    },
    {
      "id": 300,
      "handId": 200,
      "length": 70.0,
      "direction": [-1.0, 0, 1.0],
      "timeVisible": 1.0,
      "tipPosition": [-1.0, 0, 1.0],
      "tipVelocity": [-1.0, 0, 1.0],
      "stabilizedTipPosition": [-1.0, 0, 1.0],
      "touchZone": "hovering",
      "touchDistance": 0.5,
      "tool": false
    }
  ],
  "gestures": []
}
`)
