// F Window System Protocol package
package fwsprotocol

import (
	"encoding/binary"

	"github.com/nsf/termbox-go"
)

// FWS Protocol socket constatnt
const FWS_SOCKET = "/tmp/fws_server.sock"

// Message type – alias for []uint8
type Msg []uint8

// Decodes message into Request interface implementations
func (msg *Msg) Decode() Request {
	header := Header(Msg(*msg)[0])
	payload := []uint8(*msg)[1:]
	switch header {
	case NEW:
		return &NewWindowRequest{
			int(binary.LittleEndian.Uint32(payload[0:4])),
			int(binary.LittleEndian.Uint32(payload[4:8])),
			int(binary.LittleEndian.Uint32(payload[8:12])),
			int(binary.LittleEndian.Uint32(payload[12:16])),
			int(binary.LittleEndian.Uint32(payload[16:20]))}
	case GET:
		return &GetRequest{
			ID(binary.LittleEndian.Uint32(payload[0:4])),
			int(binary.LittleEndian.Uint64(payload[4:12])),
			int(binary.LittleEndian.Uint64(payload[12:20]))}
	case REPLY_CREATION:
		return &ReplyCreationRequest{ID(binary.LittleEndian.Uint32(payload[0:4]))}
	case REPLY_GET:
		return &ReplyGetRequest{decodeCell(payload)}
	case EVENT:
		return &EventRequest{ID(binary.LittleEndian.Uint32(payload[0:4])),
			decodeEvent(payload[4:])}
	case DRAW:
		return &DrawRequest{ID(binary.LittleEndian.Uint32(payload[0:4])), int(binary.LittleEndian.Uint64(payload[4:12])), int(binary.LittleEndian.Uint64(payload[12:20])), decodeCell(payload[20:34])}
	case DRAW_FILL:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		width := int(binary.LittleEndian.Uint64(payload[4:12]))
		height := int(binary.LittleEndian.Uint64(payload[12:20]))
		img := make([][]Cell, width)
		bitmapPayload := payload[20:]
		index := 0
		for i := 0; i < width; i++ {
			img[i] = make([]Cell, height)
			for j := 0; j < height; j++ {
				img[i][j] = decodeCell(bitmapPayload[index : index+14])
				index += 14
			}
		}
		return &DrawFillRequest{id, width, height, img}
	case RENDER:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &RenderRequest{Id: id}
	case DELETE:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &DeleteRequest{Id: id}
	case RESIZE:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		width := int(binary.LittleEndian.Uint64(payload[4:12]))
		height := int(binary.LittleEndian.Uint64(payload[12:20]))
		return &ResizeRequest{Id: id, Width: width, Height: height}
	case MOVE:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		x := int(binary.LittleEndian.Uint64(payload[4:12]))
		y := int(binary.LittleEndian.Uint64(payload[12:20]))
		return &MoveRequest{Id: id, X: x, Y: y}
	case FOCUS:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &FocusRequest{Id: id}
	case UNFOCUS:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &UnfocusRequest{Id: id}
	case ACK:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &AckRequest{Id: id}
	case REPEAT:
		id := ID(binary.LittleEndian.Uint32(payload[0:4]))
		return &RepeatRequest{Id: id}
	default:
		return nil
	}
}

// Message class descriptor
type Header uint8

const (
	NEW            Header = iota // Message containing new window declaration
	GET                          // Message requesting cell information from specified window on specified location
	REPLY_CREATION               // Message containing ID of created window
	REPLY_GET                    // Message containing requested cell data
	EVENT                        // Message containing event in specified window
	DRAW                         // Message containing new cell data
	DRAW_FILL                    // Message containing large rectangle image
	RENDER                       // Message saying that image is finished and can be shown
	RESIZE                       // Message requesting window resize with specified parameters
	DELETE                       // Message requesting window deletion
	MOVE                         // Message specifying window shift
	FOCUS                        // Message requesting putting window on top
	UNFOCUS                      // Message stating that window in not active now
	ACK
	REPEAT
)

// Window ID type
// (4 bytes)
type ID uint32

// 16 bit analog to termbox Attribute
type Attr uint16

const (
	Bold Attr = 1 << (iota + 9)
	Blink
	Hidden
	Dim
	Underline
	Cursive
	Reverse
)

// aRGB cell color
// (4 bytes)
type Color struct {
	A, R, G, B uint8
}

func ColorFromTermboxAttr(attr termbox.Attribute) Color {
	r, g, b := termbox.AttributeToRGB(attr)
	return Color{A: 255, R: r, G: g, B: b}
}

func (c *Color) ToTerboxAttr() termbox.Attribute {
	return termbox.RGBToAttribute(c.R, c.G, c.B)
}

// aRGB overlaying operator
func (a *Color) Over(b Color) Color {
	alphaA := float32(a.A) / 255
	alphaB := float32(b.A) / 255
	alpha0 := alphaA + alphaB*(1-alphaA)

	Ra := float32(a.R) / 255
	Ga := float32(a.G) / 255
	Ba := float32(a.B) / 255

	Rb := float32(b.R) / 255
	Gb := float32(b.G) / 255
	Bb := float32(b.B) / 255

	nr := uint8(((Ra*alphaA + Rb*(1-alphaA)*alphaB) / alpha0) * 255)
	ng := uint8(((Ga*alphaA + Gb*(1-alphaA)*alphaB) / alpha0) * 255)
	nb := uint8(((Ba*alphaA + Bb*(1-alphaA)*alphaB) / alpha0) * 255)
	return Color{uint8(alpha0 * 255), nr, ng, nb}
}

func decodeColor(encoded []uint8) Color {
	return Color{encoded[0], encoded[1], encoded[2], encoded[3]}
}

// Color type binary encoder
func (c *Color) Encode() []uint8 {
	return []uint8{c.A, c.R, c.G, c.B}
}

// Text cell decriptor
// (14 bytes)
type Cell struct {
	Ch        rune
	Fg        Color
	Bg        Color
	Attribute Attr
}

func FromTermboxCell(cell termbox.Cell) Cell {
	return Cell{
		Ch:        cell.Ch,
		Fg:        ColorFromTermboxAttr(cell.Fg),
		Bg:        ColorFromTermboxAttr(cell.Bg),
		Attribute: 0,
	}
}

func (cell *Cell) ToTerboxCell() termbox.Cell {
	return termbox.Cell{
		Ch: cell.Ch,
		Fg: cell.Fg.ToTerboxAttr() | termbox.Attribute(cell.Attribute),
		Bg: cell.Bg.ToTerboxAttr()}
}

func (over *Cell) Over(underlying Cell) Cell {
	new_ch := over.Ch
	new_fg := over.Fg
	if over.Ch == rune(" "[0]) && over.Bg.A < 255 {
		new_ch = underlying.Ch
		new_fg = over.Bg.Over(underlying.Fg)
	}
	newCell := Cell{Ch: new_ch,
		Fg:        new_fg,
		Bg:        over.Bg.Over(underlying.Bg),
		Attribute: over.Attribute}
	return newCell
}

func decodeCell(encoded []uint8) Cell {
	ch := rune(binary.LittleEndian.Uint32(encoded[0:4]))
	fg := decodeColor(encoded[4:8])
	bg := decodeColor(encoded[8:12])
	attr := Attr(binary.LittleEndian.Uint16(encoded[12:14]))
	return Cell{ch, fg, bg, attr}
}

// Cell type binary encoder
func (c Cell) Encode() []uint8 {
	code := []uint8{}
	code = binary.LittleEndian.AppendUint32(code, uint32(c.Ch))
	code = append(code, c.Fg.Encode()...)
	code = append(code, c.Bg.Encode()...)
	code = binary.LittleEndian.AppendUint16(code, uint16(c.Attribute))
	return code
}

// General interface for all FWS messages
type Request interface {
	Encode() Msg // Message structure binary encoder
}

// New window request
// (16 bytes)
type NewWindowRequest struct {
	Pid    int // Requesting application Unix pid
	X      int // Global X position
	Y      int // Global Y position
	Width  int // Window width
	Height int // Window height
}

// New window request binary encoder
func (o *NewWindowRequest) Encode() Msg {
	msg := []uint8{uint8(NEW)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Pid))
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.X))
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Y))
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Width))
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Height))
	return Msg(msg)
}

// Cell data request
// (20 bytes)
type GetRequest struct {
	Id   ID  // Window ID
	X, Y int // Local X, Y coordinates
}

func (o *GetRequest) Encode() Msg {
	msg := []uint8{uint8(GET)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.X))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Y))
	return Msg(msg)
}

// New window ID reply
// (4 bytes)
type ReplyCreationRequest struct {
	Id ID
}

// New window creation reply binary encoder
func (o *ReplyCreationRequest) Encode() Msg {
	msg := []uint8{uint8(REPLY_CREATION)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

// Cell data reply
// (14 bytes)
type ReplyGetRequest struct {
	C Cell // Request-specified cell descriptor
}

func (o *ReplyGetRequest) Encode() Msg {
	msg := []uint8{uint8(REPLY_GET)}
	msg = append(msg, o.C.Encode()...)
	return Msg(msg)
}

// Event data message descriptor
// uses termbox Event with overwritten XY local coordinates
// (48 bytes)
type EventRequest struct {
	Id            ID // Layer ID to identify window
	termbox.Event    // Dispatched termbox event
}

func decodeEvent(encoded []uint8) termbox.Event {
	typ := encoded[0]
	mod := encoded[1]
	key := binary.LittleEndian.Uint16(encoded[2:4])
	ch := binary.LittleEndian.Uint32(encoded[4:8])
	width := binary.LittleEndian.Uint64(encoded[8:16])
	height := binary.LittleEndian.Uint64(encoded[16:24])
	mousex := binary.LittleEndian.Uint64(encoded[24:32])
	mousey := binary.LittleEndian.Uint64(encoded[32:40])
	n := binary.LittleEndian.Uint64(encoded[40:48])
	return termbox.Event{
		Type:   termbox.EventType(typ),
		Mod:    termbox.Modifier(mod),
		Key:    termbox.Key(key),
		Ch:     rune(ch),
		Width:  int(width),
		Height: int(height),
		Err:    nil,
		MouseX: int(mousex),
		MouseY: int(mousey),
		N:      int(n)}
}

// Event binary encoder
func (o *EventRequest) Encode() Msg {
	msg := []uint8{uint8(EVENT)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = append(msg, uint8(o.Type))
	msg = append(msg, uint8(o.Mod))
	msg = binary.LittleEndian.AppendUint16(msg, uint16(o.Key))
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Ch))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Width))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Height))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.MouseX))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.MouseY))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.N))
	return Msg(msg)
}

// Cell draw request
type DrawRequest struct {
	Id   ID
	X, Y int
	Cell Cell
}

// Cell draw request binary encoder
func (o *DrawRequest) Encode() Msg {
	msg := []uint8{uint8(DRAW)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.X))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Y))
	msg = append(msg, o.Cell.Encode()...)
	return Msg(msg)
}

// Filled rectangle request descriptor
type DrawFillRequest struct {
	Id            ID
	Width, Height int
	Img           [][]Cell
}

// Filled rectangle request descriptor binary encoder
func (o *DrawFillRequest) Encode() Msg {
	msg := []uint8{uint8(DRAW_FILL)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Width))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Height))
	for i := 0; i < o.Width; i++ {
		for j := 0; j < o.Height; j++ {
			msg = append(msg, o.Img[i][j].Encode()...)
		}
	}
	return msg
}

type RenderRequest struct {
	Id ID
}

func (o *RenderRequest) Encode() Msg {
	msg := []uint8{uint8(RENDER)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

type DeleteRequest struct {
	Id ID
}

func (o *DeleteRequest) Encode() Msg {
	msg := []uint8{uint8(DELETE)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

type ResizeRequest struct {
	Id            ID
	Width, Height int
}

func (o *ResizeRequest) Encode() Msg {
	msg := []uint8{uint8(RESIZE)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Width))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Height))
	return msg
}

type MoveRequest struct {
	Id   ID
	X, Y int
}

func (o *MoveRequest) Encode() Msg {
	msg := []uint8{uint8(MOVE)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.X))
	msg = binary.LittleEndian.AppendUint64(msg, uint64(o.Y))
	return msg
}

type FocusRequest struct {
	Id ID
}

func (o *FocusRequest) Encode() Msg {
	msg := []uint8{uint8(FOCUS)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

type UnfocusRequest struct {
	Id ID
}

func (o *UnfocusRequest) Encode() Msg {
	msg := []uint8{uint8(UNFOCUS)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

type AckRequest struct {
	Id ID
}

func (o *AckRequest) Encode() Msg {
	msg := []uint8{uint8(ACK)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}

type RepeatRequest struct {
	Id ID
}

func (o *RepeatRequest) Encode() Msg {
	msg := []uint8{uint8(REPEAT)}
	msg = binary.LittleEndian.AppendUint32(msg, uint32(o.Id))
	return msg
}
