import 'dart:async';

import '../../models/model_running.dart';

import './runningApiStub.dart'
    if (dart.library.io) './androidRunningApi.dart'
    if (dart.library.html) './webRunningApi.dart';

abstract class RunningApiService {
  Future<RunningResponse> fetchRunningResponse() {
    return Future<RunningResponse>.value(RunningResponse(weeklyClimb: 0));
  }

  factory RunningApiService() => getRunningApiService();
}
