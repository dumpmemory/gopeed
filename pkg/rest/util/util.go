package util

import (
	"encoding/json"
	"github.com/GopeedLab/gopeed/pkg/rest/model"
	"net/http"
)

func ReadJson(w http.ResponseWriter, r *http.Request, v any) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		WriteJson(w, http.StatusInternalServerError, model.NewResultWithMsg(err.Error()))
		return false
	}
	return true
}

func WriteJsonOk(w http.ResponseWriter, v any) {
	if v == nil {
		WriteJson(w, http.StatusOK, model.NewResultWithData[any](nil))
	} else {
		WriteJson(w, http.StatusOK, v)
	}
}

func WriteJson(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
