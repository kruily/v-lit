import 'package:flutter/material.dart';
import 'content_parser.dart';

typedef SpecialContentBuilder = Widget Function(SpecialContent content);

class ContentParser {
  final Map<SpecialContentType, Map<DisplayType, SpecialContentBuilder>> _builders = {};
  
  /// 注册特殊内容的构建器
  void registerBuilder(
    SpecialContentType type,
    DisplayType displayType,
    SpecialContentBuilder builder,
  ) {
    _builders[type] ??= {};
    _builders[type]![displayType] = builder;
  }

  /// 解析文本内容
  List<dynamic> parse(String content) {
    final List<dynamic> result = [];
    final RegExp regex = RegExp(r'{([^}]+)}');
    int lastIndex = 0;

    for (final Match match in regex.allMatches(content)) {
      // 添加特殊内容之前的普通文本
      if (match.start > lastIndex) {
        result.add(content.substring(lastIndex, match.start));
      }

      // 解析并添加特殊内容
      final specialContent = SpecialContent.fromString(match.group(1)!);
      result.add(specialContent);

      lastIndex = match.end;
    }

    // 添加最后一段普通文本
    if (lastIndex < content.length) {
      result.add(content.substring(lastIndex));
    }

    return result;
  }

  /// 将解析结果转换为Widget列表
  List<Widget> buildContent(String content) {
    final parsed = parse(content);
    return parsed.map((element) {
      if (element is String) {
        return Text(element);
      } else if (element is SpecialContent) {
        final builder = _builders[element.type]?[element.displayType];
        if (builder == null) {
          // 为特殊内容添加默认样式
          return Text(
            element.displayText,
            style: const TextStyle(
              color: Colors.blue,
              decoration: TextDecoration.underline,
              fontWeight: FontWeight.bold,
            ),
          );
        }
        return builder(element);
      }
      return const SizedBox.shrink();
    }).toList();
  }
} 