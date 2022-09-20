import {sendUnaryData, ServerUnaryCall, ServerWritableStream, UntypedHandleCall} from "@grpc/grpc-js";

import {IHealthServer} from "../proto/health_grpc_pb";
import {HealthCheckRequest, HealthCheckResponse} from "../proto/health_pb";
import ServingStatus = HealthCheckResponse.ServingStatus;

export class HealthServerImpl implements IHealthServer {
    [name: string]: UntypedHandleCall;

    check(call: ServerUnaryCall<HealthCheckRequest, HealthCheckResponse>, callback: sendUnaryData<HealthCheckResponse>): void {
        const response = new HealthCheckResponse();
        response.setStatus(ServingStatus.SERVING);
        callback(null, response);
    }

    watch(call: ServerWritableStream<HealthCheckRequest, HealthCheckResponse>): void {
        call.on('data', function() {
            const response = new HealthCheckResponse();
            response.setStatus(ServingStatus.SERVING);
            call.write(response);
        });
    }
}