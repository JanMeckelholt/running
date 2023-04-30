///
//  Generated code. Do not modify.
//  source: database/database.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use activitiesRequestDescriptor instead')
const ActivitiesRequest$json = const {
  '1': 'ActivitiesRequest',
  '2': const [
    const {'1': 'ClientId', '3': 1, '4': 1, '5': 9, '10': 'ClientId'},
    const {'1': 'Since', '3': 2, '4': 1, '5': 4, '10': 'Since'},
    const {'1': 'Until', '3': 3, '4': 1, '5': 4, '10': 'Until'},
  ],
};

/// Descriptor for `ActivitiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activitiesRequestDescriptor = $convert.base64Decode('ChFBY3Rpdml0aWVzUmVxdWVzdBIaCghDbGllbnRJZBgBIAEoCVIIQ2xpZW50SWQSFAoFU2luY2UYAiABKARSBVNpbmNlEhQKBVVudGlsGAMgASgEUgVVbnRpbA==');
@$core.Deprecated('Use activitiesResponseDescriptor instead')
const ActivitiesResponse$json = const {
  '1': 'ActivitiesResponse',
  '2': const [
    const {'1': 'Activities', '3': 1, '4': 3, '5': 11, '6': '.strava.Activity', '10': 'Activities'},
  ],
};

/// Descriptor for `ActivitiesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activitiesResponseDescriptor = $convert.base64Decode('ChJBY3Rpdml0aWVzUmVzcG9uc2USMAoKQWN0aXZpdGllcxgBIAMoCzIQLnN0cmF2YS5BY3Rpdml0eVIKQWN0aXZpdGllcw==');
@$core.Deprecated('Use clientIdDescriptor instead')
const ClientId$json = const {
  '1': 'ClientId',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
  ],
};

/// Descriptor for `ClientId`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List clientIdDescriptor = $convert.base64Decode('CghDbGllbnRJZBIaCghjbGllbnRJZBgBIAEoCVIIY2xpZW50SWQ=');
@$core.Deprecated('Use activityIdDescriptor instead')
const ActivityId$json = const {
  '1': 'ActivityId',
  '2': const [
    const {'1': 'ActivityId', '3': 1, '4': 1, '5': 9, '10': 'ActivityId'},
  ],
};

/// Descriptor for `ActivityId`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activityIdDescriptor = $convert.base64Decode('CgpBY3Rpdml0eUlkEh4KCkFjdGl2aXR5SWQYASABKAlSCkFjdGl2aXR5SWQ=');
@$core.Deprecated('Use clientDescriptor instead')
const Client$json = const {
  '1': 'Client',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'clientSecret', '3': 2, '4': 1, '5': 9, '10': 'clientSecret'},
    const {'1': 'token', '3': 3, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'refreshToken', '3': 4, '4': 1, '5': 9, '10': 'refreshToken'},
    const {'1': 'athleteId', '3': 5, '4': 1, '5': 4, '10': 'athleteId'},
  ],
};

/// Descriptor for `Client`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List clientDescriptor = $convert.base64Decode('CgZDbGllbnQSGgoIY2xpZW50SWQYASABKAlSCGNsaWVudElkEiIKDGNsaWVudFNlY3JldBgCIAEoCVIMY2xpZW50U2VjcmV0EhQKBXRva2VuGAMgASgJUgV0b2tlbhIiCgxyZWZyZXNoVG9rZW4YBCABKAlSDHJlZnJlc2hUb2tlbhIcCglhdGhsZXRlSWQYBSABKARSCWF0aGxldGVJZA==');
@$core.Deprecated('Use updateRequestDescriptor instead')
const UpdateRequest$json = const {
  '1': 'UpdateRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'kvPairs', '3': 2, '4': 3, '5': 11, '6': '.database.kvPair', '10': 'kvPairs'},
  ],
};

/// Descriptor for `UpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateRequestDescriptor = $convert.base64Decode('Cg1VcGRhdGVSZXF1ZXN0EhoKCGNsaWVudElkGAEgASgJUghjbGllbnRJZBIqCgdrdlBhaXJzGAIgAygLMhAuZGF0YWJhc2Uua3ZQYWlyUgdrdlBhaXJz');
@$core.Deprecated('Use kvPairDescriptor instead')
const kvPair$json = const {
  '1': 'kvPair',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `kvPair`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kvPairDescriptor = $convert.base64Decode('CgZrdlBhaXISEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVl');
@$core.Deprecated('Use tokenResponseDescriptor instead')
const TokenResponse$json = const {
  '1': 'TokenResponse',
  '2': const [
    const {'1': 'success_response', '3': 1, '4': 1, '5': 11, '6': '.database.SuccessResponse', '9': 0, '10': 'successResponse'},
    const {'1': 'error_response', '3': 2, '4': 1, '5': 11, '6': '.database.ErrorResponse', '9': 0, '10': 'errorResponse'},
  ],
  '8': const [
    const {'1': 'response'},
  ],
};

/// Descriptor for `TokenResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tokenResponseDescriptor = $convert.base64Decode('Cg1Ub2tlblJlc3BvbnNlEkYKEHN1Y2Nlc3NfcmVzcG9uc2UYASABKAsyGS5kYXRhYmFzZS5TdWNjZXNzUmVzcG9uc2VIAFIPc3VjY2Vzc1Jlc3BvbnNlEkAKDmVycm9yX3Jlc3BvbnNlGAIgASgLMhcuZGF0YWJhc2UuRXJyb3JSZXNwb25zZUgAUg1lcnJvclJlc3BvbnNlQgoKCHJlc3BvbnNl');
@$core.Deprecated('Use successResponseDescriptor instead')
const SuccessResponse$json = const {
  '1': 'SuccessResponse',
  '2': const [
    const {'1': 'message', '3': 1, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `SuccessResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List successResponseDescriptor = $convert.base64Decode('Cg9TdWNjZXNzUmVzcG9uc2USGAoHbWVzc2FnZRgBIAEoCVIHbWVzc2FnZQ==');
@$core.Deprecated('Use errorResponseDescriptor instead')
const ErrorResponse$json = const {
  '1': 'ErrorResponse',
  '2': const [
    const {'1': 'error', '3': 1, '4': 1, '5': 9, '10': 'error'},
  ],
};

/// Descriptor for `ErrorResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List errorResponseDescriptor = $convert.base64Decode('Cg1FcnJvclJlc3BvbnNlEhQKBWVycm9yGAEgASgJUgVlcnJvcg==');
