package store_test

import (
	"fmt"
	"gostore/store"
	"testing"
)

const (
	testKey        = "k"
	testData       = "test data"
	testDataOther  = "test data other"
	nonExistingKey = "m"
)

func TestStore_Get(t *testing.T) {
	myStore := store.NewStore()

	t.Run("Test GET with an Empty Store", func(t *testing.T) {
		if _, err := myStore.Get(testKey); err != store.ErrKeyNotFound {
			t.Errorf("Expected error '%s', got '%s'", store.ErrKeyNotFound, err)
		}
	})

	t.Run("Test GET with one key", func(t *testing.T) {
		myStore.Put(testKey, testData)
		expected := testData
		if v, _ := myStore.Get(testKey); *v != expected {
			t.Errorf("Expected %s got %v", expected, *v)
		}
	})

	t.Run("Test GET with a key that does not exist", func(t *testing.T) {
		if _, err := myStore.Get(nonExistingKey); err != store.ErrKeyNotFound {
			t.Errorf("Expected error '%s', got '%s'", store.ErrKeyNotFound, err)
		}
	})

	myStore.Close()
}

func TestStore_Put(t *testing.T) {
	myStore := store.NewStore()

	t.Run("Test PUT with one key", func(t *testing.T) {
		myStore.Put(testKey, testData)
		expected := testData

		if v, _ := myStore.Get(testKey); *v != expected {
			t.Errorf("Expected %s got %v", expected, *v)
		}
	})

	t.Run("Test PUT with one key having overwritten the original value", func(t *testing.T) {
		expectedOther := testDataOther
		myStore.Put(testKey, testDataOther)

		if v, _ := myStore.Get(testKey); *v != expectedOther {
			t.Errorf("Expected %s got %v", expectedOther, *v)
		}
	})

	myStore.Close()
}

func TestStore_Delete(t *testing.T) {
	myStore := store.NewStore()

	t.Run("Test DELETE with one key", func(t *testing.T) {
		myStore.Put(testKey, testData)
		myStore.Delete(testKey)

		if _, err := myStore.Get(testKey); err != store.ErrKeyNotFound {
			t.Errorf("Expected error '%s', got '%s'", store.ErrKeyNotFound, err)
		}
	})

	myStore.Close()
}

func BenchmarkStore_Put(b *testing.B) {

	myStore := store.NewStore()

	b.Run("Benchmarking PUT on the store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			val := fmt.Sprint(i)
			myStore.Put(val, val)
		}
	})

	myStore.Close()
}

func BenchmarkStore_Get(b *testing.B) {

	myStore := store.NewStore()

	//Setup
	for i := 0; i < b.N; i++ {
		val := fmt.Sprint(i)
		myStore.Put(val, val)
	}

	b.Run("Benchmarking GET on the store", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			val := fmt.Sprint(i)
			_, _ = myStore.Get(val)
		}
	})

	myStore.Close()
}
