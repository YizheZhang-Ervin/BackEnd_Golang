# GoKit Demo

## Commands
```
# init
cd server
go mod init server
cd client
go mod init client
cd ..
go work init
go work use server
go work use client

# run
go run ./server/main.go -httpAddr :8080

# add
curl -d '{"id":"1111","Name":"Go Kit"}' -H "Content-Type: application/json" -X POST http://localhost:8080/profiles/

# search
curl localhost:8080/profiles/1111
```
