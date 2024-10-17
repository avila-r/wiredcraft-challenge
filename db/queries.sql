-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, dob, description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    dob = $3,
    description = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserAddress :one
SELECT * FROM user_address
WHERE id = $1 LIMIT 1;

-- name: ListUserAddressByUser :many
SELECT * FROM user_address
WHERE user_id = $1
ORDER BY city;

-- name: CreateUserAddress :one
INSERT INTO user_address (
  user_id, address_line1, address_line2, city, state, postal_code, country
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateUserAddress :exec
UPDATE user_address
SET address_line1 = $2,
    address_line2 = $3,
    city = $4,
    state = $5,
    postal_code = $6,
    country = $7
WHERE id = $1;

-- name: DeleteUserAddress :exec
DELETE FROM user_address
WHERE id = $1;
