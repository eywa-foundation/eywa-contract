/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { HandshakeType } from "./handshake_type";
import { Params } from "./params";
import { RegisterType } from "./register_type";
import { SendChatType } from "./send_chat_type";

export const protobufPackage = "eywacontract.eywacontract";

/** GenesisState defines the eywacontract module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  registerTypeList: RegisterType[];
  sendChatTypeList: SendChatType[];
  handshakeTypeList: HandshakeType[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, registerTypeList: [], sendChatTypeList: [], handshakeTypeList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.registerTypeList) {
      RegisterType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.sendChatTypeList) {
      SendChatType.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.handshakeTypeList) {
      HandshakeType.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.registerTypeList.push(RegisterType.decode(reader, reader.uint32()));
          break;
        case 3:
          message.sendChatTypeList.push(SendChatType.decode(reader, reader.uint32()));
          break;
        case 4:
          message.handshakeTypeList.push(HandshakeType.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      registerTypeList: Array.isArray(object?.registerTypeList)
        ? object.registerTypeList.map((e: any) => RegisterType.fromJSON(e))
        : [],
      sendChatTypeList: Array.isArray(object?.sendChatTypeList)
        ? object.sendChatTypeList.map((e: any) => SendChatType.fromJSON(e))
        : [],
      handshakeTypeList: Array.isArray(object?.handshakeTypeList)
        ? object.handshakeTypeList.map((e: any) => HandshakeType.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.registerTypeList) {
      obj.registerTypeList = message.registerTypeList.map((e) => e ? RegisterType.toJSON(e) : undefined);
    } else {
      obj.registerTypeList = [];
    }
    if (message.sendChatTypeList) {
      obj.sendChatTypeList = message.sendChatTypeList.map((e) => e ? SendChatType.toJSON(e) : undefined);
    } else {
      obj.sendChatTypeList = [];
    }
    if (message.handshakeTypeList) {
      obj.handshakeTypeList = message.handshakeTypeList.map((e) => e ? HandshakeType.toJSON(e) : undefined);
    } else {
      obj.handshakeTypeList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.registerTypeList = object.registerTypeList?.map((e) => RegisterType.fromPartial(e)) || [];
    message.sendChatTypeList = object.sendChatTypeList?.map((e) => SendChatType.fromPartial(e)) || [];
    message.handshakeTypeList = object.handshakeTypeList?.map((e) => HandshakeType.fromPartial(e)) || [];
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
