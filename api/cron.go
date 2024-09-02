package handler

import (
	"fmt"
	"net/http"
	"time"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

    formatted := now.Format("2006-01-02 15:04:05")

    // Print the formatted date and time
    fmt.Println("Current date and time:", formatted)
  fmt.Fprintf(w, "")
}