`init postgresql server`

[Setup Postgres on Arch](https://github.com/malnvenshorn/OctoPrint-FilamentManager/wiki/Setup-PostgreSQL-on-Arch-Linux)

`start postgresql server`

`sudo systemctl start postgresql.service`
`sudo systemctl enable postgresql.service`

### Install Golang and twimo's dependencies (on Mac)

brew install go
go get github.com/google/uuid
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/websocket

### Set path

`export GOPATH="/Users/louis/Documents/Engineering/go"`
