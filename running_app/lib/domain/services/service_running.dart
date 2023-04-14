import 'dart:convert';
import 'dart:html';
import 'package:http/browser_client.dart';
import 'package:running_app/constants.dart';
import 'package:running_app/domain/models/model_running.dart';
import 'package:http/http.dart' as http;

import 'dart:developer';

Future<RunningResponse> fetchRunningResponse() async {
  Map<String, String> headers = {};
  var httpUriRunningResponse = Uri(
      scheme: 'http',
      host: ApiConstants.baseURL,
      port: ApiConstants.port,
      path: ApiConstants.summaryPath,
      queryParameters: {'client': ApiConstants.clientId, 'weeks': '0'});
  log('httpUri: $httpUriRunningResponse');
  var client = BrowserClient()..withCredentials = true;
  http.Response response;
  try {
    response = await client.get(
      httpUriRunningResponse,
      headers: headers,
    );

    if (response.statusCode == 401) {
      final updatedCookie = await refreshCookie(client);
      if (updatedCookie != "") {
        headers['cookie'] = updatedCookie;
        response = await client.get(httpUriRunningResponse, headers: headers);
      }
    }
    if (response.statusCode == 200) {
      return RunningResponse.fromJson(jsonDecode("[${response.body}]"));
    }
  } finally {
    client.close();
  }
  throw Exception('Failed to get running response');
}

Future<String> refreshCookie(BrowserClient client) async {
  var httpUriLogin = Uri(
      scheme: 'http',
      host: ApiConstants.baseURL,
      port: ApiConstants.port,
      path: ApiConstants.loginPath);
  log('httpUri: $httpUriLogin');
  final body = jsonEncode({
    "username": Credentials.technicalUser,
    "password": Credentials.runningAppPassword ?? ""
  });
  final response = await client.post(httpUriLogin, body: body);
  String? rawCookie = response.headers['set-cookie'];
  if (rawCookie != null) {
    int index = rawCookie.indexOf(';');
    return (index == -1) ? rawCookie : rawCookie.substring(0, index);
  }
  return "";
}
