package repository

import (
	model "Achievement/backend/app/format/domain"
)

// FormatRepository is ..
type FormatRepository interface {
	Create(data *model.ReceiveFormatData) error
}
