package dbentities

import "time"

type GORMActivityNotesFiles struct {
	Id                uint64    `gorm:"column:id;autoIncrement;type:bigserial"`
	NoteId            uint      `gorm:"column:note_id;type:int8"`
	FileName          string    `gorm:"column:file_name;type:varchar(765)"`
	FileType          *string   `gorm:"column:file_type;type:varchar(30)"`
	UploadedDate      time.Time `gorm:"column:uploaded_date;type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedBy         uint      `gorm:"column:updated_by;type:int8"`
	ParentId          *uint     `gorm:"column:parent_id;type:int8"`
	IsCopied          *bool     `gorm:"column:iscopied;type:bool"`
	FileGuid          *string   `gorm:"column:file_guid;type:varcahr(600)"`
	FileSize          *string   `gorm:"column:file_size;type:varcahr(30)"`
	Caption           *string   `gorm:"column:caption;type:varcahr(180)"`
	ProcessedFileName *string   `gorm:"column:processed_file_name;type:varcahr(765)"`
	ThumbnailImage    *string   `gorm:"column:thumbnail_image;type:varcahr(300)"`
}

func (GORMActivityNotesFiles) TableName() string {
	return "activity_notes_file"
}
