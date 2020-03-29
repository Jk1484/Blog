package reply

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
)

// ContentType holds type
type ContentType string

// Supported content types
var (
	ContentTypeJSON ContentType = "application/json"
	ContentTypeTEXT ContentType = "text/plain"
	ContentTypeHTML ContentType = "text/html"
	ContentTypeXML  ContentType = "text/xml"
)

// Respond holds reponse format
type Respond struct {
	Payload interface{} `json:"payload"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

// JSON repsonse
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	Write(w, statusCode, ContentTypeJSON, data)
}

// Write writes to http.ResponseWriter
func Write(w http.ResponseWriter, statusCode int, contentType ContentType, data interface{}) {
	var (
		buf *bytes.Buffer
		b   []byte
		err error
	)

	switch contentType {

	case ContentTypeJSON:
		b, err = json.MarshalIndent(data, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case ContentTypeXML:
		b, err = xml.MarshalIndent(data, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case ContentTypeTEXT:
		b = []byte(fmt.Sprint(data))

	case ContentTypeHTML:
		b = []byte(template.HTMLEscaper(data))

	}

	buf = &bytes.Buffer{}
	_, err = buf.Write(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType.String())
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}

func (c ContentType) String() string {
	return string(c)
}
