-- name: FindMember :one
SELECT * FROM members
WHERE id = ? ;

-- name: CreateMember :execresult
INSERT INTO members (
    first_name,
    last_name,
    created_at,
    updated_at
) VALUES (
  ?, ?, ?, ?
);
