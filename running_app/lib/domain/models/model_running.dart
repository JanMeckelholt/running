class RunningResponse {
  final String weeklyClimb;

  const RunningResponse({
    required this.weeklyClimb,
  });

  factory RunningResponse.fromJson(List<dynamic> json) {
    return RunningResponse(
      weeklyClimb: json[0]['text'],
    );
  }
}
