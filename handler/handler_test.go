package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/CezarGarrido/nasa-sonda/handler"
	"github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/stretchr/testify/assert"
)

func TestProbeCommandHandler(t *testing.T) {
	command := handler.Command{
		[]string{"GE", "M", "M"},
	}
	payload, _ := json.Marshal(command)
	buf := new(bytes.Buffer)
	buf.Write(payload)
	req, err := http.NewRequest("POST", "/api/probe/commands", buf)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	probe := sonda.NewProbe()
	probeHandler := handler.NewHandler(probe)
	handler := http.HandlerFunc(probeHandler.Commands)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	var expectedProbe sonda.Probe
	err = json.Unmarshal(rr.Body.Bytes(), &expectedProbe)
	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, expectedProbe.X, probe.X)
	assert.Equal(t, expectedProbe.Y, probe.Y)
}

func TestInvalidProbeCommandHandler(t *testing.T) {
	command := handler.Command{
		[]string{"GR", "M", "M", "S"},
	}
	payload, _ := json.Marshal(command)
	buf := new(bytes.Buffer)
	buf.Write(payload)
	req, err := http.NewRequest("POST", "/api/probe/commands", buf)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	probe := sonda.NewProbe()
	probeHandler := handler.NewHandler(probe)
	handler := http.HandlerFunc(probeHandler.Commands)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	var expectedProbe sonda.Probe
	err = json.Unmarshal(rr.Body.Bytes(), &expectedProbe)
	assert.Nil(t, err)
	assert.NotEqual(t, rr.Code, http.StatusOK)
	assert.Equal(t, expectedProbe.X, probe.X)
	assert.Equal(t, expectedProbe.Y, probe.Y)
}

func TestProbeRestartHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/probe/restart", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	probe := sonda.NewProbe()
	probeHandler := handler.NewHandler(probe)
	handler := http.HandlerFunc(probeHandler.RestartSondaPosition)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	var expectedProbe sonda.Probe
	err = json.Unmarshal(rr.Body.Bytes(), &expectedProbe)
	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, expectedProbe.X, probe.X)
	assert.Equal(t, expectedProbe.Y, probe.Y)
}

func TestProbeFindHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/probe", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	probe := sonda.NewProbe()
	probeHandler := handler.NewHandler(probe)
	handler := http.HandlerFunc(probeHandler.RestartSondaPosition)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	var expectedProbe sonda.Probe
	err = json.Unmarshal(rr.Body.Bytes(), &expectedProbe)
	assert.Nil(t, err)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, expectedProbe.X, probe.X)
	assert.Equal(t, expectedProbe.Y, probe.Y)
}
