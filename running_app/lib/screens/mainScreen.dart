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

  Future<RunningWeek> _runningWeek = Future<RunningWeek>.value(RunningWeek(
      climb: 0,
      distance: 0,
      numberOfRuns: 0,
      numberOfOthers: 0,
      timeUnix: 0,
      timeStr: ""));

  Future<void> _callApi(RunningApiService apiService) async {
    var response = await apiService.fetchRunningResponse();
    setState(() {
      _runningWeek = response as Future<RunningWeek>;
    });
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
                return Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: <Widget>[
                      const Text(
                        'Current week: ',
                      ),
                      runnigWeekList(snapshot.data!),
                      ElevatedButton(
                        onPressed: () => _callApi(widget.apiService),
                        child: const Icon(Icons.refresh),
                      ),
                    ],
                  ),
                );
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
