package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseJson(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

type jsonRequestFunc func()

func Request(r *http.Request, w *http.ResponseWriter, method string, isJson bool, requestFunc jsonRequestFunc) {
	if isJson {
		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			(*w).WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
	}
	if r.Method != method {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
	}
	requestFunc()
}

func SetResponseException(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusBadRequest)
	_, err2 := (*w).Write([]byte(err.Error()))
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
}

func SetResponse(w *http.ResponseWriter, statusCode int, body []byte) {
	(*w).WriteHeader(statusCode)
	_, err2 := (*w).Write(body)
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
}
