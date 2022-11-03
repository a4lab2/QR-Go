package data

import "time"

type UrlData struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Url            string    `json:"url"`
	Downloadswitch bool      `json:"downloadswitch"`
}

// type PhoneData struct {
// 	ID        int64
// 	CreatedAt time.Time
// }

// type EmailData struct {
// 	ID        int64
// 	CreatedAt time.Time
// }

// type TextData struct {
// 	ID        int64
// 	CreatedAt time.Time
// }

// type QRData struct {
// 	ID        int64
// 	CreatedAt time.Time
// 	urldata   *UrlData
// 	phonedata *PhoneData
// 	emaildate *EmailData
// 	textdata  *TextData
// }
