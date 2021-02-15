package manager

import "sync"

// ObjectStore is an implementation that is basically a slice with a mutex, when it gives it deletes, when it frees it creates
// it fulfills the requirements without knowing what it's to be used
type ObjectStore struct {
	sync.Mutex
	objs []int
}

func NewObjectStore() *ObjectStore {
	return &ObjectStore{
		objs: make([]int, 0),
	}
}

// GetObject retrieves any object from the store.
// Some decisions have been made here due to missing context and information:
// * "An object cannot be given away again, unless freed" does not specify whether we care to hold the object once give
//    or not, we are assuming we don't care about it until someone gives it back
// * Objects values are not unique
// * There doesn't seem to be a way to change the value of the objects in the api, we just delete on get and create on free
// * There's no specific function to create objects so i will assume that objects are created on demand when none are
//   available in the store.
func (s *ObjectStore) GetObject() *int {
	s.Lock()
	defer s.Unlock()
	res := 0
	// if there are no objects in the store, we return a new one
	if len(s.objs) == 0 {
		return &res
	}
	// take the last object
	res = s.objs[len(s.objs)-1]
	// re-slice to a smaller obj
	// todo: should make a copy to free up the old underlying array with bigger size
	s.objs = s.objs[:len(s.objs)-1]
	return &res

}

func (s *ObjectStore) FreeObject(obj *int) {
	s.Lock()
	defer s.Unlock()
	s.objs = append(s.objs, *obj)
}
