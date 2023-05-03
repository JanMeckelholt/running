import 'dart:async';

import '../../models/runningModel.dart';

import './runningApiStub.dart'
    if (dart.library.io) './androidRunningApi.dart'
    if (dart.library.html) './webRunningApi.dart';

abstract class RunningApiService {
  Future<RunningWeek> fetchRunningResponse(int _weekIndex) {
    return Future<RunningWeek>.value(RunningWeek(
      climb: 0,
      distance: 0,
      numberOfRuns: 0,
      numberOfOthers: 0,
      timeUnix: 0,
      timeStr: "",
      startOfWeekStr: "",
    ));
  }

  factory RunningApiService() => getRunningApiService();
}
