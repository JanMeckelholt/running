class RunningWeek {
  final int climb;
  final int distance;
  final int numberOfRuns;
  final int numberOfOthers;
  final int timeUnix;
  final String timeStr;
  final String startOfWeekStr;

  const RunningWeek({
    required this.climb,
    required this.distance,
    required this.numberOfRuns,
    required this.numberOfOthers,
    required this.timeUnix,
    required this.timeStr,
    required this.startOfWeekStr,
  });

  factory RunningWeek.fromJson(Map<String, dynamic> json) {
    var currentWeek = json;
    return RunningWeek(
      climb: currentWeek['climb'] ?? 0,
      distance: currentWeek['distance'] ?? 0,
      numberOfRuns: currentWeek['numberOfRuns'] ?? 0,
      numberOfOthers: currentWeek['numberOfOthers'] ?? 0,
      timeUnix: currentWeek['timeUnix'] ?? 0,
      timeStr: currentWeek['timeStr'] ?? "",
      startOfWeekStr: currentWeek['startOfWeekStr'] ?? "",
    );
  }
}

class WebsiteResponse {
  final String LatestWebsitePing;
  const WebsiteResponse({required this.LatestWebsitePing});
  factory WebsiteResponse.fromJson(Map<String, dynamic> data) {
    final LatestWebsitePing = data['LatestWebsitePing'] as String;
    return WebsiteResponse(LatestWebsitePing: LatestWebsitePing);
  }
}
