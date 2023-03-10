package model

import "time"

type Message struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
