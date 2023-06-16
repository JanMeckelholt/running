import 'dart:async';

import 'package:flutter/material.dart';
import 'package:running_app/domain/models/runningModel.dart';
import '../domain/services/api/runningApi.dart';
import '../widgets/runningWeekWidget.dart';

class Main extends StatefulWidget {
  const Main({super.key, required this.title, required this.apiService});
  final String title;
  final RunningApiService apiService;

  @override
  State<Main> createState() => _MainState();
}

class _MainState extends State<Main> {
  @override
  void initState() {
    super.initState();
    _callApi(widget.apiService);
  }

  late Future<RunningWeek> _runningWeek;
  int _weekIndex = 0;

  Future<void> _callApi(RunningApiService apiService) async {
    // _runningWeek = Future.delayed(const Duration(seconds: 1), () {
    //   return const RunningWeek(
    //       climb: 1,
    //       distance: 2,
    //       numberOfRuns: 3,
    //       numberOfOthers: 4,
    //       timeUnix: 11111,
    //       timeStr: "2h40m10s",
    //       startOfWeekStr: "1.02.2000");
    // });

    _runningWeek = apiService.fetchRunningResponse(_weekIndex);
  }

  @override
  Widget build(BuildContext context) {
    return Material(
        child: Center(
            child: ConstrainedBox(
                constraints:
                    const BoxConstraints(maxHeight: 400, maxWidth: 400),
                child: FutureBuilder<RunningWeek>(
                    future: _runningWeek,
                    builder: (BuildContext context,
                        AsyncSnapshot<RunningWeek> snapshot) {
                      if (snapshot.data != null) {
                        return SingleChildScrollView(
                          child: Column(
                            mainAxisSize: MainAxisSize.max,
                            children: <Widget>[
                              Row(
                                  mainAxisAlignment: MainAxisAlignment.center,
                                  children: <Widget>[
                                    Padding(
                                        padding: const EdgeInsets.all(20),
                                        child: ElevatedButton(
                                          onPressed: () {
                                            setState(() {
                                              _weekIndex--;
                                              _callApi(widget.apiService);
                                            });
                                          },
                                          child: const Icon(Icons.remove),
                                        )),
                                    Text(
                                      'Current week: $_weekIndex',
                                      style: const TextStyle(
                                          fontSize: 24,
                                          fontWeight: FontWeight.bold),
                                    ),
                                    Padding(
                                        padding: const EdgeInsets.all(20),
                                        child: ElevatedButton(
                                          onPressed: () {
                                            setState(() {
                                              if (_weekIndex < 0) {
                                                _weekIndex++;
                                                _callApi(widget.apiService);
                                              }
                                            });
                                          },
                                          child: const Icon(Icons.add),
                                        )),
                                  ]),
                              Text(
                                "Week starts at ${snapshot.data!.startOfWeekStr}",
                                style: const TextStyle(
                                    fontSize: 24, fontWeight: FontWeight.bold),
                              ),
                              const Text(
                                'Summary: ',
                                style: TextStyle(
                                    fontSize: 24, fontWeight: FontWeight.bold),
                              ),
                              runnigWeekList(snapshot.data!),
                            ],
                          ),
                        );
                      } else {
                        return Column(
                          children: const [
                            CircularProgressIndicator(),
                            Text("Load data, please wait...")
                          ],
                        );
                      }
                    }))));
  }
}
