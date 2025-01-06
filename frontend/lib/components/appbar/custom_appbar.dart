import 'package:flutter/material.dart';

class CustomAppbarComponent extends StatefulWidget
    implements PreferredSizeWidget {
  const CustomAppbarComponent({super.key});

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  State<CustomAppbarComponent> createState() => _CustomAppbarComponentState();
}

class _CustomAppbarComponentState extends State<CustomAppbarComponent> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: kToolbarHeight,
      padding: const EdgeInsets.symmetric(horizontal: 16),
      decoration: BoxDecoration(
        color: const Color(0x0ff3fccf),
        boxShadow: [
          BoxShadow(
            color: Colors.grey,
            spreadRadius: 1,
            blurRadius: 3,
            offset: const Offset(0, 1),
          ),
        ],
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const SizedBox(width: 40), // 左侧占位，保持对称
          const Text(
            '首页',
            style: TextStyle(
              color: Colors.black,
              fontSize: 18,
              fontWeight: FontWeight.bold,
            ),
          ),
          Row(
            children: [
              IconButton(
                icon: const Icon(Icons.search, color: Colors.black),
                onPressed: () {
                  // 搜索功能
                },
              ),
              IconButton(
                icon: const Icon(Icons.notifications, color: Colors.black),
                onPressed: () {
                  // 通知功能
                },
              ),
              IconButton(
                icon: const Icon(Icons.person, color: Colors.black),
                onPressed: () {
                  // 个人中心
                },
              ),
            ],
          ),
        ],
      ),
    );
  }
}
