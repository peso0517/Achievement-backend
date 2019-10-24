package usecase

import (
	model "Achievement/backend/app/format/domain"
	dbinterface "Achievement/backend/app/format/repository"
	"log"
)

// FormatUsecase is ...
type FormatUsecase interface {
		Create(data *model.ReceiveFormatData) error
}

type formatUsecase struct {
 	db               dbinterface.FormatRepository
}

// NewFormatUsecase is ..
func NewFormatUsecase() FormatUsecase {
	return &formatUsecase{
		db:               dbinterface.NewFormatMysql(),
		}
	}
// Create is ..
func (fu *formatUsecase) Create(data *model.ReceiveFormatData) error{
	log.Println(data, "usecase--------data----------")
		err := fu.db.Create(data)
		return err
}