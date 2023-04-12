///
//  Generated code. Do not modify.
//  source: database/database.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'database.pb.dart' as $2;
import '../google/protobuf/empty.pb.dart' as $1;
import '../strava/strava.pb.dart' as $0;
export 'database.pb.dart';

class DatabaseClient extends $grpc.Client {
  static final _$upsertClient = $grpc.ClientMethod<$2.Client, $1.Empty>(
      '/database.Database/UpsertClient',
      ($2.Client value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$updateClient = $grpc.ClientMethod<$2.UpdateRequest, $2.Client>(
      '/database.Database/UpdateClient',
      ($2.UpdateRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Client.fromBuffer(value));
  static final _$getClient = $grpc.ClientMethod<$2.ClientId, $2.Client>(
      '/database.Database/GetClient',
      ($2.ClientId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Client.fromBuffer(value));
  static final _$upsertActivity = $grpc.ClientMethod<$0.Activity, $1.Empty>(
      '/database.Database/UpsertActivity',
      ($0.Activity value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$upsertActivityFromCSV =
      $grpc.ClientMethod<$0.Activity, $1.Empty>(
          '/database.Database/UpsertActivityFromCSV',
          ($0.Activity value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$getActivity = $grpc.ClientMethod<$2.ActivityId, $0.Activity>(
      '/database.Database/GetActivity',
      ($2.ActivityId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Activity.fromBuffer(value));
  static final _$getActivities =
      $grpc.ClientMethod<$2.ActivitiesRequest, $2.ActivitiesResponse>(
          '/database.Database/GetActivities',
          ($2.ActivitiesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $2.ActivitiesResponse.fromBuffer(value));

  DatabaseClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.Empty> upsertClient($2.Client request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$upsertClient, request, options: options);
  }

  $grpc.ResponseFuture<$2.Client> updateClient($2.UpdateRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateClient, request, options: options);
  }

  $grpc.ResponseFuture<$2.Client> getClient($2.ClientId request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getClient, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> upsertActivity($0.Activity request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$upsertActivity, request, options: options);
  }

  $grpc.ResponseFuture<$1.Empty> upsertActivityFromCSV($0.Activity request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$upsertActivityFromCSV, request, options: options);
  }

  $grpc.ResponseFuture<$0.Activity> getActivity($2.ActivityId request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActivity, request, options: options);
  }

  $grpc.ResponseFuture<$2.ActivitiesResponse> getActivities(
      $2.ActivitiesRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActivities, request, options: options);
  }
}

abstract class DatabaseServiceBase extends $grpc.Service {
  $core.String get $name => 'database.Database';

  DatabaseServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.Client, $1.Empty>(
        'UpsertClient',
        upsertClient_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Client.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.UpdateRequest, $2.Client>(
        'UpdateClient',
        updateClient_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UpdateRequest.fromBuffer(value),
        ($2.Client value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ClientId, $2.Client>(
        'GetClient',
        getClient_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ClientId.fromBuffer(value),
        ($2.Client value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Activity, $1.Empty>(
        'UpsertActivity',
        upsertActivity_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Activity.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Activity, $1.Empty>(
        'UpsertActivityFromCSV',
        upsertActivityFromCSV_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Activity.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ActivityId, $0.Activity>(
        'GetActivity',
        getActivity_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ActivityId.fromBuffer(value),
        ($0.Activity value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.ActivitiesRequest, $2.ActivitiesResponse>(
        'GetActivities',
        getActivities_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ActivitiesRequest.fromBuffer(value),
        ($2.ActivitiesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> upsertClient_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Client> request) async {
    return upsertClient(call, await request);
  }

  $async.Future<$2.Client> updateClient_Pre(
      $grpc.ServiceCall call, $async.Future<$2.UpdateRequest> request) async {
    return updateClient(call, await request);
  }

  $async.Future<$2.Client> getClient_Pre(
      $grpc.ServiceCall call, $async.Future<$2.ClientId> request) async {
    return getClient(call, await request);
  }

  $async.Future<$1.Empty> upsertActivity_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Activity> request) async {
    return upsertActivity(call, await request);
  }

  $async.Future<$1.Empty> upsertActivityFromCSV_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Activity> request) async {
    return upsertActivityFromCSV(call, await request);
  }

  $async.Future<$0.Activity> getActivity_Pre(
      $grpc.ServiceCall call, $async.Future<$2.ActivityId> request) async {
    return getActivity(call, await request);
  }

  $async.Future<$2.ActivitiesResponse> getActivities_Pre($grpc.ServiceCall call,
      $async.Future<$2.ActivitiesRequest> request) async {
    return getActivities(call, await request);
  }

  $async.Future<$1.Empty> upsertClient(
      $grpc.ServiceCall call, $2.Client request);
  $async.Future<$2.Client> updateClient(
      $grpc.ServiceCall call, $2.UpdateRequest request);
  $async.Future<$2.Client> getClient(
      $grpc.ServiceCall call, $2.ClientId request);
  $async.Future<$1.Empty> upsertActivity(
      $grpc.ServiceCall call, $0.Activity request);
  $async.Future<$1.Empty> upsertActivityFromCSV(
      $grpc.ServiceCall call, $0.Activity request);
  $async.Future<$0.Activity> getActivity(
      $grpc.ServiceCall call, $2.ActivityId request);
  $async.Future<$2.ActivitiesResponse> getActivities(
      $grpc.ServiceCall call, $2.ActivitiesRequest request);
}
