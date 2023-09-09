/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "eywacontract.eywacontract";

export interface MsgRegister {
  creator: string;
  pubkey: string;
}

export interface MsgRegisterResponse {
}

export interface MsgSendChat {
  creator: string;
  chat: string;
  receiver: string;
  time: string;
}

export interface MsgSendChatResponse {
}

export interface MsgHandshake {
  creator: string;
  receiver: string;
  serverAddress: string;
  roomId: string;
  time: string;
}

export interface MsgHandshakeResponse {
}

function createBaseMsgRegister(): MsgRegister {
  return { creator: "", pubkey: "" };
}

export const MsgRegister = {
  encode(message: MsgRegister, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.pubkey !== "") {
      writer.uint32(18).string(message.pubkey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegister {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegister();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegister {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
    };
  },

  toJSON(message: MsgRegister): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegister>, I>>(object: I): MsgRegister {
    const message = createBaseMsgRegister();
    message.creator = object.creator ?? "";
    message.pubkey = object.pubkey ?? "";
    return message;
  },
};

function createBaseMsgRegisterResponse(): MsgRegisterResponse {
  return {};
}

export const MsgRegisterResponse = {
  encode(_: MsgRegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterResponse();
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

  fromJSON(_: any): MsgRegisterResponse {
    return {};
  },

  toJSON(_: MsgRegisterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterResponse>, I>>(_: I): MsgRegisterResponse {
    const message = createBaseMsgRegisterResponse();
    return message;
  },
};

function createBaseMsgSendChat(): MsgSendChat {
  return { creator: "", chat: "", receiver: "", time: "" };
}

export const MsgSendChat = {
  encode(message: MsgSendChat, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chat !== "") {
      writer.uint32(18).string(message.chat);
    }
    if (message.receiver !== "") {
      writer.uint32(26).string(message.receiver);
    }
    if (message.time !== "") {
      writer.uint32(34).string(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendChat {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendChat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chat = reader.string();
          break;
        case 3:
          message.receiver = reader.string();
          break;
        case 4:
          message.time = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendChat {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chat: isSet(object.chat) ? String(object.chat) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      time: isSet(object.time) ? String(object.time) : "",
    };
  },

  toJSON(message: MsgSendChat): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chat !== undefined && (obj.chat = message.chat);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.time !== undefined && (obj.time = message.time);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendChat>, I>>(object: I): MsgSendChat {
    const message = createBaseMsgSendChat();
    message.creator = object.creator ?? "";
    message.chat = object.chat ?? "";
    message.receiver = object.receiver ?? "";
    message.time = object.time ?? "";
    return message;
  },
};

function createBaseMsgSendChatResponse(): MsgSendChatResponse {
  return {};
}

export const MsgSendChatResponse = {
  encode(_: MsgSendChatResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendChatResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendChatResponse();
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

  fromJSON(_: any): MsgSendChatResponse {
    return {};
  },

  toJSON(_: MsgSendChatResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendChatResponse>, I>>(_: I): MsgSendChatResponse {
    const message = createBaseMsgSendChatResponse();
    return message;
  },
};

function createBaseMsgHandshake(): MsgHandshake {
  return { creator: "", receiver: "", serverAddress: "", roomId: "", time: "" };
}

export const MsgHandshake = {
  encode(message: MsgHandshake, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.serverAddress !== "") {
      writer.uint32(26).string(message.serverAddress);
    }
    if (message.roomId !== "") {
      writer.uint32(34).string(message.roomId);
    }
    if (message.time !== "") {
      writer.uint32(42).string(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgHandshake {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgHandshake();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.serverAddress = reader.string();
          break;
        case 4:
          message.roomId = reader.string();
          break;
        case 5:
          message.time = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgHandshake {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      serverAddress: isSet(object.serverAddress) ? String(object.serverAddress) : "",
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      time: isSet(object.time) ? String(object.time) : "",
    };
  },

  toJSON(message: MsgHandshake): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.serverAddress !== undefined && (obj.serverAddress = message.serverAddress);
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.time !== undefined && (obj.time = message.time);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgHandshake>, I>>(object: I): MsgHandshake {
    const message = createBaseMsgHandshake();
    message.creator = object.creator ?? "";
    message.receiver = object.receiver ?? "";
    message.serverAddress = object.serverAddress ?? "";
    message.roomId = object.roomId ?? "";
    message.time = object.time ?? "";
    return message;
  },
};

function createBaseMsgHandshakeResponse(): MsgHandshakeResponse {
  return {};
}

export const MsgHandshakeResponse = {
  encode(_: MsgHandshakeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgHandshakeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgHandshakeResponse();
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

  fromJSON(_: any): MsgHandshakeResponse {
    return {};
  },

  toJSON(_: MsgHandshakeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgHandshakeResponse>, I>>(_: I): MsgHandshakeResponse {
    const message = createBaseMsgHandshakeResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Register(request: MsgRegister): Promise<MsgRegisterResponse>;
  SendChat(request: MsgSendChat): Promise<MsgSendChatResponse>;
  Handshake(request: MsgHandshake): Promise<MsgHandshakeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Register = this.Register.bind(this);
    this.SendChat = this.SendChat.bind(this);
    this.Handshake = this.Handshake.bind(this);
  }
  Register(request: MsgRegister): Promise<MsgRegisterResponse> {
    const data = MsgRegister.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Msg", "Register", data);
    return promise.then((data) => MsgRegisterResponse.decode(new _m0.Reader(data)));
  }

  SendChat(request: MsgSendChat): Promise<MsgSendChatResponse> {
    const data = MsgSendChat.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Msg", "SendChat", data);
    return promise.then((data) => MsgSendChatResponse.decode(new _m0.Reader(data)));
  }

  Handshake(request: MsgHandshake): Promise<MsgHandshakeResponse> {
    const data = MsgHandshake.encode(request).finish();
    const promise = this.rpc.request("eywacontract.eywacontract.Msg", "Handshake", data);
    return promise.then((data) => MsgHandshakeResponse.decode(new _m0.Reader(data)));
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
