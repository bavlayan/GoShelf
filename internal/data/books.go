package data

import "time"

type Book struct {
	ID          string
	IsRead      bool
	Author      string
	StartedAt   time.Time
	ReleaseYear int32
	Quotation   []Quotation
}

type Quotation struct {
	ID            string
	BookID        string
	QuotationNote string
}
