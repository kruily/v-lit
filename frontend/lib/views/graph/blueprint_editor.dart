import 'package:flutter/material.dart';
import 'package:graphview/GraphView.dart';

class BlueprintEditor extends StatefulWidget {
  @override
  _BlueprintEditorState createState() => _BlueprintEditorState();
}

class _BlueprintEditorState extends State<BlueprintEditor> {
  final Graph graph = Graph();
  int nodeCount = 0;

  @override
  void initState() {
    super.initState();
    graph.addNode(Node.Id(0));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('蓝图编辑器'),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Node node = Node.Id("${nodeCount+1}");
          graph.addNode(node);
        },
        child: Icon(Icons.add),
      ),
      body: InteractiveViewer(
        constrained: false,
        boundaryMargin: EdgeInsets.all(100),
        minScale: 0.01,
        maxScale: 5.6,
        child: GraphView(
          graph: graph,
          algorithm: FruchtermanReingoldAlgorithm(),
          builder: (Node node) {
            // 这个 builder 是必需的，即使没有节点
            var nodeId = node.key?.value as int;
            return nodeWidget(nodeId);
          },
        ),
      ),
    );
  }

  Widget nodeWidget(int text) {
    return Container(
      padding: EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(8),
        boxShadow: [
          BoxShadow(
            color: Colors.grey.withOpacity(0.5),
            spreadRadius: 1,
            blurRadius: 2,
            offset: Offset(0, 3),
          ),
        ],
      ),
      child: Text('节点 $text'),
    );
  }
}
