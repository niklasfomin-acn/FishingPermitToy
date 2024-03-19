package storage

import "rest-backend/types"

type Storage interface {
	Get(int) *types.User
}
