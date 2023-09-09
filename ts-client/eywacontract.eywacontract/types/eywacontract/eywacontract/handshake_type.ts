/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "eywacontract.eywacontract";

export interface HandshakeType {
  index: string;
  receiver: string;
  sender: string;
  serverAddress: string;
  roomId: string;
  time: string;
}

function createBaseHandshakeType(): HandshakeType {
  return { index: "", receiver: "", sender: "", serverAddress: "", roomId: "", time: "" };
}

export const HandshakeType = {
  encode(message: HandshakeType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.receiver !== "") {
      writer.uint32(18).string(message.receiver);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    if (message.serverAddress !== "") {
      writer.uint32(34).string(message.serverAddress);
    }
    if (message.roomId !== "") {
      writer.uint32(42).string(message.roomId);
    }
    if (message.time !== "") {
      writer.uint32(50).string(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HandshakeType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHandshakeType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.receiver = reader.string();
          break;
        case 3:
          message.sender = reader.string();
          break;
        case 4:
          message.serverAddress = reader.string();
          break;
        case 5:
          message.roomId = reader.string();
          break;
        case 6:
          message.time = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HandshakeType {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
      serverAddress: isSet(object.serverAddress) ? String(object.serverAddress) : "",
      roomId: isSet(object.roomId) ? String(object.roomId) : "",
      time: isSet(object.time) ? String(object.time) : "",
    };
  },

  toJSON(message: HandshakeType): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.sender !== undefined && (obj.sender = message.sender);
    message.serverAddress !== undefined && (obj.serverAddress = message.serverAddress);
    message.roomId !== undefined && (obj.roomId = message.roomId);
    message.time !== undefined && (obj.time = message.time);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HandshakeType>, I>>(object: I): HandshakeType {
    const message = createBaseHandshakeType();
    message.index = object.index ?? "";
    message.receiver = object.receiver ?? "";
    message.sender = object.sender ?? "";
    message.serverAddress = object.serverAddress ?? "";
    message.roomId = object.roomId ?? "";
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
