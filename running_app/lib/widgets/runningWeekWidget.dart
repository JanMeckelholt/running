import 'package:flutter/material.dart';
import 'package:running_app/domain/models/runningModel.dart';

ListView runnigWeekList(RunningWeek runningWeek) {
  const textStyle = TextStyle(color: Colors.black, fontSize: 18);
  return ListView(
    shrinkWrap: true,
    children: [
      ListTile(
        title: const Text(
          'Distance',
          style: textStyle,
        ),
        leading: const Icon(Icons.label),
        trailing: Text(
          "${runningWeek.distance / 1000} km",
          style: textStyle,
        ),
      ),
      ListTile(
        title: const Text(
          'Climb',
          style: textStyle,
        ),
        leading: const Icon(Icons.label),
        trailing: Text(
          "${runningWeek.climb} m",
          style: textStyle,
        ),
      ),
      ListTile(
        title: const Text(
          'Nubmer of Runs',
          style: textStyle,
        ),
        leading: const Icon(Icons.label),
        trailing: Text(
          "${runningWeek.numberOfRuns}",
          style: textStyle,
        ),
      ),
      ListTile(
        title: const Text(
          'Time spent running',
          style: textStyle,
        ),
        leading: const Icon(Icons.label),
        trailing: Text(
          runningWeek.timeStr,
          style: textStyle,
        ),
      ),
      ListTile(
        title: const Text(
          'Number of other Activities',
          style: textStyle,
        ),
        leading: const Icon(Icons.label),
        trailing: Text(
          "${runningWeek.numberOfOthers}",
          style: textStyle,
        ),
      ),
    ],
  );
}
