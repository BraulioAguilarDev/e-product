-- name: ListProducts :many
SELECT * FROM product;

-- name: CreateProduct :one
INSERT INTO product (
  sku, name, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM product WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;

-- name: UpdateProduct :one
UPDATE product SET sku = $1, name = $2, price = $3 WHERE id = $4 RETURNING *;
