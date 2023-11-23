package product

import (
	"context"
	"errors"
	mock_repository "github.com/milkfrogge/coffee-store/internal/repository/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

/*
 * @project coffee-store
 * @author nick
 */

func TestService_DeleteCategory(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	m.EXPECT().DeleteCategory(ctx, "0xdeadbeef").Return(nil).Times(1)

	svc := NewService(m, log)

	err := svc.DeleteCategory(ctx, "0xdeadbeef")
	require.NoError(t, err)
}

func TestService_DeleteCategoryError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	errExp := errors.New("db is down")

	m.EXPECT().DeleteCategory(ctx, "0xdeadbeef").Return(errExp).Times(1)

	svc := NewService(m, log)

	err := svc.DeleteCategory(ctx, "0xdeadbeef")
	require.EqualError(t, err, errExp.Error())
}

func TestService_DeleteProduct(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	m.EXPECT().DeleteProduct(ctx, "0xdeadbeef").Return(nil).Times(1)

	svc := NewService(m, log)

	err := svc.DeleteProduct(ctx, "0xdeadbeef")
	require.NoError(t, err)
}

func TestService_DeleteProductError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	errExp := errors.New("db is down")

	m.EXPECT().DeleteProduct(ctx, "0xdeadbeef").Return(errExp).Times(1)

	svc := NewService(m, log)

	err := svc.DeleteProduct(ctx, "0xdeadbeef")
	require.EqualError(t, err, errExp.Error())
}
