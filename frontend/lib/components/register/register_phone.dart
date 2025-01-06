import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';

class RegisterPhoneComponent extends StatefulWidget {
  const RegisterPhoneComponent({super.key, required this.onFormTrigger, required this.onLoginHandler});

  final Function(bool) onFormTrigger;
  final Function(bool) onLoginHandler;

  @override
  State<RegisterPhoneComponent> createState() => _RegisterPhoneComponentState();
}

class _RegisterPhoneComponentState extends State<RegisterPhoneComponent> {
  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text('欢迎注册',
            style: TextStyle(
              fontSize: 32,
              fontFamily: 'Helvetica',
              fontWeight: FontWeight.bold,
            )),
        SizedBox(height: 40),
        // Lottie 动画
        Lottie.asset('assets/lottie_files/cat_rolling.json',
            width: 100, height: 100),
        // 注册选项
        Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 300,
              child: TextField(
                decoration: InputDecoration(
                  labelText: '邮箱',
                  suffixIcon: Icon(Icons.email),
                  labelStyle: TextStyle(
                      color: Colors.black,
                      fontSize: 16,
                      fontFamily: 'Helvetica'),
                ),
              ),
            ),
            SizedBox(height: 20),
            SizedBox(
              width: 300,
              child: TextField(
                decoration: InputDecoration(
                  labelText: '用户名',
                  suffixIcon: Icon(Icons.person),
                  labelStyle: TextStyle(
                      color: Colors.black,
                      fontSize: 16,
                      fontFamily: 'Helvetica'),
                ),
              ),
            ),
            SizedBox(height: 20),
            SizedBox(
              width: 300,
              child: TextField(
                decoration: InputDecoration(
                  labelText: '密码',
                  suffixIcon: Icon(Icons.lock),
                  labelStyle: TextStyle(
                      color: Colors.black,
                      fontSize: 16,
                      fontFamily: 'Helvetica'),
                ),
              ),
            ),
            SizedBox(height: 20),
            SizedBox(
              width: 300,
              child: TextField(
                decoration: InputDecoration(
                  labelText: '验证码',
                  suffixIcon: Icon(Icons.code),
                  labelStyle: TextStyle(
                      color: Colors.black,
                      fontSize: 16,
                      fontFamily: 'Helvetica'),
                ),
              ),
            ),
            SizedBox(height: 20),
            SizedBox(
              width: 300,
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  TextButton(
                    onPressed: () {
                      widget.onFormTrigger(false);
                    },
                    child: Text(
                      '已有账号?立即登录',
                      style: TextStyle(
                          color: Colors.blue,
                          decoration: TextDecoration.underline,
                          fontSize: 12),
                    ),
                  ),
                  ElevatedButton(
                    onPressed: () {
                      // 注册逻辑
                      widget.onLoginHandler(true);
                    },
                    child: Text('注册',
                        style: TextStyle(
                            fontSize: 16,
                            fontFamily: 'Helvetica',
                            fontWeight: FontWeight.bold)),
                  ),
                ],
              ),
            ),
          ],
        ),
      ],
    );
  }
}