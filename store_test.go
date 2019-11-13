package ibc

import (
	"testing"

	"github.com/DE-labtory/ibc/spec"
)

func TestProvableStoreGet(t *testing.T) {
	tests := []struct {
		key      spec.Path
		expected spec.Value
	}{
		{
			key:      spec.Path("testKey"),
			expected: "testValue",
		},
		{
			key:      spec.Path("testKey2"),
			expected: "testValue2",
		},
		{
			key:      spec.Path("testKey3"),
			expected: nil,
		},
	}
	provableStore := ProvableStore{
		store: make(map[spec.Path]spec.Value),
	}

	provableStore.set(spec.Path("testKey"), "testValue")
	provableStore.set(spec.Path("testKey2"), "testValue2")

	for i, test := range tests {
		result := provableStore.get(test.key)
		if result != test.expected {
			t.Fatalf("test [%d] failed : value : %v, actual : %v", i, test.expected, result)
		}
	}
}

func TestProvableStoreDelete(t *testing.T) {
	tests := []struct {
		key   spec.Path
		value string
	}{
		{
			key:   spec.Path("testKey"),
			value: "testValue",
		},
		{
			key:   spec.Path("testKey2"),
			value: "testValue2",
		},
		{
			key:   spec.Path("testKey3"),
			value: "testValue3",
		},
	}
	provableStore := ProvableStore{
		store: make(map[spec.Path]spec.Value),
	}

	for i, test := range tests {
		provableStore.set(test.key, test.value)
		provableStore.delete(test.key)
		result := provableStore.get(test.key)
		if result != nil {
			t.Fatalf("test [%d] failed : value : nil, actual : %v", i, result)
		}
	}

	provableStore.delete(spec.Path("deletedKey"))
	if value := provableStore.get("deletedKey"); value != nil {
		t.Fatalf("test failed")
	}
}

func TestPrivateStoreGet(t *testing.T) {
	tests := []struct {
		key      spec.Path
		expected spec.Value
	}{
		{
			key:      spec.Path("testKey"),
			expected: "testValue",
		},
		{
			key:      spec.Path("testKey2"),
			expected: "testValue2",
		},
		{
			key:      spec.Path("testKey3"),
			expected: nil,
		},
	}
	privateStore := PrivateStore{
		store: make(map[spec.Path]spec.Value),
	}

	privateStore.set(spec.Path("testKey"), "testValue")
	privateStore.set(spec.Path("testKey2"), "testValue2")

	for i, test := range tests {
		result := privateStore.get(test.key)
		if result != test.expected {
			t.Fatalf("test [%d] failed : value : %v, actual : %v", i, test.expected, result)
		}
	}
}

func TestPrivateStoreDelete(t *testing.T) {
	tests := []struct {
		key   spec.Path
		value string
	}{
		{
			key:   spec.Path("testKey"),
			value: "testValue",
		},
		{
			key:   spec.Path("testKey2"),
			value: "testValue2",
		},
		{
			key:   spec.Path("testKey3"),
			value: "testValue3",
		},
	}
	privateStore := PrivateStore{
		store: make(map[spec.Path]spec.Value),
	}

	for i, test := range tests {
		privateStore.set(test.key, test.value)
		privateStore.delete(test.key)
		result := privateStore.get(test.key)
		if result != nil {
			t.Fatalf("test [%d] failed : value : nil, actual : %v", i, result)
		}
	}

	privateStore.delete(spec.Path("deletedKey"))
	if value := privateStore.get("deletedKey"); value != nil {
		t.Fatalf("test failed")
	}
}
