class RunningResponse {
  final int weeklyClimb;

  const RunningResponse({
    required this.weeklyClimb,
  });

  factory RunningResponse.fromJson(Map<String, List<dynamic>> json) {
    return RunningResponse(
      weeklyClimb: json['WeekSummaries']?[0]['Climb'],
    );
  }
}
