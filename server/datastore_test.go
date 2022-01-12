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