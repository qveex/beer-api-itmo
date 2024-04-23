LOCAL_PROTO_IN:=$(CURDIR)/api/proto
LOCAL_PROTO_OUT:=$(CURDIR)/pkg/api

build_proto:
	protoc --proto_path=$(LOCAL_PROTO_IN) $(LOCAL_PROTO_IN)/*.proto --go_out=$(LOCAL_PROTO_OUT) --go-grpc_out=$(LOCAL_PROTO_OUT)
