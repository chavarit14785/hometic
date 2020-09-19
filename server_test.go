package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/pair-device", nil)
	rec := httptest.NewRecorder()
	PairDeviceHandler(rec, req)
	if http.StatusOK != rec.Code {
		t.Error("expect 200 OK", rec.Code)
	}
	expected := `{"status":"active"}`
	if rec.Body.String() != expected {
		t.Errorf("expect %q but got %q\n", expected, rec.Body.String())
	}
}
