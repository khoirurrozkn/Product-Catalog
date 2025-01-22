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
	SELECT_COUNT_PRODUCT       = "SELECT COUNT(id) FROM product"
	UPDATE_PRODUCT_BY_ID       = "UPDATE product SET img_url = $2, price = $3, name = $4 WHERE id = $1"
	DELETE_PRODUCT_BY_ID       = "DELETE FROM product WHERE id = $1"

	INSERT_USER = `
		INSERT INTO users (
			id, 
			email,
			nickname, 
			password,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4)
	`
	SELECT_USER_WITH_PAGING = "SELECT id, email, nickname, created_at, updated_at FROM users ORDER BY %s %s LIMIT $1 OFFSET $2"
	SELECT_COUNT_USER       = "SELECT COUNT(id) FROM users"
	SELECT_USER_BY_ID       = "SELECT id, email, nickname, created_at, updated_at FROM users"
	UPDATE_USER_BY_ID       = "UPDATE users SET email = $2, nickname = $3, password = $4 WHERE id = $1"
	DELETE_USER_BY_ID       = "DELETE FROM users WHERE id = $1"

	INSERT_FAVORITE = `
		INSERT INTO favorite (
			id, 
			user_id,
			product_id,
			created_at
		) VALUES ($1, $2, $3, $4)
	`
	SELECT_FAVORITE_WITH_PAGING = "SELECT id, user_id, product_id, created_at FROM favorite ORDER BY %s %s LIMIT $1 OFFSET $2"
	SELECT_COUNT_FAVORITE       = "SELECT COUNT(id) FROM favorite"
	SELECT_FAVORITE_BY_ID       = "SELECT id, user_id, product_id, created_at FROM favorite"
	DELETE_FAVORITE_BY_ID       = "DELETE FROM favorite WHERE id = $1"
)
