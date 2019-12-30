### GoModules

export GO111MODULE=on

go mod init github.com/azul915/portfolio-api

go: creating new go.mod: module github.com/azul915/portfolio-api
コマンドを実行したディレクトリにgo.modができる

go get ~

go build

docker-compose restart

docker-compose logs -f --tail="5"