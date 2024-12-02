package utils

const (
	INSERT_PRODUCT = `
		INSERT INTO product (
			id, 
			img_url,
			price, 
			name
		) VALUES ($1, $2, $3, $4)
		RETURNING id, img_url, price, name, created_at
	`
	SELECT_PRODUCT_WITH_PAGING = "SELECT id, img_url, price, name, created_at FROM product ORDER BY %s %s LIMIT $1 OFFSET $2"
	SELECT_COUNT_PRODUCT = "SELECT COUNT(id) FROM product"
	UPDATE_PRODUCT_BY_ID = "UPDATE product SET img_url = $2, price = $3, name = $4 WHERE id = $1"
	DELETE_PRODUCT_BY_ID = "DELETE FROM product WHERE id = $1"
)