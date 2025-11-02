package service

import (
	"net/http"

	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
)


type EvenementService struct {
    repo interfaces.EvenementRepository
}


func (service EvenementService) CreateNewEvenement(req *http.Request) (*models.CreationEvenementResponse,error) {
    creationReq, httpErr := CreationEvenementMapper(req)
    if httpErr != nil {
        return nil, httpErr
    }
    evenement_id, err := service.repo.CreateNewEvenement(*creationReq)
    if err != nil {
        return nil, err
    }
    return &models.CreationEvenementResponse{
        Message: "Creaction success",
        ID: evenement_id.String(),
    }, nil
}


func NewEvenementService (repo interfaces.EvenementRepository) interfaces.EvenementService {
    return EvenementService{repo: repo}
}