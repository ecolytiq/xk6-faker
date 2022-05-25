package faker

// MIT License
//
// Copyright (c) 2021 Iván Szkiba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/modules"
)

const envSEED = "XK6_FAKER_SEED"

func init() {
	modules.Register("k6/x/faker", New())
}

type (
	RootModule struct{}

	Fake struct {
		*Faker
	}
)

var (
	_ modules.Instance = &Fake{}
	_ modules.Module   = &RootModule{}
)

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Fake{newFaker(vu, seed())}
}

func (f *Fake) Exports() modules.Exports {
	return modules.Exports{
		Default: f,
		Named: map[string]interface{}{
			"Faker": f.constructor,
		},
	}
}

func seed() int64 {
	str := os.Getenv(envSEED)
	if str == "" {
		return 0
	}

	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logrus.Error(err)
		return 0
	}

	return n
}
