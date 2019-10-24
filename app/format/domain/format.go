package domain

type ReceiveFormatData struct {
		FormatId							int								`json:"format_id" gorm:"column:format_id, primary"`
		BookTitle						string					`json:"book_title" gorm:"column:book_title"`
		Summary								string					`json:"summary" gorm:"column:summary"`
		UserId									int								`json:"user_id" gorm:"column:user_id"`
}