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

import "testing"

func TestDatastore(t *testing.T) {
	ds := Datastore{backingStore: make(map[string]string)}

	key, value := "foo", "bar"

	err := ds.Put(key, value)
	if err != nil {
		t.Errorf("failed to store key, value pair: (%s, %s)", key, value)
	}

	expected := value
	actual, err := ds.Get(key)
	if err != nil {
		t.Errorf("error retreiving value for key %s: %d", key, err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}

	key = "i am a key that doesn't exist!"
	_, err = ds.Get(key)
	if err == nil {
		t.Errorf("key %s should not exist in the datastore", key)
	}
}
