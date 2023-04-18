import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:running_app/domain/models/model_running.dart';
import 'package:running_app/domain/services/service_api.dart';

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

  int _climb = 0;
  late Future<RunningResponse> futureRunningResponse;

  Future<void> _callApi(RunningApiService apiService) async {
    futureRunningResponse = apiService.fetchRunningResponse();
    int climbResponse =
        await futureRunningResponse.then((value) => value.weeklyClimb);
    setState(() {
      _climb = climbResponse;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(widget.title),
        ),
        body: FutureBuilder<RunningResponse>(
            future: widget.apiService.fetchRunningResponse(),
            builder: (BuildContext context,
                AsyncSnapshot<RunningResponse> snapshot) {
              if (snapshot.data != null) {
                return Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: <Widget>[
                      const Text(
                        'Weekly climb since start of the week: ',
                      ),
                      Text(
                        '$_climb m',
                        style: Theme.of(context).textTheme.headlineMedium,
                      ),
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
