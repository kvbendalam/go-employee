package models

type Employee struct {
	ID           string   `gorm:"primaryKey"`
	Name         string   `gorm:"not null"`
	Salary       string   `gorm:"not null"`
	ProjectID    string   `gorm:"not null"`
	ProjectName  string   `gorm:"not null"`
	Billing      string   `gorm:"not null"`
	Duration     string   `gorm:"not null"`
	NoOfEmpWork  int      `gorm:"not null"`
	Technologies []string `gorm:"-"`
	// Manager      Manager   `gorm:"foreignKey:ManagerID"`
	// Projects     []Project `gorm:"foreignKey:ProjectID"`
}

type Project struct {
	ID          string `gorm:"primaryKey"`
	ProjectName string `gorm:"not null"`
	Billing     string `gorm:"not null"`
	Duration    string `gorm:"not null"`
	// Employees   []Employee `gorm:"foreignKey:ProjectID"`
}

type Manager struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	// Employees []Employee `gorm:"foreignKey:ManagerID"`
}
