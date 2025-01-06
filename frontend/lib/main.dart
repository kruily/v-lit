import 'package:flutter/material.dart';
import 'package:frontend/common/router/router_manager.dart';
import 'package:frontend/views/graph/graph.dart';
// import 'package:frontend/views/home/home.dart';
// import 'package:frontend/views/login/login.dart';

void main() {
  RouterManager.defineRoutes();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const GraphPage(),
      onGenerateRoute: RouterManager.router.generator,
    );
  }
}