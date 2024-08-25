import 'package:ecoe_flutter/presentation/donation/donation_view.dart';
import 'package:ecoe_flutter/presentation/home/home_view.dart';
import 'package:ecoe_flutter/presentation/land.dart';
import 'package:ecoe_flutter/presentation/my/my_view.dart';
import 'package:ecoe_flutter/presentation/point/earn_pont_view.dart';
import 'package:flutter/material.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';

import 'model/ecoe_colors.dart';

void main() {
  FlutterNativeSplash.remove();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _selectedIndex = 0;

  // 각 화면에 해당하는 위젯 리스트
  final List<Widget> _pages = <Widget>[
    Home(),
    PhotoUploadPage(),
    Donation(),
    MyPage(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: EcoeColors.white,
      body: _pages[_selectedIndex], // 현재 선택된 화면을 보여줍니다.
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.restore_from_trash_sharp),
            label: 'Point',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.recycling),
            label: 'Donation',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'My bin',
          ),
        ],
        currentIndex: _selectedIndex,
        selectedItemColor: Colors.black,
        unselectedItemColor: EcoeColors.textGrey,
        onTap: _onItemTapped,
        // 탭했을 때 실행될 함수
        selectedLabelStyle: TextStyle(fontFamily: 'PretendardSemiBold'),
        unselectedLabelStyle: TextStyle(fontFamily: 'PretendardSemiBold'),
      ),
    );
  }
}
