class ApiConstants {
  static const String scheme = 'http';
  static const String baseURL = 'localhost';
  //static const String baseURL = 'http_gateway';
  static const String baseURLAndroidEmulator = '10.0.2.2';
  //static const int port = 443;
  static const int port = 81;
  static const String summaryPath = '/weeksummary';
  static const String loginPath = '/website';
  static const String clientId = '77376';
  static const int timeout = 5;
}

class Credentials {
  static const String technicalUser = 'running_app';
  static String? runningAppPassword;
}
