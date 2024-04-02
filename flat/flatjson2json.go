package flat2j

import (
	"context"
	"encoding/json"
	"errors"
	"io"
)

type SimpleJSON2JSON struct {
	Parser
	Converter
	Serializer
}

func (f SimpleJSON2JSON) Convert(c context.Context, j []byte) ([]byte, error) {
	parsed, e := f.Parser.Parse(c, j)
	if nil != e {
		return nil, e
	}
	converted, e := f.Converter.Convert(c, parsed)
	if nil != e {
		return nil, e
	}
	return f.Serializer.Serialize(c, converted)
}

type IoJSON2JSON struct {
	*json.Decoder
	Converter
	*json.Encoder
}

func (i IoJSON2JSON) EncodeConverted(
	ctx context.Context,
	original map[string]any,
) error {
	converted, e := i.Converter.Convert(ctx, original)
	if nil != e {
		return e
	}

	return i.Encoder.Encode(converted)
}

func (i IoJSON2JSON) ConvertAll(ctx context.Context) error {
	m := make(map[string]any)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e := i.Decoder.Decode(&m)
		switch {
		case errors.Is(e, io.EOF):
			return nil
		case nil == e:
		default:
			return e
		}
		e = i.EncodeConverted(ctx, m)
		if nil != e {
			return e
		}
	}
}

type Parser interface {
	Parse(context.Context, []byte) (map[string]any, error)
}

type ParseFn func(context.Context, []byte) (map[string]any, error)

func (f ParseFn) Parse(c context.Context, j []byte) (map[string]any, error) {
	return f(c, j)
}
func (f ParseFn) AsIf() Parser { return f }

func ParserNew(f ParseFn) Parser { return f.AsIf() }

var parseFnDefault ParseFn = ParseFn(func(
	_ context.Context,
	jbytes []byte,
) (map[string]any, error) {
	m := make(map[string]any)
	e := json.Unmarshal(jbytes, &m)
	return m, e
})

var ParserDefault Parser = ParserNew(parseFnDefault)

// Converter converts an original json map to a converted json map.
type Converter interface {
	Convert(context.Context, map[string]any) (map[string]any, error)
}
type ConvFn func(context.Context, map[string]any) (map[string]any, error)

func (f ConvFn) Convert(
	ctx context.Context,
	original map[string]any,
) (map[string]any, error) {
	return f(ctx, original)
}

// ConvKeys implements Converter.
type ConvKeys struct{ Keys []string }

// Copies key/val pairs from the original map when key found in [ConvKeys].
func (k ConvKeys) Convert(
	_ context.Context,
	original map[string]any,
) (map[string]any, error) {
	neo := make(map[string]any)
	for _, key := range k.Keys {
		val, ok := original[key]
		if ok {
			neo[key] = val
		}
	}
	return neo, nil
}

func (k ConvKeys) AsIf() Converter { return k }

func (f ConvFn) AsIf() Converter { return f }

type Serializer interface {
	Serialize(context.Context, map[string]any) ([]byte, error)
}

type SerFn func(context.Context, map[string]any) ([]byte, error)

func (f SerFn) Serialize(c context.Context, m map[string]any) ([]byte, error) {
	return f(c, m)
}

func (f SerFn) AsIf() Serializer { return f }

func SerializerNew(f SerFn) Serializer { return f.AsIf() }

var serFnDefault SerFn = func(
	_ context.Context,
	m map[string]any,
) ([]byte, error) {
	return json.Marshal(m)
}

var SerializerDefault Serializer = SerializerNew(serFnDefault)
