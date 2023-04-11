///
//  Generated code. Do not modify.
//  source: strava/strava.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'strava.pb.dart' as $1;
import '../google/protobuf/empty.pb.dart' as $0;
import 'strava.pbjson.dart';

export 'strava.pb.dart';

abstract class StravaServiceBase extends $pb.GeneratedService {
  $async.Future<$1.RunnerResponse> getRunner($pb.ServerContext ctx, $1.RunnerRequest request);
  $async.Future<$1.ActivitiesResponse> getActivities($pb.ServerContext ctx, $1.ActivitiesRequest request);
  $async.Future<$0.Empty> activitiesToDB($pb.ServerContext ctx, $1.ActivitiesRequest request);
  $async.Future<$1.TokenResponse> useRefreshToken($pb.ServerContext ctx, $1.RefreshRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetRunner': return $1.RunnerRequest();
      case 'GetActivities': return $1.ActivitiesRequest();
      case 'ActivitiesToDB': return $1.ActivitiesRequest();
      case 'UseRefreshToken': return $1.RefreshRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetRunner': return this.getRunner(ctx, request as $1.RunnerRequest);
      case 'GetActivities': return this.getActivities(ctx, request as $1.ActivitiesRequest);
      case 'ActivitiesToDB': return this.activitiesToDB(ctx, request as $1.ActivitiesRequest);
      case 'UseRefreshToken': return this.useRefreshToken(ctx, request as $1.RefreshRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => StravaServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => StravaServiceBase$messageJson;
}

