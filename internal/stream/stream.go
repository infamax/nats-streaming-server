package stream

import (
	"encoding/json"
	"errors"
	"github.com/infamax/nats-streaming-server/internal/models"
	"github.com/infamax/nats-streaming-server/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type Stream struct {
	sc      stan.Conn
	service service.Service
}

func New(sc stan.Conn, service service.Service) (*Stream, error) {
	if sc == nil || service == nil {
		return nil, errors.New("cannot create stream")
	}
	return &Stream{
		sc:      sc,
		service: service,
	}, nil
}

func (s *Stream) Start(subject, durableName string) {
	s.sc.Subscribe(subject, func(msg *stan.Msg) {
		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("not valid data get service")
			_, _ = s.service.AddData(string(msg.Data))
			return
		}
		log.Println("add order in cache")
		err = s.service.AddModelCache(&order)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("add order in db")
		err = s.service.AddModelDB(&order)
		if err != nil {
			log.Println(err)
			return
		}
	}, stan.DurableName(durableName))
}
