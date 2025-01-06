import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';

class LoginComponent extends StatefulWidget {
  const LoginComponent({super.key, required this.onFormTrigger, required this.onLoginHandler});
  
  final Function(bool) onFormTrigger;
  final Function(bool) onLoginHandler;

  @override
  State<LoginComponent> createState() => _LoginComponentState();
}

class _LoginComponentState extends State<LoginComponent> {
  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        // 左侧 Lottie 动画
        Expanded(
          flex: 1,
          child: Lottie.asset('assets/lottie_files/cat_rolling.json',
              width: 100, height: 100),
        ),
        // 右侧登录选项
        Expanded(
          flex: 1,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text('欢迎登录',
                  style: TextStyle(
                    fontSize: 32,
                    fontFamily: 'Helvetica',
                    fontWeight: FontWeight.bold,
                  )),
              SizedBox(height: 40),
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
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    TextButton(
                      onPressed: () {
                        widget.onFormTrigger(true);
                      },
                      child: Text(
                        '没有账号?立即注册',
                        style: TextStyle(
                            color: Colors.blue,
                            decoration: TextDecoration.underline,
                            fontSize: 12),
                      ),
                    ),
                    ElevatedButton(
                      onPressed: () {
                        // 登录逻辑
                        widget.onLoginHandler(true);
                      },
                      child: Text('登录',
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
        ),
      ],
    );
  }
}
