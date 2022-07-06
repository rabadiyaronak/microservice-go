package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	protos "github.com/rabadiyaronak/microservice-go/currency-service/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.GetRateRequest) (*protos.GetRateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.GetRateResponse{Rate: 0.55}, nil
}
