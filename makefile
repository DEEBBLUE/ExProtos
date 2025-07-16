CC = protoc
PREF = ./api/
CONF = --proto_path=./contracts --go_out=$(PREF)$@ --go_opt=paths=source_relative --go-grpc_out=$(PREF)$@ --go-grpc_opt=paths=source_relative

COM = $(CC) $(CONF)

all: Utils Service 

Database:
	$(COM) DataBase.proto 

Exchange:
	$(COM) Exchange.proto 

Sso:
	$(COM) Sso.proto 

Message:
	$(COM) Message.proto 

Types: 
	$(COM) Types.proto 

Req:
	$(COM) Req.proto 

Utils: Req Types

Service: Exchange Database Sso Message

