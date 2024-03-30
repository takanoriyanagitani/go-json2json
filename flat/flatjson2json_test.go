package flat2j_test

import (
	"log"

	"context"
	"encoding/json"
	"os"

	fj "github.com/takanoriyanagitani/go-json2json/flat"
)

func ExampleConvKeys_Convert() {
	ck := fj.ConvKeys{Keys: []string{"seqno", "name"}}
	original := map[string]any{
		"seqno": 42,
		"name":  "JD",
		"phone": "01-2345-6789",
	}
	converted, e := ck.Convert(context.Background(), original)
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
