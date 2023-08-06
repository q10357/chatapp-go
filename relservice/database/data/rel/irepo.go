package rel

type IRelRepo[T any] interface {
	CreateRel(*T) *T
	GetAllRels() []*T
	GetRelById(uint) (*T, error)
	UpdateRel(uint, *T) (*T, error)
	AcceptRel(uint, uint) (*T, error)
	DeleteRelById(uint) (bool, error)
	GetRelsByUserId(uint) ([]*T, error)
}
