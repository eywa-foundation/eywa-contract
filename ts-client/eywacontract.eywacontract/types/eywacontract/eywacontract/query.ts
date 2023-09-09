/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { HandshakeType } from "./handshake_type";
import { Params } from "./params";
import { RegisterType } from "./register_type";
import { SendChatType } from "./send_chat_type";

export const protobufPackage = "eywacontract.eywacontract";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetRegisterTypeRequest {
  index: string;
}

export interface QueryGetRegisterTypeResponse {
  registerType: RegisterType | undefined;
}

export interface QueryAllRegisterTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegisterTypeResponse {
  registerType: RegisterType[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSendChatTypeRequest {
  index: string;
}

export interface QueryGetSendChatTypeResponse {
  sendChatType: SendChatType | undefined;
}

export interface QueryAllSendChatTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSendChatTypeResponse {
  sendChatType: SendChatType[];
  pagination: PageResponse | undefined;
}

export interface QueryGetHandshakeTypeRequest {
  index: string;
}

export interface QueryGetHandshakeTypeResponse {
  handshakeType: HandshakeType | undefined;
}

export interface QueryAllHandshakeTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllHandshakeTypeResponse {
  handshakeType: HandshakeType[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRegisterTypeRequest(): QueryGetRegisterTypeRequest {
  return { index: "" };
}

export const QueryGetRegisterTypeRequest = {
  encode(message: QueryGetRegisterTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRegisterTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRegisterTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegisterTypeRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetRegisterTypeRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRegisterTypeRequest>, I>>(object: I): QueryGetRegisterTypeRequest {
    const message = createBaseQueryGetRegisterTypeRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetRegisterTypeResponse(): QueryGetRegisterTypeResponse {
  return { registerType: undefined };
}

export const QueryGetRegisterTypeResponse = {
  encode(message: QueryGetRegisterTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.registerType !== undefined) {
      RegisterType.encode(message.registerType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRegisterTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRegisterTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registerType = RegisterType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegisterTypeResponse {
    return { registerType: isSet(object.registerType) ? RegisterType.fromJSON(object.registerType) : undefined };
  },

  toJSON(message: QueryGetRegisterTypeResponse): unknown {
    const obj: any = {};
    message.registerType !== undefined
      && (obj.registerType = message.registerType ? RegisterType.toJSON(message.registerType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRegisterTypeResponse>, I>>(object: I): QueryGetRegisterTypeResponse {
    const message = createBaseQueryGetRegisterTypeResponse();
    message.registerType = (object.registerType !== undefined && object.registerType !== null)
      ? RegisterType.fromPartial(object.registerType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRegisterTypeRequest(): QueryAllRegisterTypeRequest {
  return { pagination: undefined };
}

export const QueryAllRegisterTypeRequest = {
  encode(message: QueryAllRegisterTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRegisterTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRegisterTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRegisterTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRegisterTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRegisterTypeRequest>, I>>(object: I): QueryAllRegisterTypeRequest {
    const message = createBaseQueryAllRegisterTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRegisterTypeResponse(): QueryAllRegisterTypeResponse {
  return { registerType: [], pagination: undefined };
}

export const QueryAllRegisterTypeResponse = {
  encode(message: QueryAllRegisterTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.registerType) {
      RegisterType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRegisterTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRegisterTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registerType.push(RegisterType.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRegisterTypeResponse {
    return {
      registerType: Array.isArray(object?.registerType)
        ? object.registerType.map((e: any) => RegisterType.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRegisterTypeResponse): unknown {
    const obj: any = {};
    if (message.registerType) {
      obj.registerType = message.registerType.map((e) => e ? RegisterType.toJSON(e) : undefined);
    } else {
      obj.registerType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRegisterTypeResponse>, I>>(object: I): QueryAllRegisterTypeResponse {
    const message = createBaseQueryAllRegisterTypeResponse();
    message.registerType = object.registerType?.map((e) => RegisterType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetSendChatTypeRequest(): QueryGetSendChatTypeRequest {
  return { index: "" };
}

export const QueryGetSendChatTypeRequest = {
  encode(message: QueryGetSendChatTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSendChatTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSendChatTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSendChatTypeRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetSendChatTypeRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSendChatTypeRequest>, I>>(object: I): QueryGetSendChatTypeRequest {
    const message = createBaseQueryGetSendChatTypeRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetSendChatTypeResponse(): QueryGetSendChatTypeResponse {
  return { sendChatType: undefined };
}

export const QueryGetSendChatTypeResponse = {
  encode(message: QueryGetSendChatTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sendChatType !== undefined) {
      SendChatType.encode(message.sendChatType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetSendChatTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetSendChatTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sendChatType = SendChatType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSendChatTypeResponse {
    return { sendChatType: isSet(object.sendChatType) ? SendChatType.fromJSON(object.sendChatType) : undefined };
  },

  toJSON(message: QueryGetSendChatTypeResponse): unknown {
    const obj: any = {};
    message.sendChatType !== undefined
      && (obj.sendChatType = message.sendChatType ? SendChatType.toJSON(message.sendChatType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetSendChatTypeResponse>, I>>(object: I): QueryGetSendChatTypeResponse {
    const message = createBaseQueryGetSendChatTypeResponse();
    message.sendChatType = (object.sendChatType !== undefined && object.sendChatType !== null)
      ? SendChatType.fromPartial(object.sendChatType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSendChatTypeRequest(): QueryAllSendChatTypeRequest {
  return { pagination: undefined };
}

export const QueryAllSendChatTypeRequest = {
  encode(message: QueryAllSendChatTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSendChatTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSendChatTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSendChatTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllSendChatTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSendChatTypeRequest>, I>>(object: I): QueryAllSendChatTypeRequest {
    const message = createBaseQueryAllSendChatTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllSendChatTypeResponse(): QueryAllSendChatTypeResponse {
  return { sendChatType: [], pagination: undefined };
}

export const QueryAllSendChatTypeResponse = {
  encode(message: QueryAllSendChatTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.sendChatType) {
      SendChatType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllSendChatTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllSendChatTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sendChatType.push(SendChatType.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSendChatTypeResponse {
    return {
      sendChatType: Array.isArray(object?.sendChatType)
        ? object.sendChatType.map((e: any) => SendChatType.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllSendChatTypeResponse): unknown {
    const obj: any = {};
    if (message.sendChatType) {
      obj.sendChatType = message.sendChatType.map((e) => e ? SendChatType.toJSON(e) : undefined);
    } else {
      obj.sendChatType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllSendChatTypeResponse>, I>>(object: I): QueryAllSendChatTypeResponse {
    const message = createBaseQueryAllSendChatTypeResponse();
    message.sendChatType = object.sendChatType?.map((e) => SendChatType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetHandshakeTypeRequest(): QueryGetHandshakeTypeRequest {
  return { index: "" };
}

export const QueryGetHandshakeTypeRequest = {
  encode(message: QueryGetHandshakeTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHandshakeTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHandshakeTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHandshakeTypeRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetHandshakeTypeRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHandshakeTypeRequest>, I>>(object: I): QueryGetHandshakeTypeRequest {
    const message = createBaseQueryGetHandshakeTypeRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetHandshakeTypeResponse(): QueryGetHandshakeTypeResponse {
  return { handshakeType: undefined };
}

export const QueryGetHandshakeTypeResponse = {
  encode(message: QueryGetHandshakeTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.handshakeType !== undefined) {
      HandshakeType.encode(message.handshakeType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHandshakeTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHandshakeTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.handshakeType = HandshakeType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHandshakeTypeResponse {
    return { handshakeType: isSet(object.handshakeType) ? HandshakeType.fromJSON(object.handshakeType) : undefined };
  },

  toJSON(message: QueryGetHandshakeTypeResponse): unknown {
    const obj: any = {};
    message.handshakeType !== undefined
      && (obj.handshakeType = message.handshakeType ? HandshakeType.toJSON(message.handshakeType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHandshakeTypeResponse>, I>>(
    object: I,
  ): QueryGetHandshakeTypeResponse {
    const message = createBaseQueryGetHandshakeTypeResponse();
    message.handshakeType = (object.handshakeType !== undefined && object.handshakeType !== null)
      ? HandshakeType.fromPartial(object.handshakeType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHandshakeTypeRequest(): QueryAllHandshakeTypeRequest {
  return { pagination: undefined };
}

export const QueryAllHandshakeTypeRequest = {
  encode(message: QueryAllHandshakeTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHandshakeTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHandshakeTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHandshakeTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllHandshakeTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHandshakeTypeRequest>, I>>(object: I): QueryAllHandshakeTypeRequest {
    const message = createBaseQueryAllHandshakeTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHandshakeTypeResponse(): QueryAllHandshakeTypeResponse {
  return { handshakeType: [], pagination: undefined };
}

export const QueryAllHandshakeTypeResponse = {
  encode(message: QueryAllHandshakeTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.handshakeType) {
      HandshakeType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHandshakeTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHandshakeTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.handshakeType.push(HandshakeType.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHandshakeTypeResponse {
    return {
      handshakeType: Array.isArray(object?.handshakeType)
        ? object.handshakeType.map((e: any) => HandshakeType.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllHandshakeTypeResponse): unknown {
    const obj: any = {};
    if (message.handshakeType) {
      obj.handshakeType = message.handshakeType.map((e) => e ? HandshakeType.toJSON(e) : undefined);
    } else {
      obj.handshakeType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHandshakeTypeResponse>, I>>(
    object: I,
  ): QueryAllHandshakeTypeResponse {
    const message = createBaseQueryAllHandshakeTypeResponse();
    message.handshakeType = object.handshakeType?.map((e) => HandshakeType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of RegisterType items. */
  RegisterType(request: QueryGetRegisterTypeRequest): Promise<QueryGetRegisterTypeResponse>;
  RegisterTypeAll(request: QueryAllRegisterTypeRequest): Promise<QueryAllRegisterTypeResponse>;
  /** Queries a list of SendChatType items. */
  SendChatType(request: QueryGetSendChatTypeRequest): Promise<QueryGetSendChatTypeResponse>;
  SendChatTypeAll(request: QueryAllSendChatTypeRequest): Promise<QueryAllSendChatTypeResponse>;
  /** Queries a list of HandshakeType items. */
  HandshakeType(request: QueryGetHandshakeTypeRequest): Promise<QueryGetHandshakeTypeResponse>;
  HandshakeTypeAll(request: QueryAllHandshakeTypeRequest): Promise<QueryAllHandshakeTypeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.RegisterType = this.RegisterType.bind(this);
    this.RegisterTypeAll = this.RegisterTypeAll.bind(this);
    this.SendChatType = this.SendChatType.bind(this);
    this.SendChatTypeAll = this.SendChatTypeAll.bind(this);
    this.HandshakeType = this.HandshakeType.bind(this);
    this.HandshakeTypeAll = this.HandshakeTypeAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  RegisterType(request: QueryGetRegisterTypeRequest): Promise<QueryGetRegisterTypeResponse> {
    const data = QueryGetRegisterTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "RegisterType", data);
    return promise.then((data) => QueryGetRegisterTypeResponse.decode(new _m0.Reader(data)));
  }

  RegisterTypeAll(request: QueryAllRegisterTypeRequest): Promise<QueryAllRegisterTypeResponse> {
    const data = QueryAllRegisterTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "RegisterTypeAll", data);
    return promise.then((data) => QueryAllRegisterTypeResponse.decode(new _m0.Reader(data)));
  }

  SendChatType(request: QueryGetSendChatTypeRequest): Promise<QueryGetSendChatTypeResponse> {
    const data = QueryGetSendChatTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "SendChatType", data);
    return promise.then((data) => QueryGetSendChatTypeResponse.decode(new _m0.Reader(data)));
  }

  SendChatTypeAll(request: QueryAllSendChatTypeRequest): Promise<QueryAllSendChatTypeResponse> {
    const data = QueryAllSendChatTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "SendChatTypeAll", data);
    return promise.then((data) => QueryAllSendChatTypeResponse.decode(new _m0.Reader(data)));
  }

  HandshakeType(request: QueryGetHandshakeTypeRequest): Promise<QueryGetHandshakeTypeResponse> {
    const data = QueryGetHandshakeTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "HandshakeType", data);
    return promise.then((data) => QueryGetHandshakeTypeResponse.decode(new _m0.Reader(data)));
  }

  HandshakeTypeAll(request: QueryAllHandshakeTypeRequest): Promise<QueryAllHandshakeTypeResponse> {
    const data = QueryAllHandshakeTypeRequest.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Query", "HandshakeTypeAll", data);
    return promise.then((data) => QueryAllHandshakeTypeResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
