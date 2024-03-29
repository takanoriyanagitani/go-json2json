package main

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	f2j "github.com/takanoriyanagitani/go-json2json/flat"
)

var rdr io.Reader = os.Stdin
var wtr io.Writer = os.Stdout

var dec *json.Decoder = json.NewDecoder(rdr)
var enc *json.Encoder = json.NewEncoder(wtr)

var keys []string = strings.Split(os.Getenv("ENV_KEYS"), ",")

var cvk f2j.ConvKeys = f2j.ConvKeys{Keys: keys}
var cnv f2j.Converter = cvk.AsIf()

var ij2 f2j.IoJSON2JSON = f2j.IoJSON2JSON{
	Decoder:   dec,
	Converter: cnv,
	Encoder:   enc,
}

func mustNil(e error) {
	if nil != e {
		panic(e)
	}
}

func main() {
	mustNil(ij2.ConvertAll(context.Background()))
}
