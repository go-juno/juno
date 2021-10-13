package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if !t.IsZero() {
		dateString := t.Format("2006-01-02 15:04:05")
		return json.Marshal(dateString)
	} else {
		return json.Marshal(nil)
	}
}

func (t *LocalTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	return
}

func (t LocalTime) Value() (driver.Value, error) {

	return t.Format("2006-01-02 15:04:05"), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil && ns.String != "")
	return err
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

type NullUint struct {
	Uint  uint
	Valid bool
}

func (ns NullUint) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Uint)
}

func (ns *NullUint) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.Uint)
	ns.Valid = (err == nil || ns.Uint == 0)
	return err
}

func (ns NullUint) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Uint, nil
}

type LocalDate struct {
	time.Time
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	if !t.IsZero() {
		dateString := t.Format("2006-01-02")
		return json.Marshal(dateString)
	} else {
		return json.Marshal(nil)
	}
}

func (t *LocalDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse("2006-01-02", s)
	return
}

func (t LocalDate) Value() (driver.Value, error) {
	return t.Format("2006-01-02"), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDate{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Model struct {
	Id        uint           `json:"id" gorm:"primary_key"`
	CreatedAt LocalTime      `json:"created_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);comment:创建时间"`
	UpdatedAt LocalTime      `json:"updated_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" `
}
