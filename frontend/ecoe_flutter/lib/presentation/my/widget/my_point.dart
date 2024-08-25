import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:flutter/cupertino.dart';

import '../../button.dart';
import '../../point.dart';
import '../../title.dart';

Widget buildMyPoint() {
  return Container(
    color: EcoeColors.upperGrey,
    child: Column(
      children: [_buildTop(), SizedBox(height: 26), _buildPoint()],
    ),
  );
}

Widget _buildPoint() {
  return Column(
    children: [
      buildTitle("my point bin"),
      buildPoint("10,200", MainAxisAlignment.end),
    ],
  );
}

Widget _buildTop() {
  return Container(
    height: 98,
    child: Stack(
      children: [
        Positioned(
          left: 36,
          top: 8,
          child: Container(width: 68, height: 68),
        ),
        const Positioned(
          left: 294,
          top: 47,
          child: Text(
            'Jinnie',
            textAlign: TextAlign.right,
            style: TextStyle(
              color: Color(0xFF151A1E),
              fontSize: 20,
              fontFamily: 'Pretendard',
              fontWeight: FontWeight.w600,
              height: 0,
            ),
          ),
        ),
        const Positioned(
          left: 139,
          top: 18,
          child: Text(
            'Coupang / business planning',
            textAlign: TextAlign.right,
            style: TextStyle(
              color: Color(0xFF151A1E),
              fontSize: 16,
              fontFamily: 'Pretendard',
              fontWeight: FontWeight.w500,
              height: 0,
            ),
          ),
        ),
        Positioned(
          left: 40,
          top: 30,
          child: Container(
            width: 69,
            height: 68,
            child: Stack(
              children: [
                Positioned(
                  left: 0,
                  top: 0,
                  child: Container(
                    width: 69,
                    height: 68,
                    decoration: const BoxDecoration(
                      image: DecorationImage(
                        image:
                            NetworkImage("https://via.placeholder.com/69x68"),
                        fit: BoxFit.cover,
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ],
    ),
  );
}
