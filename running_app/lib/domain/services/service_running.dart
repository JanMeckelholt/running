import 'dart:convert';

import 'package:fixnum/fixnum.dart';

import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';

import 'package:running_app/constants.dart';
import 'package:running_app/domain/models/model_running.dart';
import 'package:http/http.dart' as http;

import 'dart:developer';
import 'package:running_app/gen/grpc/runner/runner.pb.dart' as pbRunner;
import 'package:running_app/gen/grpc/runner/runner.pbgrpc.dart';

Future<RunningResponse> fetchRunningResponse() async {
  var httpUri = Uri(
      scheme: 'http',
      host: ApiConstants.baseURL,
      port: ApiConstants.port,
      path: ApiConstants.summaryPath,
      queryParameters: {'client': ApiConstants.clientId, 'weeks': '0'});
  log('httpUri: $httpUri');
  final response = await http.get(httpUri);

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    return RunningResponse.fromJson(jsonDecode(response.body));
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load album');
  }
}

Future<RunningResponse> fetchRunningResponseGrpc() async {
  var request = pbRunner.WeekSummariesRequest(
      clientId: ApiConstants.clientId, weeks: Int64(0));
  //final bytes = await rootBundle.load('certs/running_app-cert.pem');
  //final bytes = await rootBundle.load('certs/ca-cert.pem');
  final bytes =
      await rootBundle.load('certs/running_app.includesprivatekey.pem');
  ClientChannel channel = ClientChannel(
    'runner',
    port: 666,
    options: ChannelOptions(
      credentials: ChannelCredentials.secure(
        certificates: bytes.buffer.asUint8List(),
        // authority: 'www.janmeckelholt.de',
        onBadCertificate: (_, String s) => true,
      ),
    ),
  );

  RunnerClient client = RunnerClient(channel,
      options: CallOptions(timeout: Duration(seconds: 30)));
  final response = await client.getWeekSummaries(request);
  WeekSummary week_0 = response.$_getList(0) as pbRunner.WeekSummary;
  return RunningResponse(weeklyClimb: week_0.getFieldOrNull(4));
}
