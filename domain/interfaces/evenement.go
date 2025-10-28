package interfaces

import (
	"context"

	"github.com/J2d6/reny_event/domain/models"
	"github.com/google/uuid"
)


type EvenementRepository interface {
	CreateNewEvenement(ctx context.Context, req models.CreationEvenementRequest,) (uuid.UUID, error)
}


var TypeEvenementIDMap = map[string]string{
    "Concert": "96021c38-1efd-4d02-ab6e-61dc3f154d68",
    "Conference": "fafb3782-22e6-465c-8660-3cbc77a9d6fd",
    "Seminaire": "e935ccd2-9be2-4e98-a7e0-52d165933588",
    "Foire": "1dff3ab8-f5f4-4ab7-866b-4ee7d6f4b78c",
	"Exposition": "5927389b-888f-4b72-be32-238f4da9f8d5",
}


var TypePlaceIDMap = map[string]string{
	"VIP":"27cc5f33-3e97-45b7-96a9-7861ce72ec0f",
	"Standard":"31894dbd-cc07-444a-8ca0-f64a77d313e2",
	"Premium":"bcfddb39-1689-4222-949f-5b788cae7bbe",
}
