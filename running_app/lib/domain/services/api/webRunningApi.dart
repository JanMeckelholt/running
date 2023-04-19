import 'dart:convert';
import 'dart:developer';

import 'package:http/browser_client.dart';
import 'package:http/http.dart' as http;

import '../../../constants.dart';
import '../../models/model_running.dart';
import 'runningApi.dart';

class WebRunningApiService implements RunningApiService {
  static final apiClient = http.Client() as BrowserClient
    ..withCredentials = true;

  WebRunningApiService();

  @override
  Future<RunningResponse> fetchRunningResponse() async {
    Map<String, String> headers = {};

    var httpUriRunningResponse = Uri(
        scheme: 'http',
        host: ApiConstants.baseURL,
        port: ApiConstants.port,
        path: ApiConstants.summaryPath,
        queryParameters: {'client': ApiConstants.clientId, 'weeks': '0'});
    log('httpUri: $httpUriRunningResponse');

    var response =
        await apiClient.get(httpUriRunningResponse, headers: headers);

    if (response.statusCode == 401) {
      final updatedCookie = await refreshCookie();
      if (updatedCookie != "") {
        headers['cookie'] = updatedCookie;
        response =
            await apiClient.get(httpUriRunningResponse, headers: headers);
      }
    }
    if (response.statusCode == 200) {
      return RunningResponse.fromJson(jsonDecode("[${response.body}]"));
    }
    throw Exception('Failed to get running response');
  }

  Future<String> refreshCookie() async {
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
    final response = await apiClient.post(httpUriLogin, body: body);

    String? c = response.headers['set-cookie'];
    if (c != null) {
      int index = c.indexOf(';');
      return (index == -1) ? c : c.substring(0, index);
    }
    return "";
  }
}

RunningApiService getRunningApiService() => WebRunningApiService();
