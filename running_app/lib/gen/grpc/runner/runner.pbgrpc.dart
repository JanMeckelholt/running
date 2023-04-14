///
//  Generated code. Do not modify.
//  source: runner/runner.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'runner.pb.dart' as $3;
import '../strava/strava.pb.dart' as $0;
import '../database/database.pb.dart' as $2;
import '../google/protobuf/empty.pb.dart' as $1;
export 'runner.pb.dart';

class RunnerClient extends $grpc.Client {
  static final _$getRunner =
      $grpc.ClientMethod<$3.RunnerRequest, $0.RunnerResponse>(
          '/runner.Runner/GetRunner',
          ($3.RunnerRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.RunnerResponse.fromBuffer(value));
  static final _$createClient = $grpc.ClientMethod<$2.Client, $1.Empty>(
      '/runner.Runner/CreateClient',
      ($2.Client value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$getActivities =
      $grpc.ClientMethod<$2.ActivitiesRequest, $2.ActivitiesResponse>(
          '/runner.Runner/GetActivities',
          ($2.ActivitiesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.ActivitiesResponse.fromBuffer(value));
  static final _$activitiesToDB =
      $grpc.ClientMethod<$3.ActivitiesRequest, $1.Empty>(
          '/runner.Runner/ActivitiesToDB',
          ($3.ActivitiesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$getWeekSummaries =
      $grpc.ClientMethod<$3.WeekSummariesRequest, $3.WeekSummariesResponse>(
          '/runner.Runner/GetWeekSummaries',
          ($3.WeekSummariesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $3.WeekSummariesResponse.fromBuffer(value));
  static final _$health =
      $grpc.ClientMethod<$3.HealthMessage, $3.HealthMessage>(
          '/runner.Runner/Health',
          ($3.HealthMessage value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $3.HealthMessage.fromBuffer(value));

  RunnerClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.RunnerResponse> getRunner($3.RunnerRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRunner, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> createClient($2.Client request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createClient, request, options: options);
  }

  $grpc.ResponseFuture<$2.ActivitiesResponse> getActivities(
      $2.ActivitiesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActivities, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> activitiesToDB($3.ActivitiesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$activitiesToDB, request, options: options);
  }

  $grpc.ResponseFuture<$3.WeekSummariesResponse> getWeekSummaries(
      $3.WeekSummariesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getWeekSummaries, request, options: options);
  }

  $grpc.ResponseFuture<$3.HealthMessage> health($3.HealthMessage request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$health, request, options: options);
  }
}

abstract class RunnerServiceBase extends $grpc.Service {
  $core.String get $name => 'runner.Runner';

  RunnerServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.RunnerRequest, $0.RunnerResponse>(
        'GetRunner',
        getRunner_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.RunnerRequest.fromBuffer(value),
        ($0.RunnerResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Client, $1.Empty>(
        'CreateClient',
        createClient_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Client.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ActivitiesRequest, $2.ActivitiesResponse>(
        'GetActivities',
        getActivities_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ActivitiesRequest.fromBuffer(value),
        ($2.ActivitiesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.ActivitiesRequest, $1.Empty>(
        'ActivitiesToDB',
        activitiesToDB_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.ActivitiesRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$3.WeekSummariesRequest, $3.WeekSummariesResponse>(
            'GetWeekSummaries',
            getWeekSummaries_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $3.WeekSummariesRequest.fromBuffer(value),
            ($3.WeekSummariesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.HealthMessage, $3.HealthMessage>(
        'Health',
        health_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.HealthMessage.fromBuffer(value),
        ($3.HealthMessage value) => value.writeToBuffer()));
  }

  $async.Future<$0.RunnerResponse> getRunner_Pre(
      $grpc.ServiceCall call, $async.Future<$3.RunnerRequest> request) async {
    return getRunner(call, await request);
  }

  $async.Future<$1.Empty> createClient_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Client> request) async {
    return createClient(call, await request);
  }

  $async.Future<$2.ActivitiesResponse> getActivities_Pre($grpc.ServiceCall call,
      $async.Future<$2.ActivitiesRequest> request) async {
    return getActivities(call, await request);
  }

  $async.Future<$1.Empty> activitiesToDB_Pre($grpc.ServiceCall call,
      $async.Future<$3.ActivitiesRequest> request) async {
    return activitiesToDB(call, await request);
  }

  $async.Future<$3.WeekSummariesResponse> getWeekSummaries_Pre(
      $grpc.ServiceCall call,
      $async.Future<$3.WeekSummariesRequest> request) async {
    return getWeekSummaries(call, await request);
  }

  $async.Future<$3.HealthMessage> health_Pre(
      $grpc.ServiceCall call, $async.Future<$3.HealthMessage> request) async {
    return health(call, await request);
  }

  $async.Future<$0.RunnerResponse> getRunner(
      $grpc.ServiceCall call, $3.RunnerRequest request);
  $async.Future<$1.Empty> createClient(
      $grpc.ServiceCall call, $2.Client request);
  $async.Future<$2.ActivitiesResponse> getActivities(
      $grpc.ServiceCall call, $2.ActivitiesRequest request);
  $async.Future<$1.Empty> activitiesToDB(
      $grpc.ServiceCall call, $3.ActivitiesRequest request);
  $async.Future<$3.WeekSummariesResponse> getWeekSummaries(
      $grpc.ServiceCall call, $3.WeekSummariesRequest request);
  $async.Future<$3.HealthMessage> health(
      $grpc.ServiceCall call, $3.HealthMessage request);
}
