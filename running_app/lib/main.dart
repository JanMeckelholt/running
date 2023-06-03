import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:running_app/constants.dart';
import 'package:running_app/domain/services/api/runningApi.dart';
import 'package:running_app/screens/mainScreen.dart';

//import 'domain/services/api/runningApi.dart';
//import 'domain/services/api/webRunningApi.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  //await dotenv.load(fileName: ".env.docker.running_app.secret");
  const runningAppPassword =
      String.fromEnvironment('RUNNING_APP_PASSWORD', defaultValue: '');
  log("RunningAppPassword: $runningAppPassword");
  if (runningAppPassword.length < 8) {
    print("Password too short!");
  }
  Credentials.runningAppPassword = runningAppPassword;

  //final apiService = WebRunningApiService.instance;
  final apiService = RunningApiService();
  runApp(MyApp(apiService: apiService));
}

class MyApp extends StatelessWidget {
  final RunningApiService apiService;
  const MyApp({super.key, required this.apiService});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Running App',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Main(title: 'Running Overview', apiService: apiService),
    );
  }
}
