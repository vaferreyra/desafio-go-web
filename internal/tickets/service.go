package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(context.Context, string) (int, error)
	AverageDestination(context.Context, string) (float32, error)
}

type service struct {
	rp Repository
}

func NewService(rp Repository) *service {
	return &service{rp}
}

func (s *service) GetTotalTickets(ctx context.Context, dst string) (int, error) {
	tickets, err := s.rp.GetTicketByDestination(ctx, dst)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, dst string) (avg float32, err error) {
	totaltickets, err := s.rp.GetAll(ctx)
	if err != nil {
		return 0, nil
	}

	ticketsForDestination, err := s.rp.GetTicketByDestination(ctx, dst)
	if err != nil {
		return 0, nil
	}

	result := (len(ticketsForDestination) / 100) * len(totaltickets)
	return float32(result), nil
}
