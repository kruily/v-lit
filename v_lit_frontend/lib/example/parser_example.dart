import 'package:flutter/material.dart';
import '../core/content_parser/parser.dart';
import '../core/content_parser/content_parser.dart';

class ParserExample extends StatelessWidget {
  final ContentParser parser = ContentParser();

  ParserExample({super.key}) {
    // 注册卡片文本显示构建器
    parser.registerBuilder(
      SpecialContentType.card,
      DisplayType.text,
      (content) => Container(
        padding: const EdgeInsets.all(8),
        decoration: BoxDecoration(
          border: Border.all(color: Colors.blue),
          borderRadius: BorderRadius.circular(4),
          color: Colors.blue.withOpacity(0.1),
        ),
        child: Text(
          content.displayText,
          style: const TextStyle(
            color: Colors.blue,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    const String content = '我这是一个内容，他展示了卡片的用法：{CARD|TEXT|card_id=1|水果}';
    
    return Container(
      color: Colors.white,
      child: Wrap(
        children: parser.buildContent(content),
      ),
    );
  }
} 