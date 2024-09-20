package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/pkg/db/repository"
	"example.com/mymodule/internal/service"
	"github.com/gorilla/mux"
)

type Server struct {
	Service *service.Service // Ссылка на слой сервиса
	ServiceOrder *service.ServiceOrder
}

func NewServer(service *service.Service,serviceOrder *service.ServiceOrder) *Server {
	return &Server{Service: service,ServiceOrder: serviceOrder}
}

const queryParamKey = "key"

func CreateRouter(ctx context.Context, implemotation Server) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/PVZ", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			implemotation.CreatePVZ(ctx, w, r)
		case http.MethodPut:
			implemotation.ListPVZ(ctx,w,r)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/PVZ/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			implemotation.DeletePVZ(w, r)
		default:
			fmt.Println("error")
		}
	})



	router.HandleFunc("/Order", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			implemotation.CreateOrder(ctx, w, r)
		case http.MethodPut:
			implemotation.ListOrder(ctx,w,r)
		case http.MethodGet:
			implemotation.OrderStatus(ctx,w,r)
		case http.MethodPatch:
			implemotation.OrderSerch(ctx,w,r)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/Order/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			implemotation.DeleteOrder(w, r)
		default:
			fmt.Println("error")
		}
	})

	return router
}

func (s *Server) CreatePVZ(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm model.PVZInput
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	PVZRepo := &model.PVZInput{Title: unm.Title, Address: unm.Address, ContactInformation: unm.ContactInformation}

	id, err := s.Service.CreatePVZ(ctx, *PVZRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &model.PVZ{
		ID:                 id,
		Title:              PVZRepo.Title,
		Address:            PVZRepo.Address,
		ContactInformation: PVZRepo.ContactInformation,
	}
	articleJson, _ := json.Marshal(resp)
	w.Write(articleJson)
}

func (s *Server) OrderStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm model.OrderStatus
	if err = json.Unmarshal(body, &unm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OrderRepo := &model.OrderStatus{ID: unm.ID,Status: unm.Status}
	err = s.ServiceOrder.StatusOrder(ctx, *OrderRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) OrderSerch(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm model.OrderSerch
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OrderRepo := &model.OrderSerch{FullName: unm.FullName,OrderCode: unm.OrderCode}

	resp,err := s.ServiceOrder.SearchOrder(ctx, *OrderRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}else if resp!=0{
		articleJson, _ := json.Marshal(resp)
		w.Write(articleJson)
	}
}


func (s *Server) CreateOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm model.OrderInput
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OrderRepo := &model.OrderInput{FullName: unm.FullName, OrderCode: unm.OrderCode}
	id, err := s.ServiceOrder.CreateOrder(ctx, *OrderRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &model.Order{
		ID:                 id,
		FullName: OrderRepo.FullName,
		OrderCode: OrderRepo.OrderCode,
		Status: "заказ на складе",
	}
	articleJson, _ := json.Marshal(resp)
	w.Write(articleJson)
}


func (s *Server) ListPVZ(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	list, err := s.Service.ListPVZ(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	articleJson, _ := json.Marshal(list)
	w.Write(articleJson)
}

func (s *Server) ListOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {


	list, err := s.ServiceOrder.ListOrder(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	articleJson, _ := json.Marshal(list)
	w.Write(articleJson)
}

func (s *Server) DeletePVZ(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Service.DeletePVZ(r.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrorObjectNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
}

func (s *Server) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.ServiceOrder.DeleteOrder(r.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrorObjectNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	
}
