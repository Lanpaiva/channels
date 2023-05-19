package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Acesso a pagina %d vezes", number)))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
	})
	http.ListenAndServe(":8080", nil)
}
