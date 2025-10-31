package errors

import "fmt"


type ErreurValidation struct {
    Champ   string
    Message string
}

func (e *ErreurValidation) Error() string {
    return fmt.Sprintf("Erreur validation %s: %s", e.Champ, e.Message)
}



type ErreurAuthentification struct {
	Message string
}

func (e *ErreurAuthentification) Error() string {
	return e.Message
}