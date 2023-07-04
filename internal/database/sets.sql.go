// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: sets.sql

package database

import (
	"context"
	"time"
)

const createSet = `-- name: CreateSet :one
INSERT INTO sets (exercise, count, intensity, type, weight, workout_id, created_at, updated_at)
values($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, exercise, count, intensity, type, weight, workout_id, created_at, updated_at
`

type CreateSetParams struct {
	Exercise  string
	Count     int32
	Intensity Intensity
	Type      string
	Weight    string
	WorkoutID int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateSet(ctx context.Context, arg CreateSetParams) (Set, error) {
	row := q.db.QueryRowContext(ctx, createSet,
		arg.Exercise,
		arg.Count,
		arg.Intensity,
		arg.Type,
		arg.Weight,
		arg.WorkoutID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Set
	err := row.Scan(
		&i.ID,
		&i.Exercise,
		&i.Count,
		&i.Intensity,
		&i.Type,
		&i.Weight,
		&i.WorkoutID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSetById = `-- name: GetSetById :one
SELECT id, exercise, count, intensity, type, weight, workout_id, created_at, updated_at FROM sets where id = $1
`

func (q *Queries) GetSetById(ctx context.Context, id int32) (Set, error) {
	row := q.db.QueryRowContext(ctx, getSetById, id)
	var i Set
	err := row.Scan(
		&i.ID,
		&i.Exercise,
		&i.Count,
		&i.Intensity,
		&i.Type,
		&i.Weight,
		&i.WorkoutID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSetsByWorkoutIdAsc = `-- name: GetSetsByWorkoutIdAsc :many
SELECT id, exercise, count, intensity, type, weight, workout_id, created_at, updated_at FROM sets where workout_id = $1::int
ORDER BY $2::text ASC
LIMIT $4::int
OFFSET $3::int
`

type GetSetsByWorkoutIdAscParams struct {
	WorkoutID  int32
	OrderByCol string
	Offset     int32
	Limit      int32
}

func (q *Queries) GetSetsByWorkoutIdAsc(ctx context.Context, arg GetSetsByWorkoutIdAscParams) ([]Set, error) {
	rows, err := q.db.QueryContext(ctx, getSetsByWorkoutIdAsc,
		arg.WorkoutID,
		arg.OrderByCol,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Set
	for rows.Next() {
		var i Set
		if err := rows.Scan(
			&i.ID,
			&i.Exercise,
			&i.Count,
			&i.Intensity,
			&i.Type,
			&i.Weight,
			&i.WorkoutID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSetsByWorkoutIdDesc = `-- name: GetSetsByWorkoutIdDesc :many
SELECT id, exercise, count, intensity, type, weight, workout_id, created_at, updated_at FROM sets where workout_id = $1::int
ORDER BY $2::text DESC
LIMIT $4::int
OFFSET $3::int
`

type GetSetsByWorkoutIdDescParams struct {
	WorkoutID  int32
	OrderByCol string
	Offset     int32
	Limit      int32
}

func (q *Queries) GetSetsByWorkoutIdDesc(ctx context.Context, arg GetSetsByWorkoutIdDescParams) ([]Set, error) {
	rows, err := q.db.QueryContext(ctx, getSetsByWorkoutIdDesc,
		arg.WorkoutID,
		arg.OrderByCol,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Set
	for rows.Next() {
		var i Set
		if err := rows.Scan(
			&i.ID,
			&i.Exercise,
			&i.Count,
			&i.Intensity,
			&i.Type,
			&i.Weight,
			&i.WorkoutID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
