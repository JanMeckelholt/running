///
//  Generated code. Do not modify.
//  source: database/database.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../strava/strava.pb.dart' as $0;
import '../google/protobuf/empty.pb.dart' as $1;

class ActivitiesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivitiesRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ClientId', protoName: 'ClientId')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Since', $pb.PbFieldType.OU6, protoName: 'Since', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  ActivitiesRequest._() : super();
  factory ActivitiesRequest({
    $core.String? clientId,
    $fixnum.Int64? since,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (since != null) {
      _result.since = since;
    }
    return _result;
  }
  factory ActivitiesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActivitiesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActivitiesRequest clone() => ActivitiesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActivitiesRequest copyWith(void Function(ActivitiesRequest) updates) => super.copyWith((message) => updates(message as ActivitiesRequest)) as ActivitiesRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActivitiesRequest create() => ActivitiesRequest._();
  ActivitiesRequest createEmptyInstance() => create();
  static $pb.PbList<ActivitiesRequest> createRepeated() => $pb.PbList<ActivitiesRequest>();
  @$core.pragma('dart2js:noInline')
  static ActivitiesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActivitiesRequest>(create);
  static ActivitiesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get since => $_getI64(1);
  @$pb.TagNumber(2)
  set since($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSince() => $_has(1);
  @$pb.TagNumber(2)
  void clearSince() => clearField(2);
}

class ActivitiesResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivitiesResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..pc<$0.Activity>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Activities', $pb.PbFieldType.PM, protoName: 'Activities', subBuilder: $0.Activity.create)
    ..hasRequiredFields = false
  ;

  ActivitiesResponse._() : super();
  factory ActivitiesResponse({
    $core.Iterable<$0.Activity>? activities,
  }) {
    final _result = create();
    if (activities != null) {
      _result.activities.addAll(activities);
    }
    return _result;
  }
  factory ActivitiesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActivitiesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActivitiesResponse clone() => ActivitiesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActivitiesResponse copyWith(void Function(ActivitiesResponse) updates) => super.copyWith((message) => updates(message as ActivitiesResponse)) as ActivitiesResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActivitiesResponse create() => ActivitiesResponse._();
  ActivitiesResponse createEmptyInstance() => create();
  static $pb.PbList<ActivitiesResponse> createRepeated() => $pb.PbList<ActivitiesResponse>();
  @$core.pragma('dart2js:noInline')
  static ActivitiesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActivitiesResponse>(create);
  static ActivitiesResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$0.Activity> get activities => $_getList(0);
}

class ClientId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ClientId', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..hasRequiredFields = false
  ;

  ClientId._() : super();
  factory ClientId({
    $core.String? clientId,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    return _result;
  }
  factory ClientId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ClientId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ClientId clone() => ClientId()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ClientId copyWith(void Function(ClientId) updates) => super.copyWith((message) => updates(message as ClientId)) as ClientId; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ClientId create() => ClientId._();
  ClientId createEmptyInstance() => create();
  static $pb.PbList<ClientId> createRepeated() => $pb.PbList<ClientId>();
  @$core.pragma('dart2js:noInline')
  static ClientId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ClientId>(create);
  static ClientId? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);
}

class ActivityId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivityId', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ActivityId', protoName: 'ActivityId')
    ..hasRequiredFields = false
  ;

  ActivityId._() : super();
  factory ActivityId({
    $core.String? activityId,
  }) {
    final _result = create();
    if (activityId != null) {
      _result.activityId = activityId;
    }
    return _result;
  }
  factory ActivityId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActivityId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActivityId clone() => ActivityId()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActivityId copyWith(void Function(ActivityId) updates) => super.copyWith((message) => updates(message as ActivityId)) as ActivityId; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActivityId create() => ActivityId._();
  ActivityId createEmptyInstance() => create();
  static $pb.PbList<ActivityId> createRepeated() => $pb.PbList<ActivityId>();
  @$core.pragma('dart2js:noInline')
  static ActivityId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActivityId>(create);
  static ActivityId? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get activityId => $_getSZ(0);
  @$pb.TagNumber(1)
  set activityId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasActivityId() => $_has(0);
  @$pb.TagNumber(1)
  void clearActivityId() => clearField(1);
}

class Client extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Client', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientSecret', protoName: 'clientSecret')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'refreshToken', protoName: 'refreshToken')
    ..a<$fixnum.Int64>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'athleteId', $pb.PbFieldType.OU6, protoName: 'athleteId', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  Client._() : super();
  factory Client({
    $core.String? clientId,
    $core.String? clientSecret,
    $core.String? token,
    $core.String? refreshToken,
    $fixnum.Int64? athleteId,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (clientSecret != null) {
      _result.clientSecret = clientSecret;
    }
    if (token != null) {
      _result.token = token;
    }
    if (refreshToken != null) {
      _result.refreshToken = refreshToken;
    }
    if (athleteId != null) {
      _result.athleteId = athleteId;
    }
    return _result;
  }
  factory Client.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Client.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Client clone() => Client()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Client copyWith(void Function(Client) updates) => super.copyWith((message) => updates(message as Client)) as Client; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Client create() => Client._();
  Client createEmptyInstance() => create();
  static $pb.PbList<Client> createRepeated() => $pb.PbList<Client>();
  @$core.pragma('dart2js:noInline')
  static Client getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Client>(create);
  static Client? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get clientSecret => $_getSZ(1);
  @$pb.TagNumber(2)
  set clientSecret($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasClientSecret() => $_has(1);
  @$pb.TagNumber(2)
  void clearClientSecret() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get token => $_getSZ(2);
  @$pb.TagNumber(3)
  set token($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasToken() => $_has(2);
  @$pb.TagNumber(3)
  void clearToken() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get refreshToken => $_getSZ(3);
  @$pb.TagNumber(4)
  set refreshToken($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRefreshToken() => $_has(3);
  @$pb.TagNumber(4)
  void clearRefreshToken() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get athleteId => $_getI64(4);
  @$pb.TagNumber(5)
  set athleteId($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasAthleteId() => $_has(4);
  @$pb.TagNumber(5)
  void clearAthleteId() => clearField(5);
}

class UpdateRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..pc<kvPair>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'kvPairs', $pb.PbFieldType.PM, protoName: 'kvPairs', subBuilder: kvPair.create)
    ..hasRequiredFields = false
  ;

  UpdateRequest._() : super();
  factory UpdateRequest({
    $core.String? clientId,
    $core.Iterable<kvPair>? kvPairs,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (kvPairs != null) {
      _result.kvPairs.addAll(kvPairs);
    }
    return _result;
  }
  factory UpdateRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateRequest clone() => UpdateRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateRequest copyWith(void Function(UpdateRequest) updates) => super.copyWith((message) => updates(message as UpdateRequest)) as UpdateRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateRequest create() => UpdateRequest._();
  UpdateRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateRequest> createRepeated() => $pb.PbList<UpdateRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateRequest>(create);
  static UpdateRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<kvPair> get kvPairs => $_getList(1);
}

class kvPair extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'kvPair', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'key')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'value')
    ..hasRequiredFields = false
  ;

  kvPair._() : super();
  factory kvPair({
    $core.String? key,
    $core.String? value,
  }) {
    final _result = create();
    if (key != null) {
      _result.key = key;
    }
    if (value != null) {
      _result.value = value;
    }
    return _result;
  }
  factory kvPair.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory kvPair.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  kvPair clone() => kvPair()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  kvPair copyWith(void Function(kvPair) updates) => super.copyWith((message) => updates(message as kvPair)) as kvPair; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static kvPair create() => kvPair._();
  kvPair createEmptyInstance() => create();
  static $pb.PbList<kvPair> createRepeated() => $pb.PbList<kvPair>();
  @$core.pragma('dart2js:noInline')
  static kvPair getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<kvPair>(create);
  static kvPair? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
}

enum TokenResponse_Response {
  successResponse, 
  errorResponse, 
  notSet
}

class TokenResponse extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, TokenResponse_Response> _TokenResponse_ResponseByTag = {
    1 : TokenResponse_Response.successResponse,
    2 : TokenResponse_Response.errorResponse,
    0 : TokenResponse_Response.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TokenResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..aOM<SuccessResponse>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'successResponse', subBuilder: SuccessResponse.create)
    ..aOM<ErrorResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorResponse', subBuilder: ErrorResponse.create)
    ..hasRequiredFields = false
  ;

  TokenResponse._() : super();
  factory TokenResponse({
    SuccessResponse? successResponse,
    ErrorResponse? errorResponse,
  }) {
    final _result = create();
    if (successResponse != null) {
      _result.successResponse = successResponse;
    }
    if (errorResponse != null) {
      _result.errorResponse = errorResponse;
    }
    return _result;
  }
  factory TokenResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TokenResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TokenResponse clone() => TokenResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TokenResponse copyWith(void Function(TokenResponse) updates) => super.copyWith((message) => updates(message as TokenResponse)) as TokenResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TokenResponse create() => TokenResponse._();
  TokenResponse createEmptyInstance() => create();
  static $pb.PbList<TokenResponse> createRepeated() => $pb.PbList<TokenResponse>();
  @$core.pragma('dart2js:noInline')
  static TokenResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TokenResponse>(create);
  static TokenResponse? _defaultInstance;

  TokenResponse_Response whichResponse() => _TokenResponse_ResponseByTag[$_whichOneof(0)]!;
  void clearResponse() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  SuccessResponse get successResponse => $_getN(0);
  @$pb.TagNumber(1)
  set successResponse(SuccessResponse v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSuccessResponse() => $_has(0);
  @$pb.TagNumber(1)
  void clearSuccessResponse() => clearField(1);
  @$pb.TagNumber(1)
  SuccessResponse ensureSuccessResponse() => $_ensure(0);

  @$pb.TagNumber(2)
  ErrorResponse get errorResponse => $_getN(1);
  @$pb.TagNumber(2)
  set errorResponse(ErrorResponse v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorResponse() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorResponse() => clearField(2);
  @$pb.TagNumber(2)
  ErrorResponse ensureErrorResponse() => $_ensure(1);
}

class SuccessResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SuccessResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'message')
    ..hasRequiredFields = false
  ;

  SuccessResponse._() : super();
  factory SuccessResponse({
    $core.String? message,
  }) {
    final _result = create();
    if (message != null) {
      _result.message = message;
    }
    return _result;
  }
  factory SuccessResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SuccessResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SuccessResponse clone() => SuccessResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SuccessResponse copyWith(void Function(SuccessResponse) updates) => super.copyWith((message) => updates(message as SuccessResponse)) as SuccessResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SuccessResponse create() => SuccessResponse._();
  SuccessResponse createEmptyInstance() => create();
  static $pb.PbList<SuccessResponse> createRepeated() => $pb.PbList<SuccessResponse>();
  @$core.pragma('dart2js:noInline')
  static SuccessResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SuccessResponse>(create);
  static SuccessResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get message => $_getSZ(0);
  @$pb.TagNumber(1)
  set message($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMessage() => $_has(0);
  @$pb.TagNumber(1)
  void clearMessage() => clearField(1);
}

class ErrorResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ErrorResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'database'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'error')
    ..hasRequiredFields = false
  ;

  ErrorResponse._() : super();
  factory ErrorResponse({
    $core.String? error,
  }) {
    final _result = create();
    if (error != null) {
      _result.error = error;
    }
    return _result;
  }
  factory ErrorResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ErrorResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ErrorResponse clone() => ErrorResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ErrorResponse copyWith(void Function(ErrorResponse) updates) => super.copyWith((message) => updates(message as ErrorResponse)) as ErrorResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ErrorResponse create() => ErrorResponse._();
  ErrorResponse createEmptyInstance() => create();
  static $pb.PbList<ErrorResponse> createRepeated() => $pb.PbList<ErrorResponse>();
  @$core.pragma('dart2js:noInline')
  static ErrorResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ErrorResponse>(create);
  static ErrorResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get error => $_getSZ(0);
  @$pb.TagNumber(1)
  set error($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasError() => $_has(0);
  @$pb.TagNumber(1)
  void clearError() => clearField(1);
}

class DatabaseApi {
  $pb.RpcClient _client;
  DatabaseApi(this._client);

  $async.Future<$1.Empty> upsertClient($pb.ClientContext? ctx, Client request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'Database', 'UpsertClient', request, emptyResponse);
  }
  $async.Future<Client> updateClient($pb.ClientContext? ctx, UpdateRequest request) {
    var emptyResponse = Client();
    return _client.invoke<Client>(ctx, 'Database', 'UpdateClient', request, emptyResponse);
  }
  $async.Future<Client> getClient($pb.ClientContext? ctx, ClientId request) {
    var emptyResponse = Client();
    return _client.invoke<Client>(ctx, 'Database', 'GetClient', request, emptyResponse);
  }
  $async.Future<$1.Empty> upsertActivity($pb.ClientContext? ctx, $0.Activity request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'Database', 'UpsertActivity', request, emptyResponse);
  }
  $async.Future<$1.Empty> upsertActivityFromCSV($pb.ClientContext? ctx, $0.Activity request) {
    var emptyResponse = $1.Empty();
    return _client.invoke<$1.Empty>(ctx, 'Database', 'UpsertActivityFromCSV', request, emptyResponse);
  }
  $async.Future<$0.Activity> getActivity($pb.ClientContext? ctx, ActivityId request) {
    var emptyResponse = $0.Activity();
    return _client.invoke<$0.Activity>(ctx, 'Database', 'GetActivity', request, emptyResponse);
  }
  $async.Future<ActivitiesResponse> getActivities($pb.ClientContext? ctx, ActivitiesRequest request) {
    var emptyResponse = ActivitiesResponse();
    return _client.invoke<ActivitiesResponse>(ctx, 'Database', 'GetActivities', request, emptyResponse);
  }
}

