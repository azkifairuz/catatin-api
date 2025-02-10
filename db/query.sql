-- name: GetWalletByAccount :many
select * from wallet where accounts_id = $1;