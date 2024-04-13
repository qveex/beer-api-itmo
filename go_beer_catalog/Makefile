LOCAL_PROTO_IN:=$(CURDIR)/api/proto
LOCAL_PROTO_OUT:=$(CURDIR)/pkg/api

build_proto:
	protoc -I api/proto --go-grpc_out=$(LOCAL_PROTO_OUT) $(LOCAL_PROTO_IN)/catalog.proto
	#protoc -I $(LOCAL_PROTO_OUT) --go_out=$(LOCAL_PROTO_IN)/catalog.proto
	protoc --proto_path=$(LOCAL_PROTO_IN) --go_out=$(LOCAL_PROTO_OUT) catalog.proto beer.proto get_beer.proto get_beers.proto create_beer.proto update_beer.proto delete_beer.proto