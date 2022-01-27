package dao

type Specification struct {
	SpecID int64 `gorm:"primaryKey"`
	Spec   string
}

const (
	TableName = "specification"
)
