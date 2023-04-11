import 'dart:convert';

import 'package:running_app/domain/models/model_running.dart';
import 'package:http/http.dart' as http;

Future<RunningResponse> fetchRunningResponse() async {
  final response =
      await http.get(Uri.parse('https://cat-fact.herokuapp.com/facts'));

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    return RunningResponse.fromJson(jsonDecode(response.body));
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load album');
  }
}
