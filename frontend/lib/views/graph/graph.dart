import 'package:flutter/material.dart';
import 'package:frontend/components/appbar/custom_appbar.dart';

class GraphPage extends StatefulWidget {
  const GraphPage({super.key});

  @override
  State<GraphPage> createState() => _GraphPageState();
}

class _GraphPageState extends State<GraphPage> {

  /// Notifier for the tension slider
  final segmentedTension = ValueNotifier<double>(1);
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const CustomAppbarComponent(),
      body: Text("test")
    );
  }
}
