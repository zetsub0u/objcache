package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectStore(t *testing.T) {
	impl := NewObjectStore()
	// our test number
	dummyInt := 42
	assert := assert.New(t)
	// store should be empty
	assert.Len(impl.objs, 0)
	// get our first object and assert that it's an int
	first := impl.GetObject()
	assert.IsType(&dummyInt, first)
	// store should have 0 elements still, as we created but gave it away
	assert.Len(impl.objs, 0)
	// free it and check how many we have now
	impl.FreeObject(first)
	assert.Len(impl.objs, 1)
}
