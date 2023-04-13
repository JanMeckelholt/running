import 'dart:convert';
import 'package:running_app/constants.dart';
import 'package:running_app/domain/models/model_running.dart';
import 'package:http/http.dart' as http;

import 'dart:developer';

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
