import 'package:flutter/cupertino.dart';

Widget buildBottomCircle() {
  return Container(
    width: 326,
    height: 126,
    child: Row(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        Container(
          width: 126,
          height: 126,
          child: Image.asset("assets/images/copang.png"),
        ),
        const SizedBox(width: 16),
        Container(
          width: 94,
          height: 94,
          child: Image.asset("assets/images/zkpass.png"),
        ),
        const SizedBox(width: 16),
        Container(
          width: 74,
          height: 74,
          child: Image.asset("assets/images/play.png"),
        ),
      ],
    ),
  );
}
