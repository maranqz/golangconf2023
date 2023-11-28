package handler

import (
	"net/http"
	"strconv"
	"time"

	"ddd/example/app/domain/ad"
)

type archive struct {
	adRepo ad.AdRepo
}

func (a *archive) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("idStr")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	adID, err := ad.NewAdID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	ad, err := a.adRepo.GetByID(ctx, adID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = ad.Archive(time.Now())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = a.adRepo.Save(ctx, ad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
