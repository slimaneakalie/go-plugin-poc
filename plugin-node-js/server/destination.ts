import {IDestinationPluginServer} from "../proto/destination_grpc_pb";
import {sendUnaryData, ServerUnaryCall} from "@grpc/grpc-js";
import {
    BatchDeleteUsersRequest,
    BatchDeleteUsersResponse,
    BatchUpdateUsersRequest,
    BatchUpdateUsersResponse,
    BatchUpsertUsersRequest,
    BatchUpsertUsersResponse,
    OperationStatus
} from "../proto/destination_pb";
import {UntypedHandleCall} from "@grpc/grpc-js";

export class DestinationPluginImpl implements IDestinationPluginServer {
    batchDeleteUsers(call: ServerUnaryCall<BatchDeleteUsersRequest, BatchDeleteUsersResponse>, callback: sendUnaryData<BatchDeleteUsersResponse>): void {
    }

    batchUpdateUsers(call: ServerUnaryCall<BatchUpdateUsersRequest, BatchUpdateUsersResponse>, callback: sendUnaryData<BatchUpdateUsersResponse>): void {
    }

    batchUpsertUsers(call: ServerUnaryCall<BatchUpsertUsersRequest, BatchUpsertUsersResponse>, callback: sendUnaryData<BatchUpsertUsersResponse>): void {
        console.log("request: ", call.request);
        const resp = new BatchUpsertUsersResponse()
        resp.setStatus(OperationStatus.SUCCESS)
        callback(null, resp)
    }

    [name: string]: UntypedHandleCall;
}