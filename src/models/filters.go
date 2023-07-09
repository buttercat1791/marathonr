package models

import (
	"time"
)

type Filters struct {
	Ids     [][32]byte `json:"ids"`
	Authors [][32]byte `json:"authors"`
	Kinds   []int8     `json:"kinds"`
	E       [][32]byte `json:"#e"`
	P       [][32]byte `json:"#p"`
	Since   time.Time  `json:"since"`
	Until   time.Time  `json:"until"`
	Limit   int        `json:"limit"`
}
