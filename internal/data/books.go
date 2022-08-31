package data

import "time"

type Book struct {
	ID          string      `json:"id"`
	IsRead      bool        `json:"is_read"`
	Author      string      `json:"author"`
	StartedAt   time.Time   `json:"started_at"`
	ReleaseYear int32       `json:"release_year"`
	Quotation   []Quotation `json:"queations"`
}

type Quotation struct {
	ID            string `json:"id"`
	BookID        string `json:"book_id"`
	QuotationNote string `json:"quation_note"`
}
