package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PushRequest struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

type ReceivedBodyRequest struct {
	Title string  `json:"title"`
	Body  string  `json:"body"`
	Value float64 `json:"value"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//TO DO

	pubSubMsg := &PushRequest{}
	errDecode := json.NewDecoder(r.Body).Decode(pubSubMsg)
	if errDecode != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", errDecode), http.StatusBadRequest)
		return
	}

	body := ReceivedBodyRequest{}
	errUnmarshal := json.Unmarshal(pubSubMsg.Message.Data, &body)
	if errUnmarshal != nil {
		http.Error(w, fmt.Sprintf("Decode Data Error: %v. Wrong payload format: %v", errUnmarshal, body), http.StatusBadRequest)
		return
	}

	fmt.Printf("\n%+v\n", body)

	http.ResponseWriter(w).WriteHeader(http.StatusOK)
}
