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

func TestService_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	tstamp := time.Now().UTC()
	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "test",
		Description:   "test",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "0xdeadbeef",
			Name: "test",
		},
		Pics:      nil,
		CreatedAt: tstamp,
	}
	svc := NewService(m, log)

	m.EXPECT().UpdateProduct(ctx, p).Return(nil).Times(1)

	err := svc.UpdateProduct(ctx, p)

	require.NoError(t, err)

}

func TestService_UpdateProductError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	tstamp := time.Now().UTC()
	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "test",
		Description:   "test",
		Price:         10,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "0xdeadbeef",
			Name: "test",
		},
		Pics:      nil,
		CreatedAt: tstamp,
	}
	svc := NewService(m, log)

	m.EXPECT().UpdateProduct(ctx, p).Return(repo_product.ErrNoSuchProduct).Times(1)

	err := svc.UpdateProduct(ctx, p)

	require.EqualError(t, err, repo_product.ErrNoSuchProduct.Error())

}

func TestService_UpdateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.Category{
		Id:   "0xdeadbeef",
		Name: "test",
	}
	svc := NewService(m, log)

	m.EXPECT().UpdateCategory(ctx, c).Return(nil).Times(1)

	err := svc.UpdateCategory(ctx, c)

	require.NoError(t, err)

}

func TestService_UpdateCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.Category{
		Id:   "0xdeadbeef",
		Name: "test",
	}
	svc := NewService(m, log)

	m.EXPECT().UpdateCategory(ctx, c).Return(repo_product.ErrNoSuchCategory).Times(1)

	err := svc.UpdateCategory(ctx, c)

	require.EqualError(t, err, repo_product.ErrNoSuchCategory.Error())

}

func TestService_AddCountToProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.CountToProductDTO{
		Id:    "0xdeadbeef",
		Count: 1,
	}

	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "",
		Description:   "",
		Price:         0,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Time{},
	}

	svc := NewService(m, log)

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, nil).Times(1)
	m.EXPECT().UpdateCountOfProduct(ctx, c.Id, c.Count+p.Count).Return(nil).Times(1)
	err := svc.AddCountToProduct(ctx, c)

	require.NoError(t, err)
}

func TestService_AddCountToProductError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.CountToProductDTO{
		Id:    "0xdeadbeef",
		Count: 1,
	}

	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "",
		Description:   "",
		Price:         0,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Time{},
	}

	svc := NewService(m, log)

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, repo_product.ErrNoSuchProduct).Times(1)
	//m.EXPECT().UpdateCountOfProduct(ctx, c.Id, c.Count+p.Count).Return(nil).Times(1)
	err := svc.AddCountToProduct(ctx, c)

	require.EqualError(t, err, repo_product.ErrNoSuchProduct.Error())

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, nil).Times(1)
	m.EXPECT().UpdateCountOfProduct(ctx, c.Id, c.Count+p.Count).Return(repo_product.ErrDbIsDown).Times(1)
	err = svc.AddCountToProduct(ctx, c)

	require.EqualError(t, err, repo_product.ErrDbIsDown.Error())
}

func TestService_SubtractCountToProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.CountToProductDTO{
		Id:    "0xdeadbeef",
		Count: 11,
	}

	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "",
		Description:   "",
		Price:         0,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Time{},
	}

	svc := NewService(m, log)

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, nil).Times(1)
	m.EXPECT().UpdateCountOfProduct(ctx, c.Id, p.Count-c.Count).Return(nil).Times(1)
	err := svc.SubtractCountToProduct(ctx, c)

	require.NoError(t, err)
}

func TestService_SubtractCountToProductError(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	c := model.CountToProductDTO{
		Id:    "0xdeadbeef",
		Count: 11,
	}

	p := model.Product{
		Id:            "0xdeadbeef",
		Name:          "",
		Description:   "",
		Price:         0,
		Count:         10,
		BaristaNeeded: false,
		KitchenNeeded: false,
		Category: model.Category{
			Id:   "",
			Name: "",
		},
		Pics:      nil,
		CreatedAt: time.Time{},
	}

	svc := NewService(m, log)

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, repo_product.ErrNoSuchProduct).Times(1)
	err := svc.SubtractCountToProduct(ctx, c)

	require.EqualError(t, err, repo_product.ErrNoSuchProduct.Error())

	m.EXPECT().FindOneProduct(ctx, c.Id).Return(p, nil).Times(1)
	m.EXPECT().UpdateCountOfProduct(ctx, c.Id, p.Count-c.Count).Return(repo_product.ErrDbIsDown).Times(1)
	err = svc.SubtractCountToProduct(ctx, c)

	require.EqualError(t, err, repo_product.ErrDbIsDown.Error())
}

func TestService_SubtractCountToProducts(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	svc := NewService(m, log)

	dto := []model.CountToProductDTO{
		{
			Id:    "0x001",
			Count: 10,
		},
		{
			Id:    "0x002",
			Count: 1,
		},
	}

	in := make(map[string]uint64)

	for i := 0; i < len(dto); i++ {
		m.EXPECT().FindOneProduct(ctx, dto[i].Id).Return(model.Product{Id: dto[i].Id, Count: dto[i].Count}, nil).Times(1)
		in[dto[i].Id] = 0
	}

	m.EXPECT().UpdateManyCountsOfProduct(ctx, in).Return(nil)

	err := svc.SubtractCountToProducts(ctx, dto)

	require.NoError(t, err)

}

func TestService_SubtractCountToProductsError(t *testing.T) {

	ctrl := gomock.NewController(t)

	m := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	svc := NewService(m, log)

	dto := []model.CountToProductDTO{
		{
			Id:    "0x001",
			Count: 10,
		},
		{
			Id:    "0x002",
			Count: 1,
		},
	}

	in := make(map[string]uint64)

	m.EXPECT().FindOneProduct(ctx, dto[0].Id).Return(model.Product{}, repo_product.ErrNoSuchProduct).Times(1)

	err := svc.SubtractCountToProducts(ctx, dto)

	require.EqualError(t, err, repo_product.ErrNoSuchProduct.Error())

	for i := 0; i < len(dto); i++ {
		m.EXPECT().FindOneProduct(ctx, dto[i].Id).Return(model.Product{Id: dto[i].Id, Count: dto[i].Count}, nil).Times(1)
		in[dto[i].Id] = 0
	}

	m.EXPECT().UpdateManyCountsOfProduct(ctx, in).Return(repo_product.ErrDbIsDown)

	err = svc.SubtractCountToProducts(ctx, dto)

	require.EqualError(t, err, repo_product.ErrDbIsDown.Error())
}
