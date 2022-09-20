import { Server, ServerCredentials } from "@grpc/grpc-js";
import {DestinationPluginService} from "../proto/destination_grpc_pb";
import { DestinationPluginImpl } from "./destination";
import {HealthService} from "../proto/health_grpc_pb";
import {HealthServerImpl} from "./health";

const server = new Server();
server.addService(DestinationPluginService, new DestinationPluginImpl());
server.addService(HealthService, new HealthServerImpl());

// for more info: check https://github.com/hashicorp/go-plugin/blob/master/docs/guide-plugin-write-non-go.md
const PLUGIN_CORE_PROTOCOL_VERSION = "1";
const PLUGIN_APP_PROTOCOL_VERSION = "1";
const PLUGIN_NETWORK_TYPE = "tcp";
const PLUGIN_PROTOCOL = "grpc";

const host = '127.0.0.1';
server.bindAsync(`${host}:0`, ServerCredentials.createInsecure(), (error: Error | null, port: number) => {
    if (error) {
        throw (error);
    }

    const PLUGIN_NETWORK_ADDRESS = `${host}:${port}`;
    const PLUGIN_HANDSHAKE_INFO = `${PLUGIN_CORE_PROTOCOL_VERSION}|${PLUGIN_APP_PROTOCOL_VERSION}|${PLUGIN_NETWORK_TYPE}|${PLUGIN_NETWORK_ADDRESS}|${PLUGIN_PROTOCOL}`;
    console.log(PLUGIN_HANDSHAKE_INFO);
    server.start();
});