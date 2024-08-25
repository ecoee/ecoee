import 'package:flutter/cupertino.dart';

Widget buildEcoeButton(double width, double height, String text,
    Color backgroundColor, Color textColor) {
  return Container(
    width: width,
    height: height,
    clipBehavior: Clip.antiAlias,
    decoration: ShapeDecoration(
      color: backgroundColor,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(24),
      ),
    ),
    child: Row(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text(
          text,
          textAlign: TextAlign.right,
          style: TextStyle(
            color: textColor,
            fontSize: 18,
            fontFamily: 'GlutenSemiBold',
            fontWeight: FontWeight.w600,
            height: 0,
          ),
        ),
      ],
    ),
  );
}
