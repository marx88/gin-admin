package orm

import (
	"bytes"
	"database/sql"
	"fmt"
	"time"
)

// Time time类型
type Time struct {
	sql.NullTime
}

// MarshalJSON 重写json序列化
func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.NullTime.Valid {
		return []byte("null"), nil
	}
	return []byte(t.NullTime.Time.Format(`"2006-01-02 15:04:05"`)), nil
}

// UnMarshalJSON 重写json反序列化
func (t *Time) UnMarshalJSON(b []byte) (err error) {
	value, err := time.Parse("2006-01-02 15:04:05", string(bytes.Trim(b, "\"")))
	if err == nil {
		err = t.Scan(value)
	}
	return
}

// String string类型
type String struct {
	sql.NullString
}

// MarshalJSON 重写json序列化
func (t *String) MarshalJSON() ([]byte, error) {
	if !t.NullString.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.NullString.String)), nil
}

// UnmarshalJSON 重写json反序列化
func (t *String) UnmarshalJSON(b []byte) (err error) {
	return t.Scan(string(bytes.Trim(b, "\"")))
}
