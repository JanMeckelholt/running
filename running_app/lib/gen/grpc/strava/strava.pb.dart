///
//  Generated code. Do not modify.
//  source: strava/strava.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class RunnerRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RunnerRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  RunnerRequest._() : super();
  factory RunnerRequest({
    $core.String? clientId,
    $core.String? token,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (token != null) {
      _result.token = token;
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

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class RefreshRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RefreshRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientSecret', protoName: 'clientSecret')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'refresToken', protoName: 'refresToken')
    ..hasRequiredFields = false
  ;

  RefreshRequest._() : super();
  factory RefreshRequest({
    $core.String? clientId,
    $core.String? clientSecret,
    $core.String? refresToken,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (clientSecret != null) {
      _result.clientSecret = clientSecret;
    }
    if (refresToken != null) {
      _result.refresToken = refresToken;
    }
    return _result;
  }
  factory RefreshRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RefreshRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RefreshRequest clone() => RefreshRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RefreshRequest copyWith(void Function(RefreshRequest) updates) => super.copyWith((message) => updates(message as RefreshRequest)) as RefreshRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RefreshRequest create() => RefreshRequest._();
  RefreshRequest createEmptyInstance() => create();
  static $pb.PbList<RefreshRequest> createRepeated() => $pb.PbList<RefreshRequest>();
  @$core.pragma('dart2js:noInline')
  static RefreshRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RefreshRequest>(create);
  static RefreshRequest? _defaultInstance;

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
  $core.String get refresToken => $_getSZ(2);
  @$pb.TagNumber(3)
  set refresToken($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRefresToken() => $_has(2);
  @$pb.TagNumber(3)
  void clearRefresToken() => clearField(3);
}

class TokenResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TokenResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  TokenResponse._() : super();
  factory TokenResponse({
    $core.String? clientId,
    $core.String? token,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (token != null) {
      _result.token = token;
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

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class athleteType extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'athleteType', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aInt64(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'resourceState')
    ..hasRequiredFields = false
  ;

  athleteType._() : super();
  factory athleteType({
    $fixnum.Int64? id,
    $fixnum.Int64? resourceState,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (resourceState != null) {
      _result.resourceState = resourceState;
    }
    return _result;
  }
  factory athleteType.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory athleteType.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  athleteType clone() => athleteType()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  athleteType copyWith(void Function(athleteType) updates) => super.copyWith((message) => updates(message as athleteType)) as athleteType; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static athleteType create() => athleteType._();
  athleteType createEmptyInstance() => create();
  static $pb.PbList<athleteType> createRepeated() => $pb.PbList<athleteType>();
  @$core.pragma('dart2js:noInline')
  static athleteType getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<athleteType>(create);
  static athleteType? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get resourceState => $_getI64(1);
  @$pb.TagNumber(2)
  set resourceState($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasResourceState() => $_has(1);
  @$pb.TagNumber(2)
  void clearResourceState() => clearField(2);
}

class coordinate extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'coordinate', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  coordinate._() : super();
  factory coordinate() => create();
  factory coordinate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory coordinate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  coordinate clone() => coordinate()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  coordinate copyWith(void Function(coordinate) updates) => super.copyWith((message) => updates(message as coordinate)) as coordinate; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static coordinate create() => coordinate._();
  coordinate createEmptyInstance() => create();
  static $pb.PbList<coordinate> createRepeated() => $pb.PbList<coordinate>();
  @$core.pragma('dart2js:noInline')
  static coordinate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<coordinate>(create);
  static coordinate? _defaultInstance;
}

class RunnerResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'RunnerResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'username')
    ..a<$fixnum.Int64>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'resourceState', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'firstname')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lastname')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'bio')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'city')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'state')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'country')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sex')
    ..aOB(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'premium')
    ..aOB(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'summit')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'createdAt')
    ..aOS(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updatedAt')
    ..a<$fixnum.Int64>(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'badgeTypeId', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$core.double>(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'weight', $pb.PbFieldType.OD)
    ..aOS(17, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'profileMedium')
    ..aOS(18, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'profile')
    ..hasRequiredFields = false
  ;

  RunnerResponse._() : super();
  factory RunnerResponse({
    $fixnum.Int64? id,
    $core.String? username,
    $fixnum.Int64? resourceState,
    $core.String? firstname,
    $core.String? lastname,
    $core.String? bio,
    $core.String? city,
    $core.String? state,
    $core.String? country,
    $core.String? sex,
    $core.bool? premium,
    $core.bool? summit,
    $core.String? createdAt,
    $core.String? updatedAt,
    $fixnum.Int64? badgeTypeId,
    $core.double? weight,
    $core.String? profileMedium,
    $core.String? profile,
  }) {
    final _result = create();
    if (id != null) {
      _result.id = id;
    }
    if (username != null) {
      _result.username = username;
    }
    if (resourceState != null) {
      _result.resourceState = resourceState;
    }
    if (firstname != null) {
      _result.firstname = firstname;
    }
    if (lastname != null) {
      _result.lastname = lastname;
    }
    if (bio != null) {
      _result.bio = bio;
    }
    if (city != null) {
      _result.city = city;
    }
    if (state != null) {
      _result.state = state;
    }
    if (country != null) {
      _result.country = country;
    }
    if (sex != null) {
      _result.sex = sex;
    }
    if (premium != null) {
      _result.premium = premium;
    }
    if (summit != null) {
      _result.summit = summit;
    }
    if (createdAt != null) {
      _result.createdAt = createdAt;
    }
    if (updatedAt != null) {
      _result.updatedAt = updatedAt;
    }
    if (badgeTypeId != null) {
      _result.badgeTypeId = badgeTypeId;
    }
    if (weight != null) {
      _result.weight = weight;
    }
    if (profileMedium != null) {
      _result.profileMedium = profileMedium;
    }
    if (profile != null) {
      _result.profile = profile;
    }
    return _result;
  }
  factory RunnerResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RunnerResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RunnerResponse clone() => RunnerResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RunnerResponse copyWith(void Function(RunnerResponse) updates) => super.copyWith((message) => updates(message as RunnerResponse)) as RunnerResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RunnerResponse create() => RunnerResponse._();
  RunnerResponse createEmptyInstance() => create();
  static $pb.PbList<RunnerResponse> createRepeated() => $pb.PbList<RunnerResponse>();
  @$core.pragma('dart2js:noInline')
  static RunnerResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RunnerResponse>(create);
  static RunnerResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get username => $_getSZ(1);
  @$pb.TagNumber(2)
  set username($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUsername() => $_has(1);
  @$pb.TagNumber(2)
  void clearUsername() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get resourceState => $_getI64(2);
  @$pb.TagNumber(3)
  set resourceState($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasResourceState() => $_has(2);
  @$pb.TagNumber(3)
  void clearResourceState() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstname => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstname($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstname() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstname() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastname => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastname($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastname() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastname() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get bio => $_getSZ(5);
  @$pb.TagNumber(6)
  set bio($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasBio() => $_has(5);
  @$pb.TagNumber(6)
  void clearBio() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get city => $_getSZ(6);
  @$pb.TagNumber(7)
  set city($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasCity() => $_has(6);
  @$pb.TagNumber(7)
  void clearCity() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get state => $_getSZ(7);
  @$pb.TagNumber(8)
  set state($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasState() => $_has(7);
  @$pb.TagNumber(8)
  void clearState() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get country => $_getSZ(8);
  @$pb.TagNumber(9)
  set country($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasCountry() => $_has(8);
  @$pb.TagNumber(9)
  void clearCountry() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get sex => $_getSZ(9);
  @$pb.TagNumber(10)
  set sex($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasSex() => $_has(9);
  @$pb.TagNumber(10)
  void clearSex() => clearField(10);

  @$pb.TagNumber(11)
  $core.bool get premium => $_getBF(10);
  @$pb.TagNumber(11)
  set premium($core.bool v) { $_setBool(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasPremium() => $_has(10);
  @$pb.TagNumber(11)
  void clearPremium() => clearField(11);

  @$pb.TagNumber(12)
  $core.bool get summit => $_getBF(11);
  @$pb.TagNumber(12)
  set summit($core.bool v) { $_setBool(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasSummit() => $_has(11);
  @$pb.TagNumber(12)
  void clearSummit() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get createdAt => $_getSZ(12);
  @$pb.TagNumber(13)
  set createdAt($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasCreatedAt() => $_has(12);
  @$pb.TagNumber(13)
  void clearCreatedAt() => clearField(13);

  @$pb.TagNumber(14)
  $core.String get updatedAt => $_getSZ(13);
  @$pb.TagNumber(14)
  set updatedAt($core.String v) { $_setString(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasUpdatedAt() => $_has(13);
  @$pb.TagNumber(14)
  void clearUpdatedAt() => clearField(14);

  @$pb.TagNumber(15)
  $fixnum.Int64 get badgeTypeId => $_getI64(14);
  @$pb.TagNumber(15)
  set badgeTypeId($fixnum.Int64 v) { $_setInt64(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasBadgeTypeId() => $_has(14);
  @$pb.TagNumber(15)
  void clearBadgeTypeId() => clearField(15);

  @$pb.TagNumber(16)
  $core.double get weight => $_getN(15);
  @$pb.TagNumber(16)
  set weight($core.double v) { $_setDouble(15, v); }
  @$pb.TagNumber(16)
  $core.bool hasWeight() => $_has(15);
  @$pb.TagNumber(16)
  void clearWeight() => clearField(16);

  @$pb.TagNumber(17)
  $core.String get profileMedium => $_getSZ(16);
  @$pb.TagNumber(17)
  set profileMedium($core.String v) { $_setString(16, v); }
  @$pb.TagNumber(17)
  $core.bool hasProfileMedium() => $_has(16);
  @$pb.TagNumber(17)
  void clearProfileMedium() => clearField(17);

  @$pb.TagNumber(18)
  $core.String get profile => $_getSZ(17);
  @$pb.TagNumber(18)
  set profile($core.String v) { $_setString(17, v); }
  @$pb.TagNumber(18)
  $core.bool hasProfile() => $_has(17);
  @$pb.TagNumber(18)
  void clearProfile() => clearField(18);
}

class ActivityRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivityRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..hasRequiredFields = false
  ;

  ActivityRequest._() : super();
  factory ActivityRequest({
    $core.String? clientId,
    $core.String? token,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (token != null) {
      _result.token = token;
    }
    return _result;
  }
  factory ActivityRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ActivityRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ActivityRequest clone() => ActivityRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ActivityRequest copyWith(void Function(ActivityRequest) updates) => super.copyWith((message) => updates(message as ActivityRequest)) as ActivityRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ActivityRequest create() => ActivityRequest._();
  ActivityRequest createEmptyInstance() => create();
  static $pb.PbList<ActivityRequest> createRepeated() => $pb.PbList<ActivityRequest>();
  @$core.pragma('dart2js:noInline')
  static ActivityRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ActivityRequest>(create);
  static ActivityRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get clientId => $_getSZ(0);
  @$pb.TagNumber(1)
  set clientId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasClientId() => $_has(0);
  @$pb.TagNumber(1)
  void clearClientId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);
}

class ActivitiesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivitiesRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientId', protoName: 'clientId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..a<$fixnum.Int64>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'since', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  ActivitiesRequest._() : super();
  factory ActivitiesRequest({
    $core.String? clientId,
    $core.String? token,
    $fixnum.Int64? since,
  }) {
    final _result = create();
    if (clientId != null) {
      _result.clientId = clientId;
    }
    if (token != null) {
      _result.token = token;
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
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get since => $_getI64(2);
  @$pb.TagNumber(3)
  set since($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSince() => $_has(2);
  @$pb.TagNumber(3)
  void clearSince() => clearField(3);
}

class ActivitiesResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ActivitiesResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..pc<Activity>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Activities', $pb.PbFieldType.PM, protoName: 'Activities', subBuilder: Activity.create)
    ..hasRequiredFields = false
  ;

  ActivitiesResponse._() : super();
  factory ActivitiesResponse({
    $core.Iterable<Activity>? activities,
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
  $core.List<Activity> get activities => $_getList(0);
}

class Activity extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Activity', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'strava'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'resourceState', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOM<athleteType>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'athlete', subBuilder: athleteType.create)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..a<$core.double>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'distance', $pb.PbFieldType.OD)
    ..a<$fixnum.Int64>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'movingTime', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'elapsedTime', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$core.double>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'totalElevationGain', $pb.PbFieldType.OD)
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'type')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sportType')
    ..a<$fixnum.Int64>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startDate')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startDateLocale')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timezone')
    ..a<$core.double>(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'utcOffset', $pb.PbFieldType.OD)
    ..aOS(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationCity')
    ..aOS(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationState')
    ..aOS(17, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'locationCountry')
    ..a<$fixnum.Int64>(18, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'achievementCount', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(19, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'kudosCount', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(20, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'commentCount', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOB(21, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'manual')
    ..aOS(22, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'visibility')
    ..p<$core.double>(23, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startLatlng', $pb.PbFieldType.KD)
    ..p<$core.double>(24, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'endLatlng', $pb.PbFieldType.KD)
    ..a<$core.double>(25, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'averageSpeed', $pb.PbFieldType.OD)
    ..a<$core.double>(26, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'maxSpeed', $pb.PbFieldType.OD)
    ..a<$core.double>(27, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'averageHeartrate', $pb.PbFieldType.OD)
    ..a<$core.double>(28, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'maxHeartrate', $pb.PbFieldType.OD)
    ..a<$core.double>(29, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'elevHigh', $pb.PbFieldType.OD)
    ..a<$core.double>(30, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'elevLow', $pb.PbFieldType.OD)
    ..a<$fixnum.Int64>(40, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'startDateUnix', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  Activity._() : super();
  factory Activity({
    $fixnum.Int64? resourceState,
    athleteType? athlete,
    $core.String? name,
    $core.double? distance,
    $fixnum.Int64? movingTime,
    $fixnum.Int64? elapsedTime,
    $core.double? totalElevationGain,
    $core.String? type,
    $core.String? sportType,
    $fixnum.Int64? id,
    $core.String? startDate,
    $core.String? startDateLocale,
    $core.String? timezone,
    $core.double? utcOffset,
    $core.String? locationCity,
    $core.String? locationState,
    $core.String? locationCountry,
    $fixnum.Int64? achievementCount,
    $fixnum.Int64? kudosCount,
    $fixnum.Int64? commentCount,
    $core.bool? manual,
    $core.String? visibility,
    $core.Iterable<$core.double>? startLatlng,
    $core.Iterable<$core.double>? endLatlng,
    $core.double? averageSpeed,
    $core.double? maxSpeed,
    $core.double? averageHeartrate,
    $core.double? maxHeartrate,
    $core.double? elevHigh,
    $core.double? elevLow,
    $fixnum.Int64? startDateUnix,
  }) {
    final _result = create();
    if (resourceState != null) {
      _result.resourceState = resourceState;
    }
    if (athlete != null) {
      _result.athlete = athlete;
    }
    if (name != null) {
      _result.name = name;
    }
    if (distance != null) {
      _result.distance = distance;
    }
    if (movingTime != null) {
      _result.movingTime = movingTime;
    }
    if (elapsedTime != null) {
      _result.elapsedTime = elapsedTime;
    }
    if (totalElevationGain != null) {
      _result.totalElevationGain = totalElevationGain;
    }
    if (type != null) {
      _result.type = type;
    }
    if (sportType != null) {
      _result.sportType = sportType;
    }
    if (id != null) {
      _result.id = id;
    }
    if (startDate != null) {
      _result.startDate = startDate;
    }
    if (startDateLocale != null) {
      _result.startDateLocale = startDateLocale;
    }
    if (timezone != null) {
      _result.timezone = timezone;
    }
    if (utcOffset != null) {
      _result.utcOffset = utcOffset;
    }
    if (locationCity != null) {
      _result.locationCity = locationCity;
    }
    if (locationState != null) {
      _result.locationState = locationState;
    }
    if (locationCountry != null) {
      _result.locationCountry = locationCountry;
    }
    if (achievementCount != null) {
      _result.achievementCount = achievementCount;
    }
    if (kudosCount != null) {
      _result.kudosCount = kudosCount;
    }
    if (commentCount != null) {
      _result.commentCount = commentCount;
    }
    if (manual != null) {
      _result.manual = manual;
    }
    if (visibility != null) {
      _result.visibility = visibility;
    }
    if (startLatlng != null) {
      _result.startLatlng.addAll(startLatlng);
    }
    if (endLatlng != null) {
      _result.endLatlng.addAll(endLatlng);
    }
    if (averageSpeed != null) {
      _result.averageSpeed = averageSpeed;
    }
    if (maxSpeed != null) {
      _result.maxSpeed = maxSpeed;
    }
    if (averageHeartrate != null) {
      _result.averageHeartrate = averageHeartrate;
    }
    if (maxHeartrate != null) {
      _result.maxHeartrate = maxHeartrate;
    }
    if (elevHigh != null) {
      _result.elevHigh = elevHigh;
    }
    if (elevLow != null) {
      _result.elevLow = elevLow;
    }
    if (startDateUnix != null) {
      _result.startDateUnix = startDateUnix;
    }
    return _result;
  }
  factory Activity.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Activity.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Activity clone() => Activity()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Activity copyWith(void Function(Activity) updates) => super.copyWith((message) => updates(message as Activity)) as Activity; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Activity create() => Activity._();
  Activity createEmptyInstance() => create();
  static $pb.PbList<Activity> createRepeated() => $pb.PbList<Activity>();
  @$core.pragma('dart2js:noInline')
  static Activity getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Activity>(create);
  static Activity? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get resourceState => $_getI64(0);
  @$pb.TagNumber(1)
  set resourceState($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasResourceState() => $_has(0);
  @$pb.TagNumber(1)
  void clearResourceState() => clearField(1);

  @$pb.TagNumber(2)
  athleteType get athlete => $_getN(1);
  @$pb.TagNumber(2)
  set athlete(athleteType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasAthlete() => $_has(1);
  @$pb.TagNumber(2)
  void clearAthlete() => clearField(2);
  @$pb.TagNumber(2)
  athleteType ensureAthlete() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get distance => $_getN(3);
  @$pb.TagNumber(4)
  set distance($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDistance() => $_has(3);
  @$pb.TagNumber(4)
  void clearDistance() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get movingTime => $_getI64(4);
  @$pb.TagNumber(5)
  set movingTime($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasMovingTime() => $_has(4);
  @$pb.TagNumber(5)
  void clearMovingTime() => clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get elapsedTime => $_getI64(5);
  @$pb.TagNumber(6)
  set elapsedTime($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasElapsedTime() => $_has(5);
  @$pb.TagNumber(6)
  void clearElapsedTime() => clearField(6);

  @$pb.TagNumber(7)
  $core.double get totalElevationGain => $_getN(6);
  @$pb.TagNumber(7)
  set totalElevationGain($core.double v) { $_setDouble(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasTotalElevationGain() => $_has(6);
  @$pb.TagNumber(7)
  void clearTotalElevationGain() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get type => $_getSZ(7);
  @$pb.TagNumber(8)
  set type($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasType() => $_has(7);
  @$pb.TagNumber(8)
  void clearType() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get sportType => $_getSZ(8);
  @$pb.TagNumber(9)
  set sportType($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasSportType() => $_has(8);
  @$pb.TagNumber(9)
  void clearSportType() => clearField(9);

  @$pb.TagNumber(10)
  $fixnum.Int64 get id => $_getI64(9);
  @$pb.TagNumber(10)
  set id($fixnum.Int64 v) { $_setInt64(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasId() => $_has(9);
  @$pb.TagNumber(10)
  void clearId() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get startDate => $_getSZ(10);
  @$pb.TagNumber(11)
  set startDate($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasStartDate() => $_has(10);
  @$pb.TagNumber(11)
  void clearStartDate() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get startDateLocale => $_getSZ(11);
  @$pb.TagNumber(12)
  set startDateLocale($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasStartDateLocale() => $_has(11);
  @$pb.TagNumber(12)
  void clearStartDateLocale() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get timezone => $_getSZ(12);
  @$pb.TagNumber(13)
  set timezone($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasTimezone() => $_has(12);
  @$pb.TagNumber(13)
  void clearTimezone() => clearField(13);

  @$pb.TagNumber(14)
  $core.double get utcOffset => $_getN(13);
  @$pb.TagNumber(14)
  set utcOffset($core.double v) { $_setDouble(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasUtcOffset() => $_has(13);
  @$pb.TagNumber(14)
  void clearUtcOffset() => clearField(14);

  @$pb.TagNumber(15)
  $core.String get locationCity => $_getSZ(14);
  @$pb.TagNumber(15)
  set locationCity($core.String v) { $_setString(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasLocationCity() => $_has(14);
  @$pb.TagNumber(15)
  void clearLocationCity() => clearField(15);

  @$pb.TagNumber(16)
  $core.String get locationState => $_getSZ(15);
  @$pb.TagNumber(16)
  set locationState($core.String v) { $_setString(15, v); }
  @$pb.TagNumber(16)
  $core.bool hasLocationState() => $_has(15);
  @$pb.TagNumber(16)
  void clearLocationState() => clearField(16);

  @$pb.TagNumber(17)
  $core.String get locationCountry => $_getSZ(16);
  @$pb.TagNumber(17)
  set locationCountry($core.String v) { $_setString(16, v); }
  @$pb.TagNumber(17)
  $core.bool hasLocationCountry() => $_has(16);
  @$pb.TagNumber(17)
  void clearLocationCountry() => clearField(17);

  @$pb.TagNumber(18)
  $fixnum.Int64 get achievementCount => $_getI64(17);
  @$pb.TagNumber(18)
  set achievementCount($fixnum.Int64 v) { $_setInt64(17, v); }
  @$pb.TagNumber(18)
  $core.bool hasAchievementCount() => $_has(17);
  @$pb.TagNumber(18)
  void clearAchievementCount() => clearField(18);

  @$pb.TagNumber(19)
  $fixnum.Int64 get kudosCount => $_getI64(18);
  @$pb.TagNumber(19)
  set kudosCount($fixnum.Int64 v) { $_setInt64(18, v); }
  @$pb.TagNumber(19)
  $core.bool hasKudosCount() => $_has(18);
  @$pb.TagNumber(19)
  void clearKudosCount() => clearField(19);

  @$pb.TagNumber(20)
  $fixnum.Int64 get commentCount => $_getI64(19);
  @$pb.TagNumber(20)
  set commentCount($fixnum.Int64 v) { $_setInt64(19, v); }
  @$pb.TagNumber(20)
  $core.bool hasCommentCount() => $_has(19);
  @$pb.TagNumber(20)
  void clearCommentCount() => clearField(20);

  @$pb.TagNumber(21)
  $core.bool get manual => $_getBF(20);
  @$pb.TagNumber(21)
  set manual($core.bool v) { $_setBool(20, v); }
  @$pb.TagNumber(21)
  $core.bool hasManual() => $_has(20);
  @$pb.TagNumber(21)
  void clearManual() => clearField(21);

  @$pb.TagNumber(22)
  $core.String get visibility => $_getSZ(21);
  @$pb.TagNumber(22)
  set visibility($core.String v) { $_setString(21, v); }
  @$pb.TagNumber(22)
  $core.bool hasVisibility() => $_has(21);
  @$pb.TagNumber(22)
  void clearVisibility() => clearField(22);

  @$pb.TagNumber(23)
  $core.List<$core.double> get startLatlng => $_getList(22);

  @$pb.TagNumber(24)
  $core.List<$core.double> get endLatlng => $_getList(23);

  @$pb.TagNumber(25)
  $core.double get averageSpeed => $_getN(24);
  @$pb.TagNumber(25)
  set averageSpeed($core.double v) { $_setDouble(24, v); }
  @$pb.TagNumber(25)
  $core.bool hasAverageSpeed() => $_has(24);
  @$pb.TagNumber(25)
  void clearAverageSpeed() => clearField(25);

  @$pb.TagNumber(26)
  $core.double get maxSpeed => $_getN(25);
  @$pb.TagNumber(26)
  set maxSpeed($core.double v) { $_setDouble(25, v); }
  @$pb.TagNumber(26)
  $core.bool hasMaxSpeed() => $_has(25);
  @$pb.TagNumber(26)
  void clearMaxSpeed() => clearField(26);

  @$pb.TagNumber(27)
  $core.double get averageHeartrate => $_getN(26);
  @$pb.TagNumber(27)
  set averageHeartrate($core.double v) { $_setDouble(26, v); }
  @$pb.TagNumber(27)
  $core.bool hasAverageHeartrate() => $_has(26);
  @$pb.TagNumber(27)
  void clearAverageHeartrate() => clearField(27);

  @$pb.TagNumber(28)
  $core.double get maxHeartrate => $_getN(27);
  @$pb.TagNumber(28)
  set maxHeartrate($core.double v) { $_setDouble(27, v); }
  @$pb.TagNumber(28)
  $core.bool hasMaxHeartrate() => $_has(27);
  @$pb.TagNumber(28)
  void clearMaxHeartrate() => clearField(28);

  @$pb.TagNumber(29)
  $core.double get elevHigh => $_getN(28);
  @$pb.TagNumber(29)
  set elevHigh($core.double v) { $_setDouble(28, v); }
  @$pb.TagNumber(29)
  $core.bool hasElevHigh() => $_has(28);
  @$pb.TagNumber(29)
  void clearElevHigh() => clearField(29);

  @$pb.TagNumber(30)
  $core.double get elevLow => $_getN(29);
  @$pb.TagNumber(30)
  set elevLow($core.double v) { $_setDouble(29, v); }
  @$pb.TagNumber(30)
  $core.bool hasElevLow() => $_has(29);
  @$pb.TagNumber(30)
  void clearElevLow() => clearField(30);

  @$pb.TagNumber(40)
  $fixnum.Int64 get startDateUnix => $_getI64(30);
  @$pb.TagNumber(40)
  set startDateUnix($fixnum.Int64 v) { $_setInt64(30, v); }
  @$pb.TagNumber(40)
  $core.bool hasStartDateUnix() => $_has(30);
  @$pb.TagNumber(40)
  void clearStartDateUnix() => clearField(40);
}

