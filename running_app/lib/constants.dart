class ApiConstants {
  static const String scheme = 'https';
  static String baseURL = const String.fromEnvironment('HTTP_GATEWAY_BASE_URL',
      defaultValue: 'janmeckelholt.de');
  static int port =
      const int.fromEnvironment('HTTP_GATEWAY_PORT', defaultValue: 443);
  //static const String baseURL = 'localhost';
  //static const String baseURL = 'homepage-server';
  //static const String prefix = '';
  //static const int port = 444;

  //static const int port = 81;

  static const String baseURLAndroidEmulator = '10.0.2.2';
  static const String prefix = '/run';
  static const String summaryPath = '$prefix/weeksummary';
  static const String loginPath = '$prefix/website';
  static const String clientId = '77376';
  static const int timeout = 5;
}

class Credentials {
  static const String technicalUser = 'running_app';
  static const String runningAppPassword =
      String.fromEnvironment('RUNNING_APP_PASSWORD', defaultValue: '');
}
