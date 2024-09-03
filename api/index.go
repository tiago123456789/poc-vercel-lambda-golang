package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

const URL = "https://typedwebhook.tools/webhook/cdd3bed5-1e63-4875-a00d-a95eadfe8eea"

func makePostRequest(data interface{}, wg *sync.WaitGroup) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err;
	}

	r, err := http.NewRequest("POST", URL, bytes.NewBuffer(body))
	if err != nil {
		return err;
	}

	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err;
	}
	
	defer func() {
		res.Body.Close()
		wg.Done()
	}()
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not implemented!"))
        return
	}

	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var wg sync.WaitGroup

	for i := 0; i < 500; i ++ {
		wg.Add(1)
		go makePostRequest(person, &wg)
	}

	wg.Wait()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}