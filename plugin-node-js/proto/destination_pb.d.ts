// package: proto
// file: destination.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class BatchUpsertUsersRequest extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<User>;
    setUsersList(value: Array<User>): BatchUpsertUsersRequest;
    addUsers(value?: User, index?: number): User;
    getMappings(): Uint8Array | string;
    getMappings_asU8(): Uint8Array;
    getMappings_asB64(): string;
    setMappings(value: Uint8Array | string): BatchUpsertUsersRequest;
    getSettings(): Uint8Array | string;
    getSettings_asU8(): Uint8Array;
    getSettings_asB64(): string;
    setSettings(value: Uint8Array | string): BatchUpsertUsersRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchUpsertUsersRequest.AsObject;
    static toObject(includeInstance: boolean, msg: BatchUpsertUsersRequest): BatchUpsertUsersRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchUpsertUsersRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchUpsertUsersRequest;
    static deserializeBinaryFromReader(message: BatchUpsertUsersRequest, reader: jspb.BinaryReader): BatchUpsertUsersRequest;
}

export namespace BatchUpsertUsersRequest {
    export type AsObject = {
        usersList: Array<User.AsObject>,
        mappings: Uint8Array | string,
        settings: Uint8Array | string,
    }
}

export class BatchUpsertUsersResponse extends jspb.Message { 
    getStatus(): OperationStatus;
    setStatus(value: OperationStatus): BatchUpsertUsersResponse;
    clearErrorsList(): void;
    getErrorsList(): Array<ResponseError>;
    setErrorsList(value: Array<ResponseError>): BatchUpsertUsersResponse;
    addErrors(value?: ResponseError, index?: number): ResponseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchUpsertUsersResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BatchUpsertUsersResponse): BatchUpsertUsersResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchUpsertUsersResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchUpsertUsersResponse;
    static deserializeBinaryFromReader(message: BatchUpsertUsersResponse, reader: jspb.BinaryReader): BatchUpsertUsersResponse;
}

export namespace BatchUpsertUsersResponse {
    export type AsObject = {
        status: OperationStatus,
        errorsList: Array<ResponseError.AsObject>,
    }
}

export class BatchUpdateUsersRequest extends jspb.Message { 
    clearUsersList(): void;
    getUsersList(): Array<User>;
    setUsersList(value: Array<User>): BatchUpdateUsersRequest;
    addUsers(value?: User, index?: number): User;
    getMappings(): Uint8Array | string;
    getMappings_asU8(): Uint8Array;
    getMappings_asB64(): string;
    setMappings(value: Uint8Array | string): BatchUpdateUsersRequest;
    getSettings(): Uint8Array | string;
    getSettings_asU8(): Uint8Array;
    getSettings_asB64(): string;
    setSettings(value: Uint8Array | string): BatchUpdateUsersRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchUpdateUsersRequest.AsObject;
    static toObject(includeInstance: boolean, msg: BatchUpdateUsersRequest): BatchUpdateUsersRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchUpdateUsersRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchUpdateUsersRequest;
    static deserializeBinaryFromReader(message: BatchUpdateUsersRequest, reader: jspb.BinaryReader): BatchUpdateUsersRequest;
}

export namespace BatchUpdateUsersRequest {
    export type AsObject = {
        usersList: Array<User.AsObject>,
        mappings: Uint8Array | string,
        settings: Uint8Array | string,
    }
}

export class BatchUpdateUsersResponse extends jspb.Message { 
    getStatus(): OperationStatus;
    setStatus(value: OperationStatus): BatchUpdateUsersResponse;
    clearErrorsList(): void;
    getErrorsList(): Array<ResponseError>;
    setErrorsList(value: Array<ResponseError>): BatchUpdateUsersResponse;
    addErrors(value?: ResponseError, index?: number): ResponseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchUpdateUsersResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BatchUpdateUsersResponse): BatchUpdateUsersResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchUpdateUsersResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchUpdateUsersResponse;
    static deserializeBinaryFromReader(message: BatchUpdateUsersResponse, reader: jspb.BinaryReader): BatchUpdateUsersResponse;
}

export namespace BatchUpdateUsersResponse {
    export type AsObject = {
        status: OperationStatus,
        errorsList: Array<ResponseError.AsObject>,
    }
}

export class BatchDeleteUsersRequest extends jspb.Message { 
    clearUserIdsList(): void;
    getUserIdsList(): Array<string>;
    setUserIdsList(value: Array<string>): BatchDeleteUsersRequest;
    addUserIds(value: string, index?: number): string;
    getSettings(): Uint8Array | string;
    getSettings_asU8(): Uint8Array;
    getSettings_asB64(): string;
    setSettings(value: Uint8Array | string): BatchDeleteUsersRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchDeleteUsersRequest.AsObject;
    static toObject(includeInstance: boolean, msg: BatchDeleteUsersRequest): BatchDeleteUsersRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchDeleteUsersRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchDeleteUsersRequest;
    static deserializeBinaryFromReader(message: BatchDeleteUsersRequest, reader: jspb.BinaryReader): BatchDeleteUsersRequest;
}

export namespace BatchDeleteUsersRequest {
    export type AsObject = {
        userIdsList: Array<string>,
        settings: Uint8Array | string,
    }
}

export class BatchDeleteUsersResponse extends jspb.Message { 
    getStatus(): OperationStatus;
    setStatus(value: OperationStatus): BatchDeleteUsersResponse;
    clearErrorsList(): void;
    getErrorsList(): Array<ResponseError>;
    setErrorsList(value: Array<ResponseError>): BatchDeleteUsersResponse;
    addErrors(value?: ResponseError, index?: number): ResponseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BatchDeleteUsersResponse.AsObject;
    static toObject(includeInstance: boolean, msg: BatchDeleteUsersResponse): BatchDeleteUsersResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BatchDeleteUsersResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BatchDeleteUsersResponse;
    static deserializeBinaryFromReader(message: BatchDeleteUsersResponse, reader: jspb.BinaryReader): BatchDeleteUsersResponse;
}

export namespace BatchDeleteUsersResponse {
    export type AsObject = {
        status: OperationStatus,
        errorsList: Array<ResponseError.AsObject>,
    }
}

export class User extends jspb.Message { 
    getId(): string;
    setId(value: string): User;
    getFullName(): string;
    setFullName(value: string): User;
    getFirstName(): string;
    setFirstName(value: string): User;
    getLastName(): string;
    setLastName(value: string): User;
    getEmail(): string;
    setEmail(value: string): User;
    getPhone(): string;
    setPhone(value: string): User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): User.AsObject;
    static toObject(includeInstance: boolean, msg: User): User.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): User;
    static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
    export type AsObject = {
        id: string,
        fullName: string,
        firstName: string,
        lastName: string,
        email: string,
        phone: string,
    }
}

export class ResponseError extends jspb.Message { 
    getCode(): string;
    setCode(value: string): ResponseError;
    getUserId(): string;
    setUserId(value: string): ResponseError;
    getMessage(): string;
    setMessage(value: string): ResponseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResponseError.AsObject;
    static toObject(includeInstance: boolean, msg: ResponseError): ResponseError.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResponseError, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResponseError;
    static deserializeBinaryFromReader(message: ResponseError, reader: jspb.BinaryReader): ResponseError;
}

export namespace ResponseError {
    export type AsObject = {
        code: string,
        userId: string,
        message: string,
    }
}

export enum OperationStatus {
    SUCCESS = 0,
    PARTIAL_FAILURE = 1,
    TOTAL_FAILURE = 2,
}
