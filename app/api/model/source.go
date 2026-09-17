package model

import "time"

type Source struct {
	ID          int64      `db:"id" json:"id"`
	URL         string     `db:"url" json:"url"`
	Title       *string    `db:"title" json:"title,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	PublishedAt *time.Time `db:"published_at" json:"published_at,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
