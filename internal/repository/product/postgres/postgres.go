package postgres

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid/v5"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/milkfrogge/coffee-store/internal/model"
	"github.com/milkfrogge/coffee-store/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type Repository struct {
	conn   *pgxpool.Pool
	log    *slog.Logger
	tracer trace.TracerProvider
}

func (r *Repository) UpdateCountOfProduct(ctx context.Context, id string, count uint64) error {
	const op = "Product.Repo.Postgres.UpdateCountOfProduct"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(id)
	productId := pgxuuid.UUID(temp)

	_, err := r.conn.Exec(ctx, "UPDATE product SET counter = ($1) WHERE id=($2);", count, productId)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil

}

func (r *Repository) UpdateManyCountsOfProduct(ctx context.Context, products map[string]uint64) error {
	const op = "Product.Repo.Postgres.UpdateManyCountsOfProduct"
	r.log.Debug(op)

	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	for k, v := range products {
		err := r.UpdateCountOfProduct(ctx, k, v)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil

}

func (r *Repository) CreateCategory(ctx context.Context, category model.CreateCategoryDTO) (string, error) {
	const op = "Product.Repo.Postgres.CreateCategory"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID

	err := r.conn.QueryRow(ctx, "INSERT INTO category (name) VALUES ($1) RETURNING id", category.Name).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}

	b, _ := id.UUIDValue()
	return fmt.Sprintf("%x", b.Bytes), nil
}

func (r *Repository) CreateProduct(ctx context.Context, product model.CreateProductDTO) (string, error) {
	const op = "Product.Repo.Postgres.CreateProduct"
	r.log.Debug(op)
	var id pgxuuid.UUID

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(product.CategoryId)
	categoryId := pgxuuid.UUID(temp)

	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			r.log.Error(err.Error())
		}
	}(tx, context.Background())

	// insert product
	err = tx.QueryRow(
		ctx,
		"INSERT INTO product (name,description,price,category, counter, barista_needed, kitchen_needed) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id",
		product.Name,
		product.Description,
		product.Price,
		categoryId,
		product.Count,
		product.BaristaNeeded,
		product.KitchenNeeded,
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
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
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}

	b, _ := id.UUIDValue()
	return fmt.Sprintf("%x", b.Bytes), nil
}

func (r *Repository) FindAllCategories(ctx context.Context) ([]model.Category, error) {
	const op = "Product.Repo.Postgres.FindAllCategories"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	var id pgxuuid.UUID
	var category model.Category
	out := make([]model.Category, 0)

	rows, err := r.conn.Query(ctx, "SELECT name, id from category")
	if err != nil {
		return out, fmt.Errorf("%s: %s", op, err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&category.Name, &id)
		if err != nil {
			return []model.Category{}, fmt.Errorf("%s: %s", op, err.Error())
		}

		b, _ := id.UUIDValue()
		category.Id = fmt.Sprintf("%x", b.Bytes)

		out = append(out, category)
	}

	return out, nil
}

func (r *Repository) FindAllProducts(ctx context.Context) ([]model.Product, error) {
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
		return out, fmt.Errorf("%s: %s", op, err.Error())
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
			return []model.Product{}, fmt.Errorf("%s: %s", op, err.Error())
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

func (r *Repository) FindProductsByCategory(ctx context.Context, categoryIdS string, limit, offset uint32) ([]model.Product, error) {
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

	q := ""

	if limit == 0 {
		q = "SELECT product.id, product.name, product.description, product.counter, product.price,product.created_at,  category.id, category.name, su.arr\nfrom product\n         join category on product.category = category.id\n         left join (select array_agg(url) as arr, product_images.product_id as id\n                    from product_images\n                             join product on product_images.product_id = product.id\n                    group by product_images.product_id) as su on product.id = su.id where product.category = $1  ORDER BY product.name OFFSET $2"
	} else {
		q = fmt.Sprintf("SELECT "+
			"product.id, "+
			"product.name, "+
			"product.description, "+
			"product.counter, "+
			"product.price, "+
			"product.created_at,  "+
			"category.id, "+
			"category.name, "+
			"su.arr\nfrom product\n         join category on product.category = category.id\n         left join (select array_agg(url) as arr, product_images.product_id as id\n                    from product_images\n                             join product on product_images.product_id = product.id\n                    group by product_images.product_id) as su on product.id = su.id where product.category = $1 ORDER BY product.name LIMIT %d OFFSET $2", limit)
	}

	fmt.Println(q)

	rows, err := r.conn.Query(
		ctx, q, idUUID, offset,
	)
	if err != nil {
		return out, fmt.Errorf("%s: %s", op, err.Error())
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
			return []model.Product{}, fmt.Errorf("%s: %s", op, err.Error())
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

func (r *Repository) FindOneProduct(ctx context.Context, idS string) (model.Product, error) {
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
		return model.Product{}, fmt.Errorf("%s: %s", op, err.Error())
	}

	b, _ := categoryId.UUIDValue()
	category.Id = fmt.Sprintf("%x", b.Bytes)
	b, _ = id.UUIDValue()
	product.Id = fmt.Sprintf("%x", b.Bytes)

	product.Category = category

	return product, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, category model.Category) error {
	const op = "Product.Repo.Postgres.UpdateCategory"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(category.Id)
	idUUID := pgxuuid.UUID(temp)

	_, err := r.conn.Exec(ctx, "UPDATE category SET name=$1 where id=$2", category.Name, idUUID)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil
}

func (r *Repository) UpdateProduct(ctx context.Context, product model.Product) error {
	const op = "Product.Repo.Postgres.UpdateProduct"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(product.Id)
	idUUID := pgxuuid.UUID(temp)

	//create transaction
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	//need to delete all pics associated with this id

	_, err = tx.Exec(ctx, "DELETE FROM product_images where product_id=$1", idUUID)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	//insert pictures to this product
	for i := 0; i < len(product.Pics); i++ {
		_, err = tx.Exec(
			ctx,
			"INSERT INTO product_images (product_id, url) VALUES ($1,$2)",
			idUUID,
			product.Pics[i],
		)

		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}
	}

	_, err = tx.Exec(
		ctx,
		"UPDATE product SET "+
			"name = $1,"+
			"description = $2,"+
			"price = $3,"+
			"barista_needed = $4,"+
			"kitchen_needed = $5,"+
			"counter = $6,"+
			"category = $7 "+
			"where id=$8",
		product.Name,
		product.Description,
		product.Price,
		product.BaristaNeeded,
		product.KitchenNeeded,
		product.Count,
		product.Category.Id,
		idUUID,
	)

	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil

}

func (r *Repository) DeleteCategory(ctx context.Context, categoryIdS string) error {
	const op = "Product.Repo.Postgres.DeleteCategory"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(categoryIdS)
	idUUID := pgxuuid.UUID(temp)

	_, err := r.conn.Exec(ctx, "Delete from category where id=$1", idUUID)

	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil

}

func (r *Repository) DeleteProduct(ctx context.Context, idS string) error {
	const op = "Product.Repo.Postgres.DeleteProduct"
	r.log.Debug(op)

	ctx, span := r.tracer.Tracer(op).Start(ctx, op)
	defer span.End()

	temp, _ := uuid.FromString(idS)
	idUUID := pgxuuid.UUID(temp)

	_, err := r.conn.Exec(ctx, "Delete from product where id=$1", idUUID)

	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}

	return nil
}

func NewProductPostgresRepository(dsn string, log *slog.Logger) (*Repository, error) {

	ctx := context.Background()

	a, err := utils.WithAttempts(func() (any, error) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		conn, err := pgxpool.New(ctxTimeout, dsn)
		if err != nil {
			log.Error(fmt.Sprintf("Can`t establish connect to db: %s", err.Error()))
			return nil, err
		}

		tracerProvider := otel.GetTracerProvider()

		if err != nil {
			log.Error(fmt.Sprintf("Can`t init tracer: %s", err.Error()))
			return nil, err
		}

		return &Repository{
			conn:   conn,
			log:    log,
			tracer: tracerProvider,
		}, nil
	}, 5, log)

	if err != nil {
		return nil, err
	}

	return a.(*Repository), nil

}
