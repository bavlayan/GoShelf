package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bavlayan/GoShelf/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new book")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	id = "a578f203-08b0-4ffd-a637-31a2f0ccb5ad"

	books := data.Book{
		ID:          id,
		IsRead:      true,
		Author:      "Mustafa Kemal ATATÜRK",
		StartedAt:   time.Now(),
		ReleaseYear: 1927,
		Quotation: []data.Quotation{
			{
				ID:            "1",
				BookID:        id,
				QuotationNote: "Temel ilke, Türk milletinin haysiyetli ve şerefli bir millet olarak yaşamasıdır.",
			},
			{
				ID:            "2",
				BookID:        id,
				QuotationNote: "Büyük ölülere matem gerekmez, fikirlerine bağlılık gerekir.",
			},
		},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem..", http.StatusInternalServerError)
	}
}
