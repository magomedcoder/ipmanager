// @generated by protoc-gen-es v2.4.0 with parameter "target=ts"
// @generated from file subnet.proto (package subnet, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file subnet.proto.
 */
export const file_subnet: GenFile = /*@__PURE__*/
  fileDesc("CgxzdWJuZXQucHJvdG8SBnN1Ym5ldCJHChNDcmVhdGVTdWJuZXRSZXF1ZXN0EgoKAmlwGAEgASgJEg8KB3ZsYW5faWQYAiABKAMSEwoLZGVzY3JpcHRpb24YAyABKAkiIgoUQ3JlYXRlU3VibmV0UmVzcG9uc2USCgoCaWQYASABKAMiMwoRR2V0U3VibmV0c1JlcXVlc3QSDAoEcGFnZRgBIAEoAxIQCghwYWdlU2l6ZRgCIAEoAyJGChJHZXRTdWJuZXRzUmVzcG9uc2USDQoFdG90YWwYASABKAMSIQoFaXRlbXMYAiADKAsyEi5zdWJuZXQuU3VibmV0SXRlbSKJAQoKU3VibmV0SXRlbRIKCgJpZBgBIAEoAxIKCgJpcBgCIAEoCRIPCgd2bGFuX2lkGAMgASgDEhEKCXZsYW5fbmFtZRgEIAEoCRITCgtjdXN0b21lcl9pZBgFIAEoAxIVCg1jdXN0b21lcl9uYW1lGAYgASgJEhMKC2Rlc2NyaXB0aW9uGAcgASgJIh4KEEdldFN1Ym5ldFJlcXVlc3QSCgoCaWQYASABKAMiggEKEUdldFN1Ym5ldFJlc3BvbnNlEgoKAmlkGAEgASgDEgoKAmlwGAIgASgJEgwKBG5hbWUYAyABKAkSDwoHdmxhbl9pZBgEIAEoAxIRCgl2bGFuX25hbWUYBSABKAkSEwoLZGVzY3JpcHRpb24YBiABKAkSDgoGY2hhcnRzGAcgAygDIjQKFUVkaXRTdWJuZXRWbGFuUmVxdWVzdBIKCgJpZBgBIAEoAxIPCgd2bGFuX2lkGAIgASgDIikKFkVkaXRTdWJuZXRWbGFuUmVzcG9uc2USDwoHc3VjY2VzcxgBIAEoCCI/ChxFZGl0U3VibmV0RGVzY3JpcHRpb25SZXF1ZXN0EgoKAmlkGAEgASgDEhMKC2Rlc2NyaXB0aW9uGAIgASgJIjAKHUVkaXRTdWJuZXREZXNjcmlwdGlvblJlc3BvbnNlEg8KB3N1Y2Nlc3MYASABKAgynAMKDVN1Ym5ldFNlcnZpY2USSQoMQ3JlYXRlU3VibmV0Ehsuc3VibmV0LkNyZWF0ZVN1Ym5ldFJlcXVlc3QaHC5zdWJuZXQuQ3JlYXRlU3VibmV0UmVzcG9uc2USQwoKR2V0U3VibmV0cxIZLnN1Ym5ldC5HZXRTdWJuZXRzUmVxdWVzdBoaLnN1Ym5ldC5HZXRTdWJuZXRzUmVzcG9uc2USRAoNR2V0U3VibmV0QnlJZBIYLnN1Ym5ldC5HZXRTdWJuZXRSZXF1ZXN0Ghkuc3VibmV0LkdldFN1Ym5ldFJlc3BvbnNlEk8KDkVkaXRTdWJuZXRWbGFuEh0uc3VibmV0LkVkaXRTdWJuZXRWbGFuUmVxdWVzdBoeLnN1Ym5ldC5FZGl0U3VibmV0VmxhblJlc3BvbnNlEmQKFUVkaXRTdWJuZXREZXNjcmlwdGlvbhIkLnN1Ym5ldC5FZGl0U3VibmV0RGVzY3JpcHRpb25SZXF1ZXN0GiUuc3VibmV0LkVkaXRTdWJuZXREZXNjcmlwdGlvblJlc3BvbnNlQipaKGdpdGh1Yi5jb20vbWFnb21lZGNvZGVyL2lwbWFuYWdlci9hcGkvcGJiBnByb3RvMw");

/**
 * @generated from message subnet.CreateSubnetRequest
 */
export type CreateSubnetRequest = Message<"subnet.CreateSubnetRequest"> & {
  /**
   * @generated from field: string ip = 1;
   */
  ip: string;

  /**
   * @generated from field: int64 vlan_id = 2;
   */
  vlanId: bigint;

  /**
   * @generated from field: string description = 3;
   */
  description: string;
};

/**
 * Describes the message subnet.CreateSubnetRequest.
 * Use `create(CreateSubnetRequestSchema)` to create a new message.
 */
export const CreateSubnetRequestSchema: GenMessage<CreateSubnetRequest> = /*@__PURE__*/
  messageDesc(file_subnet, 0);

/**
 * @generated from message subnet.CreateSubnetResponse
 */
export type CreateSubnetResponse = Message<"subnet.CreateSubnetResponse"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message subnet.CreateSubnetResponse.
 * Use `create(CreateSubnetResponseSchema)` to create a new message.
 */
export const CreateSubnetResponseSchema: GenMessage<CreateSubnetResponse> = /*@__PURE__*/
  messageDesc(file_subnet, 1);

/**
 * @generated from message subnet.GetSubnetsRequest
 */
export type GetSubnetsRequest = Message<"subnet.GetSubnetsRequest"> & {
  /**
   * @generated from field: int64 page = 1;
   */
  page: bigint;

  /**
   * @generated from field: int64 pageSize = 2;
   */
  pageSize: bigint;
};

/**
 * Describes the message subnet.GetSubnetsRequest.
 * Use `create(GetSubnetsRequestSchema)` to create a new message.
 */
export const GetSubnetsRequestSchema: GenMessage<GetSubnetsRequest> = /*@__PURE__*/
  messageDesc(file_subnet, 2);

/**
 * @generated from message subnet.GetSubnetsResponse
 */
export type GetSubnetsResponse = Message<"subnet.GetSubnetsResponse"> & {
  /**
   * @generated from field: int64 total = 1;
   */
  total: bigint;

  /**
   * @generated from field: repeated subnet.SubnetItem items = 2;
   */
  items: SubnetItem[];
};

/**
 * Describes the message subnet.GetSubnetsResponse.
 * Use `create(GetSubnetsResponseSchema)` to create a new message.
 */
export const GetSubnetsResponseSchema: GenMessage<GetSubnetsResponse> = /*@__PURE__*/
  messageDesc(file_subnet, 3);

/**
 * @generated from message subnet.SubnetItem
 */
export type SubnetItem = Message<"subnet.SubnetItem"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string ip = 2;
   */
  ip: string;

  /**
   * @generated from field: int64 vlan_id = 3;
   */
  vlanId: bigint;

  /**
   * @generated from field: string vlan_name = 4;
   */
  vlanName: string;

  /**
   * @generated from field: int64 customer_id = 5;
   */
  customerId: bigint;

  /**
   * @generated from field: string customer_name = 6;
   */
  customerName: string;

  /**
   * @generated from field: string description = 7;
   */
  description: string;
};

/**
 * Describes the message subnet.SubnetItem.
 * Use `create(SubnetItemSchema)` to create a new message.
 */
export const SubnetItemSchema: GenMessage<SubnetItem> = /*@__PURE__*/
  messageDesc(file_subnet, 4);

/**
 * @generated from message subnet.GetSubnetRequest
 */
export type GetSubnetRequest = Message<"subnet.GetSubnetRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message subnet.GetSubnetRequest.
 * Use `create(GetSubnetRequestSchema)` to create a new message.
 */
export const GetSubnetRequestSchema: GenMessage<GetSubnetRequest> = /*@__PURE__*/
  messageDesc(file_subnet, 5);

/**
 * @generated from message subnet.GetSubnetResponse
 */
export type GetSubnetResponse = Message<"subnet.GetSubnetResponse"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string ip = 2;
   */
  ip: string;

  /**
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * @generated from field: int64 vlan_id = 4;
   */
  vlanId: bigint;

  /**
   * @generated from field: string vlan_name = 5;
   */
  vlanName: string;

  /**
   * @generated from field: string description = 6;
   */
  description: string;

  /**
   * @generated from field: repeated int64 charts = 7;
   */
  charts: bigint[];
};

/**
 * Describes the message subnet.GetSubnetResponse.
 * Use `create(GetSubnetResponseSchema)` to create a new message.
 */
export const GetSubnetResponseSchema: GenMessage<GetSubnetResponse> = /*@__PURE__*/
  messageDesc(file_subnet, 6);

/**
 * @generated from message subnet.EditSubnetVlanRequest
 */
export type EditSubnetVlanRequest = Message<"subnet.EditSubnetVlanRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: int64 vlan_id = 2;
   */
  vlanId: bigint;
};

/**
 * Describes the message subnet.EditSubnetVlanRequest.
 * Use `create(EditSubnetVlanRequestSchema)` to create a new message.
 */
export const EditSubnetVlanRequestSchema: GenMessage<EditSubnetVlanRequest> = /*@__PURE__*/
  messageDesc(file_subnet, 7);

/**
 * @generated from message subnet.EditSubnetVlanResponse
 */
export type EditSubnetVlanResponse = Message<"subnet.EditSubnetVlanResponse"> & {
  /**
   * @generated from field: bool success = 1;
   */
  success: boolean;
};

/**
 * Describes the message subnet.EditSubnetVlanResponse.
 * Use `create(EditSubnetVlanResponseSchema)` to create a new message.
 */
export const EditSubnetVlanResponseSchema: GenMessage<EditSubnetVlanResponse> = /*@__PURE__*/
  messageDesc(file_subnet, 8);

/**
 * @generated from message subnet.EditSubnetDescriptionRequest
 */
export type EditSubnetDescriptionRequest = Message<"subnet.EditSubnetDescriptionRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string description = 2;
   */
  description: string;
};

/**
 * Describes the message subnet.EditSubnetDescriptionRequest.
 * Use `create(EditSubnetDescriptionRequestSchema)` to create a new message.
 */
export const EditSubnetDescriptionRequestSchema: GenMessage<EditSubnetDescriptionRequest> = /*@__PURE__*/
  messageDesc(file_subnet, 9);

/**
 * @generated from message subnet.EditSubnetDescriptionResponse
 */
export type EditSubnetDescriptionResponse = Message<"subnet.EditSubnetDescriptionResponse"> & {
  /**
   * @generated from field: bool success = 1;
   */
  success: boolean;
};

/**
 * Describes the message subnet.EditSubnetDescriptionResponse.
 * Use `create(EditSubnetDescriptionResponseSchema)` to create a new message.
 */
export const EditSubnetDescriptionResponseSchema: GenMessage<EditSubnetDescriptionResponse> = /*@__PURE__*/
  messageDesc(file_subnet, 10);

/**
 * @generated from service subnet.SubnetService
 */
export const SubnetService: GenService<{
  /**
   * @generated from rpc subnet.SubnetService.CreateSubnet
   */
  createSubnet: {
    methodKind: "unary";
    input: typeof CreateSubnetRequestSchema;
    output: typeof CreateSubnetResponseSchema;
  },
  /**
   * @generated from rpc subnet.SubnetService.GetSubnets
   */
  getSubnets: {
    methodKind: "unary";
    input: typeof GetSubnetsRequestSchema;
    output: typeof GetSubnetsResponseSchema;
  },
  /**
   * @generated from rpc subnet.SubnetService.GetSubnetById
   */
  getSubnetById: {
    methodKind: "unary";
    input: typeof GetSubnetRequestSchema;
    output: typeof GetSubnetResponseSchema;
  },
  /**
   * @generated from rpc subnet.SubnetService.EditSubnetVlan
   */
  editSubnetVlan: {
    methodKind: "unary";
    input: typeof EditSubnetVlanRequestSchema;
    output: typeof EditSubnetVlanResponseSchema;
  },
  /**
   * @generated from rpc subnet.SubnetService.EditSubnetDescription
   */
  editSubnetDescription: {
    methodKind: "unary";
    input: typeof EditSubnetDescriptionRequestSchema;
    output: typeof EditSubnetDescriptionResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_subnet, 0);

