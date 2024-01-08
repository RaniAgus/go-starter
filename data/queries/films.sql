-- name: GetFilm :one
SELECT * FROM films
WHERE id = $1 LIMIT 1;

-- name: ListFilms :many
SELECT * FROM films
ORDER BY title;

-- name: CreateFilms :copyfrom
INSERT INTO films (
  title, director
) VALUES (
  $1, $2
);

-- name: CreateFilm :one
INSERT INTO films (
  title, director
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateFilm :one
UPDATE films
  set title = $2,
  director = $3
WHERE id = $1
RETURNING *;

-- name: DeleteFilm :exec
DELETE FROM films
WHERE id = $1;
