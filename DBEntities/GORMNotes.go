package dbentities

import "time"

type GORMNotes struct {
	Id               uint       `gorm:"column:id;type:bigserial"`
	ActivityId       *uint      `gorm:"column:activity_id;type:int8"`
	Description      string     `gorm:"column:description;type:text"`
	Resolution       *uint      `gorm:"column:resolution;type:int2"`
	CreatedBy        uint       `gorm:"column:created_by;type:int8"`
	CreatedDate      time.Time  `gorm:"column:created_date;type:timestamp"`
	UpdatedBy        *uint      `gorm:"column:updated_by;type:int8"`
	UpdatedDate      *time.Time `gorm:"column:updated_date;type:timestamp"`
	Type             string     `gorm:"column:Type;type:varchar"`
	IsActive         bool       `gorm:"column:is_active;type:bool"`
	IsDeleted        bool       `gorm:"column:is_deleted;type:bool"`
	ParentId         *uint      `gorm:"column:parent_id;type:int8"`
	OrganizationId   uint       `gorm:"column:organization_id;type:int8"`
	FromResolution   *string    `gorm:"column:from_resolution;type:varchar(150)"`
	ToResolution     *string    `gorm:"column:to_resolution;type:varchar(150)"`
	TemplateId       *uint      `gorm:"column:template_id;type:int8"`
	InternalToPublic bool       `gorm:"column:internal_to_public;type:bool"`
}

func (GORMNotes) TableName() string {
	return "notes"
}
