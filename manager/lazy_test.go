package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLazyCache(t *testing.T) {
	impl := NewLazyCache()
	// our test number
	dummyInt := 42
	assert := assert.New(t)
	// get our first object and assert that it's an int
	first := impl.GetObject()
	assert.IsType(&dummyInt, first)
	// assign it a value so we can check on the next time
	*first = dummyInt
	impl.FreeObject(first)
	firstAgain := impl.GetObject()
	// this test checks that the memory address of the pointers are exactly the same, thus reusing the old object
	assert.True(first == firstAgain)
	// the obj we just got from the cache still has the value set previously
	assert.Equal(&dummyInt, firstAgain)
	// test a new object, should be 0 as that's the int's default value
	second := impl.GetObject()
	assert.Equal(0, *second)
}
