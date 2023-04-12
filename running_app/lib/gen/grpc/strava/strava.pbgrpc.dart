///
//  Generated code. Do not modify.
//  source: strava/strava.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'strava.pb.dart' as $0;
import '../google/protobuf/empty.pb.dart' as $1;
export 'strava.pb.dart';

class StravaClient extends $grpc.Client {
  static final _$getRunner =
      $grpc.ClientMethod<$0.RunnerRequest, $0.RunnerResponse>(
          '/strava.Strava/GetRunner',
          ($0.RunnerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.RunnerResponse.fromBuffer(value));
  static final _$getActivities =
      $grpc.ClientMethod<$0.ActivitiesRequest, $0.ActivitiesResponse>(
          '/strava.Strava/GetActivities',
          ($0.ActivitiesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.ActivitiesResponse.fromBuffer(value));
  static final _$activitiesToDB =
      $grpc.ClientMethod<$0.ActivitiesRequest, $1.Empty>(
          '/strava.Strava/ActivitiesToDB',
          ($0.ActivitiesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$useRefreshToken =
      $grpc.ClientMethod<$0.RefreshRequest, $0.TokenResponse>(
          '/strava.Strava/UseRefreshToken',
          ($0.RefreshRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.TokenResponse.fromBuffer(value));

  StravaClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.RunnerResponse> getRunner($0.RunnerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRunner, request, options: options);
  }

  $grpc.ResponseFuture<$0.ActivitiesResponse> getActivities(
      $0.ActivitiesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActivities, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> activitiesToDB($0.ActivitiesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$activitiesToDB, request, options: options);
  }

  $grpc.ResponseFuture<$0.TokenResponse> useRefreshToken(
      $0.RefreshRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$useRefreshToken, request, options: options);
  }
}

abstract class StravaServiceBase extends $grpc.Service {
  $core.String get $name => 'strava.Strava';

  StravaServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.RunnerRequest, $0.RunnerResponse>(
        'GetRunner',
        getRunner_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RunnerRequest.fromBuffer(value),
        ($0.RunnerResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ActivitiesRequest, $0.ActivitiesResponse>(
        'GetActivities',
        getActivities_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ActivitiesRequest.fromBuffer(value),
        ($0.ActivitiesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ActivitiesRequest, $1.Empty>(
        'ActivitiesToDB',
        activitiesToDB_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ActivitiesRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RefreshRequest, $0.TokenResponse>(
        'UseRefreshToken',
        useRefreshToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RefreshRequest.fromBuffer(value),
        ($0.TokenResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.RunnerResponse> getRunner_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RunnerRequest> request) async {
    return getRunner(call, await request);
  }

  $async.Future<$0.ActivitiesResponse> getActivities_Pre($grpc.ServiceCall call,
      $async.Future<$0.ActivitiesRequest> request) async {
    return getActivities(call, await request);
  }

  $async.Future<$1.Empty> activitiesToDB_Pre($grpc.ServiceCall call,
      $async.Future<$0.ActivitiesRequest> request) async {
    return activitiesToDB(call, await request);
  }

  $async.Future<$0.TokenResponse> useRefreshToken_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RefreshRequest> request) async {
    return useRefreshToken(call, await request);
  }

  $async.Future<$0.RunnerResponse> getRunner(
      $grpc.ServiceCall call, $0.RunnerRequest request);
  $async.Future<$0.ActivitiesResponse> getActivities(
      $grpc.ServiceCall call, $0.ActivitiesRequest request);
  $async.Future<$1.Empty> activitiesToDB(
      $grpc.ServiceCall call, $0.ActivitiesRequest request);
  $async.Future<$0.TokenResponse> useRefreshToken(
      $grpc.ServiceCall call, $0.RefreshRequest request);
}
