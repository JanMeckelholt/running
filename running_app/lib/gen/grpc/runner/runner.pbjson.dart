///
//  Generated code. Do not modify.
//  source: runner/runner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use healthMessageDescriptor instead')
const HealthMessage$json = const {
  '1': 'HealthMessage',
  '2': const [
    const {'1': 'Health', '3': 1, '4': 1, '5': 9, '10': 'Health'},
  ],
};

/// Descriptor for `HealthMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List healthMessageDescriptor = $convert.base64Decode('Cg1IZWFsdGhNZXNzYWdlEhYKBkhlYWx0aBgBIAEoCVIGSGVhbHRo');
@$core.Deprecated('Use runnerRequestDescriptor instead')
const RunnerRequest$json = const {
  '1': 'RunnerRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
  ],
};

/// Descriptor for `RunnerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List runnerRequestDescriptor = $convert.base64Decode('Cg1SdW5uZXJSZXF1ZXN0EhoKCGNsaWVudElkGAEgASgJUghjbGllbnRJZA==');
@$core.Deprecated('Use activitiesRequestDescriptor instead')
const ActivitiesRequest$json = const {
  '1': 'ActivitiesRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'since', '3': 2, '4': 1, '5': 4, '10': 'since'},
  ],
};

/// Descriptor for `ActivitiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activitiesRequestDescriptor = $convert.base64Decode('ChFBY3Rpdml0aWVzUmVxdWVzdBIaCghjbGllbnRJZBgBIAEoCVIIY2xpZW50SWQSFAoFc2luY2UYAiABKARSBXNpbmNl');
@$core.Deprecated('Use weekSummariesRequestDescriptor instead')
const WeekSummariesRequest$json = const {
  '1': 'WeekSummariesRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'weeks', '3': 2, '4': 1, '5': 4, '10': 'weeks'},
  ],
};

/// Descriptor for `WeekSummariesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List weekSummariesRequestDescriptor = $convert.base64Decode('ChRXZWVrU3VtbWFyaWVzUmVxdWVzdBIaCghjbGllbnRJZBgBIAEoCVIIY2xpZW50SWQSFAoFd2Vla3MYAiABKARSBXdlZWtz');
@$core.Deprecated('Use weekSummaryDescriptor instead')
const WeekSummary$json = const {
  '1': 'WeekSummary',
  '2': const [
    const {'1': 'distance', '3': 1, '4': 1, '5': 4, '10': 'distance'},
    const {'1': 'timeUnix', '3': 2, '4': 1, '5': 4, '10': 'timeUnix'},
    const {'1': 'timeStr', '3': 3, '4': 1, '5': 9, '10': 'timeStr'},
    const {'1': 'climb', '3': 4, '4': 1, '5': 4, '10': 'climb'},
    const {'1': 'numberOfRuns', '3': 5, '4': 1, '5': 4, '10': 'numberOfRuns'},
    const {'1': 'numberOfOthers', '3': 6, '4': 1, '5': 4, '10': 'numberOfOthers'},
  ],
};

/// Descriptor for `WeekSummary`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List weekSummaryDescriptor = $convert.base64Decode('CgtXZWVrU3VtbWFyeRIaCghkaXN0YW5jZRgBIAEoBFIIZGlzdGFuY2USGgoIdGltZVVuaXgYAiABKARSCHRpbWVVbml4EhgKB3RpbWVTdHIYAyABKAlSB3RpbWVTdHISFAoFY2xpbWIYBCABKARSBWNsaW1iEiIKDG51bWJlck9mUnVucxgFIAEoBFIMbnVtYmVyT2ZSdW5zEiYKDm51bWJlck9mT3RoZXJzGAYgASgEUg5udW1iZXJPZk90aGVycw==');
@$core.Deprecated('Use weekSummariesResponseDescriptor instead')
const WeekSummariesResponse$json = const {
  '1': 'WeekSummariesResponse',
  '2': const [
    const {'1': 'weeksummaries', '3': 1, '4': 3, '5': 11, '6': '.runner.WeekSummary', '10': 'weeksummaries'},
  ],
};

/// Descriptor for `WeekSummariesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List weekSummariesResponseDescriptor = $convert.base64Decode('ChVXZWVrU3VtbWFyaWVzUmVzcG9uc2USOQoNd2Vla3N1bW1hcmllcxgBIAMoCzITLnJ1bm5lci5XZWVrU3VtbWFyeVINd2Vla3N1bW1hcmllcw==');
