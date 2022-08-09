package types

import (
	"database/sql/driver"
	"errors"
)

// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (bit BitBool) Value() (driver.Value, error) {
	if bit {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (bit *BitBool) Scan(src interface{}) error {
	buffer, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*bit = buffer[0] == 1
	return nil
}

const DatetimeFormat = "2006-01-02 03:04:05"
