package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        int            `gorm:"primary_key" json:"id"`
	CreatedAt UnixTime       `json:"createdAt"`
	UpdatedAt UnixTime       `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// For gorm time storage
type UnixTime struct {
	time.Time
}

// UnixTime implement gorm interfaces
func (t UnixTime) MarshalJSON() ([]byte, error) {
	microSec := t.UnixNano() / int64(time.Millisecond)
	return []byte(strconv.FormatInt(microSec, 10)), nil
}

func (t UnixTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *UnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UnixTime{Time: value}
		return nil
	}
	return fmt.Errorf("cannot convert %v to timestamp", v)
}