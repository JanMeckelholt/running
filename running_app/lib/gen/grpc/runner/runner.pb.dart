///
//  Generated code. Do not modify.
//  source: runner/runner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class HealthMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'HealthMessage', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Health', protoName: 'Health')
    ..hasRequiredFields = false
  ;

  HealthMessage._() : super();
  factory HealthMessage({
    $core.String? health,
  }) {
    final _result = create();
    if (health != null) {
      _result.health = health;
    }
    return _result;
  }
  factory HealthMessage.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HealthMessage.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  HealthMessage clone() => HealthMessage()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  HealthMessage copyWith(void Function(HealthMessage) updates) => super.copyWith((message) => updates(message as HealthMessage)) as HealthMessage; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HealthMessage create() => HealthMessage._();
  HealthMessage createEmptyInstance() => create();
  static $pb.PbList<HealthMessage> createRepeated() => $pb.PbList<HealthMessage>();
  @$core.pragma('dart2js:noInline')
  static HealthMessage getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HealthMessage>(create);
  static HealthMessage? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get health => $_getSZ(0);
  @$pb.TagNumber(1)
  set health($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHealth() => $_has(0);
  @$pb.TagNumber(1)
  void clearHealth() => clearField(1);
}

class RunnerRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RunnerRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..hasRequiredFields = false
  ;

  RunnerRequest._() : super();
  factory RunnerRequest({
    $core.String? clientId,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    return _result;
  }
  factory RunnerRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RunnerRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RunnerRequest clone() => RunnerRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RunnerRequest copyWith(void Function(RunnerRequest) updates) => super.copyWith((message) => updates(message as RunnerRequest)) as RunnerRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RunnerRequest create() => RunnerRequest._();
  RunnerRequest createEmptyInstance() => create();
  static $pb.PbList<RunnerRequest> createRepeated() => $pb.PbList<RunnerRequest>();
  @$core.pragma('dart2js:noInline')
  static RunnerRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RunnerRequest>(create);
  static RunnerRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);
}

class ActivitiesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivitiesRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'since', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
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

class WeekSummariesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WeekSummariesRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'weeks', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  WeekSummariesRequest._() : super();
  factory WeekSummariesRequest({
    $core.String? clientId,
    $fixnum.Int64? weeks,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (weeks != null) {
      _result.weeks = weeks;
    }
    return _result;
  }
  factory WeekSummariesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WeekSummariesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WeekSummariesRequest clone() => WeekSummariesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WeekSummariesRequest copyWith(void Function(WeekSummariesRequest) updates) => super.copyWith((message) => updates(message as WeekSummariesRequest)) as WeekSummariesRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WeekSummariesRequest create() => WeekSummariesRequest._();
  WeekSummariesRequest createEmptyInstance() => create();
  static $pb.PbList<WeekSummariesRequest> createRepeated() => $pb.PbList<WeekSummariesRequest>();
  @$core.pragma('dart2js:noInline')
  static WeekSummariesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WeekSummariesRequest>(create);
  static WeekSummariesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get weeks => $_getI64(1);
  @$pb.TagNumber(2)
  set weeks($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasWeeks() => $_has(1);
  @$pb.TagNumber(2)
  void clearWeeks() => clearField(2);
}

class WeekSummary extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WeekSummary', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'distance', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timeUnix', $pb.PbFieldType.OU6, protoName: 'timeUnix', defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timeStr', protoName: 'timeStr')
    ..a<$fixnum.Int64>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'climb', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'numberOfRuns', $pb.PbFieldType.OU6, protoName: 'numberOfRuns', defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'numberOfOthers', $pb.PbFieldType.OU6, protoName: 'numberOfOthers', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  WeekSummary._() : super();
  factory WeekSummary({
    $fixnum.Int64? distance,
    $fixnum.Int64? timeUnix,
    $core.String? timeStr,
    $fixnum.Int64? climb,
    $fixnum.Int64? numberOfRuns,
    $fixnum.Int64? numberOfOthers,
  }) {
    final _result = create();
    if (distance != null) {
      _result.distance = distance;
    }
    if (timeUnix != null) {
      _result.timeUnix = timeUnix;
    }
    if (timeStr != null) {
      _result.timeStr = timeStr;
    }
    if (climb != null) {
      _result.climb = climb;
    }
    if (numberOfRuns != null) {
      _result.numberOfRuns = numberOfRuns;
    }
    if (numberOfOthers != null) {
      _result.numberOfOthers = numberOfOthers;
    }
    return _result;
  }
  factory WeekSummary.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WeekSummary.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WeekSummary clone() => WeekSummary()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WeekSummary copyWith(void Function(WeekSummary) updates) => super.copyWith((message) => updates(message as WeekSummary)) as WeekSummary; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WeekSummary create() => WeekSummary._();
  WeekSummary createEmptyInstance() => create();
  static $pb.PbList<WeekSummary> createRepeated() => $pb.PbList<WeekSummary>();
  @$core.pragma('dart2js:noInline')
  static WeekSummary getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WeekSummary>(create);
  static WeekSummary? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get distance => $_getI64(0);
  @$pb.TagNumber(1)
  set distance($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDistance() => $_has(0);
  @$pb.TagNumber(1)
  void clearDistance() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get timeUnix => $_getI64(1);
  @$pb.TagNumber(2)
  set timeUnix($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTimeUnix() => $_has(1);
  @$pb.TagNumber(2)
  void clearTimeUnix() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get timeStr => $_getSZ(2);
  @$pb.TagNumber(3)
  set timeStr($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTimeStr() => $_has(2);
  @$pb.TagNumber(3)
  void clearTimeStr() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get climb => $_getI64(3);
  @$pb.TagNumber(4)
  set climb($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasClimb() => $_has(3);
  @$pb.TagNumber(4)
  void clearClimb() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get numberOfRuns => $_getI64(4);
  @$pb.TagNumber(5)
  set numberOfRuns($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasNumberOfRuns() => $_has(4);
  @$pb.TagNumber(5)
  void clearNumberOfRuns() => clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get numberOfOthers => $_getI64(5);
  @$pb.TagNumber(6)
  set numberOfOthers($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasNumberOfOthers() => $_has(5);
  @$pb.TagNumber(6)
  void clearNumberOfOthers() => clearField(6);
}

class WeekSummariesResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'WeekSummariesResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'runner'), createEmptyInstance: create)
    ..pc<WeekSummary>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'weeksummaries', $pb.PbFieldType.PM, subBuilder: WeekSummary.create)
    ..hasRequiredFields = false
  ;

  WeekSummariesResponse._() : super();
  factory WeekSummariesResponse({
    $core.Iterable<WeekSummary>? weeksummaries,
  }) {
    final _result = create();
    if (weeksummaries != null) {
      _result.weeksummaries.addAll(weeksummaries);
    }
    return _result;
  }
  factory WeekSummariesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WeekSummariesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WeekSummariesResponse clone() => WeekSummariesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WeekSummariesResponse copyWith(void Function(WeekSummariesResponse) updates) => super.copyWith((message) => updates(message as WeekSummariesResponse)) as WeekSummariesResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WeekSummariesResponse create() => WeekSummariesResponse._();
  WeekSummariesResponse createEmptyInstance() => create();
  static $pb.PbList<WeekSummariesResponse> createRepeated() => $pb.PbList<WeekSummariesResponse>();
  @$core.pragma('dart2js:noInline')
  static WeekSummariesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WeekSummariesResponse>(create);
  static WeekSummariesResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<WeekSummary> get weeksummaries => $_getList(0);
}

