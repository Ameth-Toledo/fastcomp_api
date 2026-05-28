package adapters

import (
	"database/sql"
	"errors"
	"fastcomp_api/features/products/domain/entities"
	"fmt"

	"github.com/lib/pq"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

func scan(row *sql.Row, p *entities.Product) error {
	return row.Scan(
		&p.ID, &p.Brand, &p.Name, &p.Category, &p.Price,
		&p.Stock, &p.Condition, &p.Featured, &p.Description,
		pq.Array(&p.Specs), &p.ImageURL, &p.AddedAt,
	)
}

func scanRow(rows *sql.Rows, p *entities.Product) error {
	return rows.Scan(
		&p.ID, &p.Brand, &p.Name, &p.Category, &p.Price,
		&p.Stock, &p.Condition, &p.Featured, &p.Description,
		pq.Array(&p.Specs), &p.ImageURL, &p.AddedAt,
	)
}

const selectFields = `
	SELECT id, brand, name, category, price, stock, condition,
	       featured, description, specs, image_url, added_at
	FROM products`

func (ps *PostgreSQL) Save(p *entities.Product) (*entities.Product, error) {
	query := `
		INSERT INTO products (id, brand, name, category, price, stock, condition, featured, description, specs, image_url, added_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := ps.conn.Exec(query,
		p.ID, p.Brand, p.Name, string(p.Category), p.Price,
		p.Stock, string(p.Condition), p.Featured, p.Description,
		pq.Array(p.Specs), p.ImageURL, p.AddedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error al guardar producto: %v", err)
	}
	return p, nil
}

func (ps *PostgreSQL) GetByID(id string) (*entities.Product, error) {
	query := selectFields + ` WHERE id = $1`
	var p entities.Product
	err := scan(ps.conn.QueryRow(query, id), &p)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al buscar producto: %v", err)
	}
	return &p, nil
}

func (ps *PostgreSQL) GetAll() ([]*entities.Product, error) {
	rows, err := ps.conn.Query(selectFields + ` ORDER BY added_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("error al obtener productos: %v", err)
	}
	defer rows.Close()
	return scanRows(rows)
}

func (ps *PostgreSQL) Update(p *entities.Product) error {
	query := `
		UPDATE products SET
			brand = $2, name = $3, category = $4, price = $5,
			stock = $6, condition = $7, featured = $8, description = $9, specs = $10,
			image_url = $11
		WHERE id = $1
	`
	result, err := ps.conn.Exec(query,
		p.ID, p.Brand, p.Name, string(p.Category), p.Price,
		p.Stock, string(p.Condition), p.Featured, p.Description,
		pq.Array(p.Specs), p.ImageURL,
	)
	if err != nil {
		return fmt.Errorf("error al actualizar producto: %v", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}

func (ps *PostgreSQL) Delete(id string) error {
	result, err := ps.conn.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error al eliminar producto: %v", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}

func (ps *PostgreSQL) GetFeatured() ([]*entities.Product, error) {
	rows, err := ps.conn.Query(selectFields + ` WHERE featured = true ORDER BY added_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("error al obtener destacados: %v", err)
	}
	defer rows.Close()
	return scanRows(rows)
}

func (ps *PostgreSQL) GetByCategory(category entities.Category) ([]*entities.Product, error) {
	rows, err := ps.conn.Query(selectFields+` WHERE category = $1 ORDER BY added_at DESC`, string(category))
	if err != nil {
		return nil, fmt.Errorf("error al obtener por categoría: %v", err)
	}
	defer rows.Close()
	return scanRows(rows)
}

func scanRows(rows *sql.Rows) ([]*entities.Product, error) {
	var products []*entities.Product
	for rows.Next() {
		var p entities.Product
		if err := scanRow(rows, &p); err != nil {
			return nil, fmt.Errorf("error al escanear producto: %v", err)
		}
		products = append(products, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar productos: %v", err)
	}
	return products, nil
}
