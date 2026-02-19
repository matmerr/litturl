package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetStatusHandler(t *testing.T) {
	SetServerStatus("test ready", true)

	r := mux.NewRouter()
	r.Handle("/api/status", GetStatus).Methods("GET")

	req := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var s serverStatus
	if err := json.NewDecoder(w.Body).Decode(&s); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if s.Comment != "test ready" {
		t.Errorf("unexpected comment: %q", s.Comment)
	}
	if !s.Ready {
		t.Error("expected Ready=true")
	}
}

func TestMuxRouteVariableExtraction(t *testing.T) {
	// Verify that gorilla/mux correctly extracts path variables.
	var captured string
	r := mux.NewRouter()
	r.HandleFunc("/{target}", func(w http.ResponseWriter, r *http.Request) {
		captured = mux.Vars(r)["target"]
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	req := httptest.NewRequest(http.MethodGet, "/testkey", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if captured != "testkey" {
		t.Errorf("expected captured variable %q, got %q", "testkey", captured)
	}
}
