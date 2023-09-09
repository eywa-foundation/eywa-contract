/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "eywacontract.eywacontract";

export interface RegisterType {
  index: string;
  submitter: string;
  pubkey: string;
}

function createBaseRegisterType(): RegisterType {
  return { index: "", submitter: "", pubkey: "" };
}

export const RegisterType = {
  encode(message: RegisterType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.submitter !== "") {
      writer.uint32(18).string(message.submitter);
    }
    if (message.pubkey !== "") {
      writer.uint32(26).string(message.pubkey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.submitter = reader.string();
          break;
        case 3:
          message.pubkey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterType {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      submitter: isSet(object.submitter) ? String(object.submitter) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
    };
  },

  toJSON(message: RegisterType): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.submitter !== undefined && (obj.submitter = message.submitter);
    message.pubkey !== undefined && (obj.pubkey = message.pubkey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterType>, I>>(object: I): RegisterType {
    const message = createBaseRegisterType();
    message.index = object.index ?? "";
    message.submitter = object.submitter ?? "";
    message.pubkey = object.pubkey ?? "";
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
