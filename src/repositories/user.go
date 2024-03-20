package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"time"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

type CreateUserModelResponse struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (repository users) Create(user models.User) (CreateUserModelResponse, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at")
	if err != nil {
		return CreateUserModelResponse{}, err
	}
	defer statement.Close()
	var response CreateUserModelResponse
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&response.ID, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		return CreateUserModelResponse{}, err
	}

	return response, nil
}

func (repository users) List(query string) ([]models.User, error) {
	query = fmt.Sprintf("%%%s%%", query)
	rows, err := repository.db.Query(
		"SELECT * FROM users WHERE name LIKE $1 or nick LIKE $2", query, query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) ListOne(id string) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT * FROM users WHERE id = $1", id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Password,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(id string, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(id string) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository users) FindByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Follow(userId string, followerId string) error {
	statement, err := repository.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) Unfollow(userId string, followerId string) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = $1 AND follower_id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository users) SearchFollowers(userId string) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.created_at, u.updated_at
        FROM users u inner join followers s on u.id = s.follower_id where s.user_id = $1
        `, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository users) SearchFollowings(userId string) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.created_at, u.updated_at
          FROM users u inner join followers s on u.id = s.user_id where s.follower_id = $1
          `, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
