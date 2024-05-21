package handlers

import (
	"fmt"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Foo"))
	return fmt.Errorf("AHHHHHH")
}
