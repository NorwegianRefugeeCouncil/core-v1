/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Any } from "./google/protobuf/any";
import { Duration } from "./google/protobuf/duration";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "proto";

export interface FormField {
  readonly id: string;
  readonly typedConfig: Any | undefined;
}

export interface FormTextField {
  readonly label: string;
  readonly value?: string | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormDateField {
  readonly label: string;
  readonly value?: Date | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormTimeField {
  readonly label: string;
  readonly value?: Duration | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormSelectOption {
  readonly label: string;
  readonly value?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormSelectField {
  readonly label: string;
  readonly value?: string | undefined;
  readonly options: readonly FormSelectOption[];
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormCheckboxField {
  readonly label: string;
  readonly value?: boolean | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormRadioField {
  readonly label: string;
  readonly value: string;
  readonly options: readonly FormSelectOption[];
}

export interface FormTextareaField {
  readonly label: string;
  readonly value?: string | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
  readonly rows?: number | undefined;
}

export interface FormFileField {
  readonly label: string;
  readonly value?: string | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface FormDurationField {
  readonly label: string;
  readonly value?: Duration | undefined;
  readonly placeholder?: string | undefined;
  readonly disabled?: boolean | undefined;
}

export interface Form {
  readonly id: string;
  readonly title: string;
  readonly description: string;
  readonly fields: readonly FormField[];
  readonly fieldSections: { [key: string]: string };
}

export interface Form_FieldSectionsEntry {
  readonly key: string;
  readonly value: string;
}

export interface FormSection {
  readonly id: string;
  readonly title: string;
  readonly description: string;
  readonly collapsed: boolean;
  readonly collapsible: boolean;
}

function createBaseFormField(): FormField {
  return { id: "", typedConfig: undefined };
}

export const FormField = {
  encode(message: FormField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.typedConfig !== undefined) {
      Any.encode(message.typedConfig, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.typedConfig = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormField {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      typedConfig: isSet(object.typedConfig) ? Any.fromJSON(object.typedConfig) : undefined,
    };
  },

  toJSON(message: FormField): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.typedConfig !== undefined &&
      (obj.typedConfig = message.typedConfig ? Any.toJSON(message.typedConfig) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormField>, I>>(object: I): FormField {
    const message = createBaseFormField() as any;
    message.id = object.id ?? "";
    message.typedConfig = (object.typedConfig !== undefined && object.typedConfig !== null)
      ? Any.fromPartial(object.typedConfig)
      : undefined;
    return message;
  },
};

function createBaseFormTextField(): FormTextField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined };
}

export const FormTextField = {
  encode(message: FormTextField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormTextField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormTextField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormTextField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormTextField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormTextField>, I>>(object: I): FormTextField {
    const message = createBaseFormTextField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormDateField(): FormDateField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined };
}

export const FormDateField = {
  encode(message: FormDateField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      Timestamp.encode(toTimestamp(message.value), writer.uint32(18).fork()).ldelim();
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormDateField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormDateField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormDateField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? fromJsonTimestamp(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormDateField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value.toISOString());
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormDateField>, I>>(object: I): FormDateField {
    const message = createBaseFormDateField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormTimeField(): FormTimeField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined };
}

export const FormTimeField = {
  encode(message: FormTimeField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      Duration.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormTimeField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormTimeField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormTimeField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? Duration.fromJSON(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormTimeField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value ? Duration.toJSON(message.value) : undefined);
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormTimeField>, I>>(object: I): FormTimeField {
    const message = createBaseFormTimeField() as any;
    message.label = object.label ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? Duration.fromPartial(object.value)
      : undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormSelectOption(): FormSelectOption {
  return { label: "", value: undefined, disabled: undefined };
}

export const FormSelectOption = {
  encode(message: FormSelectOption, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    if (message.disabled !== undefined) {
      writer.uint32(24).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormSelectOption {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormSelectOption() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormSelectOption {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormSelectOption): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormSelectOption>, I>>(object: I): FormSelectOption {
    const message = createBaseFormSelectOption() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormSelectField(): FormSelectField {
  return { label: "", value: undefined, options: [], placeholder: undefined, disabled: undefined };
}

export const FormSelectField = {
  encode(message: FormSelectField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    for (const v of message.options) {
      FormSelectOption.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.placeholder !== undefined) {
      writer.uint32(34).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(40).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormSelectField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormSelectField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.options.push(FormSelectOption.decode(reader, reader.uint32()));
          break;
        case 4:
          message.placeholder = reader.string();
          break;
        case 5:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormSelectField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : undefined,
      options: Array.isArray(object?.options) ? object.options.map((e: any) => FormSelectOption.fromJSON(e)) : [],
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormSelectField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    if (message.options) {
      obj.options = message.options.map((e) => e ? FormSelectOption.toJSON(e) : undefined);
    } else {
      obj.options = [];
    }
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormSelectField>, I>>(object: I): FormSelectField {
    const message = createBaseFormSelectField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.options = object.options?.map((e) => FormSelectOption.fromPartial(e)) || [];
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormCheckboxField(): FormCheckboxField {
  return { label: "", value: undefined, disabled: undefined };
}

export const FormCheckboxField = {
  encode(message: FormCheckboxField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(16).bool(message.value);
    }
    if (message.disabled !== undefined) {
      writer.uint32(24).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormCheckboxField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormCheckboxField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.bool();
          break;
        case 3:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormCheckboxField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? Boolean(object.value) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormCheckboxField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormCheckboxField>, I>>(object: I): FormCheckboxField {
    const message = createBaseFormCheckboxField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormRadioField(): FormRadioField {
  return { label: "", value: "", options: [] };
}

export const FormRadioField = {
  encode(message: FormRadioField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    for (const v of message.options) {
      FormSelectOption.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormRadioField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormRadioField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.options.push(FormSelectOption.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormRadioField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : "",
      options: Array.isArray(object?.options) ? object.options.map((e: any) => FormSelectOption.fromJSON(e)) : [],
    };
  },

  toJSON(message: FormRadioField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    if (message.options) {
      obj.options = message.options.map((e) => e ? FormSelectOption.toJSON(e) : undefined);
    } else {
      obj.options = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormRadioField>, I>>(object: I): FormRadioField {
    const message = createBaseFormRadioField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? "";
    message.options = object.options?.map((e) => FormSelectOption.fromPartial(e)) || [];
    return message;
  },
};

function createBaseFormTextareaField(): FormTextareaField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined, rows: undefined };
}

export const FormTextareaField = {
  encode(message: FormTextareaField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    if (message.rows !== undefined) {
      writer.uint32(40).int32(message.rows);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormTextareaField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormTextareaField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        case 5:
          message.rows = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormTextareaField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
      rows: isSet(object.rows) ? Number(object.rows) : undefined,
    };
  },

  toJSON(message: FormTextareaField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    message.rows !== undefined && (obj.rows = Math.round(message.rows));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormTextareaField>, I>>(object: I): FormTextareaField {
    const message = createBaseFormTextareaField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    message.rows = object.rows ?? undefined;
    return message;
  },
};

function createBaseFormFileField(): FormFileField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined };
}

export const FormFileField = {
  encode(message: FormFileField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormFileField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormFileField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormFileField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? String(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormFileField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value);
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormFileField>, I>>(object: I): FormFileField {
    const message = createBaseFormFileField() as any;
    message.label = object.label ?? "";
    message.value = object.value ?? undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseFormDurationField(): FormDurationField {
  return { label: "", value: undefined, placeholder: undefined, disabled: undefined };
}

export const FormDurationField = {
  encode(message: FormDurationField, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.label !== "") {
      writer.uint32(10).string(message.label);
    }
    if (message.value !== undefined) {
      Duration.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    if (message.placeholder !== undefined) {
      writer.uint32(26).string(message.placeholder);
    }
    if (message.disabled !== undefined) {
      writer.uint32(32).bool(message.disabled);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormDurationField {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormDurationField() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.label = reader.string();
          break;
        case 2:
          message.value = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.placeholder = reader.string();
          break;
        case 4:
          message.disabled = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormDurationField {
    return {
      label: isSet(object.label) ? String(object.label) : "",
      value: isSet(object.value) ? Duration.fromJSON(object.value) : undefined,
      placeholder: isSet(object.placeholder) ? String(object.placeholder) : undefined,
      disabled: isSet(object.disabled) ? Boolean(object.disabled) : undefined,
    };
  },

  toJSON(message: FormDurationField): unknown {
    const obj: any = {};
    message.label !== undefined && (obj.label = message.label);
    message.value !== undefined && (obj.value = message.value ? Duration.toJSON(message.value) : undefined);
    message.placeholder !== undefined && (obj.placeholder = message.placeholder);
    message.disabled !== undefined && (obj.disabled = message.disabled);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormDurationField>, I>>(object: I): FormDurationField {
    const message = createBaseFormDurationField() as any;
    message.label = object.label ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? Duration.fromPartial(object.value)
      : undefined;
    message.placeholder = object.placeholder ?? undefined;
    message.disabled = object.disabled ?? undefined;
    return message;
  },
};

function createBaseForm(): Form {
  return { id: "", title: "", description: "", fields: [], fieldSections: {} };
}

export const Form = {
  encode(message: Form, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    for (const v of message.fields) {
      FormField.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    Object.entries(message.fieldSections).forEach(([key, value]) => {
      Form_FieldSectionsEntry.encode({ key: key as any, value }, writer.uint32(42).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Form {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForm() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.fields.push(FormField.decode(reader, reader.uint32()));
          break;
        case 5:
          const entry5 = Form_FieldSectionsEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.fieldSections[entry5.key] = entry5.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Form {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      fields: Array.isArray(object?.fields) ? object.fields.map((e: any) => FormField.fromJSON(e)) : [],
      fieldSections: isObject(object.fieldSections)
        ? Object.entries(object.fieldSections).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: Form): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    if (message.fields) {
      obj.fields = message.fields.map((e) => e ? FormField.toJSON(e) : undefined);
    } else {
      obj.fields = [];
    }
    obj.fieldSections = {};
    if (message.fieldSections) {
      Object.entries(message.fieldSections).forEach(([k, v]) => {
        obj.fieldSections[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Form>, I>>(object: I): Form {
    const message = createBaseForm() as any;
    message.id = object.id ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.fields = object.fields?.map((e) => FormField.fromPartial(e)) || [];
    message.fieldSections = Object.entries(object.fieldSections ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    return message;
  },
};

function createBaseForm_FieldSectionsEntry(): Form_FieldSectionsEntry {
  return { key: "", value: "" };
}

export const Form_FieldSectionsEntry = {
  encode(message: Form_FieldSectionsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Form_FieldSectionsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForm_FieldSectionsEntry() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Form_FieldSectionsEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: Form_FieldSectionsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Form_FieldSectionsEntry>, I>>(object: I): Form_FieldSectionsEntry {
    const message = createBaseForm_FieldSectionsEntry() as any;
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseFormSection(): FormSection {
  return { id: "", title: "", description: "", collapsed: false, collapsible: false };
}

export const FormSection = {
  encode(message: FormSection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.collapsed === true) {
      writer.uint32(32).bool(message.collapsed);
    }
    if (message.collapsible === true) {
      writer.uint32(40).bool(message.collapsible);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FormSection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFormSection() as any;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.collapsed = reader.bool();
          break;
        case 5:
          message.collapsible = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FormSection {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      collapsed: isSet(object.collapsed) ? Boolean(object.collapsed) : false,
      collapsible: isSet(object.collapsible) ? Boolean(object.collapsible) : false,
    };
  },

  toJSON(message: FormSection): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.collapsed !== undefined && (obj.collapsed = message.collapsed);
    message.collapsible !== undefined && (obj.collapsible = message.collapsible);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FormSection>, I>>(object: I): FormSection {
    const message = createBaseFormSection() as any;
    message.id = object.id ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.collapsed = object.collapsed ?? false;
    message.collapsible = object.collapsible ?? false;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
