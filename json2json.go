package json2json

import (
	"context"
)

type JSON2JSON interface {
	Convert(ctx context.Context, original []byte) (converted []byte, e error)
}
