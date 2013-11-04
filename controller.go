// Copyright (c) 2013 David Lehmann <http://github.com/dtl>
// Use of this source code is governed by the MIT Licence
// that can be found in the LICENSE file.

package leapmotion

// Leap Motion controller access for Go.

import (
	"code.google.com/p/go.net/websocket"
	"errors"
)

// Constants for websocket handling.
const (
	LEAPVERSION = 4
	URL         = "ws://localhost:6437/v4.json"
	ORIGIN      = "http://localhost"
)

// A controller holds a connection to the Leap Motion websocket server and is
// used for receiving tracking frames.
type Controller struct {
	ws *websocket.Conn
}

// Sets up and returns a working Controller instance.
func NewController() (Controller, error) {
	ws, err := websocket.Dial(URL, "", ORIGIN)
	if err != nil {
		return Controller{}, err
	}

	var version = map[string]int{"version": -1}
	websocket.JSON.Receive(ws, &version)

	if version["version"] != LEAPVERSION {
		return Controller{}, errors.New("Invalid Leap Motion version")
	}

	c := Controller{ws}
	websocket.Message.Send(ws, "{\"background\":true}")
	websocket.Message.Send(ws, "{\"gestures\":true}")

	return c, nil
}

// Receive a single tracking frame.
func (c Controller) Receive(frame *Frame) {
	websocket.JSON.Receive(c.ws, frame)
}

// Receive a raw (unparsed) message.
func (c Controller) ReceiveRaw(msg *string) {
	websocket.Message.Receive(c.ws, msg)
}
