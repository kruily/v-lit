
/// 特殊内容类型枚举
enum SpecialContentType {
  card,
  node,
}

/// 展示形式枚举
enum DisplayType {
  text,
  person,
  item,
}

/// 特殊内容数据模型
class SpecialContent {
  final SpecialContentType type;
  final DisplayType displayType;
  final String id;
  final String displayText;

  SpecialContent({
    required this.type,
    required this.displayType,
    required this.id,
    required this.displayText,
  });

  factory SpecialContent.fromString(String content) {
    final parts = content.split('|');
    if (parts.length < 4) throw FormatException('Invalid special content format');

    return SpecialContent(
      type: _parseContentType(parts[0]),
      displayType: _parseDisplayType(parts[1]),
      id: parts[2].split('=')[1],
      displayText: parts[3],
    );
  }

  static SpecialContentType _parseContentType(String type) {
    switch (type.toUpperCase()) {
      case 'CARD':
        return SpecialContentType.card;
      case 'NODE':
        return SpecialContentType.node;
      default:
        throw FormatException('Unknown content type: $type');
    }
  }

  static DisplayType _parseDisplayType(String type) {
    switch (type.toUpperCase()) {
      case 'TEXT':
        return DisplayType.text;
      case 'PERSON':
        return DisplayType.person;
      case 'ITEM':
        return DisplayType.item;
      default:
        throw FormatException('Unknown display type: $type');
    }
  }
} 