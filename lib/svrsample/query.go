package svrsample

import (
	"errors"
	"fmt"
	"github.com/multiplay/go-svrquery/lib/svrsample/protocol/prom"

	"github.com/multiplay/go-svrquery/lib/svrsample/common"

	"github.com/multiplay/go-svrquery/lib/svrsample/protocol/sqp"
)

var (
	// ErrProtoNotSupported returned when a protocol is not supported
	ErrProtoNotSupported = errors.New("protocol not supported")
)

// GetResponder gets the appropriate responder for the protocol provided
func GetResponder(proto string, state common.QueryState) (common.QueryResponder, error) {
	switch proto {
	case "sqp":
		return sqp.NewQueryResponder(state)
	case "prom":
		return prom.NewQueryResponder(state)
	}

	return nil, fmt.Errorf("%w: %s", ErrProtoNotSupported, proto)
}

func GetTransport(proto, address string) (Transport, error) {
	switch proto {
	case "sqp":
		return NewUDPTransport(address), nil
	case "prom":
		return NewHTTPTransport(address), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrProtoNotSupported, proto)
	}
}
