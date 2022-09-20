// package: proto
// file: destination.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as destination_pb from "./destination_pb";

interface IDestinationPluginService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    batchUpsertUsers: IDestinationPluginService_IBatchUpsertUsers;
    batchUpdateUsers: IDestinationPluginService_IBatchUpdateUsers;
    batchDeleteUsers: IDestinationPluginService_IBatchDeleteUsers;
}

interface IDestinationPluginService_IBatchUpsertUsers extends grpc.MethodDefinition<destination_pb.BatchUpsertUsersRequest, destination_pb.BatchUpsertUsersResponse> {
    path: "/proto.DestinationPlugin/BatchUpsertUsers";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<destination_pb.BatchUpsertUsersRequest>;
    requestDeserialize: grpc.deserialize<destination_pb.BatchUpsertUsersRequest>;
    responseSerialize: grpc.serialize<destination_pb.BatchUpsertUsersResponse>;
    responseDeserialize: grpc.deserialize<destination_pb.BatchUpsertUsersResponse>;
}
interface IDestinationPluginService_IBatchUpdateUsers extends grpc.MethodDefinition<destination_pb.BatchUpdateUsersRequest, destination_pb.BatchUpdateUsersResponse> {
    path: "/proto.DestinationPlugin/BatchUpdateUsers";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<destination_pb.BatchUpdateUsersRequest>;
    requestDeserialize: grpc.deserialize<destination_pb.BatchUpdateUsersRequest>;
    responseSerialize: grpc.serialize<destination_pb.BatchUpdateUsersResponse>;
    responseDeserialize: grpc.deserialize<destination_pb.BatchUpdateUsersResponse>;
}
interface IDestinationPluginService_IBatchDeleteUsers extends grpc.MethodDefinition<destination_pb.BatchDeleteUsersRequest, destination_pb.BatchDeleteUsersResponse> {
    path: "/proto.DestinationPlugin/BatchDeleteUsers";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<destination_pb.BatchDeleteUsersRequest>;
    requestDeserialize: grpc.deserialize<destination_pb.BatchDeleteUsersRequest>;
    responseSerialize: grpc.serialize<destination_pb.BatchDeleteUsersResponse>;
    responseDeserialize: grpc.deserialize<destination_pb.BatchDeleteUsersResponse>;
}

export const DestinationPluginService: IDestinationPluginService;

export interface IDestinationPluginServer extends grpc.UntypedServiceImplementation {
    batchUpsertUsers: grpc.handleUnaryCall<destination_pb.BatchUpsertUsersRequest, destination_pb.BatchUpsertUsersResponse>;
    batchUpdateUsers: grpc.handleUnaryCall<destination_pb.BatchUpdateUsersRequest, destination_pb.BatchUpdateUsersResponse>;
    batchDeleteUsers: grpc.handleUnaryCall<destination_pb.BatchDeleteUsersRequest, destination_pb.BatchDeleteUsersResponse>;
}

export interface IDestinationPluginClient {
    batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
    batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
    batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
}

export class DestinationPluginClient extends grpc.Client implements IDestinationPluginClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    public batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    public batchUpsertUsers(request: destination_pb.BatchUpsertUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpsertUsersResponse) => void): grpc.ClientUnaryCall;
    public batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    public batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    public batchUpdateUsers(request: destination_pb.BatchUpdateUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchUpdateUsersResponse) => void): grpc.ClientUnaryCall;
    public batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
    public batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
    public batchDeleteUsers(request: destination_pb.BatchDeleteUsersRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: destination_pb.BatchDeleteUsersResponse) => void): grpc.ClientUnaryCall;
}
