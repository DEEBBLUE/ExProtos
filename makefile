CC = protoc
PREF = ./api/
CONF = --proto_path=./contracts --go_out=$(PREF)$@ --go_opt=paths=source_relative --go-grpc_out=$(PREF)$@ --go-grpc_opt=paths=source_relative

COM = $(CC) $(CONF)

all: Utils Service Adapters

Database:
	$(COM) DataBase.proto 

Exchange:
	$(COM) ExchangeProxy.proto 

SpecialExchange:
	$(COM) SpecialExchange.proto 

Sso:
	$(COM) Sso.proto 

Message:
	$(COM) Message.proto 

Rate:
	$(COM) Rate.proto 

Stock:
	$(COM) Stock.proto 

Types: 
	$(COM) Types.proto 

Req:
	$(COM) Req.proto 

Utils: Req Types

Service: Exchange SpecialExchange Database Sso Message Rate 

Adapters: Stock
