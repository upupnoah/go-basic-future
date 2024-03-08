//go:build wireinject

package wire_demo

import (
	"context"
	"errors"

	"github.com/google/wire"
)

// Guide:
// https://github.com/google/wire/blob/main/docs/guide.md#advanced-features

type Foo struct {
	X int
}

// ProvideFoo returns a Foo.
func ProvideFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	X int
}

// ProvideBar returns a Bar: a negative Foo.
func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

type Baz struct {
	X int
}

// ProvideBaz returns a value if Bar is not zero.
func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
var MegaSet = wire.NewSet(SuperSet)

func initializeBaz(ctx context.Context) (Baz, error) {
	wire.Build(MegaSet)
	return Baz{}, nil
}
