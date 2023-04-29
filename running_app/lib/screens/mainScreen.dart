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
    //_callApi(widget.apiService);
  }

  late Future<RunningWeek> _runningWeek =
      widget.apiService.fetchRunningResponse();

  Future<void> _callApi(RunningApiService apiService) async {
    var response = await apiService.fetchRunningResponse();
    setState(() {
      _runningWeek = Future<RunningWeek>.value(response);
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
                return Container(
                    color: Colors.amber,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      mainAxisSize: MainAxisSize.min,
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: <Widget>[
                        Expanded(child: Container(color: Colors.green)),
                        SingleChildScrollView(
                          child: ConstrainedBox(
                              constraints: const BoxConstraints(
                                  minHeight: 600, maxHeight: 800),
                              child: SingleChildScrollView(
                                  scrollDirection: Axis.horizontal,
                                  child: ConstrainedBox(
                                      constraints: const BoxConstraints(
                                          minWidth: 300, maxWidth: 400),
                                      child: Column(
                                        mainAxisSize: MainAxisSize.max,
                                        children: <Widget>[
                                          const Text(
                                            'Current week: ',
                                            style: TextStyle(
                                                fontSize: 24,
                                                fontWeight: FontWeight.bold),
                                          ),
                                          runnigWeekList(snapshot.data!),
                                          ElevatedButton(
                                            onPressed: () =>
                                                _callApi(widget.apiService),
                                            child: const Icon(Icons.refresh),
                                          ),
                                          Flexible(
                                              child: Container(
                                            color: Colors.blueGrey,
                                          ))
                                        ],
                                      )))),
                        ),
                        Expanded(child: Container(color: Colors.lightBlue)),
                      ],
                    ));
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
