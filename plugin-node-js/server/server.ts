import {IDestinationPluginServer} from "../proto/destination_grpc_pb";
import {sendUnaryData, ServerUnaryCall} from "grpc";
import {
    BatchDeleteUsersRequest,
    BatchDeleteUsersResponse, BatchUpdateUsersRequest,
    BatchUpdateUsersResponse, BatchUpsertUsersRequest,
    BatchUpsertUsersResponse
} from "../proto/destination_pb";

export class DestinationPluginService implements IDestinationPluginServer {
    batchDeleteUsers(call: ServerUnaryCall<BatchDeleteUsersRequest>, callback: sendUnaryData<BatchDeleteUsersResponse>): void {
    }

    batchUpdateUsers(call: ServerUnaryCall<BatchUpdateUsersRequest>, callback: sendUnaryData<BatchUpdateUsersResponse>): void {
    }

    batchUpsertUsers(call: ServerUnaryCall<BatchUpsertUsersRequest>, callback: sendUnaryData<BatchUpsertUsersResponse>): void {
    }

}