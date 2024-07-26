package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/models/task"
	"go.uber.org/zap"
)

func (r *Repository) GetTasks(searchStr string) ([]task.Task, error) {
	const op = "storage.db.GetTasks"
	builderGet := sq.
		Select("*").
		From(tableName)

	searchDate, err := time.Parse("02.01.2006", searchStr)
	if err == nil {
		builderGet = builderGet.Where(sq.Eq{"date": searchDate.Format("20060102")})
	} else {
		builderGet = builderGet.Where(
			sq.Or{
				sq.Like{"title": "%" + searchStr + "%"},
				sq.Like{"comment": "%" + searchStr + "%"},
			},
		)
	}

	builderGet = builderGet.OrderBy("date ASC").
		Limit(10)

	query, args, err := builderGet.ToSql()
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return nil, err
	}

	rows, err := r.DB().Query(query, args...)
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		err = rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			logger.Error(op, zap.Any("error:", err.Error()))
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) GetById(id string) (*task.Task, error) {
	const op = "storage.db.GetById"
	query := sq.
		Select("*").
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return nil, err
	}

	row := r.DB().QueryRow(sql, args...)

	var t task.Task
	err = row.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return nil, err
	}
	return &t, nil
}
