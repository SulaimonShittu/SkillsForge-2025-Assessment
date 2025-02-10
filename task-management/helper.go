package task_management

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func WriteJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetIntParam(r *http.Request, param string, defaultValue int) int {
	val := r.URL.Query().Get(param)
	if val == "" {
		return defaultValue
	}
	if intValue, err := strconv.Atoi(val); err == nil {
		return intValue
	}
	return defaultValue
}

func GetTaskIDFromURL(r *http.Request) int {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return 0
	}
	if id, err := strconv.Atoi(parts[2]); err == nil {
		return id
	}
	return 0
}
