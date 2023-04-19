import '../../models/model_running.dart';
import 'package:http/http.dart' as http;

import './runningApiStub.dart'
    // xignore: uri_does_not_exist
    //if (dart.library.io) 'package:flutter_conditional_dependencies_example/storage/shared_pref_key_finder.dart'

    if (dart.library.html) './webRunningApi.dart';

abstract class RunningApiService {
  //final http.Client client;

  //RunningApiService(this.client);

  Future<RunningResponse> fetchRunningResponse();

  factory RunningApiService() => getRunningApiService();
}
