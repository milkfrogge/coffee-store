package product

import (
	"context"
	"github.com/milkfrogge/coffee-store/internal/model"
	mock_repository "github.com/milkfrogge/coffee-store/internal/repository/mocks"
	repo_product "github.com/milkfrogge/coffee-store/internal/repository/product"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestService_GetAllCategories(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ret := []model.Category{{
		Id:   "x",
		Name: "xxxxx",
	}}

	m.EXPECT().FindAllCategories(ctx).Return(ret, nil).Times(1)

	uCase := NewService(m, log)

	cat, err := uCase.GetAllCategories(ctx)

	require.NoError(t, err)
	require.ElementsMatch(t, ret, cat)

}

func TestService_GetAllCategoriesError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ret := []model.Category{}

	m.EXPECT().FindAllCategories(ctx).Return(ret, repo_product.ErrDbIsDown).Times(1)

	uCase := NewService(m, log)

	cat, err := uCase.GetAllCategories(ctx)

	require.EqualError(t, err, repo_product.ErrDbIsDown.Error())
	require.ElementsMatch(t, ret, cat)

}

func TestService_GetAllProducts(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ret := []model.Product{{
		Id:            "0xff",
		Name:          "s",
		Description:   "s",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "0xff",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Now(),
	}}

	m.EXPECT().FindAllProducts(ctx).Return(ret, nil).Times(1)

	uCase := NewService(m, log)

	prod, err := uCase.GetAllProducts(ctx)

	require.NoError(t, err)
	require.ElementsMatch(t, ret, prod)

}

func TestService_GetAllProductsError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ret := []model.Product{}

	m.EXPECT().FindAllProducts(ctx).Return(ret, repo_product.ErrDbIsDown).Times(1)

	uCase := NewService(m, log)

	prod, err := uCase.GetAllProducts(ctx)

	require.EqualError(t, err, repo_product.ErrDbIsDown.Error())
	require.ElementsMatch(t, ret, prod)

}

func TestService_GetAllProductsByCategory(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	catId := "0xff"
	ret := []model.Product{{
		Id:            "0xff",
		Name:          "s",
		Description:   "s",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   catId,
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Now(),
	}}

	m.EXPECT().FindProductsByCategory(ctx, catId).Return(ret, nil).Times(1)

	uCase := NewService(m, log)

	prod, err := uCase.GetAllProductsByCategory(ctx, catId)

	require.NoError(t, err)
	require.ElementsMatch(t, ret, prod)

}

func TestService_GetAllProductsByCategoryError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	catId := []string{"0xff", "0xee"}
	ret := [][]model.Product{
		[]model.Product{},
		[]model.Product{},
	}

	expErrors := []error{repo_product.ErrDbIsDown, repo_product.ErrNoSuchCategory}

	for i := 0; i < len(catId); i++ {
		m.EXPECT().FindProductsByCategory(ctx, catId[i]).Return(ret[i], expErrors[i]).Times(1)

		uCase := NewService(m, log)

		prod, err := uCase.GetAllProductsByCategory(ctx, catId[i])

		require.EqualError(t, err, expErrors[i].Error())
		require.ElementsMatch(t, ret[i], prod)
	}

}

func TestService_GetSingleProduct(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	id := "0xff"
	ret := model.Product{
		Id:            id,
		Name:          "s",
		Description:   "s",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "0xff",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Now(),
	}

	m.EXPECT().FindOneProduct(ctx, id).Return(ret, nil).Times(1)

	uCase := NewService(m, log)

	retId, err := uCase.GetSingleProduct(ctx, id)

	require.NoError(t, err)
	require.Equal(t, ret, retId)

}

func TestService_GetSingleProductError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ret := model.Product{}
	id := "0xff"
	expErrors := []error{repo_product.ErrDbIsDown, repo_product.ErrNoSuchProduct}

	uCase := NewService(m, log)

	for i := 0; i < len(expErrors); i++ {

		m.EXPECT().FindOneProduct(ctx, id).Return(ret, expErrors[i]).Times(1)

		prod, err := uCase.GetSingleProduct(ctx, id)

		require.EqualError(t, err, expErrors[i].Error())
		require.ElementsMatch(t, ret, prod)
	}

}
