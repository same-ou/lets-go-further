package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"strings"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q (at character %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			}
			return fmt.Errorf("body contains JSON value that cannot be stored in the target destination (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains EOF")
		
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown field %s", fieldName)
		
		case err.Error() == "http: request body too large":
			return errors.New("body must not be larger than 1MB")
		
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		default:
			return err
		}
	}
	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("invalid id parameter")
	}

	return id, nil
}

type envlope map[string]interface{}

func (app *application) writeJSON(rw http.ResponseWriter, status int, data envlope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		rw.Header()[key] = value
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(js)

	return nil
}
