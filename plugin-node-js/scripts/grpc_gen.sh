#!/bin/bash

IN_PROTO_DIR=../proto
OUT_PROTO_DIR=./proto

# Generate JavaScript code
grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:${OUT_PROTO_DIR} \
    --grpc_out=${OUT_PROTO_DIR} \
    --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
    -I ${IN_PROTO_DIR} \
    ${IN_PROTO_DIR}/*.proto

# Generate TypeScript code (d.ts)
grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=${OUT_PROTO_DIR} \
    -I ${IN_PROTO_DIR} \
    ${IN_PROTO_DIR}/*.proto