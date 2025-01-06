import 'package:flutter/material.dart';
import 'package:frontend/common/router/router_manager.dart';
import 'package:frontend/components/login/login.dart';
import 'package:frontend/components/login/login_phone.dart';
import 'package:frontend/components/register/register.dart';
import 'package:frontend/components/register/register_phone.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  bool isRegister = false;
  bool isLogin = false;

  void onFormTriggerHandler(bool reg) {
    setState(() {
      isRegister = reg;
    });
  }

  void onLoginHandler(bool login) {
    if (login) {
      RouterManager.router.navigateTo(context, '/home');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SizedBox(
          width: 800,
          height: 600,
          child: AnimatedSwitcher(
            duration: const Duration(milliseconds: 300),
            child: isRegister
                ? MediaQuery.of(context).size.width > 800
                    ? RegisterComponent(
                        key: const ValueKey('RegisterPage'),
                        onFormTrigger: onFormTriggerHandler,
                        onLoginHandler: onLoginHandler,
                      )
                    : RegisterPhoneComponent(
                        key: ValueKey('RegisterPhoneComponent'),
                        onFormTrigger: onFormTriggerHandler,
                        onLoginHandler: onLoginHandler,
                      )
                : MediaQuery.of(context).size.width > 800
                    ? LoginComponent(
                        key: ValueKey('LoginComponent'),
                        onFormTrigger: onFormTriggerHandler,
                        onLoginHandler: onLoginHandler,
                      )
                    : LoginPhoneComponent(
                        key: ValueKey('LoginPhoneComponent'),
                        onFormTrigger: onFormTriggerHandler,
                        onLoginHandler: onLoginHandler,
                      ),
          ),
        ),
      ),
    );
  }
}
