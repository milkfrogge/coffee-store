package product

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid/v5"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/milkfrogge/coffee-store/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type PostgresRepository struct {
	conn   *pgx.Conn
	log    *slog.Logger
	tracer trace.TracerProvider
}

func (r *PostgresRepository) CreateCategory(ctx context.Context, category model.CreateCategoryDTO) (string, error) {
	const op = "Product.Repo.Postgres.CreateCategory"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID

	err := r.conn.QueryRow(ctx, "INSERT INTO category (name) VALUES ($1) RETURNING id", category.Name).Scan(&id)
	if err != nil {
		return "", err
	}

	b, _ := id.UUIDValue()
	return fmt.Sprintf("%x", b.Bytes), nil
}

func (r *PostgresRepository) CreateProduct(ctx context.Context, product model.CreateProductDTO) (string, error) {
	const op = "Product.Repo.Postgres.CreateProduct"
	r.log.Debug(op)
	var id pgxuuid.UUID

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(product.CategoryId)
	categoryId := pgxuuid.UUID(temp)

	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(context.Background())

	// insert product
	err = tx.QueryRow(
		ctx,
		"INSERT INTO product (name,description,price,category, counter) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		product.Name,
		product.Description,
		product.Price,
		categoryId,
		product.Count,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	//insert pictures to this product
	for i := 0; i < len(product.Pics); i++ {
		_, err = tx.Exec(
			ctx,
			"INSERT INTO product_images (product_id, url) VALUES ($1,$2)",
			id,
			product.Pics[i],
		)

		if err != nil {
			return "", err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	b, _ := id.UUIDValue()
	return fmt.Sprintf("%x", b.Bytes), nil
}

func (r *PostgresRepository) FindAllCategories(ctx context.Context) ([]model.Category, error) {
	const op = "Product.Repo.Postgres.FindAllCategories"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID
	var category model.Category
	out := make([]model.Category, 0)

	rows, err := r.conn.Query(ctx, "SELECT name, id from category")
	if err != nil {
		return out, err
	}

	for rows.Next() {
		err = rows.Scan(&category.Name, &id)
		if err != nil {
			return []model.Category{}, err
		}

		b, _ := id.UUIDValue()
		category.Id = fmt.Sprintf("%x", b.Bytes)

		out = append(out, category)
	}

	return out, nil
}

func (r *PostgresRepository) FindAllProducts(ctx context.Context) ([]model.Product, error) {
	const op = "Product.Repo.Postgres.FindAllProducts"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID
	var categoryId pgxuuid.UUID
	var product model.Product
	var category model.Category
	out := make([]model.Product, 0)

	rows, err := r.conn.Query(
		ctx,
		"SELECT product.id, product.name, product.description, product.counter, product.price,product.created_at,  category.id, category.name, su.arr\nfrom product\n         join category on product.category = category.id\n         left join (select array_agg(url) as arr, product_images.product_id as id\n                    from product_images\n                             join product on product_images.product_id = product.id\n                    group by product_images.product_id) as su on product.id = su.id",
	)
	if err != nil {
		return out, err
	}

	for rows.Next() {
		links := make([]string, 0)
		err = rows.Scan(
			&id,
			&product.Name,
			&product.Description,
			&product.Count,
			&product.Price,
			&product.CreatedAt,
			&categoryId,
			&category.Name,
			&links,
		)
		if err != nil {
			return []model.Product{}, err
		}

		b, _ := categoryId.UUIDValue()
		category.Id = fmt.Sprintf("%x", b.Bytes)
		b, _ = id.UUIDValue()
		product.Id = fmt.Sprintf("%x", b.Bytes)
		product.Category = category
		product.Pics = links

		out = append(out, product)
	}

	return out, nil
}

func (r *PostgresRepository) FindProductsByCategory(ctx context.Context, categoryIdS string) ([]model.Product, error) {
	const op = "Product.Repo.Postgres.FindProductsByCategory"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID
	var categoryId pgxuuid.UUID
	var product model.Product
	var category model.Category
	out := make([]model.Product, 0)

	temp, _ := uuid.FromString(categoryIdS)
	idUUID := pgxuuid.UUID(temp)

	rows, err := r.conn.Query(
		ctx,
		"SELECT product.id, product.name, product.description, product.counter, product.price,product.created_at,  category.id, category.name, su.arr\nfrom product\n         join category on product.category = category.id\n         left join (select array_agg(url) as arr, product_images.product_id as id\n                    from product_images\n                             join product on product_images.product_id = product.id\n                    group by product_images.product_id) as su on product.id = su.id where product.category = $1", idUUID,
	)
	if err != nil {
		return out, err
	}

	for rows.Next() {
		links := make([]string, 0)
		err = rows.Scan(
			&id,
			&product.Name,
			&product.Description,
			&product.Count,
			&product.Price,
			&product.CreatedAt,
			&categoryId,
			&category.Name,
			&links,
		)
		if err != nil {
			return []model.Product{}, err
		}

		b, _ := categoryId.UUIDValue()
		category.Id = fmt.Sprintf("%x", b.Bytes)
		b, _ = id.UUIDValue()
		product.Id = fmt.Sprintf("%x", b.Bytes)
		product.Category = category
		product.Pics = links

		out = append(out, product)
	}

	return out, nil
}

func (r *PostgresRepository) FindOneProduct(ctx context.Context, idS string) (model.Product, error) {
	const op = "Product.Repo.Postgres.FindOneProduct"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID
	var categoryId pgxuuid.UUID
	var product model.Product
	var category model.Category

	temp, _ := uuid.FromString(idS)
	idUUID := pgxuuid.UUID(temp)

	err := r.conn.QueryRow(ctx, "SELECT product.id, product.name, product.description, product.counter, product.price,product.created_at,  category.id, category.name, su.arr\nfrom product\n         join category on product.category = category.id\n         left join (select array_agg(url) as arr, product_images.product_id as id\n                    from product_images\n                             join product on product_images.product_id = product.id\n                    group by product_images.product_id) as su on product.id = su.id where product.id = $1", idUUID).Scan(
		&id,
		&product.Name,
		&product.Description,
		&product.Count,
		&product.Price,
		&product.CreatedAt,
		&categoryId,
		&category.Name,
		&product.Pics,
	)
	if err != nil {
		return model.Product{}, err
	}

	b, _ := categoryId.UUIDValue()
	category.Id = fmt.Sprintf("%x", b.Bytes)
	b, _ = id.UUIDValue()
	product.Id = fmt.Sprintf("%x", b.Bytes)

	product.Category = category

	return product, nil
}

func (r *PostgresRepository) UpdateCategory(ctx context.Context, category model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) UpdateProduct(ctx context.Context, product model.Product) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) DeleteCategory(ctx context.Context, categoryId string) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) DeleteProduct(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewProductPostgresRepository(dsn string, log *slog.Logger) (*PostgresRepository, error) {

	ctx := context.Background()

	a, err := utils.WithAttempts(func() (any, error) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		conn, err := pgx.Connect(ctxTimeout, dsn)
		if err != nil {
			log.Error(fmt.Sprintf("Can`t establish connect to db: %s", err.Error()))
			return nil, err
		}

		tracerProvider := otel.GetTracerProvider()

		if err != nil {
			log.Error(fmt.Sprintf("Can`t init tracer: %s", err.Error()))
			return nil, err
		}

		return &PostgresRepository{
			conn:   conn,
			log:    log,
			tracer: tracerProvider,
		}, nil
	}, 5, log)

	if err != nil {
		return nil, err
	}

	return a.(*PostgresRepository), nil

}
