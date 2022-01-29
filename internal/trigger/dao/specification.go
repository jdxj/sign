package dao

type Specification struct {
	SpecID int64 `gorm:"primaryKey"`
	Spec   string
}

func (Specification) TableName() string {
	return TableName
}

const (
	TableName = "specification"
)
