package dbentities

import "time"

type GORMWorkkardFiles struct {
	Id                uint       `gorm:"column:id;type:bigserial"`
	WorkkardId        *uint      `gorm:"column:workkard_id;type:int8"`
	FileName          *string    `gorm:"column:file_name;type:varchar(765)"`
	FileType          *string    `gorm:"column:file_type;type:varhcar(30)"`
	UploadedDate      *time.Time `gorm:"column:uploaded_date;type:timestamp"`
	UploadedBy        *uint      `gorm:"column:uploaded_by;type:int8"`
	FileSize          *string    `gorm:"column:file_size;type:varchar(30)"`
	Type              *string    `gorm:"column:type;type:varchar(30)"`
	NoteId            *uint      `gorm:"column:note_id;type:int8"`
	UploadedByOrgId   uint       `gorm:"column:uploaded_by_org_id;type:int8"`
	FileGuid          *string    `gorm:"column:file_guid;type:varchar(300)"`
	Caption           *string    `gorm:"column:caption;type:varchar(180)"`
	ProcessedFileName *string    `gorm:"column:processed_file_name;type:varchar(765)"`
	ThumbnailImage    *string    `gorm:"column:thumbnail_image;type:varchar(300)"`
}

func (GORMWorkkardFiles) TableName() string {
	return "workkard_files"
}
