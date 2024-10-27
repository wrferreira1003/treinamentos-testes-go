# Criar comando para rodar a aplicacao
# Agora com o make, e so rodar o comando: make createmigration, make migrateup, make migratedown

createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrateup:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

.PHONY: createmigration migrateup migratedown # Ignore o arquivo se ele ja existir

