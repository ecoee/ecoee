import 'package:flutter/cupertino.dart';

Widget buildPoint(String point, MainAxisAlignment align) {
  return Row(
    mainAxisAlignment: align,
    children: [
      Text(point, style: TextStyle(fontSize: 48, fontFamily: 'GlutenMedium')),
      SizedBox(width: 4),
      Text("p", style: TextStyle(fontSize: 24, fontFamily: 'GlutenRegular')),
    ],
  );
}
