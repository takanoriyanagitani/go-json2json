package flat2j_test

import (
	"log"

	"context"
	"encoding/json"
	"os"

	"bytes"
	"fmt"

	fj "github.com/takanoriyanagitani/go-json2json/flat"
)

func ExampleConvKeys_Convert() {
	ck := fj.ConvKeys{Keys: []string{"seqno", "name"}}
	var cv fj.Converter = ck.AsIf()
	original := map[string]any{
		"seqno": 42,
		"name":  "JD",
		"phone": "01-2345-6789",
	}
	converted, e := cv.Convert(context.Background(), original)
	if nil != e {
		log.Fatal(e)
	}
	var dec *json.Encoder = json.NewEncoder(os.Stdout)
	e = dec.Encode(&converted)
	if nil != e {
		log.Fatal(e)
	}
	// Output: {"name":"JD","seqno":42}
}

func ExampleIoJSON2JSON_ConvertAll() {
	var buf bytes.Buffer
	var dec *json.Decoder = json.NewDecoder(&buf)
	var enc *json.Encoder = json.NewEncoder(os.Stdout)
	ck := fj.ConvKeys{Keys: []string{"seqno", "name"}}
	var cv fj.Converter = ck.AsIf()
	_, _ = fmt.Fprintf(&buf, `
		{"seqno":42,"name":"JD","phone":"01-2345-6789"}
		{"seqno":43,"name":"DJ","phone":"ab-cdef-0123"}
	`)
	ij2j := fj.IoJSON2JSON{
		Decoder:   dec,
		Converter: cv,
		Encoder:   enc,
	}
	e := ij2j.ConvertAll(context.Background())
	if nil != e {
		log.Fatal(e)
	}
	// Output:
	// {"name":"JD","seqno":42}
	// {"name":"DJ","seqno":43}
}
