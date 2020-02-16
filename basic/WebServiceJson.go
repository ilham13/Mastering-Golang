package basic

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"01", "ilham", 25},
	student{"02", "sarah", 24},
	student{"03", "Al-Kautsar", 1},
}

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
