import 'dart:convert';
import 'dart:developer';

import 'package:http/browser_client.dart';
import 'package:http/http.dart' as http;
import 'package:http/http.dart';
import 'package:running_app/utils/utils.dart';

import '../../../constants.dart';
import '../../models/runningModel.dart';
import 'runningApi.dart';

class WebRunningApiService implements RunningApiService {
  static final apiClient = http.Client() as BrowserClient
    ..withCredentials = true;
  WebRunningApiService._privateConstructor();
  static final RunningApiService instance =
      WebRunningApiService._privateConstructor();

  @override
  Future<RunningWeek> fetchRunningResponse(int _weekIndex) async {
    Map<String, String> headers = {};

    var httpUriRunningResponse = Uri(
        scheme: 'http',
        host: ApiConstants.baseURL,
        port: ApiConstants.port,
        path: ApiConstants.summaryPath,
        queryParameters: {
          'client': ApiConstants.clientId,
          'week': '$_weekIndex'
        });
    log('httpUri: $httpUriRunningResponse');

    Response response =
        await apiClient.get(httpUriRunningResponse, headers: headers);
    log("ResponseCode: ${response.statusCode}");
    if (response.statusCode == 401) {
      String updatedCookie = await _refreshCookie();
      if (updatedCookie != "") {
        headers['cookie'] = updatedCookie;

        response =
            await apiClient.get(httpUriRunningResponse, headers: headers);
        log("Response with updated coockie ${response.body} - ${response.statusCode}");
      }
      log("ResponseCode after 401: ${response.statusCode} -> updatedCookie: $updatedCookie");
    }
    if (response.statusCode == 200) {
      return RunningWeek.fromJson(jsonDecode("[${response.body}]"));
    }
    throw Exception('Failed to get running response');
  }

  Future<String> _refreshCookie() async {
    log("refresCookie");
    var httpUriLogin = Uri(
        scheme: 'http',
        host: ApiConstants.baseURL,
        port: ApiConstants.port,
        path: ApiConstants.loginPath);
    log('httpUri: $httpUriLogin');
    Response response = await apiClient.get(httpUriLogin);

    WebsiteResponse wsResponse =
        WebsiteResponse.fromJson(jsonDecode(response.body));
    if (wsResponse.LatestWebsitePing == "") {
      log("did not get LatestWebsitePing");
      return "";
    }
    EncryptData ed = EncryptData();
    String latestPingEncrypted = ed.encryptAES(
        wsResponse.LatestWebsitePing, Credentials.runningAppPassword);

    final body = jsonEncode({
      "username": Credentials.technicalUser,
      "LatestPingEncrypted": latestPingEncrypted
    });
    response = await apiClient.post(httpUriLogin, body: body);
    log("postResponse ${response.body} - ${response.statusCode} - ${response.headers}");

    String? c = response.headers['set-cookie'];
    if (c != null) {
      int index = c.indexOf(';');
      return (index == -1) ? c : c.substring(0, index);
    }
    return "";
  }
}

@override
RunningApiService getRunningApiService() => WebRunningApiService.instance;
