package grom

import (
	"database/sql"
	"gorm.io/gorm"
)

type Roles []string

type Techer struct {
	gorm.Model
	Name         string `gorm:"column:techer_name"`
	Email        string
	Age          uint8 `gorm:"check:age > 20"`
	WorkingYears uint8
	Birthday     int64 `gorm:"serializer:unixtime;type:time"`
	StuNumber    sql.NullString
	Roles        Roles `gorm:"serializer:json"`
	JobInfo      Job   `gorm:"embedded;embeddedPrefix:job_"`
	JobInfo2     Job   `gorm:"type:bytes;serializer:gob"`
}

type Job struct {
	Title    string
	Location string
}
