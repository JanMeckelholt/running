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
    _runningWeek = apiService.fetchRunningResponse(_weekIndex);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(widget.title),
        ),
        body: FutureBuilder<RunningWeek>(
            future: _runningWeek,
            builder:
                (BuildContext context, AsyncSnapshot<RunningWeek> snapshot) {
              if (snapshot.data != null) {
                return Container(
                    color: Colors.white70,
                    child: Column(children: [
                      Row(
                        //mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                        mainAxisSize: MainAxisSize.max,
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: <Widget>[
                          Expanded(child: Container(color: Colors.green)),
                          Container(
                            decoration: BoxDecoration(
                                border:
                                    Border.all(color: Colors.black, width: 2)),
                            child: SingleChildScrollView(
                              child: ConstrainedBox(
                                  constraints:
                                      const BoxConstraints(maxHeight: 400),
                                  child: SingleChildScrollView(
                                      scrollDirection: Axis.horizontal,
                                      child: ConstrainedBox(
                                          constraints: const BoxConstraints(
                                              maxWidth: 400),
                                          child: Column(
                                            mainAxisSize: MainAxisSize.max,
                                            children: <Widget>[
                                              Row(
                                                  mainAxisAlignment:
                                                      MainAxisAlignment.center,
                                                  children: <Widget>[
                                                    Padding(
                                                        padding:
                                                            const EdgeInsets
                                                                .all(20),
                                                        child: ElevatedButton(
                                                          onPressed: () {
                                                            setState(() {
                                                              _weekIndex--;
                                                              _callApi(widget
                                                                  .apiService);
                                                            });
                                                          },
                                                          child: const Icon(
                                                              Icons.remove),
                                                        )),
                                                    Text(
                                                      'Current week: $_weekIndex',
                                                      style: const TextStyle(
                                                          fontSize: 24,
                                                          fontWeight:
                                                              FontWeight.bold),
                                                    ),
                                                    Padding(
                                                        padding:
                                                            const EdgeInsets
                                                                .all(20),
                                                        child: ElevatedButton(
                                                          onPressed: () {
                                                            setState(() {
                                                              if (_weekIndex <
                                                                  0) {
                                                                _weekIndex++;
                                                                _callApi(widget
                                                                    .apiService);
                                                              }
                                                            });
                                                          },
                                                          child: const Icon(
                                                              Icons.add),
                                                        )),
                                                  ]),
                                              Text(
                                                "Week starts at ${snapshot.data!.startOfWeekStr}",
                                                style: const TextStyle(
                                                    fontSize: 24,
                                                    fontWeight:
                                                        FontWeight.bold),
                                              ),
                                              const Text(
                                                'Summary: ',
                                                style: TextStyle(
                                                    fontSize: 24,
                                                    fontWeight:
                                                        FontWeight.bold),
                                              ),
                                              runnigWeekList(snapshot.data!),
                                            ],
                                          )))),
                            ),
                          ),
                          Expanded(child: Container(color: Colors.lightBlue)),
                        ],
                      )
                    ]));
              } else {
                return Center(
                  child: Column(
                    children: const [
                      CircularProgressIndicator(),
                      Text("Load data, please wait...")
                    ],
                  ),
                );
              }
            }));
  }
}
