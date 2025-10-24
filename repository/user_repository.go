package repository

import (
	"database/sql"
	"fmt"
	"go-crud-starter/config"
	"go-crud-starter/models"
	"strings"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.DB}
}

func (r *UserRepository) Create(user *models.CreateUserRequest) (*models.User, error) {
	query := `
		INSERT INTO users (name, email, phone)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, phone, created_at, updated_at
	`

	var result models.User
	err := r.db.QueryRow(query, user.Name, user.Email, user.Phone).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Phone,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *UserRepository) FindAll(page, limit int) (*models.PaginatedResponse, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, name, email, phone, created_at, updated_at
		FROM users
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return nil, err
	}

	totalPages := (total + limit - 1) / limit

	return &models.PaginatedResponse{
		Data:       users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	query := `
		SELECT id, name, email, phone, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Search(query string) ([]models.User, error) {
	searchQuery := `
		SELECT id, name, email, phone, created_at, updated_at
		FROM users
		WHERE name ILIKE $1 OR email ILIKE $1
		ORDER BY id DESC
		LIMIT 50
	`

	rows, err := r.db.Query(searchQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Update(id int, req *models.UpdateUserRequest) (*models.User, error) {
	var fields []string
	var args []interface{}
	argCount := 1

	if req.Name != nil {
		fields = append(fields, fmt.Sprintf("name = $%d", argCount))
		args = append(args, *req.Name)
		argCount++
	}
	if req.Email != nil {
		fields = append(fields, fmt.Sprintf("email = $%d", argCount))
		args = append(args, *req.Email)
		argCount++
	}
	if req.Phone != nil {
		fields = append(fields, fmt.Sprintf("phone = $%d", argCount))
		args = append(args, *req.Phone)
		argCount++
	}

	if len(fields) == 0 {
		return r.FindByID(id)
	}

	fields = append(fields, "updated_at = CURRENT_TIMESTAMP")
	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
		RETURNING id, name, email, phone, created_at, updated_at
	`, strings.Join(fields, ", "), argCount)

	var user models.User
	err := r.db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id int) (bool, error) {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *UserRepository) EmailExists(email string, excludeID *int) (bool, error) {
	var query string
	var args []interface{}

	if excludeID != nil {
		query = "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND id != $2)"
		args = []interface{}{email, *excludeID}
	} else {
		query = "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
		args = []interface{}{email}
	}

	var exists bool
	err := r.db.QueryRow(query, args...).Scan(&exists)
	return exists, err
}
