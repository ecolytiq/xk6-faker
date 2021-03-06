package faker

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
	"lukechampine.com/frand"
)

type Faker struct {
	*gofakeit.Faker
	vu modules.VU
}

func newFaker(vu modules.VU, seed int64) *Faker {
	src := frand.NewSource()

	if seed != 0 {
		src.Seed(seed)
	}

	return &Faker{
		Faker: gofakeit.NewCustom(src),
		vu:    vu,
	}
}

func (f *Faker) constructor(c goja.ConstructorCall) *goja.Object {
	seed := int64(0)
	rt := f.vu.Runtime()
	argSeed, ok := c.Argument(0).Export().(int64)
	// if !ok {
	// 	common.Throw(rt, fmt.Errorf("Faker constructor expects int64 as it's seed argument"))
	// }
	if ok {
		seed = argSeed
	}

	return rt.ToValue(newFaker(f.vu, seed)).ToObject(rt)
}

func (f *Faker) Ipv4Address() string {
	return f.IPv4Address()
}

func (f *Faker) Ipv6Address() string {
	return f.IPv6Address()
}

func (f *Faker) HttpStatusCodeSimple() int {
	return f.HTTPStatusCodeSimple()
}

func (f *Faker) HttpStatusCode() int {
	return f.HTTPStatusCode()
}

func (f *Faker) HttpMethod() string {
	return f.HTTPMethod()
}

func (f *Faker) Bs() string {
	return f.BS()
}

func (f *Faker) Uuid() string {
	return f.UUID()
}

func (f *Faker) RgbColor() []int {
	return f.RGBColor()
}

func (f *Faker) ImageJpeg(ctx context.Context, width int, height int) goja.ArrayBuffer {
	return f.vu.Runtime().NewArrayBuffer(f.Faker.ImageJpeg(width, height))
}

func (f *Faker) ImagePng(ctx context.Context, width int, height int) goja.ArrayBuffer {
	return f.vu.Runtime().NewArrayBuffer(f.Faker.ImagePng(width, height))
}
