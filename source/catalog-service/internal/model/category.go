package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	Id        string     `bun:"id,pk"`
	Name      string     `bun:"name,notnull"`
	CreatedAt *time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt *time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
