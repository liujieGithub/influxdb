// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: stream_reader.gen.go.tmpl

package reads

import (
	"fmt"

	"github.com/influxdata/platform/storage/reads/datatypes"
	"github.com/influxdata/platform/tsdb/cursors"
)

type floatCursorStreamReader struct {
	fr *frameReader
	a  cursors.FloatArray
}

func (c *floatCursorStreamReader) Close() {
	for c.fr.state == stateReadFloatPoints {
		c.readFrame()
	}
}

func (c *floatCursorStreamReader) Err() error { return c.fr.err }

func (c *floatCursorStreamReader) Next() *cursors.FloatArray {
	if c.fr.state == stateReadFloatPoints {
		c.readFrame()
	}
	return &c.a
}

func (c *floatCursorStreamReader) readFrame() {
	c.a.Timestamps = nil
	c.a.Values = nil

	if f := c.fr.peekFrame(); f != nil {
		switch ff := f.Data.(type) {
		case *datatypes.ReadResponse_Frame_FloatPoints:
			c.a.Timestamps = ff.FloatPoints.Timestamps
			c.a.Values = ff.FloatPoints.Values
			c.fr.nextFrame()

		case *datatypes.ReadResponse_Frame_Series:
			c.fr.state = stateReadSeries

		case *datatypes.ReadResponse_Frame_Group:
			c.fr.state = stateReadGroup

		default:
			c.fr.setErr(fmt.Errorf("floatCursorStreamReader: unexpected frame type %T", f.Data))
		}
	}
}

type integerCursorStreamReader struct {
	fr *frameReader
	a  cursors.IntegerArray
}

func (c *integerCursorStreamReader) Close() {
	for c.fr.state == stateReadIntegerPoints {
		c.readFrame()
	}
}

func (c *integerCursorStreamReader) Err() error { return c.fr.err }

func (c *integerCursorStreamReader) Next() *cursors.IntegerArray {
	if c.fr.state == stateReadIntegerPoints {
		c.readFrame()
	}
	return &c.a
}

func (c *integerCursorStreamReader) readFrame() {
	c.a.Timestamps = nil
	c.a.Values = nil

	if f := c.fr.peekFrame(); f != nil {
		switch ff := f.Data.(type) {
		case *datatypes.ReadResponse_Frame_IntegerPoints:
			c.a.Timestamps = ff.IntegerPoints.Timestamps
			c.a.Values = ff.IntegerPoints.Values
			c.fr.nextFrame()

		case *datatypes.ReadResponse_Frame_Series:
			c.fr.state = stateReadSeries

		case *datatypes.ReadResponse_Frame_Group:
			c.fr.state = stateReadGroup

		default:
			c.fr.setErr(fmt.Errorf("integerCursorStreamReader: unexpected frame type %T", f.Data))
		}
	}
}

type unsignedCursorStreamReader struct {
	fr *frameReader
	a  cursors.UnsignedArray
}

func (c *unsignedCursorStreamReader) Close() {
	for c.fr.state == stateReadUnsignedPoints {
		c.readFrame()
	}
}

func (c *unsignedCursorStreamReader) Err() error { return c.fr.err }

func (c *unsignedCursorStreamReader) Next() *cursors.UnsignedArray {
	if c.fr.state == stateReadUnsignedPoints {
		c.readFrame()
	}
	return &c.a
}

func (c *unsignedCursorStreamReader) readFrame() {
	c.a.Timestamps = nil
	c.a.Values = nil

	if f := c.fr.peekFrame(); f != nil {
		switch ff := f.Data.(type) {
		case *datatypes.ReadResponse_Frame_UnsignedPoints:
			c.a.Timestamps = ff.UnsignedPoints.Timestamps
			c.a.Values = ff.UnsignedPoints.Values
			c.fr.nextFrame()

		case *datatypes.ReadResponse_Frame_Series:
			c.fr.state = stateReadSeries

		case *datatypes.ReadResponse_Frame_Group:
			c.fr.state = stateReadGroup

		default:
			c.fr.setErr(fmt.Errorf("unsignedCursorStreamReader: unexpected frame type %T", f.Data))
		}
	}
}

type stringCursorStreamReader struct {
	fr *frameReader
	a  cursors.StringArray
}

func (c *stringCursorStreamReader) Close() {
	for c.fr.state == stateReadStringPoints {
		c.readFrame()
	}
}

func (c *stringCursorStreamReader) Err() error { return c.fr.err }

func (c *stringCursorStreamReader) Next() *cursors.StringArray {
	if c.fr.state == stateReadStringPoints {
		c.readFrame()
	}
	return &c.a
}

func (c *stringCursorStreamReader) readFrame() {
	c.a.Timestamps = nil
	c.a.Values = nil

	if f := c.fr.peekFrame(); f != nil {
		switch ff := f.Data.(type) {
		case *datatypes.ReadResponse_Frame_StringPoints:
			c.a.Timestamps = ff.StringPoints.Timestamps
			c.a.Values = ff.StringPoints.Values
			c.fr.nextFrame()

		case *datatypes.ReadResponse_Frame_Series:
			c.fr.state = stateReadSeries

		case *datatypes.ReadResponse_Frame_Group:
			c.fr.state = stateReadGroup

		default:
			c.fr.setErr(fmt.Errorf("stringCursorStreamReader: unexpected frame type %T", f.Data))
		}
	}
}

type booleanCursorStreamReader struct {
	fr *frameReader
	a  cursors.BooleanArray
}

func (c *booleanCursorStreamReader) Close() {
	for c.fr.state == stateReadBooleanPoints {
		c.readFrame()
	}
}

func (c *booleanCursorStreamReader) Err() error { return c.fr.err }

func (c *booleanCursorStreamReader) Next() *cursors.BooleanArray {
	if c.fr.state == stateReadBooleanPoints {
		c.readFrame()
	}
	return &c.a
}

func (c *booleanCursorStreamReader) readFrame() {
	c.a.Timestamps = nil
	c.a.Values = nil

	if f := c.fr.peekFrame(); f != nil {
		switch ff := f.Data.(type) {
		case *datatypes.ReadResponse_Frame_BooleanPoints:
			c.a.Timestamps = ff.BooleanPoints.Timestamps
			c.a.Values = ff.BooleanPoints.Values
			c.fr.nextFrame()

		case *datatypes.ReadResponse_Frame_Series:
			c.fr.state = stateReadSeries

		case *datatypes.ReadResponse_Frame_Group:
			c.fr.state = stateReadGroup

		default:
			c.fr.setErr(fmt.Errorf("booleanCursorStreamReader: unexpected frame type %T", f.Data))
		}
	}
}
