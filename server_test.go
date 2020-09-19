package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockInsertPairDevice struct {
}

func (i mockInsertPairDevice) insertPairDevice(pd PairDevice) {
}

func TestPairDeviceHandler(t *testing.T) {

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(
		PairDevice{
			DeviceID: 1234,
			UserID:   4433,
		})
	req := httptest.NewRequest(http.MethodPost, "/pair-device", payload)

	rec := httptest.NewRecorder()

	h := PairDeviceHandler(mockInsertPairDevice{})
	h.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expect 200 OK", rec.Code)
	}
	expected := `{"status":"active"}`
	if rec.Body.String() != expected {
		t.Errorf("expect %q but got %q\n", expected, rec.Body.String())
	}
}
