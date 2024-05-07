package dbentities

import (
	"time"
)

type GORMActivityFiles struct {
	id                  uint64     `gorm:"column:id;type:autoIncrement"`
	Iscopied            *bool      `gorm:"column:iscopied;type:bool"`
	Activity_id         *uint      `gorm:"column:activity_id;type:int8"`
	Uploaded_by         *uint      `gorm:"column:uploaded_by;type:int8"`
	Parent_id           *uint      `gorm:"column:parent_id;type:int8"`
	Uploaded_date       *time.Time `gorm:"column:uploaded_date;type:timestamp"`
	Caption             *string    `gorm:"column:caption;type:varchar(180)"`
	File_type           *string    `gorm:"column:file_type;type:varchar(30)"`
	File_size           *string    `gorm:"column:file_size;type:varchar(30)"`
	Type                *string    `gorm:"column:Type;type:varchar(30)"`
	File_guid           *string    `gorm:"column:file_guid;type:varchar(300)"`
	Thumbnail_image     *string    `gorm:"column:thumbnail_image;type:varchar(300)"`
	File_category       *string    `gorm:"column:file_category;type:varchar(60)"`
	File_name           *string    `gorm:"column:file_name;type:varchar(765)"`
	Processed_file_name *string    `gorm:"column:processed_file_name;type:varchar(765)"`
}

func (GORMActivityFiles) TableName() string {
	return "activity_files"
}
