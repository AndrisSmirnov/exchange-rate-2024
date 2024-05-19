package vo

import "github.com/google/uuid"

type UUID string

func (u UUID) String() string {
	return string(u)
}

func NewID() UUID {
	return UUID(uuid.NewString())
}
