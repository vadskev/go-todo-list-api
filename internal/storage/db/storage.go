package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/storage"
	"go.uber.org/zap"
)

const (
	tableName  = "scheduler"
	colTitle   = "title"
	colDate    = "date"
	colComment = "comment"
	colRepeat  = "repeat"
)

type Repository struct {
	db *sql.DB
}

var _ storage.TaskRepository = (*Repository)(nil)

/**/

func NewRepository(filepath string) (Repository, error) {
	const op = "storage.db.storage.NewRepository"
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", filepath)
		if err != nil {
			logger.Error(op, zap.Any("error:", err.Error()))
		}

		_, err = db.Exec(` CREATE TABLE scheduler (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date TEXT,
            title TEXT,
            comment TEXT,
            repeat TEXT
        );
        CREATE INDEX idx_scheduler_date ON scheduler(date);`)
		if err != nil {
			logger.Error(op, zap.Any("error:", err.Error()))
			return Repository{}, err
		}

		return Repository{db: db}, nil

	} else if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
		return Repository{}, err
	}

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		logger.Error(op, zap.Any("error:", err.Error()))
	}

	return Repository{
		db: db,
	}, nil
}

func (r *Repository) DB() *sql.DB {
	return r.db
}

func (r *Repository) Close() error {
	const op = "storage.db.storage.Close"
	if r.db != nil {
		if err := r.db.Close(); err != nil {
			logger.Error(op, zap.Any("error:", err.Error()))
		}
	}
	return nil
}
