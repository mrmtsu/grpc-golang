package domain

type TodoId string

func (id TodoId) String() string {
	return string(id)
}
