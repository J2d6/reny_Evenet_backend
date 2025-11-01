// package interfaces

// import (
// 	"context"

// 	"github.com/J2d6/reny_event/domain/models"
// 	"github.com/google/uuid"
// )

// type EvenementRepository interface {
//     CreateNewEvenement(ctx context.Context, req models.CreationEvenementRequest) (uuid.UUID, error)
//     GetEvenementByID(ctx context.Context, evenementID uuid.UUID) (*models.EvenementRow, error)
//     GetFichierContenu(ctx context.Context, evenementID uuid.UUID, fichierID uuid.UUID) (*models.FichierContenu, error)
//     CloseConn()
// }

// var TypeEvenementIDMap = map[string]string{
//     "Concert": "59754700-0a98-4fdc-86b8-a064198b8981",
//     "Conference": "5680e0c6-baec-43ec-937d-d4aba054f8ee",
// 	"Spectacle": "93f86c4a-5b9d-4282-b09b-56d0130e25fa",
//     "Seminaire": "929150c9-011d-4f70-99c9-e5a37ed31da4",
//     "Foire": "4a5c9c8b-d8a4-4f3f-8666-7325bb82c62c",
// 	"Exposition": "af70749f-3245-4169-a8ed-528d23ae9378",
// }


// var TypePlaceIDMap = map[string]string{
// 	"VIP":"b7e5c1e6-c1a2-4f26-8aa5-1ae631e3b9d3",
// 	"Standard":"8c316efe-da87-42fc-856c-c2b27db473e7",
// 	"Premium":"dd77a83c-172a-4e2a-9140-0ee0ceada351",
// 	"Economique":"7566d9a8-42bf-4d2c-b949-f7703b6cad26",
// }





