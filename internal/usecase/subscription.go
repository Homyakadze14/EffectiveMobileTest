package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"test/internal/entity"
)

type SubStorage interface {
	Create(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error)
	GetByID(ctx context.Context, id int) (*entity.Subscritpion, error)
	Update(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]entity.Subscritpion, error)
	GetSum(ctx context.Context, f entity.Filter) (int, error)
}

type SubService struct {
	log *slog.Logger
	st  SubStorage
}

func NewSubscriptionService(log *slog.Logger, st SubStorage) *SubService {
	return &SubService{
		log: log,
		st:  st,
	}
}

func (s *SubService) Create(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error) {
	const op = "SubService.Create"
	log := s.log.With(slog.String("op", op))

	res, err := s.st.Create(ctx, sub)
	if err != nil {
		log.Error(fmt.Sprintf("failed to create subscription: %s", err))
		return nil, err
	}
	return res, nil
}

func (s *SubService) GetByID(ctx context.Context, id int) (*entity.Subscritpion, error) {
	const op = "SubService.GetByID"
	log := s.log.With(slog.String("op", op))

	res, err := s.st.GetByID(ctx, id)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get subscription by id: %s", err))
		return nil, err
	}
	return res, nil
}

func (s *SubService) Update(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error) {
	const op = "SubService.Update"
	log := s.log.With(slog.String("op", op))

	res, err := s.st.Update(ctx, sub)
	if err != nil {
		log.Error(fmt.Sprintf("failed to update subscription: %s", err))
		return nil, err
	}
	return res, nil
}

func (s *SubService) Delete(ctx context.Context, id int) error {
	const op = "SubService.Delete"
	log := s.log.With(slog.String("op", op))

	err := s.st.Delete(ctx, id)
	if err != nil {
		log.Error(fmt.Sprintf("failed to delete subscription: %s", err))
		return err
	}
	return nil
}

func (s *SubService) GetAll(ctx context.Context) ([]entity.Subscritpion, error) {
	const op = "SubService.Delete"
	log := s.log.With(slog.String("op", op))

	res, err := s.st.GetAll(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get all subscription: %s", err))
		return nil, err
	}
	return res, nil
}

func (s *SubService) GetSum(ctx context.Context, filter entity.Filter) (int, error) {
	const op = "SubService.GetSum"
	log := s.log.With(slog.String("op", op))

	res, err := s.st.GetSum(ctx, filter)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get sum of subscription prices: %s", err))
		return 0, err
	}
	return res, nil
}
