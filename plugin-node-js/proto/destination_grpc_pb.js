// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var destination_pb = require('./destination_pb.js');

function serialize_proto_BatchDeleteUsersRequest(arg) {
  if (!(arg instanceof destination_pb.BatchDeleteUsersRequest)) {
    throw new Error('Expected argument of type proto.BatchDeleteUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchDeleteUsersRequest(buffer_arg) {
  return destination_pb.BatchDeleteUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BatchDeleteUsersResponse(arg) {
  if (!(arg instanceof destination_pb.BatchDeleteUsersResponse)) {
    throw new Error('Expected argument of type proto.BatchDeleteUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchDeleteUsersResponse(buffer_arg) {
  return destination_pb.BatchDeleteUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BatchUpdateUsersRequest(arg) {
  if (!(arg instanceof destination_pb.BatchUpdateUsersRequest)) {
    throw new Error('Expected argument of type proto.BatchUpdateUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchUpdateUsersRequest(buffer_arg) {
  return destination_pb.BatchUpdateUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BatchUpdateUsersResponse(arg) {
  if (!(arg instanceof destination_pb.BatchUpdateUsersResponse)) {
    throw new Error('Expected argument of type proto.BatchUpdateUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchUpdateUsersResponse(buffer_arg) {
  return destination_pb.BatchUpdateUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BatchUpsertUsersRequest(arg) {
  if (!(arg instanceof destination_pb.BatchUpsertUsersRequest)) {
    throw new Error('Expected argument of type proto.BatchUpsertUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchUpsertUsersRequest(buffer_arg) {
  return destination_pb.BatchUpsertUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_BatchUpsertUsersResponse(arg) {
  if (!(arg instanceof destination_pb.BatchUpsertUsersResponse)) {
    throw new Error('Expected argument of type proto.BatchUpsertUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_BatchUpsertUsersResponse(buffer_arg) {
  return destination_pb.BatchUpsertUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var DestinationPluginService = exports.DestinationPluginService = {
  batchUpsertUsers: {
    path: '/proto.DestinationPlugin/BatchUpsertUsers',
    requestStream: false,
    responseStream: false,
    requestType: destination_pb.BatchUpsertUsersRequest,
    responseType: destination_pb.BatchUpsertUsersResponse,
    requestSerialize: serialize_proto_BatchUpsertUsersRequest,
    requestDeserialize: deserialize_proto_BatchUpsertUsersRequest,
    responseSerialize: serialize_proto_BatchUpsertUsersResponse,
    responseDeserialize: deserialize_proto_BatchUpsertUsersResponse,
  },
  batchUpdateUsers: {
    path: '/proto.DestinationPlugin/BatchUpdateUsers',
    requestStream: false,
    responseStream: false,
    requestType: destination_pb.BatchUpdateUsersRequest,
    responseType: destination_pb.BatchUpdateUsersResponse,
    requestSerialize: serialize_proto_BatchUpdateUsersRequest,
    requestDeserialize: deserialize_proto_BatchUpdateUsersRequest,
    responseSerialize: serialize_proto_BatchUpdateUsersResponse,
    responseDeserialize: deserialize_proto_BatchUpdateUsersResponse,
  },
  batchDeleteUsers: {
    path: '/proto.DestinationPlugin/BatchDeleteUsers',
    requestStream: false,
    responseStream: false,
    requestType: destination_pb.BatchDeleteUsersRequest,
    responseType: destination_pb.BatchDeleteUsersResponse,
    requestSerialize: serialize_proto_BatchDeleteUsersRequest,
    requestDeserialize: deserialize_proto_BatchDeleteUsersRequest,
    responseSerialize: serialize_proto_BatchDeleteUsersResponse,
    responseDeserialize: deserialize_proto_BatchDeleteUsersResponse,
  },
};

exports.DestinationPluginClient = grpc.makeGenericClientConstructor(DestinationPluginService);
