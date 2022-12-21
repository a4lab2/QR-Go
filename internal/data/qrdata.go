package data

import "time"

type QRData struct {
	DataKind     Qtype     `json:"datakind"`
	CreatedAt    time.Time `json:"created_at"`
	Qr_code_text string    `json:"text"`
}
