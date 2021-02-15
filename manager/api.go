package manager

type ObjectMgr interface {
	GetObject() *int
	FreeObject(obj *int)
}
