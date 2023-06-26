package models

type Department struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	ManagerID string    // 部门经理ID
	Manager   *Employee `gorm:"foreignKey:ManagerID"`
}

func (Department) TableName() string {
	return "department"
}
