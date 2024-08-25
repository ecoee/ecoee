import 'package:flutter/cupertino.dart';

Widget buildTitle(String title) {
  return Row(
    children: [
      SizedBox(width: 30),
      Image.asset('assets/images/bin.png'),
      SizedBox(width: 10),
      Text(title, style: TextStyle(fontSize: 24, fontFamily: 'GlutenBold')),
    ],
  );
}
