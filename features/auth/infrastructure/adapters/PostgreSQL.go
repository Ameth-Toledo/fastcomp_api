package adapters

import (
	"database/sql"
	"errors"
	"fastcomp_api/features/auth/domain/entities"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func (ps *PostgreSQL) Save(user *entities.User) (*entities.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, business_name, email, password, phone, website, profile_photo, registration_date, role_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`
	err := ps.conn.QueryRow(query,
		user.FirstName, user.LastName, user.BusinessName,
		user.Email, user.Password, user.Phone,
		user.Website, user.ProfilePhoto,
		user.RegistrationDate, user.RoleID,
	).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error al guardar usuario: %v", err)
	}
	return user, nil
}

func (ps *PostgreSQL) GetByEmail(email string) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, business_name, email, password, phone, website, profile_photo, registration_date, role_id
		FROM users WHERE email = $1
	`
	var u entities.User
	err := ps.conn.QueryRow(query, email).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.BusinessName,
		&u.Email, &u.Password, &u.Phone,
		&u.Website, &u.ProfilePhoto, &u.RegistrationDate, &u.RoleID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}
	return &u, nil
}

func (ps *PostgreSQL) GetByID(id int) (*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, business_name, email, password, phone, website, profile_photo, registration_date, role_id
		FROM users WHERE id = $1
	`
	var u entities.User
	err := ps.conn.QueryRow(query, id).Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.BusinessName,
		&u.Email, &u.Password, &u.Phone,
		&u.Website, &u.ProfilePhoto, &u.RegistrationDate, &u.RoleID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar usuario: %v", err)
	}
	return &u, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.User, error) {
	query := `
		SELECT id, first_name, last_name, business_name, email, password, phone, website, profile_photo, registration_date, role_id
		FROM users ORDER BY registration_date DESC
	`
	rows, err := ps.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var u entities.User
		if err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.BusinessName,
			&u.Email, &u.Password, &u.Phone,
			&u.Website, &u.ProfilePhoto, &u.RegistrationDate, &u.RoleID,
		); err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		users = append(users, &u)
	}
	return users, rows.Err()
}

func (ps *PostgreSQL) Update(user *entities.User) error {
	query := `
		UPDATE users SET
			first_name = $2, last_name = $3, business_name = $4,
			email = $5, phone = $6, website = $7, profile_photo = $8
		WHERE id = $1
	`
	result, err := ps.conn.Exec(query,
		user.ID, user.FirstName, user.LastName, user.BusinessName,
		user.Email, user.Phone, user.Website, user.ProfilePhoto,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("usuario no encontrado")
	}
	return nil
}

func (ps *PostgreSQL) Delete(id int) error {
	result, err := ps.conn.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("usuario no encontrado")
	}
	return nil
}

func (ps *PostgreSQL) GetTotal() (int, error) {
	var total int
	err := ps.conn.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error al obtener total: %v", err)
	}
	return total, nil
}
