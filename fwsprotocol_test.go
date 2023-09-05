package fwsprotocol

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestNewWindowRequest(t *testing.T) {
	windowRequest := NewWindowRequest{
		Pid:       0,
		X:         1,
		Y:         2,
		Width:     10,
		Height:    20,
		LayerAttr: BOTTOM}
	encoded := windowRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *NewWindowRequest:
		if tdecode.Pid != windowRequest.Pid {
			t.Errorf("Pid field decoding failed: expected %d, got %d\n", tdecode.X, windowRequest.X)
		}
		if tdecode.X != windowRequest.X {
			t.Errorf("X field decoding failed: expected %d, got %d\n", tdecode.X, windowRequest.X)
		}
		if tdecode.Y != windowRequest.Y {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", tdecode.Y, windowRequest.Y)
		}
		if tdecode.Width != windowRequest.Width {
			t.Errorf("Width field decoding failed: expected %d, got %d\n", tdecode.Width, windowRequest.Width)
		}
		if tdecode.Height != windowRequest.Height {
			t.Errorf("Height field decoding failed: expected %d, got %d\n", tdecode.Height, windowRequest.Height)
		}
		if tdecode.LayerAttr != windowRequest.LayerAttr {
			t.Errorf("Layer attreibute field decoding failed: expected %d, got %d\n", tdecode.LayerAttr, windowRequest.LayerAttr)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestGetRequest(t *testing.T) {
	getRequest := GetRequest{
		Id: 12345,
		X:  10,
		Y:  20}
	encoded := getRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *GetRequest:
		if tdecode.Id != getRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", getRequest.Id, tdecode.Id)
		}
		if tdecode.X != getRequest.X {
			t.Errorf("X field decoding failed: expected %d, got %d\n", getRequest.X, tdecode.X)
		}
		if tdecode.Y != getRequest.Y {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", getRequest.Y, tdecode.Y)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestReplyCreationRequest(t *testing.T) {
	replyCreationRequest := ReplyCreationRequest{12345}
	encoded := replyCreationRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *ReplyCreationRequest:
		if tdecode.Id != replyCreationRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", replyCreationRequest.Id, tdecode.Id)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestReplyGetRequest(t *testing.T) {
	replyGetRequest := ReplyGetRequest{Cell{rune(456), Color{1, 2, 3, 4}, Color{5, 6, 7, 8}, Bold}}
	encoded := replyGetRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *ReplyGetRequest:
		if tdecode.C.Ch != replyGetRequest.C.Ch {
			t.Errorf("Ch field decoding failed: expected %d, got %d\n", replyGetRequest.C.Ch, tdecode.C.Ch)
		}
		if tdecode.C.Fg.A != replyGetRequest.C.Fg.A {
			t.Errorf("Fg.A field decoding failed: expected %d, got %d\n", replyGetRequest.C.Fg.A, tdecode.C.Fg.A)
		}
		if tdecode.C.Fg.R != replyGetRequest.C.Fg.R {
			t.Errorf("Fg.R field decoding failed: expected %d, got %d\n", replyGetRequest.C.Fg.R, tdecode.C.Fg.R)
		}
		if tdecode.C.Fg.G != replyGetRequest.C.Fg.G {
			t.Errorf("Fg.G field decoding failed: expected %d, got %d\n", replyGetRequest.C.Fg.G, tdecode.C.Fg.G)
		}
		if tdecode.C.Fg.B != replyGetRequest.C.Fg.B {
			t.Errorf("Fg.B field decoding failed: expected %d, got %d\n", replyGetRequest.C.Fg.B, tdecode.C.Fg.B)
		}
		if tdecode.C.Bg.A != replyGetRequest.C.Bg.A {
			t.Errorf("Bg.A field decoding failed: expected %d, got %d\n", replyGetRequest.C.Bg.A, tdecode.C.Bg.A)
		}
		if tdecode.C.Bg.R != replyGetRequest.C.Bg.R {
			t.Errorf("Bg.R field decoding failed: expected %d, got %d\n", replyGetRequest.C.Bg.R, tdecode.C.Bg.R)
		}
		if tdecode.C.Bg.G != replyGetRequest.C.Bg.G {
			t.Errorf("Bg.G field decoding failed: expected %d, got %d\n", replyGetRequest.C.Bg.G, tdecode.C.Bg.G)
		}
		if tdecode.C.Bg.B != replyGetRequest.C.Bg.B {
			t.Errorf("Bg.B field decoding failed: expected %d, got %d\n", replyGetRequest.C.Bg.B, tdecode.C.Bg.B)
		}
		if tdecode.C.Attribute != replyGetRequest.C.Attribute {
			t.Errorf("Attribute field decoding failed: expected %d, got %d\n", replyGetRequest.C.Attribute, tdecode.C.Attribute)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestEventRequest(t *testing.T) {
	eventRequest := EventRequest{
		Id: 123,
		Event: termbox.Event{
			Type:   termbox.EventKey,
			Mod:    termbox.ModAlt,
			Key:    termbox.KeyArrowDown,
			Ch:     rune(12),
			Width:  10,
			Height: 20,
			Err:    nil,
			MouseX: 30,
			MouseY: 40,
			N:      100}}
	encoded := eventRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *EventRequest:
		if tdecode.Id != eventRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", eventRequest.Type, tdecode.Type)
		}
		if tdecode.Type != eventRequest.Type {
			t.Errorf("Type field decoding failed: expected %d, got %d\n", eventRequest.Type, tdecode.Type)
		}
		if tdecode.Mod != eventRequest.Mod {
			t.Errorf("Mod field decoding failed: expected %d, got %d\n", eventRequest.Mod, tdecode.Mod)
		}
		if tdecode.Key != eventRequest.Key {
			t.Errorf("Key field decoding failed: expected %d, got %d\n", eventRequest.Key, tdecode.Key)
		}
		if tdecode.Ch != eventRequest.Ch {
			t.Errorf("Ch field decoding failed: expected %d, got %d\n", eventRequest.Ch, tdecode.Ch)
		}
		if tdecode.Width != eventRequest.Width {
			t.Errorf("Width field decoding failed: expected %d, got %d\n", eventRequest.Width, tdecode.Width)
		}
		if tdecode.Height != eventRequest.Height {
			t.Errorf("Height field decoding failed: expected %d, got %d\n", eventRequest.Height, tdecode.Height)
		}
		if tdecode.MouseX != eventRequest.MouseX {
			t.Errorf("MouseX field decoding failed: expected %d, got %d\n", eventRequest.MouseX, tdecode.MouseX)
		}
		if tdecode.MouseY != eventRequest.MouseY {
			t.Errorf("MouseY field decoding failed: expected %d, got %d\n", eventRequest.MouseY, tdecode.MouseY)
		}
		if tdecode.N != eventRequest.N {
			t.Errorf("N field decoding failed: expected %d, got %d\n", eventRequest.N, tdecode.N)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestDrawRequest(t *testing.T) {
	replyGetRequest := DrawRequest{123, 10, 20, Cell{rune(456), Color{1, 2, 3, 4}, Color{5, 6, 7, 8}, Bold}}
	encoded := replyGetRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *DrawRequest:
		if tdecode.Id != replyGetRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", replyGetRequest.Id, tdecode.Id)
		}
		if tdecode.X != replyGetRequest.X {
			t.Errorf("X field decoding failed: expected %d, got %d\n", replyGetRequest.X, tdecode.X)
		}
		if tdecode.Y != replyGetRequest.Y {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", replyGetRequest.Y, tdecode.Y)
		}
		if tdecode.Cell.Ch != replyGetRequest.Cell.Ch {
			t.Errorf("Ch field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Ch, tdecode.Cell.Ch)
		}
		if tdecode.Cell.Fg.A != replyGetRequest.Cell.Fg.A {
			t.Errorf("Fg.A field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Fg.A, tdecode.Cell.Fg.A)
		}
		if tdecode.Cell.Fg.R != replyGetRequest.Cell.Fg.R {
			t.Errorf("Fg.R field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Fg.R, tdecode.Cell.Fg.R)
		}
		if tdecode.Cell.Fg.G != replyGetRequest.Cell.Fg.G {
			t.Errorf("Fg.G field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Fg.G, tdecode.Cell.Fg.G)
		}
		if tdecode.Cell.Fg.B != replyGetRequest.Cell.Fg.B {
			t.Errorf("Fg.B field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Fg.B, tdecode.Cell.Fg.B)
		}
		if tdecode.Cell.Bg.A != replyGetRequest.Cell.Bg.A {
			t.Errorf("Bg.A field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Bg.A, tdecode.Cell.Bg.A)
		}
		if tdecode.Cell.Bg.R != replyGetRequest.Cell.Bg.R {
			t.Errorf("Bg.R field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Bg.R, tdecode.Cell.Bg.R)
		}
		if tdecode.Cell.Bg.G != replyGetRequest.Cell.Bg.G {
			t.Errorf("Bg.G field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Bg.G, tdecode.Cell.Bg.G)
		}
		if tdecode.Cell.Bg.B != replyGetRequest.Cell.Bg.B {
			t.Errorf("Bg.B field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Bg.B, tdecode.Cell.Bg.B)
		}
		if tdecode.Cell.Attribute != replyGetRequest.Cell.Attribute {
			t.Errorf("Attribute field decoding failed: expected %d, got %d\n", replyGetRequest.Cell.Attribute, tdecode.Cell.Attribute)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestDrawFillRequest(t *testing.T) {
	drawFillRequest := DrawFillRequest{}
	drawFillRequest.Id = 123
	drawFillRequest.Width = 10
	drawFillRequest.Height = 15
	drawFillRequest.Img = make([][]Cell, drawFillRequest.Width)
	for i := 0; i < drawFillRequest.Width; i++ {
		drawFillRequest.Img[i] = make([]Cell, drawFillRequest.Height)
		for j := 0; j < drawFillRequest.Height; j++ {
			drawFillRequest.Img[i][j] = Cell{
				Ch:        rune(456),
				Fg:        Color{10, 20, 30, 40},
				Bg:        Color{50, 60, 70, 80},
				Attribute: Bold}
		}
	}
	encoded := drawFillRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *DrawFillRequest:
		if tdecode.Id != drawFillRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", drawFillRequest.Id, tdecode.Id)
		}
		if tdecode.Width != drawFillRequest.Width {
			t.Errorf("X field decoding failed: expected %d, got %d\n", drawFillRequest.Width, tdecode.Width)
		}
		if tdecode.Height != drawFillRequest.Height {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", drawFillRequest.Height, tdecode.Height)
		}
		for i := 0; i < tdecode.Width; i++ {
			for j := 0; j < tdecode.Height; j++ {
				expected := drawFillRequest.Img[i][j]
				recieved := tdecode.Img[i][j]

				if expected.Ch != recieved.Ch {
					t.Errorf("[%d][%d] Ch field decoding failed: expected %d, got %d\n", i, j, expected.Ch, recieved.Ch)
				}
				if expected.Fg.A != recieved.Fg.A {
					t.Errorf("[%d][%d] Fg.A field decoding failed: expected %d, got %d\n", i, j, expected.Fg.A, recieved.Fg.A)
				}
				if expected.Fg.R != recieved.Fg.R {
					t.Errorf("[%d][%d] Fg.R field decoding failed: expected %d, got %d\n", i, j, expected.Fg.R, recieved.Fg.R)
				}
				if expected.Fg.G != recieved.Fg.G {
					t.Errorf("[%d][%d] Fg.G field decoding failed: expected %d, got %d\n", i, j, expected.Fg.G, recieved.Fg.G)
				}
				if expected.Fg.B != recieved.Fg.B {
					t.Errorf("[%d][%d] Fg.B field decoding failed: expected %d, got %d\n", i, j, expected.Fg.B, recieved.Fg.B)
				}
				if expected.Bg.A != recieved.Bg.A {
					t.Errorf("[%d][%d] Bg.A field decoding failed: expected %d, got %d\n", i, j, expected.Bg.A, recieved.Bg.A)
				}
				if expected.Bg.R != recieved.Bg.R {
					t.Errorf("[%d][%d] Bg.R field decoding failed: expected %d, got %d\n", i, j, expected.Bg.R, recieved.Bg.R)
				}
				if expected.Bg.G != recieved.Bg.G {
					t.Errorf("[%d][%d] Bg.G field decoding failed: expected %d, got %d\n", i, j, expected.Bg.G, recieved.Bg.G)
				}
				if expected.Bg.B != recieved.Bg.B {
					t.Errorf("[%d][%d] Bg.B field decoding failed: expected %d, got %d\n", i, j, expected.Bg.B, recieved.Bg.B)
				}
				if expected.Attribute != recieved.Attribute {
					t.Errorf("[%d][%d] Attribute field decoding failed: expected %d, got %d\n", i, j, expected.Attribute, recieved.Attribute)
				}
			}
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestRender(t *testing.T) {
	renderRequest := RenderRequest{Id: 1234}
	encoded := renderRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *RenderRequest:
		if tdecode.Id != renderRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", renderRequest.Id, tdecode.Id)
		}
	}
}

func TestDelete(t *testing.T) {
	renderRequest := DeleteRequest{Id: 1234}
	encoded := renderRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *RenderRequest:
		if tdecode.Id != renderRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", renderRequest.Id, tdecode.Id)
		}
	}
}

func TestResizeRequest(t *testing.T) {
	getRequest := ResizeRequest{
		Id:     12345,
		Width:  10,
		Height: 20}
	encoded := getRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *ResizeRequest:
		if tdecode.Id != getRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", getRequest.Id, tdecode.Id)
		}
		if tdecode.Width != getRequest.Width {
			t.Errorf("X field decoding failed: expected %d, got %d\n", getRequest.Width, tdecode.Width)
		}
		if tdecode.Height != getRequest.Height {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", getRequest.Height, tdecode.Height)
		}
	default:
		t.Errorf("Wrong decoded type: %v\n", tdecode)
	}
}

func TestMove(t *testing.T) {
	moveRequest := MoveRequest{Id: 1234, X: 56, Y: 78}
	encoded := moveRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *MoveRequest:
		if tdecode.Id != moveRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", moveRequest.Id, tdecode.Id)
		}
		if tdecode.X != moveRequest.X {
			t.Errorf("X field decoding failed: expected %d, got %d\n", moveRequest.X, tdecode.X)
		}
		if tdecode.Y != moveRequest.Y {
			t.Errorf("Y field decoding failed: expected %d, got %d\n", moveRequest.Y, tdecode.Y)
		}
	}
}

func TestFocus(t *testing.T) {
	topRequest := FocusRequest{Id: 1234}
	encoded := topRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *FocusRequest:
		if tdecode.Id != topRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", topRequest.Id, tdecode.Id)
		}
	}
}

func TestUnfocus(t *testing.T) {
	topRequest := UnfocusRequest{Id: 1234}
	encoded := topRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *UnfocusRequest:
		if tdecode.Id != topRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", topRequest.Id, tdecode.Id)
		}
	}
}

func TestAck(t *testing.T) {
	ackRequest := AckRequest{Id: 1234}
	encoded := ackRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *AckRequest:
		if tdecode.Id != ackRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", ackRequest.Id, tdecode.Id)
		}
	}
}

func TestRepeatRequest(t *testing.T) {
	repeatRequest := RepeatRequest{Id: 1234}
	encoded := repeatRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *RepeatRequest:
		if tdecode.Id != repeatRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", repeatRequest.Id, tdecode.Id)
		}
	}
}

func TestScreenRequest(t *testing.T) {
	screenRequest := ScreenRequest{Id: 1234}
	encoded := screenRequest.Encode()
	decoded := encoded.Decode()
	switch tdecode := decoded.(type) {
	case *ScreenRequest:
		if tdecode.Id != screenRequest.Id {
			t.Errorf("Id field decoding failed: expected %d, got %d\n", screenRequest.Id, tdecode.Id)
		}
	}
}

func TestReplyScreenRequest(t *testing.T) {
	screenRequest := ReplyScreenRequest{Width: 80, Height: 25, Mode: termbox.Output256}
	encoded := screenRequest.Encode()
	decoded := encoded.Decode()

	switch tdecode := decoded.(type) {
	case *ReplyScreenRequest:
		if tdecode.Width != screenRequest.Width {
			t.Errorf("Width field decoding failed: expected %d, got %d\n", screenRequest.Width, tdecode.Height)
		}
		if tdecode.Height != screenRequest.Height {
			t.Errorf("Height field decoding failed: expected %d, got %d\n", screenRequest.Height, tdecode.Height)
		}
		if tdecode.Mode != screenRequest.Mode {
			t.Errorf("Mode field decoding failed: expected %d, got %d\n", screenRequest.Mode, tdecode.Mode)
		}
	}
}
