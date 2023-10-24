package responsewriter

import (
	"encoding/json"
	"net/http"
)

// ResponseFormat stands for our default response in API
type ResponseFormat struct {
	Data     interface{} `json:"data,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
	Jsonapi  interface{} `json:"jsonapi,omitempty"`
	Links    interface{} `json:"links,omitempty"`
	Included interface{} `json:"included,omitempty"`
}

// ResponseOK is function to return json OK
func (rf *ResponseFormat) ResponseOK(code int, data interface{}, w http.ResponseWriter) {
	if w.Header().Get("Content-Type") != "" {
		w.WriteHeader(code)
		w.Write(data.([]byte))
		return
	}

	// default with json response
	rf.Data = data
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(rf)
	if err != nil {
		resErr := ResponseFormat{
			Errors: err,
		}
		jsonErr, _ := json.Marshal(resErr)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonErr)
		return
	}
	w.WriteHeader(code)
	w.Write(resp)
}

// ResponseNOK is function to return json NOK
func (rf *ResponseFormat) ResponseNOK(code int, errors interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(rf)
	if err != nil {
		resErr := ResponseFormat{
			Errors: err.Error(),
		}
		jsonErr, _ := json.Marshal(resErr)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonErr)
		return
	}

	w.WriteHeader(code)
	w.Write(resp)
}
