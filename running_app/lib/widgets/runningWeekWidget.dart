import 'package:flutter/material.dart';
import 'package:running_app/domain/models/runningModel.dart';

Expanded runnigWeekList(RunningWeek runningWeek) {
  return Expanded(
      child: (ListView(
    children: [
      ListTile(
        title: Text('Distance'),
        leading: Icon(Icons.label),
        trailing: Text("${runningWeek.distance / 1000} km"),
      ),
      ListTile(
        title: Text('Climb'),
        leading: Icon(Icons.label),
        trailing: Text("${runningWeek.climb} m"),
      ),
      ListTile(
        title: Text('Nubmer of Runs'),
        leading: Icon(Icons.label),
        trailing: Text("${runningWeek.numberOfRuns}"),
      ),
      ListTile(
        title: Text('Time spent running'),
        leading: Icon(Icons.label),
        trailing: Text(runningWeek.timeStr),
      ),
      ListTile(
        title: Text('Number of other Activities'),
        leading: Icon(Icons.label),
        trailing: Text("${runningWeek.numberOfOthers}"),
      ),
    ],
  )));
}
