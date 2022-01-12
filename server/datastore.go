/*
Copyright © 2022 William Cygan <wcygan.io@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package server

import (
	"errors"
	"fmt"
)

type Datastore struct {
	backingStore map[string]string
}

func (d *Datastore) Put(key, value string) error {
	d.backingStore[key] = value
	return nil
}

func (d *Datastore) Get(key string) (string, error) {
	value, ok := d.backingStore[key]
	if ok {
		return value, nil
	} else {
		return "", errKeyDoesNotExist(key)
	}
}

func errKeyDoesNotExist(key string) error {
	return errors.New(fmt.Sprintf("key `%s` does not exist in the datastore", key))
}
