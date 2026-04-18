package domain

import "time"

// Segment represents one transcribed subtitle chunk.
type Segment struct {
	Index int           `json:"index"`
	Start time.Duration `json:"start"`
	End   time.Duration `json:"end"`
	Text  string        `json:"text"`
}
