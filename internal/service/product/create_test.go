package product

import (
	"context"
	"errors"
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/milkfrogge/coffee-store/internal/repository"
	mock_repository "github.com/milkfrogge/coffee-store/internal/repository/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestService_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	in := model.CreateCategoryDTO{Name: "minimum"}
	ret := "0xdeadbeef"

	m.EXPECT().CreateCategory(ctx, in).Return(ret, nil)

	uCase := NewService(m, log)

	id, err := uCase.CreateCategory(ctx, in)

	require.NoError(t, err)
	require.Equal(t, ret, id)

}

func TestService_CreateCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	in := []model.CreateCategoryDTO{
		model.CreateCategoryDTO{Name: "1"},
		model.CreateCategoryDTO{Name: "norm"}}
	ret := ""
	expErrors := []error{ErrTooSmallCategoryName, errors.New("db is down")}

	uCase := NewService(m, log)

	m.EXPECT().CreateCategory(ctx, in[1]).Return(ret, errors.New("db is down")).Times(1)

	for i := 0; i < len(in); i++ {

		id, err := uCase.CreateCategory(ctx, in[i])

		require.EqualError(t, err, expErrors[i].Error())
		require.Equal(t, ret, id)
	}

}

func TestService_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	in := model.CreateProductDTO{
		Name:          "Latte",
		Description:   "Description of latte",
		Price:         100,
		Count:         100,
		CategoryId:    "deadbeefdeadbeefdeadbeefdeadbeef",
		BaristaNeeded: true,
		KitchenNeeded: false,
		Pics:          nil,
	}
	ret := "0xdeadbeef"

	m.EXPECT().CreateProduct(ctx, in).Return(ret, nil)

	uCase := NewService(m, log)

	id, err := uCase.CreateProduct(ctx, in)

	require.NoError(t, err)
	require.Equal(t, ret, id)
}

func TestService_CreateProductError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	in := []model.CreateProductDTO{
		model.CreateProductDTO{
			Name:          "Latte",
			Description:   "Description of latte",
			Price:         100,
			Count:         100,
			CategoryId:    "deadbeefdeadbeefdeadbeefdeadbeef",
			BaristaNeeded: true,
			KitchenNeeded: true,
			Pics:          nil,
		},
		model.CreateProductDTO{
			Name:          "Latte",
			Description:   "Description of latte",
			Price:         0,
			Count:         100,
			CategoryId:    "deadbeefdeadbeefdeadbeefdeadbeef",
			BaristaNeeded: true,
			KitchenNeeded: false,
			Pics:          nil,
		},
		model.CreateProductDTO{
			Name:          "Latte",
			Description:   "Description of latte",
			Price:         100,
			Count:         0,
			CategoryId:    "deadbeefdeadbeefdeadbeefdeadbeef",
			BaristaNeeded: true,
			KitchenNeeded: false,
			Pics:          nil,
		},
		model.CreateProductDTO{
			Name:          "Latte",
			Description:   "Description of latte",
			Price:         100,
			Count:         100,
			CategoryId:    "no such category",
			BaristaNeeded: true,
			KitchenNeeded: false,
			Pics:          nil,
		},
	}
	ret := ""

	expErrors := []error{ErrBaristaAndKitchenNeeded, ErrWrongPrice, ErrWrongCount, repository.ErrNoSuchCategory}

	m.EXPECT().CreateProduct(ctx, in[3]).Return(ret, repository.ErrNoSuchCategory).Times(1)

	uCase := NewService(m, log)

	for i := 0; i < len(in); i++ {
		id, err := uCase.CreateProduct(ctx, in[i])
		require.EqualError(t, err, expErrors[i].Error())
		require.Equal(t, ret, id)
	}

}
