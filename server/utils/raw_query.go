package utils

const (
	INSERT_PRODUCT = `
		INSERT INTO product (
			id, 
			price, 
			name
		) VALUES ($1, $2, $3)
		RETURNING id, price, name
	`
	SELECT_PRODUCT = "SELECT * FROM product"
	UPDATE_PRODUCT_BY_ID = "UPDATE product SET price = $2, name = $3 WHERE id = $1"
	DELETE_PRODUCT_BY_ID = "DELETE FROM product WHERE id = $1"
)