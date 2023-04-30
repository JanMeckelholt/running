import 'dart:convert';
import 'dart:developer';
import 'dart:io';

import 'package:flutter/services.dart';
import 'package:http/io_client.dart';
import 'package:running_app/utils/utils.dart';

import '../../../constants.dart';
import '../../models/runningModel.dart';
import 'runningApi.dart';

class AndroidRunningApiService implements RunningApiService {
  AndroidRunningApiService._privateConstructor();
  static final AndroidRunningApiService instance =
      AndroidRunningApiService._privateConstructor();

  @override
  Future<RunningWeek> fetchRunningResponse() async {
    ByteData rootCACertificate = await rootBundle.load("certs/ca-cert.pem");
    ByteData clientCertificate =
        await rootBundle.load("certs/running_app-cert.pem");
    ByteData privateKey = await rootBundle.load("certs/running_app-key.pem");
    SecurityContext context = SecurityContext(withTrustedRoots: true);

    context.setTrustedCertificatesBytes(rootCACertificate.buffer.asUint8List());

    //context.useCertificateChainBytes(clientCertificate.buffer.asUint8List());

    //context.usePrivateKeyBytes(privateKey.buffer.asUint8List());
    IOClient apiClient = IOClient(HttpClient(context: context));

    Map<String, String> headers = {};

    var httpUriRunningResponse = Uri(
        scheme: 'http',
        host: ApiConstants.baseURLAndroidEmulator,
        port: ApiConstants.port,
        path: ApiConstants.summaryPath,
        queryParameters: {'client': ApiConstants.clientId, 'week': '0'});
    log('httpUri: $httpUriRunningResponse');

    var response =
        await apiClient.get(httpUriRunningResponse, headers: headers);

    if (response.statusCode == 401) {
      final updatedCookie = await refreshCookie(apiClient);
      if (updatedCookie != "") {
        headers['cookie'] = updatedCookie;
        response =
            await apiClient.get(httpUriRunningResponse, headers: headers);
      }
    }
    if (response.statusCode == 200) {
      return RunningWeek.fromJson(jsonDecode("[${response.body}]"));
    }
    throw Exception('Failed to get running response');
  }

  Future<String> refreshCookie(IOClient apiClient) async {
    var httpUriLogin = Uri(
        scheme: 'http',
        host: ApiConstants.baseURLAndroidEmulator,
        port: ApiConstants.port,
        path: ApiConstants.loginPath);
    log('httpUri: $httpUriLogin');
    var response = await apiClient.get(httpUriLogin);

    WebsiteResponse wsResponse =
        WebsiteResponse.fromJson(jsonDecode(response.body));
    if (wsResponse.LatestWebsitePing == "") {
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

    String? c = response.headers['set-cookie'];
    if (c != null) {
      int index = c.indexOf(';');
      return (index == -1) ? c : c.substring(0, index);
    }
    return "";
  }
}

@override
RunningApiService getRunningApiService() => AndroidRunningApiService.instance;
