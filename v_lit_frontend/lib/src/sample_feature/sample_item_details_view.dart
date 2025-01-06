import 'package:flutter/material.dart';
import 'package:v_lit_frontend/core/content_parser/parser.dart';

/// Displays detailed information about a SampleItem.
class SampleItemDetailsView extends StatelessWidget {
  const SampleItemDetailsView({super.key});

  static const routeName = '/sample_item';
  static final ContentParser parser = ContentParser();
  static const String content = '我这是一个内容，他展示了卡片的用法：{CARD|TEXT|card_id=1|水果},我是一个节点：{NODE|TEXT|node_id=1|节点}';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Item Details'),
      ),
      body: Center(
        child: Container(
          color: Colors.white,
          child: Wrap(
            children: parser.buildContent(content),
          ),
        ),
      ),
    );
  }
}
