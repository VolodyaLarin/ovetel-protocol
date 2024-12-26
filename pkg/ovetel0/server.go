package ovetel0

import (
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_uc"
	"github.com/go-bson/bson"
	"github.com/pkg/errors"
	"io"
	"log/slog"
	"net/http"
)

type Server struct {
	usecase ovetel0_uc.IServerUsecase
}

func NewServer(usecase ovetel0_uc.IServerUsecase) *Server {
	return &Server{usecase}
}

func (server *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/vehicle", server.HandleVehicle)
	mux.HandleFunc("/config", server.HandleConfig)
	mux.HandleFunc("/changes", server.HandleChanges)
	mux.HandleFunc("/telemetry", server.HandleSendTelemetry)
}

func parseData[T interface{}](r *http.Request) (error, *ovetel0_if.Request[T]) {
	if r.Method != http.MethodPost {
		return errors.New("not supported method"), nil
	}

	id := r.Header.Get(ovetel0_if.AuthHeader)
	// @todo add validation

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "can't read body"), nil
	}

	var data T

	err = bson.Unmarshal(bytes, &data)
	if err != nil {
		return errors.Wrap(err, "can't parse body"), nil
	}

	return nil, &ovetel0_if.Request[T]{
		VehicleID: id,
		Data:      data,
	}
}

func sendData[T interface{}](w http.ResponseWriter, resp *T) {
	data, err := bson.Marshal(resp)
	if err != nil {
		slog.Error("can't marshal data", "error", errors.WithStack(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(data)
	if err != nil {
		slog.Error("can't write response", "error", errors.WithStack(err))
	}
}

func (server *Server) HandleVehicle(w http.ResponseWriter, r *http.Request) {
	err, req := parseData[ovetel0_if.Vehicle](r)
	if err != nil {
		slog.Error("HandleVehicle error", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := server.usecase.HandleVehicle(r.Context(), req)
	if err != nil {
		slog.Error("HandleVehicle error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	sendData(w, resp)
}

func (server *Server) HandleConfig(w http.ResponseWriter, r *http.Request) {
	err, req := parseData[ovetel0_if.DefaultReq](r)
	if err != nil {
		slog.Error("HandleConfig error", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := server.usecase.GetConfig(r.Context(), req)
	if err != nil {
		slog.Error("HandleConfig error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendData(w, resp)
}

func (server *Server) HandleChanges(w http.ResponseWriter, r *http.Request) {
	err, req := parseData[ovetel0_if.DefaultReq](r)
	if err != nil {
		slog.Error("HandleChanges error", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := server.usecase.GetChanges(r.Context(), req)
	if err != nil {
		slog.Error("HandleChanges error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendData(w, resp)
}

func (server *Server) HandleSendTelemetry(w http.ResponseWriter, r *http.Request) {
	err, req := parseData[ovetel0_if.SendTelemetryReq](r)
	if err != nil {
		slog.Error("HandleSendTelemetry error", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := server.usecase.HandleSendTelemetry(r.Context(), req)
	if err != nil {
		slog.Error("HandleSendTelemetry error", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	sendData(w, resp)
}
