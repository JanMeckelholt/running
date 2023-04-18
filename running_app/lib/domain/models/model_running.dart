class RunningResponse {
  final int weeklyClimb;

  const RunningResponse({
    required this.weeklyClimb,
  });

  factory RunningResponse.fromJson(List<dynamic> json) {
    return RunningResponse(
      weeklyClimb: json[0]['weeksummaries']?[0]['climb'] ?? "0",
    );
  }
}
