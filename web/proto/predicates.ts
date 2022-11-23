/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Any } from "./google/protobuf/any";
import { Value } from "./google/protobuf/struct";

export const protobufPackage = "no.nrc.core.form";

export interface RecordPredicate {
  readonly predicate: Any | undefined;
}

export interface AllOfPredicate {
  readonly predicates: readonly RecordPredicate[];
}

export interface AnyOfPredicate {
  readonly predicates: readonly RecordPredicate[];
}

export interface NotPredicate {
  readonly predicate: RecordPredicate | undefined;
}

export interface EqualityPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface InequalityPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface ContainsPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface StartsWithPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface EndsWithPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface GreaterThanPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface GreaterThanOrEqualPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface LessThanPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface LessThanOrEqualPredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface LikePredicate {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface RangePredicate {
  readonly fieldId: string;
  readonly min: any | undefined;
  readonly max: any | undefined;
}

export interface InPredicate {
  readonly fieldId: string;
  readonly values: readonly any[];
}

function createBaseRecordPredicate(): RecordPredicate {
  return { predicate: undefined };
}

export const RecordPredicate = {
  encode(message: RecordPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.predicate !== undefined) {
      Any.encode(message.predicate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RecordPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRecordPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.predicate = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RecordPredicate {
    return { predicate: isSet(object.predicate) ? Any.fromJSON(object.predicate) : undefined };
  },

  toJSON(message: RecordPredicate): unknown {
    const obj: any = {};
    message.predicate !== undefined && (obj.predicate = message.predicate ? Any.toJSON(message.predicate) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RecordPredicate>, I>>(object: I): RecordPredicate {
    const message = createBaseRecordPredicate() as any;
    message.predicate = (object.predicate !== undefined && object.predicate !== null)
      ? Any.fromPartial(object.predicate)
      : undefined;
    return message;
  },
};

function createBaseAllOfPredicate(): AllOfPredicate {
  return { predicates: [] };
}

export const AllOfPredicate = {
  encode(message: AllOfPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.predicates) {
      RecordPredicate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AllOfPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAllOfPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.predicates.push(RecordPredicate.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AllOfPredicate {
    return {
      predicates: Array.isArray(object?.predicates)
        ? object.predicates.map((e: any) => RecordPredicate.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AllOfPredicate): unknown {
    const obj: any = {};
    if (message.predicates) {
      obj.predicates = message.predicates.map((e) => e ? RecordPredicate.toJSON(e) : undefined);
    } else {
      obj.predicates = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AllOfPredicate>, I>>(object: I): AllOfPredicate {
    const message = createBaseAllOfPredicate() as any;
    message.predicates = object.predicates?.map((e) => RecordPredicate.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAnyOfPredicate(): AnyOfPredicate {
  return { predicates: [] };
}

export const AnyOfPredicate = {
  encode(message: AnyOfPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.predicates) {
      RecordPredicate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AnyOfPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAnyOfPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.predicates.push(RecordPredicate.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AnyOfPredicate {
    return {
      predicates: Array.isArray(object?.predicates)
        ? object.predicates.map((e: any) => RecordPredicate.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AnyOfPredicate): unknown {
    const obj: any = {};
    if (message.predicates) {
      obj.predicates = message.predicates.map((e) => e ? RecordPredicate.toJSON(e) : undefined);
    } else {
      obj.predicates = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AnyOfPredicate>, I>>(object: I): AnyOfPredicate {
    const message = createBaseAnyOfPredicate() as any;
    message.predicates = object.predicates?.map((e) => RecordPredicate.fromPartial(e)) || [];
    return message;
  },
};

function createBaseNotPredicate(): NotPredicate {
  return { predicate: undefined };
}

export const NotPredicate = {
  encode(message: NotPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.predicate !== undefined) {
      RecordPredicate.encode(message.predicate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NotPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNotPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.predicate = RecordPredicate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NotPredicate {
    return { predicate: isSet(object.predicate) ? RecordPredicate.fromJSON(object.predicate) : undefined };
  },

  toJSON(message: NotPredicate): unknown {
    const obj: any = {};
    message.predicate !== undefined &&
      (obj.predicate = message.predicate ? RecordPredicate.toJSON(message.predicate) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NotPredicate>, I>>(object: I): NotPredicate {
    const message = createBaseNotPredicate() as any;
    message.predicate = (object.predicate !== undefined && object.predicate !== null)
      ? RecordPredicate.fromPartial(object.predicate)
      : undefined;
    return message;
  },
};

function createBaseEqualityPredicate(): EqualityPredicate {
  return { fieldId: "", value: undefined };
}

export const EqualityPredicate = {
  encode(message: EqualityPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EqualityPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEqualityPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EqualityPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: EqualityPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EqualityPredicate>, I>>(object: I): EqualityPredicate {
    const message = createBaseEqualityPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseInequalityPredicate(): InequalityPredicate {
  return { fieldId: "", value: undefined };
}

export const InequalityPredicate = {
  encode(message: InequalityPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InequalityPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInequalityPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InequalityPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: InequalityPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InequalityPredicate>, I>>(object: I): InequalityPredicate {
    const message = createBaseInequalityPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseContainsPredicate(): ContainsPredicate {
  return { fieldId: "", value: undefined };
}

export const ContainsPredicate = {
  encode(message: ContainsPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContainsPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContainsPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContainsPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: ContainsPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ContainsPredicate>, I>>(object: I): ContainsPredicate {
    const message = createBaseContainsPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseStartsWithPredicate(): StartsWithPredicate {
  return { fieldId: "", value: undefined };
}

export const StartsWithPredicate = {
  encode(message: StartsWithPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartsWithPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartsWithPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StartsWithPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: StartsWithPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StartsWithPredicate>, I>>(object: I): StartsWithPredicate {
    const message = createBaseStartsWithPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseEndsWithPredicate(): EndsWithPredicate {
  return { fieldId: "", value: undefined };
}

export const EndsWithPredicate = {
  encode(message: EndsWithPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EndsWithPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEndsWithPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EndsWithPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: EndsWithPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EndsWithPredicate>, I>>(object: I): EndsWithPredicate {
    const message = createBaseEndsWithPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseGreaterThanPredicate(): GreaterThanPredicate {
  return { fieldId: "", value: undefined };
}

export const GreaterThanPredicate = {
  encode(message: GreaterThanPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GreaterThanPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGreaterThanPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GreaterThanPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: GreaterThanPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GreaterThanPredicate>, I>>(object: I): GreaterThanPredicate {
    const message = createBaseGreaterThanPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseGreaterThanOrEqualPredicate(): GreaterThanOrEqualPredicate {
  return { fieldId: "", value: undefined };
}

export const GreaterThanOrEqualPredicate = {
  encode(message: GreaterThanOrEqualPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GreaterThanOrEqualPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGreaterThanOrEqualPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GreaterThanOrEqualPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: GreaterThanOrEqualPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GreaterThanOrEqualPredicate>, I>>(object: I): GreaterThanOrEqualPredicate {
    const message = createBaseGreaterThanOrEqualPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseLessThanPredicate(): LessThanPredicate {
  return { fieldId: "", value: undefined };
}

export const LessThanPredicate = {
  encode(message: LessThanPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LessThanPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLessThanPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LessThanPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: LessThanPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LessThanPredicate>, I>>(object: I): LessThanPredicate {
    const message = createBaseLessThanPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseLessThanOrEqualPredicate(): LessThanOrEqualPredicate {
  return { fieldId: "", value: undefined };
}

export const LessThanOrEqualPredicate = {
  encode(message: LessThanOrEqualPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LessThanOrEqualPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLessThanOrEqualPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LessThanOrEqualPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: LessThanOrEqualPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LessThanOrEqualPredicate>, I>>(object: I): LessThanOrEqualPredicate {
    const message = createBaseLessThanOrEqualPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseLikePredicate(): LikePredicate {
  return { fieldId: "", value: undefined };
}

export const LikePredicate = {
  encode(message: LikePredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LikePredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLikePredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.value = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LikePredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: LikePredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LikePredicate>, I>>(object: I): LikePredicate {
    const message = createBaseLikePredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseRangePredicate(): RangePredicate {
  return { fieldId: "", min: undefined, max: undefined };
}

export const RangePredicate = {
  encode(message: RangePredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.min !== undefined) {
      Value.encode(Value.wrap(message.min), writer.uint32(18).fork()).ldelim();
    }
    if (message.max !== undefined) {
      Value.encode(Value.wrap(message.max), writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RangePredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRangePredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.min = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        case 3:
          message.max = Value.unwrap(Value.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RangePredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      min: isSet(object?.min) ? object.min : undefined,
      max: isSet(object?.max) ? object.max : undefined,
    };
  },

  toJSON(message: RangePredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.min !== undefined && (obj.min = message.min);
    message.max !== undefined && (obj.max = message.max);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RangePredicate>, I>>(object: I): RangePredicate {
    const message = createBaseRangePredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.min = object.min ?? undefined;
    message.max = object.max ?? undefined;
    return message;
  },
};

function createBaseInPredicate(): InPredicate {
  return { fieldId: "", values: [] };
}

export const InPredicate = {
  encode(message: InPredicate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    for (const v of message.values) {
      Value.encode(Value.wrap(v!), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InPredicate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInPredicate() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fieldId = reader.string();
          break;
        case 2:
          message.values.push(Value.unwrap(Value.decode(reader, reader.uint32())));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InPredicate {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      values: Array.isArray(object?.values) ? [...object.values] : [],
    };
  },

  toJSON(message: InPredicate): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    if (message.values) {
      obj.values = message.values.map((e) => e);
    } else {
      obj.values = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InPredicate>, I>>(object: I): InPredicate {
    const message = createBaseInPredicate() as any;
    message.fieldId = object.fieldId ?? "";
    message.values = object.values?.map((e) => e) || [];
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
