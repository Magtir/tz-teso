package response

import (
	"net/http"
	"encoding/json"
	"log"
)

func SendResponse(w http.ResponseWriter, out interface{}, name string) {
	var outJson []byte
	outJson, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(outJson)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("out %s: %+v", name, out)
}
