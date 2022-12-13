package port

import (
	"context"

	"github.com/puizeabix/appstack-service/internal/domain"
)

type AppStackRepository interface {
	Create(ctx context.Context, s *domain.AppStack) (*domain.AppStack, error)
	Get(ctx context.Context, id string) (*domain.AppStack, error)
	List(ctx context.Context) ([]domain.AppStack, error)
}
