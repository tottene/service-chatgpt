# Criar Projeto
- go mod init nome_do_projeto (github.com/tottene/fclx/chatservice)

# Atualizar Libs
- go mod tidy

# SQLC Migrate
- migrate create -ext=mysql -dir=sql/migrations -seq init

# Build Rust
- rust build --execute

# Docker
- docker-compose exec mysql-chatgpt bash
- docker exec -ti chatservice_app /bin/bash
- docker stop $(docker ps -a -q)