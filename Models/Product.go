package Models

import (
	"context"
	"fmt"
	"go-mssql/Config"
)

// GetAllUsers Fetch all data
func GetAllProducts() (product []Product, err error) {
	// Check DB Alive
	ctx := context.Background()

	if err = Config.DB.PingContext(ctx); err != nil {
		return nil, err
	}

	// Execute query
	tsql := "SELECT * FROM product"
	rows, err := Config.DB.QueryContext(ctx, tsql)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productstuct Product

	for rows.Next() {
		// Get values from row.
		err := rows.Scan(&productstuct.Id, &productstuct.Name, &productstuct.Price, &productstuct.Detail)
		if err != nil {
			return nil, err
		}
		product = append(product, productstuct)
	}

	return product, nil

}

// CreateUser ... Insert New data
func CreateProduct(product *Product) (status int64, err error) {
	// Check DB Alive
	ctx := context.Background()

	if err = Config.DB.PingContext(ctx); err != nil {
		return -1, err
	}

	// Execute Insert by Procedure
	tsql := fmt.Sprintf("insert_product @p_name ='%s', @p_price =%d, @p_detail='%s'",
		product.Name, product.Price, product.Detail)

	result, err := Config.DB.ExecContext(ctx, tsql)

	if err != nil {
		fmt.Println("error2")
		return -1, err
	}

	return result.RowsAffected()
}

// GetUserByID ... Fetch only one data by Id
func GetProductByID(product *Product, id string) (err error) {
	// Check DB Alive
	ctx := context.Background()

	if err = Config.DB.PingContext(ctx); err != nil {
		return err
	}

	// Execute query
	tsql := fmt.Sprintf("SELECT * FROM product where id=%s", id)
	row := Config.DB.QueryRowContext(ctx, tsql)

	var productstuct Product

	err = row.Scan(&productstuct.Id, &productstuct.Name, &productstuct.Price, &productstuct.Detail)
	if err != nil {
		return err
	}
	product.Id = productstuct.Id
	product.Name = productstuct.Name
	product.Price = productstuct.Price
	product.Detail = productstuct.Detail

	return nil
}

// UpdateUser ... Update data
func UpdateProduct(product *Product, id string) (status int64, err error) {
	// Check DB Alive
	ctx := context.Background()

	if err = Config.DB.PingContext(ctx); err != nil {
		return -1, err
	}

	// Execute Insert by Procedure
	tsql := fmt.Sprintf("update_product @p_id=%s, @p_name ='%s', @p_price =%d, @p_detail='%s'",
		id, product.Name, product.Price, product.Detail)

	result, err := Config.DB.ExecContext(ctx, tsql)

	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// DeleteUser ... Delete data
func DeleteProduct(product *Product, id string) (status int64, err error) {
	// Check DB Alive
	ctx := context.Background()

	if err = Config.DB.PingContext(ctx); err != nil {
		return -1, err
	}
	// Execute Delete
	tsql := fmt.Sprintf("DELETE FROM product WHERE id=%s",
		id)

	result, err := Config.DB.ExecContext(ctx, tsql)

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
