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
	SELECT_CUSTOMER_WITH_PAGING = "SELECT * FROM product LIMIT $1 OFFSET $2"
	UPDATE_CUSTOMER_BY_ID = "UPDATE product SET price = $2, name = $3 WHERE id = $1"
	DELETE_CUSTOMER_BY_ID = "DELETE FROM product WHERE id = $1"
)