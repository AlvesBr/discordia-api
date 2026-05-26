package repository

import (
	"api-go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		println(err)
		return []model.Product{}, err
	}
	var productList []model.Product
	var productObj model.Product
	for rows.Next() {
		err := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil

}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product " +
		"(name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT id, name, price FROM product WHERE id = $1")
	if err != nil {
		return nil, err
	}
	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)

	if err == sql.ErrNoRows {
		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	_, err := pr.connection.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (model.Product, error) {
	_, err := pr.connection.Exec("UPDATE product SET name = $1, price = $2 WHERE id = $3",
		product.Name, product.Price, product.ID)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
