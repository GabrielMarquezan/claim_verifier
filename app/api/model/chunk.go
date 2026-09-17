package model

import "time"

type Chunk struct {
	ID        int64      `db:"id" json:"id"`
	Content   string     `db:"content" json:"content"`
	SourceID  int64      `db:"source_id" json:"source_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
