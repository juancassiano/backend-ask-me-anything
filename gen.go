package gen

//go:generate tern migrate --migrations ./internal/store/pgstore/migrations/ --config ./internal/store/pgstore/migrations/tern.conf
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yaml
