///
//  Generated code. Do not modify.
//  source: strava/strava.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../google/protobuf/empty.pbjson.dart' as $0;

@$core.Deprecated('Use runnerRequestDescriptor instead')
const RunnerRequest$json = const {
  '1': 'RunnerRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `RunnerRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List runnerRequestDescriptor = $convert.base64Decode('Cg1SdW5uZXJSZXF1ZXN0EhoKCGNsaWVudElkGAEgASgJUghjbGllbnRJZBIUCgV0b2tlbhgCIAEoCVIFdG9rZW4=');
@$core.Deprecated('Use refreshRequestDescriptor instead')
const RefreshRequest$json = const {
  '1': 'RefreshRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'clientSecret', '3': 2, '4': 1, '5': 9, '10': 'clientSecret'},
    const {'1': 'refresToken', '3': 3, '4': 1, '5': 9, '10': 'refresToken'},
  ],
};

/// Descriptor for `RefreshRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List refreshRequestDescriptor = $convert.base64Decode('Cg5SZWZyZXNoUmVxdWVzdBIaCghjbGllbnRJZBgBIAEoCVIIY2xpZW50SWQSIgoMY2xpZW50U2VjcmV0GAIgASgJUgxjbGllbnRTZWNyZXQSIAoLcmVmcmVzVG9rZW4YAyABKAlSC3JlZnJlc1Rva2Vu');
@$core.Deprecated('Use tokenResponseDescriptor instead')
const TokenResponse$json = const {
  '1': 'TokenResponse',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `TokenResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List tokenResponseDescriptor = $convert.base64Decode('Cg1Ub2tlblJlc3BvbnNlEhoKCGNsaWVudElkGAEgASgJUghjbGllbnRJZBIUCgV0b2tlbhgCIAEoCVIFdG9rZW4=');
@$core.Deprecated('Use runnerResponseDescriptor instead')
const RunnerResponse$json = const {
  '1': 'RunnerResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 4, '10': 'id'},
    const {'1': 'username', '3': 2, '4': 1, '5': 9, '10': 'username'},
    const {'1': 'resource_state', '3': 3, '4': 1, '5': 4, '10': 'resourceState'},
    const {'1': 'firstname', '3': 4, '4': 1, '5': 9, '10': 'firstname'},
    const {'1': 'lastname', '3': 5, '4': 1, '5': 9, '10': 'lastname'},
    const {'1': 'bio', '3': 6, '4': 1, '5': 9, '10': 'bio'},
    const {'1': 'city', '3': 7, '4': 1, '5': 9, '10': 'city'},
    const {'1': 'state', '3': 8, '4': 1, '5': 9, '10': 'state'},
    const {'1': 'country', '3': 9, '4': 1, '5': 9, '10': 'country'},
    const {'1': 'sex', '3': 10, '4': 1, '5': 9, '10': 'sex'},
    const {'1': 'premium', '3': 11, '4': 1, '5': 8, '10': 'premium'},
    const {'1': 'summit', '3': 12, '4': 1, '5': 8, '10': 'summit'},
    const {'1': 'created_at', '3': 13, '4': 1, '5': 9, '10': 'createdAt'},
    const {'1': 'updated_at', '3': 14, '4': 1, '5': 9, '10': 'updatedAt'},
    const {'1': 'badge_type_id', '3': 15, '4': 1, '5': 4, '10': 'badgeTypeId'},
    const {'1': 'weight', '3': 16, '4': 1, '5': 1, '10': 'weight'},
    const {'1': 'profile_medium', '3': 17, '4': 1, '5': 9, '10': 'profileMedium'},
    const {'1': 'profile', '3': 18, '4': 1, '5': 9, '10': 'profile'},
    const {'1': 'friend', '3': 19, '4': 3, '5': 11, '6': '.strava.friend', '10': 'friend'},
    const {'1': 'follower', '3': 20, '4': 3, '5': 11, '6': '.strava.follower', '10': 'follower'},
  ],
};

/// Descriptor for `RunnerResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List runnerResponseDescriptor = $convert.base64Decode('Cg5SdW5uZXJSZXNwb25zZRIOCgJpZBgBIAEoBFICaWQSGgoIdXNlcm5hbWUYAiABKAlSCHVzZXJuYW1lEiUKDnJlc291cmNlX3N0YXRlGAMgASgEUg1yZXNvdXJjZVN0YXRlEhwKCWZpcnN0bmFtZRgEIAEoCVIJZmlyc3RuYW1lEhoKCGxhc3RuYW1lGAUgASgJUghsYXN0bmFtZRIQCgNiaW8YBiABKAlSA2JpbxISCgRjaXR5GAcgASgJUgRjaXR5EhQKBXN0YXRlGAggASgJUgVzdGF0ZRIYCgdjb3VudHJ5GAkgASgJUgdjb3VudHJ5EhAKA3NleBgKIAEoCVIDc2V4EhgKB3ByZW1pdW0YCyABKAhSB3ByZW1pdW0SFgoGc3VtbWl0GAwgASgIUgZzdW1taXQSHQoKY3JlYXRlZF9hdBgNIAEoCVIJY3JlYXRlZEF0Eh0KCnVwZGF0ZWRfYXQYDiABKAlSCXVwZGF0ZWRBdBIiCg1iYWRnZV90eXBlX2lkGA8gASgEUgtiYWRnZVR5cGVJZBIWCgZ3ZWlnaHQYECABKAFSBndlaWdodBIlCg5wcm9maWxlX21lZGl1bRgRIAEoCVINcHJvZmlsZU1lZGl1bRIYCgdwcm9maWxlGBIgASgJUgdwcm9maWxlEiYKBmZyaWVuZBgTIAMoCzIOLnN0cmF2YS5mcmllbmRSBmZyaWVuZBIsCghmb2xsb3dlchgUIAMoCzIQLnN0cmF2YS5mb2xsb3dlclIIZm9sbG93ZXI=');
@$core.Deprecated('Use friendDescriptor instead')
const friend$json = const {
  '1': 'friend',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `friend`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List friendDescriptor = $convert.base64Decode('CgZmcmllbmQSEgoEbmFtZRgBIAEoCVIEbmFtZQ==');
@$core.Deprecated('Use followerDescriptor instead')
const follower$json = const {
  '1': 'follower',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `follower`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List followerDescriptor = $convert.base64Decode('Cghmb2xsb3dlchISCgRuYW1lGAEgASgJUgRuYW1l');
@$core.Deprecated('Use activityRequestDescriptor instead')
const ActivityRequest$json = const {
  '1': 'ActivityRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `ActivityRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activityRequestDescriptor = $convert.base64Decode('Cg9BY3Rpdml0eVJlcXVlc3QSGgoIY2xpZW50SWQYASABKAlSCGNsaWVudElkEhQKBXRva2VuGAIgASgJUgV0b2tlbg==');
@$core.Deprecated('Use activitiesRequestDescriptor instead')
const ActivitiesRequest$json = const {
  '1': 'ActivitiesRequest',
  '2': const [
    const {'1': 'clientId', '3': 1, '4': 1, '5': 9, '10': 'clientId'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'since', '3': 3, '4': 1, '5': 4, '10': 'since'},
  ],
};

/// Descriptor for `ActivitiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activitiesRequestDescriptor = $convert.base64Decode('ChFBY3Rpdml0aWVzUmVxdWVzdBIaCghjbGllbnRJZBgBIAEoCVIIY2xpZW50SWQSFAoFdG9rZW4YAiABKAlSBXRva2VuEhQKBXNpbmNlGAMgASgEUgVzaW5jZQ==');
@$core.Deprecated('Use activitiesResponseDescriptor instead')
const ActivitiesResponse$json = const {
  '1': 'ActivitiesResponse',
  '2': const [
    const {'1': 'Activities', '3': 1, '4': 3, '5': 11, '6': '.strava.Activity', '10': 'Activities'},
  ],
};

/// Descriptor for `ActivitiesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activitiesResponseDescriptor = $convert.base64Decode('ChJBY3Rpdml0aWVzUmVzcG9uc2USMAoKQWN0aXZpdGllcxgBIAMoCzIQLnN0cmF2YS5BY3Rpdml0eVIKQWN0aXZpdGllcw==');
@$core.Deprecated('Use activityDescriptor instead')
const Activity$json = const {
  '1': 'Activity',
  '2': const [
    const {'1': 'resource_state', '3': 1, '4': 1, '5': 4, '10': 'resourceState'},
    const {'1': 'athlete', '3': 2, '4': 1, '5': 11, '6': '.strava.athlete', '10': 'athlete'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'distance', '3': 4, '4': 1, '5': 1, '10': 'distance'},
    const {'1': 'moving_time', '3': 5, '4': 1, '5': 4, '10': 'movingTime'},
    const {'1': 'elapsed_time', '3': 6, '4': 1, '5': 4, '10': 'elapsedTime'},
    const {'1': 'total_elevation_gain', '3': 7, '4': 1, '5': 1, '10': 'totalElevationGain'},
    const {'1': 'type', '3': 8, '4': 1, '5': 9, '10': 'type'},
    const {'1': 'sport_type', '3': 9, '4': 1, '5': 9, '10': 'sportType'},
    const {'1': 'id', '3': 10, '4': 1, '5': 4, '10': 'id'},
    const {'1': 'start_date', '3': 11, '4': 1, '5': 9, '10': 'startDate'},
    const {'1': 'start_date_locale', '3': 12, '4': 1, '5': 9, '10': 'startDateLocale'},
    const {'1': 'timezone', '3': 13, '4': 1, '5': 9, '10': 'timezone'},
    const {'1': 'utc_offset', '3': 14, '4': 1, '5': 1, '10': 'utcOffset'},
    const {'1': 'location_city', '3': 15, '4': 1, '5': 9, '10': 'locationCity'},
    const {'1': 'location_state', '3': 16, '4': 1, '5': 9, '10': 'locationState'},
    const {'1': 'location_country', '3': 17, '4': 1, '5': 9, '10': 'locationCountry'},
    const {'1': 'achievement_count', '3': 18, '4': 1, '5': 4, '10': 'achievementCount'},
    const {'1': 'kudos_count', '3': 19, '4': 1, '5': 4, '10': 'kudosCount'},
    const {'1': 'comment_count', '3': 20, '4': 1, '5': 4, '10': 'commentCount'},
    const {'1': 'manual', '3': 21, '4': 1, '5': 8, '10': 'manual'},
    const {'1': 'visibility', '3': 22, '4': 1, '5': 9, '10': 'visibility'},
    const {'1': 'start_latlng', '3': 23, '4': 3, '5': 1, '10': 'startLatlng'},
    const {'1': 'end_latlng', '3': 24, '4': 3, '5': 1, '10': 'endLatlng'},
    const {'1': 'average_speed', '3': 25, '4': 1, '5': 1, '10': 'averageSpeed'},
    const {'1': 'max_speed', '3': 26, '4': 1, '5': 1, '10': 'maxSpeed'},
    const {'1': 'average_heartrate', '3': 27, '4': 1, '5': 1, '10': 'averageHeartrate'},
    const {'1': 'max_heartrate', '3': 28, '4': 1, '5': 1, '10': 'maxHeartrate'},
    const {'1': 'elev_high', '3': 29, '4': 1, '5': 1, '10': 'elevHigh'},
    const {'1': 'elev_low', '3': 30, '4': 1, '5': 1, '10': 'elevLow'},
    const {'1': 'start_date_unix', '3': 40, '4': 1, '5': 4, '10': 'startDateUnix'},
  ],
};

/// Descriptor for `Activity`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List activityDescriptor = $convert.base64Decode('CghBY3Rpdml0eRIlCg5yZXNvdXJjZV9zdGF0ZRgBIAEoBFINcmVzb3VyY2VTdGF0ZRIpCgdhdGhsZXRlGAIgASgLMg8uc3RyYXZhLmF0aGxldGVSB2F0aGxldGUSEgoEbmFtZRgDIAEoCVIEbmFtZRIaCghkaXN0YW5jZRgEIAEoAVIIZGlzdGFuY2USHwoLbW92aW5nX3RpbWUYBSABKARSCm1vdmluZ1RpbWUSIQoMZWxhcHNlZF90aW1lGAYgASgEUgtlbGFwc2VkVGltZRIwChR0b3RhbF9lbGV2YXRpb25fZ2FpbhgHIAEoAVISdG90YWxFbGV2YXRpb25HYWluEhIKBHR5cGUYCCABKAlSBHR5cGUSHQoKc3BvcnRfdHlwZRgJIAEoCVIJc3BvcnRUeXBlEg4KAmlkGAogASgEUgJpZBIdCgpzdGFydF9kYXRlGAsgASgJUglzdGFydERhdGUSKgoRc3RhcnRfZGF0ZV9sb2NhbGUYDCABKAlSD3N0YXJ0RGF0ZUxvY2FsZRIaCgh0aW1lem9uZRgNIAEoCVIIdGltZXpvbmUSHQoKdXRjX29mZnNldBgOIAEoAVIJdXRjT2Zmc2V0EiMKDWxvY2F0aW9uX2NpdHkYDyABKAlSDGxvY2F0aW9uQ2l0eRIlCg5sb2NhdGlvbl9zdGF0ZRgQIAEoCVINbG9jYXRpb25TdGF0ZRIpChBsb2NhdGlvbl9jb3VudHJ5GBEgASgJUg9sb2NhdGlvbkNvdW50cnkSKwoRYWNoaWV2ZW1lbnRfY291bnQYEiABKARSEGFjaGlldmVtZW50Q291bnQSHwoLa3Vkb3NfY291bnQYEyABKARSCmt1ZG9zQ291bnQSIwoNY29tbWVudF9jb3VudBgUIAEoBFIMY29tbWVudENvdW50EhYKBm1hbnVhbBgVIAEoCFIGbWFudWFsEh4KCnZpc2liaWxpdHkYFiABKAlSCnZpc2liaWxpdHkSIQoMc3RhcnRfbGF0bG5nGBcgAygBUgtzdGFydExhdGxuZxIdCgplbmRfbGF0bG5nGBggAygBUgllbmRMYXRsbmcSIwoNYXZlcmFnZV9zcGVlZBgZIAEoAVIMYXZlcmFnZVNwZWVkEhsKCW1heF9zcGVlZBgaIAEoAVIIbWF4U3BlZWQSKwoRYXZlcmFnZV9oZWFydHJhdGUYGyABKAFSEGF2ZXJhZ2VIZWFydHJhdGUSIwoNbWF4X2hlYXJ0cmF0ZRgcIAEoAVIMbWF4SGVhcnRyYXRlEhsKCWVsZXZfaGlnaBgdIAEoAVIIZWxldkhpZ2gSGQoIZWxldl9sb3cYHiABKAFSB2VsZXZMb3cSJgoPc3RhcnRfZGF0ZV91bml4GCggASgEUg1zdGFydERhdGVVbml4');
@$core.Deprecated('Use athleteDescriptor instead')
const athlete$json = const {
  '1': 'athlete',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 4, '10': 'id'},
    const {'1': 'resource_state', '3': 2, '4': 1, '5': 3, '10': 'resourceState'},
  ],
};

/// Descriptor for `athlete`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List athleteDescriptor = $convert.base64Decode('CgdhdGhsZXRlEg4KAmlkGAEgASgEUgJpZBIlCg5yZXNvdXJjZV9zdGF0ZRgCIAEoA1INcmVzb3VyY2VTdGF0ZQ==');
@$core.Deprecated('Use coordinateDescriptor instead')
const coordinate$json = const {
  '1': 'coordinate',
};

/// Descriptor for `coordinate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List coordinateDescriptor = $convert.base64Decode('Cgpjb29yZGluYXRl');
const $core.Map<$core.String, $core.dynamic> StravaServiceBase$json = const {
  '1': 'Strava',
  '2': const [
    const {'1': 'GetRunner', '2': '.strava.RunnerRequest', '3': '.strava.RunnerResponse', '4': const {}},
    const {'1': 'GetActivities', '2': '.strava.ActivitiesRequest', '3': '.strava.ActivitiesResponse', '4': const {}},
    const {'1': 'ActivitiesToDB', '2': '.strava.ActivitiesRequest', '3': '.google.protobuf.Empty', '4': const {}},
    const {'1': 'UseRefreshToken', '2': '.strava.RefreshRequest', '3': '.strava.TokenResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use stravaServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> StravaServiceBase$messageJson = const {
  '.strava.RunnerRequest': RunnerRequest$json,
  '.strava.RunnerResponse': RunnerResponse$json,
  '.strava.friend': friend$json,
  '.strava.follower': follower$json,
  '.strava.ActivitiesRequest': ActivitiesRequest$json,
  '.strava.ActivitiesResponse': ActivitiesResponse$json,
  '.strava.Activity': Activity$json,
  '.strava.athlete': athlete$json,
  '.google.protobuf.Empty': $0.Empty$json,
  '.strava.RefreshRequest': RefreshRequest$json,
  '.strava.TokenResponse': TokenResponse$json,
};

/// Descriptor for `Strava`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List stravaServiceDescriptor = $convert.base64Decode('CgZTdHJhdmESPAoJR2V0UnVubmVyEhUuc3RyYXZhLlJ1bm5lclJlcXVlc3QaFi5zdHJhdmEuUnVubmVyUmVzcG9uc2UiABJICg1HZXRBY3Rpdml0aWVzEhkuc3RyYXZhLkFjdGl2aXRpZXNSZXF1ZXN0Ghouc3RyYXZhLkFjdGl2aXRpZXNSZXNwb25zZSIAEkUKDkFjdGl2aXRpZXNUb0RCEhkuc3RyYXZhLkFjdGl2aXRpZXNSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgASQgoPVXNlUmVmcmVzaFRva2VuEhYuc3RyYXZhLlJlZnJlc2hSZXF1ZXN0GhUuc3RyYXZhLlRva2VuUmVzcG9uc2UiAA==');
