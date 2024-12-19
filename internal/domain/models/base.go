package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
TimestampUUID represents a database model for storing base fields of every model instance, with
primary key 'ID' of type uuid.UUID()
*/
type TimestampUUID struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

/*
TimestampUUID represents a database model for storing base fields of every model instance, with
primary key 'ID' of type uint
*/
type Timestamp struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:nano" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

/*
BeforeCreate is a GORM callback method for the TimestampUUID struct that is automatically invoked
before creating a new database record. It ensures that the ID field is populated with a new UUID
if its value is set to the zero UUID (uuid.Nil) before the creation process. This function is
designed to be used with GORM's hooks to customize behavior before certain database operations.
Parameters:
  - t: a pointer to the Timestamp instance.
  - tx: a GORM database transaction.

Returns an error, if any.
*/
func (t *TimestampUUID) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return nil
}
