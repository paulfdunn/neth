package httph

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/paulfdunn/osh/runtimeh"
)

func BodyUnmarshal(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return runtimeh.SourceInfoError("reading body", err)
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return runtimeh.SourceInfoError("unmarshal body", err)
	}

	return nil
}
