package dbentities

import "time"

// int  gorm:"column:lity

// int  gorm:"type:varchar(100)": Specifies the data type and size of the column in the database.

// int  gorm:"primary_key": Specifies the primary key field in the struct.

// int  gorm:"autoIncrement": Specifies that the field is an auto-incrementing primary key.

// int  gorm:"unique": Specifies that the field should have a unique constraint in the database.

// int  gorm:"not null": Specifies that the field cannot have a null value in the database.

// int  gorm:"default:'some_default_value'": Specifies the default value for the column in the database.

// int  gorm:"index": Specifies that an index should be created for the column in the database.

// int  gorm:"size:255": Specifies the size of the column in the database.

type GORMActivities struct {
	Id              uint       `gorm:"column:id;autoIncrement; not null"`
	SkillId         uint       `gorm:"column:skill_id;type:int4; not null"`
	PercentProgress uint       `gorm:"column:percent_progress;type:int4 not null"`
	CreatedBy       uint       `gorm:"column:created_by;type:int8 not null"`
	OwnerId         uint       `gorm:"column:owner_id;type:int8; not null"`
	AssigneeId      *uint      `gorm:"column:assignee_id;type:int8"`
	WorkkardId      uint       `gorm:"column:workkard_id;type:int8; not null"`
	OwnerOrgId      *uint      `gorm:"column:owner_org_id;type:int8"`
	AssigneeOrgId   *uint      `gorm:"column:assignee_org_id ;type:int8"`
	ParentId        *uint      `gorm:"column:parent_id;type:int8"`
	TemplateId      uint       `gorm:"column:template_id;type:int8; not null"`
	Description     *string    `gorm:"column:description;type:text;"`
	DueDate         time.Time  `gorm:"column:due_date;type:timestamp"`
	UpdatedDate     time.Time  `gorm:"column:updated_date;type:timestamp; default:CURRENT_TIMESTAMP;"`
	CreatedDate     time.Time  `gorm:"column:created_date;type:timestamp;"`
	CompletedDate   *time.Time `gorm:"column:completed_date;type:timestamp"`
	ClosedDate      *time.Time `gorm:"column:closed_date;type:timestamp"`
	DueDateTz       string     `gorm:"column:due_date_tz;type:varchar"`
	Status          string     `gorm:"column:status;type:varchar(36)"`
	Resolution      *string    `gorm:"column:resolution;type:varchar(60)"`
	Title           string     `gorm:"column:title;type:varchar(600)"`
}

func (GORMActivities) TableName() string {
	return "activities"
}
