package repository

import (
	"database/sql"
	"github.com/rg-km/final-project-engineering-70/backend/repository"
)

type UserRepository struct {
	db *sql.DB
}
