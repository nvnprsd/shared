package dbentities

import "time"

type GORMWorkkards struct {
	Id               uint       `gorm:"column:id;type:bigserial"`
	Org_id           uint       `gorm:"column:org_id;type:int8"`
	Workkard_number  *string    `gorm:"column:workkard_number;type:varchar(150)"`
	Workkard_summary *string    `gorm:"column:workkard_summary;type:text"`
	Created_by       uint       `gorm:"column:created_by;type:int8"`
	Created_date     time.Time  `gorm:"column:created_date;type:timestamp"`
	Updated_by       *uint      `gorm:"column:updated_by;type:int8"`
	Updated_date     *time.Time `gorm:"column:updated_date;type:timestamp"`
	Sponsor          uint       `gorm:"column:sponsor;type:int8"`
	Workkard_title   *string    `gorm:"column:workkard_title;type:varchar(750)"`
	Template_id      uint       `gorm:"column:template_id;type:int8"`
	Carrier_id       *uint      `gorm:"column:carrier_id;type:int8"`
}

func (GORMWorkkards) TableName() string {
	return "workkards"
}
