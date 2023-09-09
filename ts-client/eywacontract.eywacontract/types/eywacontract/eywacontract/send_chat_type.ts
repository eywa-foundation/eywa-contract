/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "eywacontract.eywacontract";

export interface SendChatType {
  index: string;
  chat: string;
  sender: string;
  receiver: string;
  time: string;
}

function createBaseSendChatType(): SendChatType {
  return { index: "", chat: "", sender: "", receiver: "", time: "" };
}

export const SendChatType = {
  encode(message: SendChatType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.chat !== "") {
      writer.uint32(18).string(message.chat);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    if (message.time !== "") {
      writer.uint32(42).string(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendChatType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendChatType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.chat = reader.string();
          break;
        case 3:
          message.sender = reader.string();
          break;
        case 4:
          message.receiver = reader.string();
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

  fromJSON(object: any): SendChatType {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      chat: isSet(object.chat) ? String(object.chat) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      time: isSet(object.time) ? String(object.time) : "",
    };
  },

  toJSON(message: SendChatType): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.chat !== undefined && (obj.chat = message.chat);
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.time !== undefined && (obj.time = message.time);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SendChatType>, I>>(object: I): SendChatType {
    const message = createBaseSendChatType();
    message.index = object.index ?? "";
    message.chat = object.chat ?? "";
    message.sender = object.sender ?? "";
    message.receiver = object.receiver ?? "";
    message.time = object.time ?? "";
    return message;
  },
};

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
