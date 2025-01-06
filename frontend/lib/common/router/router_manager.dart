import 'package:fluro/fluro.dart';
import 'package:frontend/views/home/home.dart';
import 'package:frontend/views/login/login.dart';

class RouterManager {
  static final FluroRouter router = FluroRouter();

  static void defineRoutes() {
    router.define(
      '/',
      handler: Handler(handlerFunc: (context, params) => LoginPage()),
    );

    router.define(
      '/login',
      handler: Handler(handlerFunc: (context, params) => LoginPage()),
    );

    router.define(
      '/home',
      handler: Handler(handlerFunc: (context, params) => HomePage()),
    );
  }
}
