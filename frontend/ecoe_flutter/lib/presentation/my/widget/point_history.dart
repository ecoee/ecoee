import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:flutter/cupertino.dart';

Widget buildHistoryItem(String date, String title, String point) {
  return Container(
    color: EcoeColors.upperGrey,
    padding: EdgeInsets.fromLTRB(30, 14, 33, 14),
    child: Column(children: [
      Row(
        children: [
          Text(date,
              style: TextStyle(
                  color: EcoeColors.textGrey,
                  fontSize: 14,
                  fontFamily: 'PretendardBold')),
        ],
      ),
      SizedBox(height: 12),
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(title,
              style: TextStyle(fontSize: 18, fontFamily: 'PretendardBold')),
          Text("$point p",
              style: TextStyle(
                  color: EcoeColors.textGrey,
                  fontSize: 18,
                  fontFamily: 'PretendardSemiBold')),
        ],
      )
    ]),
  );
}
