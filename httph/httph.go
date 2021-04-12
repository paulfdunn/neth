package httph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPBodyUnmarshal(w http.ResponseWriter, r *http.Request, obj interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("reading body, error:%v", err)
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return fmt.Errorf("unmarshal body, error:%v", err)
	}

	return nil
}
