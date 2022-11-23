/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Form } from "./forms";
import { Empty } from "./google/protobuf/empty";
import { Value } from "./google/protobuf/struct";
import { RecordPredicate } from "./predicates";

export const protobufPackage = "no.nrc.core.form";

export interface GetFormRequest {
  readonly formId: string;
}

export interface GetFormResponse {
  readonly form: Form | undefined;
}

export interface CreateFormRequest {
  readonly form: Form | undefined;
}

export interface CreateFormResponse {
  readonly form: Form | undefined;
}

export interface UpdateFormRequest {
  readonly form: Form | undefined;
}

export interface UpdateFormResponse {
  readonly form: Form | undefined;
}

export interface DeleteFormRequest {
  readonly formId: string;
}

export interface FieldValue {
  readonly fieldId: string;
  readonly value: any | undefined;
}

export interface Record {
  readonly recordId: string;
  readonly formId: string;
  readonly fieldValues: readonly FieldValue[];
}

export interface GetRecordRequest {
  readonly formId: string;
  readonly recordId: string;
}

export interface GetRecordResponse {
  readonly record: Record | undefined;
}

export interface UpdateRecordRequest {
  readonly record: Record | undefined;
}

export interface UpdateRecordResponse {
  readonly record: Record | undefined;
}

export interface DeleteRecordRequest {
  readonly formId: string;
  readonly recordId: string;
}

export interface SearchRecordsRequest {
  readonly formId: string;
  readonly predicate: RecordPredicate | undefined;
}

export interface ListPagination {
  readonly pageSize: number;
  readonly pageToken: string;
}

export interface SearchRecordsResponse {
  readonly pagination: ListPagination | undefined;
  readonly records: readonly Record[];
}

function createBaseGetFormRequest(): GetFormRequest {
  return { formId: "" };
}

export const GetFormRequest = {
  encode(message: GetFormRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.formId !== "") {
      writer.uint32(10).string(message.formId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetFormRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetFormRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.formId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetFormRequest {
    return { formId: isSet(object.formId) ? String(object.formId) : "" };
  },

  toJSON(message: GetFormRequest): unknown {
    const obj: any = {};
    message.formId !== undefined && (obj.formId = message.formId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetFormRequest>, I>>(object: I): GetFormRequest {
    const message = createBaseGetFormRequest() as any;
    message.formId = object.formId ?? "";
    return message;
  },
};

function createBaseGetFormResponse(): GetFormResponse {
  return { form: undefined };
}

export const GetFormResponse = {
  encode(message: GetFormResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetFormResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetFormResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetFormResponse {
    return { form: isSet(object.form) ? Form.fromJSON(object.form) : undefined };
  },

  toJSON(message: GetFormResponse): unknown {
    const obj: any = {};
    message.form !== undefined && (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetFormResponse>, I>>(object: I): GetFormResponse {
    const message = createBaseGetFormResponse() as any;
    message.form = (object.form !== undefined && object.form !== null) ? Form.fromPartial(object.form) : undefined;
    return message;
  },
};

function createBaseCreateFormRequest(): CreateFormRequest {
  return { form: undefined };
}

export const CreateFormRequest = {
  encode(message: CreateFormRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFormRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFormRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateFormRequest {
    return { form: isSet(object.form) ? Form.fromJSON(object.form) : undefined };
  },

  toJSON(message: CreateFormRequest): unknown {
    const obj: any = {};
    message.form !== undefined && (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateFormRequest>, I>>(object: I): CreateFormRequest {
    const message = createBaseCreateFormRequest() as any;
    message.form = (object.form !== undefined && object.form !== null) ? Form.fromPartial(object.form) : undefined;
    return message;
  },
};

function createBaseCreateFormResponse(): CreateFormResponse {
  return { form: undefined };
}

export const CreateFormResponse = {
  encode(message: CreateFormResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateFormResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateFormResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateFormResponse {
    return { form: isSet(object.form) ? Form.fromJSON(object.form) : undefined };
  },

  toJSON(message: CreateFormResponse): unknown {
    const obj: any = {};
    message.form !== undefined && (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateFormResponse>, I>>(object: I): CreateFormResponse {
    const message = createBaseCreateFormResponse() as any;
    message.form = (object.form !== undefined && object.form !== null) ? Form.fromPartial(object.form) : undefined;
    return message;
  },
};

function createBaseUpdateFormRequest(): UpdateFormRequest {
  return { form: undefined };
}

export const UpdateFormRequest = {
  encode(message: UpdateFormRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFormRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFormRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateFormRequest {
    return { form: isSet(object.form) ? Form.fromJSON(object.form) : undefined };
  },

  toJSON(message: UpdateFormRequest): unknown {
    const obj: any = {};
    message.form !== undefined && (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFormRequest>, I>>(object: I): UpdateFormRequest {
    const message = createBaseUpdateFormRequest() as any;
    message.form = (object.form !== undefined && object.form !== null) ? Form.fromPartial(object.form) : undefined;
    return message;
  },
};

function createBaseUpdateFormResponse(): UpdateFormResponse {
  return { form: undefined };
}

export const UpdateFormResponse = {
  encode(message: UpdateFormResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.form !== undefined) {
      Form.encode(message.form, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateFormResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateFormResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.form = Form.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateFormResponse {
    return { form: isSet(object.form) ? Form.fromJSON(object.form) : undefined };
  },

  toJSON(message: UpdateFormResponse): unknown {
    const obj: any = {};
    message.form !== undefined && (obj.form = message.form ? Form.toJSON(message.form) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateFormResponse>, I>>(object: I): UpdateFormResponse {
    const message = createBaseUpdateFormResponse() as any;
    message.form = (object.form !== undefined && object.form !== null) ? Form.fromPartial(object.form) : undefined;
    return message;
  },
};

function createBaseDeleteFormRequest(): DeleteFormRequest {
  return { formId: "" };
}

export const DeleteFormRequest = {
  encode(message: DeleteFormRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.formId !== "") {
      writer.uint32(10).string(message.formId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteFormRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteFormRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.formId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteFormRequest {
    return { formId: isSet(object.formId) ? String(object.formId) : "" };
  },

  toJSON(message: DeleteFormRequest): unknown {
    const obj: any = {};
    message.formId !== undefined && (obj.formId = message.formId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteFormRequest>, I>>(object: I): DeleteFormRequest {
    const message = createBaseDeleteFormRequest() as any;
    message.formId = object.formId ?? "";
    return message;
  },
};

function createBaseFieldValue(): FieldValue {
  return { fieldId: "", value: undefined };
}

export const FieldValue = {
  encode(message: FieldValue, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fieldId !== "") {
      writer.uint32(10).string(message.fieldId);
    }
    if (message.value !== undefined) {
      Value.encode(Value.wrap(message.value), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FieldValue {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldValue() as any;
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

  fromJSON(object: any): FieldValue {
    return {
      fieldId: isSet(object.fieldId) ? String(object.fieldId) : "",
      value: isSet(object?.value) ? object.value : undefined,
    };
  },

  toJSON(message: FieldValue): unknown {
    const obj: any = {};
    message.fieldId !== undefined && (obj.fieldId = message.fieldId);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FieldValue>, I>>(object: I): FieldValue {
    const message = createBaseFieldValue() as any;
    message.fieldId = object.fieldId ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

function createBaseRecord(): Record {
  return { recordId: "", formId: "", fieldValues: [] };
}

export const Record = {
  encode(message: Record, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.recordId !== "") {
      writer.uint32(10).string(message.recordId);
    }
    if (message.formId !== "") {
      writer.uint32(18).string(message.formId);
    }
    for (const v of message.fieldValues) {
      FieldValue.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Record {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRecord() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.recordId = reader.string();
          break;
        case 2:
          message.formId = reader.string();
          break;
        case 3:
          message.fieldValues.push(FieldValue.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Record {
    return {
      recordId: isSet(object.recordId) ? String(object.recordId) : "",
      formId: isSet(object.formId) ? String(object.formId) : "",
      fieldValues: Array.isArray(object?.fieldValues) ? object.fieldValues.map((e: any) => FieldValue.fromJSON(e)) : [],
    };
  },

  toJSON(message: Record): unknown {
    const obj: any = {};
    message.recordId !== undefined && (obj.recordId = message.recordId);
    message.formId !== undefined && (obj.formId = message.formId);
    if (message.fieldValues) {
      obj.fieldValues = message.fieldValues.map((e) => e ? FieldValue.toJSON(e) : undefined);
    } else {
      obj.fieldValues = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Record>, I>>(object: I): Record {
    const message = createBaseRecord() as any;
    message.recordId = object.recordId ?? "";
    message.formId = object.formId ?? "";
    message.fieldValues = object.fieldValues?.map((e) => FieldValue.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetRecordRequest(): GetRecordRequest {
  return { formId: "", recordId: "" };
}

export const GetRecordRequest = {
  encode(message: GetRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.formId !== "") {
      writer.uint32(10).string(message.formId);
    }
    if (message.recordId !== "") {
      writer.uint32(18).string(message.recordId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRecordRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.formId = reader.string();
          break;
        case 2:
          message.recordId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetRecordRequest {
    return {
      formId: isSet(object.formId) ? String(object.formId) : "",
      recordId: isSet(object.recordId) ? String(object.recordId) : "",
    };
  },

  toJSON(message: GetRecordRequest): unknown {
    const obj: any = {};
    message.formId !== undefined && (obj.formId = message.formId);
    message.recordId !== undefined && (obj.recordId = message.recordId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetRecordRequest>, I>>(object: I): GetRecordRequest {
    const message = createBaseGetRecordRequest() as any;
    message.formId = object.formId ?? "";
    message.recordId = object.recordId ?? "";
    return message;
  },
};

function createBaseGetRecordResponse(): GetRecordResponse {
  return { record: undefined };
}

export const GetRecordResponse = {
  encode(message: GetRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.record !== undefined) {
      Record.encode(message.record, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRecordResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.record = Record.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetRecordResponse {
    return { record: isSet(object.record) ? Record.fromJSON(object.record) : undefined };
  },

  toJSON(message: GetRecordResponse): unknown {
    const obj: any = {};
    message.record !== undefined && (obj.record = message.record ? Record.toJSON(message.record) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetRecordResponse>, I>>(object: I): GetRecordResponse {
    const message = createBaseGetRecordResponse() as any;
    message.record = (object.record !== undefined && object.record !== null)
      ? Record.fromPartial(object.record)
      : undefined;
    return message;
  },
};

function createBaseUpdateRecordRequest(): UpdateRecordRequest {
  return { record: undefined };
}

export const UpdateRecordRequest = {
  encode(message: UpdateRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.record !== undefined) {
      Record.encode(message.record, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateRecordRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.record = Record.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateRecordRequest {
    return { record: isSet(object.record) ? Record.fromJSON(object.record) : undefined };
  },

  toJSON(message: UpdateRecordRequest): unknown {
    const obj: any = {};
    message.record !== undefined && (obj.record = message.record ? Record.toJSON(message.record) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateRecordRequest>, I>>(object: I): UpdateRecordRequest {
    const message = createBaseUpdateRecordRequest() as any;
    message.record = (object.record !== undefined && object.record !== null)
      ? Record.fromPartial(object.record)
      : undefined;
    return message;
  },
};

function createBaseUpdateRecordResponse(): UpdateRecordResponse {
  return { record: undefined };
}

export const UpdateRecordResponse = {
  encode(message: UpdateRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.record !== undefined) {
      Record.encode(message.record, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateRecordResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.record = Record.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateRecordResponse {
    return { record: isSet(object.record) ? Record.fromJSON(object.record) : undefined };
  },

  toJSON(message: UpdateRecordResponse): unknown {
    const obj: any = {};
    message.record !== undefined && (obj.record = message.record ? Record.toJSON(message.record) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateRecordResponse>, I>>(object: I): UpdateRecordResponse {
    const message = createBaseUpdateRecordResponse() as any;
    message.record = (object.record !== undefined && object.record !== null)
      ? Record.fromPartial(object.record)
      : undefined;
    return message;
  },
};

function createBaseDeleteRecordRequest(): DeleteRecordRequest {
  return { formId: "", recordId: "" };
}

export const DeleteRecordRequest = {
  encode(message: DeleteRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.formId !== "") {
      writer.uint32(10).string(message.formId);
    }
    if (message.recordId !== "") {
      writer.uint32(18).string(message.recordId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRecordRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.formId = reader.string();
          break;
        case 2:
          message.recordId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteRecordRequest {
    return {
      formId: isSet(object.formId) ? String(object.formId) : "",
      recordId: isSet(object.recordId) ? String(object.recordId) : "",
    };
  },

  toJSON(message: DeleteRecordRequest): unknown {
    const obj: any = {};
    message.formId !== undefined && (obj.formId = message.formId);
    message.recordId !== undefined && (obj.recordId = message.recordId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteRecordRequest>, I>>(object: I): DeleteRecordRequest {
    const message = createBaseDeleteRecordRequest() as any;
    message.formId = object.formId ?? "";
    message.recordId = object.recordId ?? "";
    return message;
  },
};

function createBaseSearchRecordsRequest(): SearchRecordsRequest {
  return { formId: "", predicate: undefined };
}

export const SearchRecordsRequest = {
  encode(message: SearchRecordsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.formId !== "") {
      writer.uint32(10).string(message.formId);
    }
    if (message.predicate !== undefined) {
      RecordPredicate.encode(message.predicate, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchRecordsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchRecordsRequest() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.formId = reader.string();
          break;
        case 2:
          message.predicate = RecordPredicate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchRecordsRequest {
    return {
      formId: isSet(object.formId) ? String(object.formId) : "",
      predicate: isSet(object.predicate) ? RecordPredicate.fromJSON(object.predicate) : undefined,
    };
  },

  toJSON(message: SearchRecordsRequest): unknown {
    const obj: any = {};
    message.formId !== undefined && (obj.formId = message.formId);
    message.predicate !== undefined &&
      (obj.predicate = message.predicate ? RecordPredicate.toJSON(message.predicate) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchRecordsRequest>, I>>(object: I): SearchRecordsRequest {
    const message = createBaseSearchRecordsRequest() as any;
    message.formId = object.formId ?? "";
    message.predicate = (object.predicate !== undefined && object.predicate !== null)
      ? RecordPredicate.fromPartial(object.predicate)
      : undefined;
    return message;
  },
};

function createBaseListPagination(): ListPagination {
  return { pageSize: 0, pageToken: "" };
}

export const ListPagination = {
  encode(message: ListPagination, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pageSize !== 0) {
      writer.uint32(8).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(18).string(message.pageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPagination {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPagination() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pageSize = reader.int32();
          break;
        case 2:
          message.pageToken = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListPagination {
    return {
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? String(object.pageToken) : "",
    };
  },

  toJSON(message: ListPagination): unknown {
    const obj: any = {};
    message.pageSize !== undefined && (obj.pageSize = Math.round(message.pageSize));
    message.pageToken !== undefined && (obj.pageToken = message.pageToken);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListPagination>, I>>(object: I): ListPagination {
    const message = createBaseListPagination() as any;
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseSearchRecordsResponse(): SearchRecordsResponse {
  return { pagination: undefined, records: [] };
}

export const SearchRecordsResponse = {
  encode(message: SearchRecordsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      ListPagination.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.records) {
      Record.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchRecordsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchRecordsResponse() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = ListPagination.decode(reader, reader.uint32());
          break;
        case 2:
          message.records.push(Record.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchRecordsResponse {
    return {
      pagination: isSet(object.pagination) ? ListPagination.fromJSON(object.pagination) : undefined,
      records: Array.isArray(object?.records) ? object.records.map((e: any) => Record.fromJSON(e)) : [],
    };
  },

  toJSON(message: SearchRecordsResponse): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination ? ListPagination.toJSON(message.pagination) : undefined);
    if (message.records) {
      obj.records = message.records.map((e) => e ? Record.toJSON(e) : undefined);
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SearchRecordsResponse>, I>>(object: I): SearchRecordsResponse {
    const message = createBaseSearchRecordsResponse() as any;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? ListPagination.fromPartial(object.pagination)
      : undefined;
    message.records = object.records?.map((e) => Record.fromPartial(e)) || [];
    return message;
  },
};

export interface FormService {
  GetForm(request: GetFormRequest): Promise<GetFormResponse>;
  CreateForm(request: CreateFormRequest): Promise<CreateFormResponse>;
  UpdateForm(request: UpdateFormRequest): Promise<UpdateFormResponse>;
  DeleteForm(request: DeleteFormRequest): Promise<Empty>;
}

export class FormServiceClientImpl implements FormService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "no.nrc.core.form.FormService";
    this.rpc = rpc;
    this.GetForm = this.GetForm.bind(this);
    this.CreateForm = this.CreateForm.bind(this);
    this.UpdateForm = this.UpdateForm.bind(this);
    this.DeleteForm = this.DeleteForm.bind(this);
  }
  GetForm(request: GetFormRequest): Promise<GetFormResponse> {
    const data = GetFormRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetForm", data);
    return promise.then((data) => GetFormResponse.decode(new _m0.Reader(data)));
  }

  CreateForm(request: CreateFormRequest): Promise<CreateFormResponse> {
    const data = CreateFormRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateForm", data);
    return promise.then((data) => CreateFormResponse.decode(new _m0.Reader(data)));
  }

  UpdateForm(request: UpdateFormRequest): Promise<UpdateFormResponse> {
    const data = UpdateFormRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateForm", data);
    return promise.then((data) => UpdateFormResponse.decode(new _m0.Reader(data)));
  }

  DeleteForm(request: DeleteFormRequest): Promise<Empty> {
    const data = DeleteFormRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteForm", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }
}

export interface RecordService {
  GetRecord(request: GetRecordRequest): Promise<GetRecordResponse>;
  UpdateRecord(request: UpdateRecordRequest): Promise<UpdateRecordResponse>;
  DeleteRecord(request: DeleteRecordRequest): Promise<Empty>;
  SearchRecords(request: SearchRecordsRequest): Promise<SearchRecordsResponse>;
}

export class RecordServiceClientImpl implements RecordService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "no.nrc.core.form.RecordService";
    this.rpc = rpc;
    this.GetRecord = this.GetRecord.bind(this);
    this.UpdateRecord = this.UpdateRecord.bind(this);
    this.DeleteRecord = this.DeleteRecord.bind(this);
    this.SearchRecords = this.SearchRecords.bind(this);
  }
  GetRecord(request: GetRecordRequest): Promise<GetRecordResponse> {
    const data = GetRecordRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetRecord", data);
    return promise.then((data) => GetRecordResponse.decode(new _m0.Reader(data)));
  }

  UpdateRecord(request: UpdateRecordRequest): Promise<UpdateRecordResponse> {
    const data = UpdateRecordRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateRecord", data);
    return promise.then((data) => UpdateRecordResponse.decode(new _m0.Reader(data)));
  }

  DeleteRecord(request: DeleteRecordRequest): Promise<Empty> {
    const data = DeleteRecordRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteRecord", data);
    return promise.then((data) => Empty.decode(new _m0.Reader(data)));
  }

  SearchRecords(request: SearchRecordsRequest): Promise<SearchRecordsResponse> {
    const data = SearchRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SearchRecords", data);
    return promise.then((data) => SearchRecordsResponse.decode(new _m0.Reader(data)));
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
