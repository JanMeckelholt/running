///
//  Generated code. Do not modify.
//  source: database/database.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'database.pb.dart' as $2;
import '../google/protobuf/empty.pb.dart' as $1;
import '../strava/strava.pb.dart' as $0;
import 'database.pbjson.dart';

export 'database.pb.dart';

abstract class DatabaseServiceBase extends $pb.GeneratedService {
  $async.Future<$1.Empty> upsertClient($pb.ServerContext ctx, $2.Client request);
  $async.Future<$2.Client> updateClient($pb.ServerContext ctx, $2.UpdateRequest request);
  $async.Future<$2.Client> getClient($pb.ServerContext ctx, $2.ClientId request);
  $async.Future<$1.Empty> upsertActivity($pb.ServerContext ctx, $0.Activity request);
  $async.Future<$1.Empty> upsertActivityFromCSV($pb.ServerContext ctx, $0.Activity request);
  $async.Future<$0.Activity> getActivity($pb.ServerContext ctx, $2.ActivityId request);
  $async.Future<$2.ActivitiesResponse> getActivities($pb.ServerContext ctx, $2.ActivitiesRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'UpsertClient': return $2.Client();
      case 'UpdateClient': return $2.UpdateRequest();
      case 'GetClient': return $2.ClientId();
      case 'UpsertActivity': return $0.Activity();
      case 'UpsertActivityFromCSV': return $0.Activity();
      case 'GetActivity': return $2.ActivityId();
      case 'GetActivities': return $2.ActivitiesRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'UpsertClient': return this.upsertClient(ctx, request as $2.Client);
      case 'UpdateClient': return this.updateClient(ctx, request as $2.UpdateRequest);
      case 'GetClient': return this.getClient(ctx, request as $2.ClientId);
      case 'UpsertActivity': return this.upsertActivity(ctx, request as $0.Activity);
      case 'UpsertActivityFromCSV': return this.upsertActivityFromCSV(ctx, request as $0.Activity);
      case 'GetActivity': return this.getActivity(ctx, request as $2.ActivityId);
      case 'GetActivities': return this.getActivities(ctx, request as $2.ActivitiesRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => DatabaseServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => DatabaseServiceBase$messageJson;
}

