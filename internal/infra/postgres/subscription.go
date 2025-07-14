package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"test/internal/common"
	"test/internal/entity"
	"test/pkg/postgres"

	"github.com/jackc/pgx/v5"
)

const subsDefaultSliceCap = 50

type SubRepo struct {
	*postgres.Postgres
}

func NewSubscriptionRepository(pg *postgres.Postgres) *SubRepo {
	return &SubRepo{pg}
}

func (r *SubRepo) Create(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error) {
	const op = "infra.postgres.SubRepo.Create"

	err := r.Pool.QueryRow(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, begin_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;`, sub.ServiceName, sub.Price, sub.UserID, sub.BeginDate, sub.EndDate).Scan(&sub.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sub, nil
}

func (r *SubRepo) GetByID(ctx context.Context, id int) (*entity.Subscritpion, error) {
	const op = "infra.postgres.SubRepo.GetByID"

	row := r.Pool.QueryRow(ctx,
		"SELECT id, service_name, price, user_id, begin_date, end_date FROM subscriptions WHERE id=$1",
		id)

	var sub entity.Subscritpion
	err := row.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.BeginDate, &sub.EndDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, common.ErrSubscriptionNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sub, nil
}

func (r *SubRepo) Update(ctx context.Context, sub entity.Subscritpion) (*entity.Subscritpion, error) {
	const op = "infra.postgres.SubRepo.Update"

	_, err := r.Pool.Exec(ctx,
		`UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, begin_date=$4, end_date=$5 WHERE id=$6`,
		sub.ServiceName, sub.Price, sub.UserID, sub.BeginDate, sub.EndDate, sub.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sub, nil
}

func (r *SubRepo) Delete(ctx context.Context, id int) error {
	const op = "infra.postgres.SubRepo.Delete"

	_, err := r.Pool.Exec(ctx,
		"DELETE FROM subscriptions WHERE id=$1",
		id)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *SubRepo) GetAll(ctx context.Context) ([]entity.Subscritpion, error) {
	const op = "infra.postgres.SubRepo.GetAll"

	rows, err := r.Pool.Query(ctx,
		"SELECT id, service_name, price, user_id, begin_date, end_date FROM subscriptions")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	subs := make([]entity.Subscritpion, 0, subsDefaultSliceCap)
	for rows.Next() {
		var sub entity.Subscritpion

		err := rows.Scan(
			&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.BeginDate, &sub.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		subs = append(subs, sub)
	}

	return subs, nil
}

func (r *SubRepo) GetSum(ctx context.Context, f entity.Filter) (int, error) {
	const op = "infra.postgres.SubRepo.GetSum"

	filterCount := 4
	params := make([]interface{}, 0, filterCount)

	query := strings.Builder{}
	query.WriteString("SELECT SUM(price) FROM subscriptions WHERE begin_date>=$1 AND begin_date<=$2 ")
	params = append(params, f.BeginDate)
	params = append(params, f.EndDate)

	if f.UserID != "" {
		params = append(params, f.UserID)
		query.WriteString(fmt.Sprintf("AND user_id = $%d ", len(params)))
	}

	if f.ServiceName != "" {
		params = append(params, f.ServiceName)
		query.WriteString(fmt.Sprintf("AND service_name LIKE $%d ", len(params)))
	}

	row := r.Pool.QueryRow(ctx, query.String(), params...)

	var sum *int
	err := row.Scan(&sum)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if sum == nil {
		return 0, nil
	}

	return *sum, nil
}
