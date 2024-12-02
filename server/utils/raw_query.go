package utils

const (
	INSERT_PRODUCT = `
		INSERT INTO product (
			id, 
			price, 
			name
		) VALUES ($1, $2, $3)
		RETURNING id, price, name, created_at
	`
	SELECT_PRODUCT_WITH_PAGING = "SELECT * FROM product ORDER BY %s %s LIMIT $1 OFFSET $2"
	SELECT_COUNT_PRODUCT = "SELECT COUNT(id) FROM product"
	UPDATE_PRODUCT_BY_ID = "UPDATE product SET price = $2, name = $3 WHERE id = $1"
	DELETE_PRODUCT_BY_ID = "DELETE FROM product WHERE id = $1"
)