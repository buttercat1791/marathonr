package models

import (
	"time"
)

type Tag struct {
	Kind string
	Hex  [32]byte
	Url  string
}

type Event struct {
	Id        [32]byte  `json:"id"`
	Pubkey    [32]byte  `json:"pubkey"`
	CreatedAt time.Time `json:"created_at"`
	Kind      int8      `json:"kind"`
	Tags      []Tag     `json:"tags"`
	Content   string    `json:"content"`
	Sig       [64]byte  `json:"sig"`
}
