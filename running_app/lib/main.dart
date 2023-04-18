import 'dart:io';

import 'package:flutter/material.dart';
import 'package:http/browser_client.dart';
import 'package:http/io_client.dart';
import 'package:running_app/constants.dart';
import 'package:running_app/screens/screen_main.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

import 'domain/services/service_api.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await dotenv.load(fileName: ".env.docker.running_app.secret");
  Credentials.runningAppPassword = dotenv.env['RUNNING_APP_PASSWORD'];
  final client = http.Client() as BrowserClient..withCredentials = true;
  final apiService = RunningApiService(client);
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
