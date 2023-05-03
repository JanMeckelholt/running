class RunningWeek {
  final int climb;
  final int distance;
  final int numberOfRuns;
  final int numberOfOthers;
  final int timeUnix;
  final String timeStr;

  const RunningWeek({
    required this.climb,
    required this.distance,
    required this.numberOfRuns,
    required this.numberOfOthers,
    required this.timeUnix,
    required this.timeStr,
  });

  factory RunningWeek.fromJson(List<dynamic> json) {
    var currentWeek = json[0];
    return RunningWeek(
      climb: currentWeek['climb'] ?? 0,
      distance: currentWeek['distance'] ?? 0,
      numberOfRuns: currentWeek['numberOfRuns'] ?? 0,
      numberOfOthers: currentWeek['numberOfOthers'] ?? 0,
      timeUnix: currentWeek['timeUnix'] ?? 0,
      timeStr: currentWeek['timeStr'] ?? "",
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
