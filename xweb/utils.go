package xweb

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// MatchesContentType validates the content type against the expected one
func MatchesContentType(contentType, expectedType string) bool {
	mimetype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Errorf("Error parsing media type: %s error: %v", contentType, err)
	}
	return err == nil && mimetype == expectedType
}

// CheckForJSON makes sure that the request's Content-Type is application/json.
func CheckForJSON(r *http.Request) error {
	ct := r.Header.Get("Content-Type")
	// No Content-Type header is ok as long as there's no Body
	if ct == "" {
		if r.Body == nil || r.ContentLength == 0 {
			return nil
		}
	}
	// Otherwise it better be json
	if MatchesContentType(ct, "application/json") {
		return nil
	}
	return fmt.Errorf("Content-Type specified (%s) must be 'application/json'", ct)
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

// ParseForm ensures the request form is parsed even with invalid content types.
// If we don't do this, POST method without Content-type (even with empty body) will fail.
func ParseForm(r *http.Request) error {
	if r == nil {
		return nil
	}
	if err := r.ParseForm(); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}

// ParseMultipartForm ensure the request form is parsed, even with invalid content types.
func ParseMultipartForm(r *http.Request) error {
	if err := r.ParseMultipartForm(4096); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}

// WriteError decodes a specific docker error and sends it in the response.
func WriteError(w http.ResponseWriter, err error) {
	if err == nil || w == nil {
		log.WithFields(log.Fields{"error": err, "writer": w}).Error("unexpected HTTP error handling")
		return
	}

	statusCode := http.StatusInternalServerError
	errMsg := err.Error()

	errStr := strings.ToLower(err.Error())

	for keyword, status := range map[string]int{
		"not found":             http.StatusNotFound,
		"no such":               http.StatusNotFound,
		"bad parameter":         http.StatusBadRequest,
		"conflict":              http.StatusConflict,
		"impossible":            http.StatusNotAcceptable,
		"wrong login/password":  http.StatusUnauthorized,
		"hasn't been activated": http.StatusForbidden,
	} {
		if strings.Contains(errStr, keyword) {
			statusCode = status
			break
		}
	}

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	http.Error(w, errMsg, statusCode)
}
