package wishlist

import "github.com/google/uuid"

type Model struct {
	id           uuid.UUID
	characterId  uint32
	serialNumber uint32
}
