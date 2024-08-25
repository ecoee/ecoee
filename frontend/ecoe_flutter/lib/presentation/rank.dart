import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:ecoe_flutter/presentation/button.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/cupertino.dart';

import '../util.dart';
import 'font.dart';

Widget RankBoardHome() {
  return Container(
    margin: const EdgeInsets.fromLTRB(14, 10, 14, 10),
    height: 487,
    decoration: ShapeDecoration(
      shape: RoundedRectangleBorder(
        side: BorderSide(width: 1, color: Color(0xFF151A1E)),
        borderRadius: BorderRadius.circular(20),
      ),
    ),
    child: Column(
      children: [
        SizedBox(height: 17),
        Row(
          children: [
            SizedBox(width: 17),
            Text(
              "Let’s do it! Coupang",
              style: TextStyle(fontSize: 18, fontFamily: Font.PretendardBold),
            ),
          ],
        ),
        SizedBox(
          height: 45,
        ),
        Row(
          children: [
            const SizedBox(
              width: 45,
            ),
            buildChart(90000),
          ],
        ),
        SizedBox(height: 40),
        _buildRankCard(),
        SizedBox(
          height: 30,
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            buildEcoeButton(250, 35, "Go to donation bin >", EcoeColors.black,
                EcoeColors.white),
            SizedBox(
              width: 20,
            )
          ],
        )
      ],
    ),
  );
}

Widget RankBoardDonation() {
  return Container(
    margin: const EdgeInsets.fromLTRB(14, 10, 14, 10),
    child: Column(
      children: [
        SizedBox(height: 17),
        Row(
          children: [
            SizedBox(width: 17),
            Text(
              "Let’s do it! Coupang",
              style: TextStyle(fontSize: 18, fontFamily: Font.PretendardBold),
            ),
          ],
        ),
        SizedBox(
          height: 45,
        ),
        Row(
          children: [
            const SizedBox(
              width: 45,
            ),
            buildChart(90000),
          ],
        ),
        SizedBox(height: 40),
        _buildRankCard(),
      ],
    ),
  );
}

Widget buildChatInfoTitle(Color color, String title, String point) {
  return Column(
    crossAxisAlignment: CrossAxisAlignment.end,
    children: [
      Container(
        height: 19,
        child: Row(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Container(
              width: 8,
              height: 8,
              decoration: ShapeDecoration(
                color: color,
                shape: OvalBorder(),
              ),
            ),
            const SizedBox(width: 6),
            Text(
              title,
              textAlign: TextAlign.right,
              style: TextStyle(
                color: Color(0xFF6D6D6D),
                fontSize: 16,
                fontFamily: 'Inter',
                fontWeight: FontWeight.w600,
                height: 0,
              ),
            ),
          ],
        ),
      ),
      Text(
        point,
        style: TextStyle(fontFamily: Font.PretendardSemiBold, fontSize: 20),
      )
    ],
  );
}

Widget buildChart(double currentPoint) {
  const int total = 100000;
  final List<double> data = [currentPoint, total - currentPoint]; // 두 개의 값
  var remain = '${(data[0] / total * 100).toStringAsFixed(1)}%';

  return Row(
    mainAxisAlignment: MainAxisAlignment.spaceBetween,
    children: [
      Container(
        child: Stack(
          alignment: Alignment.center,
          children: [
            Container(
              width: 80, // 차트의 너비
              height: 80, // 차트의 높이
              child: PieChart(
                PieChartData(
                  sections: [
                    PieChartSectionData(
                      color: EcoeColors.green,
                      value: data[0],
                      // 첫 번째 섹션의 퍼센트
                      title: '',
                      radius: 30,
                      titleStyle: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.bold,
                        color: EcoeColors.white,
                      ),
                    ),
                    PieChartSectionData(
                      color: EcoeColors.red,
                      value: data[1],
                      // 두 번째 섹션의 퍼센트
                      radius: 30,
                      title: '',
                      titleStyle: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.bold,
                        color: EcoeColors.white,
                      ),
                    ),
                  ],
                  centerSpaceRadius: 40, // 중앙의 빈 공간 크기
                  sectionsSpace: 0, // 섹션 간 간격을 없앰
                ),
              ),
            ),
            Positioned(
              child: Text(
                remain, // 중앙에 표시할 텍스트 (총합)
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                  color: EcoeColors.black,
                ),
              ),
            ),
          ],
        ),
      ),
      SizedBox(
        width: 40,
      ),
      Column(
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          buildChatInfoTitle(EcoeColors.red, 'Our remaining point',
              Util.formatNumberWithComma((total - currentPoint).toInt())),
          buildChatInfoTitle(EcoeColors.green, 'Our current point',
              Util.formatNumberWithComma((currentPoint).toInt()))
        ],
      ),
    ],
  );
}

Widget _buildRankCard() {
  return Container(
    child: Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        _buildCard(),
        _buildBlackCard(),
        _buildCard(),
      ],
    ),
  );
}

Widget _buildCard() {
  return Container(
    width: 110,
    height: 150,
    child: Stack(
      children: [
        Positioned(
          left: 0,
          top: 0,
          child: Container(
            width: 110,
            height: 150,
            decoration: ShapeDecoration(
              color: EcoeColors.white,
              shape: RoundedRectangleBorder(
                side: BorderSide(width: 1, color: Color(0xFFB3B3B3)),
                borderRadius: BorderRadius.circular(10),
              ),
            ),
          ),
        ),
        Positioned(
          left: 22,
          top: 109,
          child: Text(
            'Taeyoung',
            textAlign: TextAlign.right,
            style: TextStyle(
              color: Color(0xFF6D6D6D),
              fontSize: 14,
              fontFamily: 'Inter',
              fontWeight: FontWeight.w600,
              height: 0,
            ),
          ),
        ),
        Positioned(
          left: 42,
          top: 21,
          child: Text(
            '2nd',
            textAlign: TextAlign.right,
            style: TextStyle(
              color: Color(0xFF151A1E),
              fontSize: 14,
              fontFamily: 'Gluten',
              fontWeight: FontWeight.w600,
              height: 0,
            ),
          ),
        ),
        Positioned(
          left: 31,
          top: 63,
          child: Text(
            '130p',
            textAlign: TextAlign.right,
            style: TextStyle(
              color: Color(0xFF151A1E),
              fontSize: 20,
              fontFamily: 'Inter',
              fontWeight: FontWeight.w600,
              height: 0,
            ),
          ),
        ),
      ],
    ),
  );
}

Widget _buildBlackCard() {
  return Container(
    width: 130.53,
    height: 178,
    child: Stack(
      children: [
        Positioned(
          left: 0,
          top: 0,
          child: Container(
            width: 130.53,
            height: 178,
            decoration: ShapeDecoration(
              color: Color(0xFF151A1E),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(10),
              ),
            ),
          ),
        ),
        Positioned(
          left: 27,
          top: 129,
          child: SizedBox(
            width: 76,
            height: 21,
            child: Text(
              'Gyeongjun',
              textAlign: TextAlign.right,
              style: TextStyle(
                color: EcoeColors.white,
                fontSize: 14,
                fontFamily: 'Inter',
                fontWeight: FontWeight.w600,
                height: 0,
              ),
            ),
          ),
        ),
        Positioned(
          left: 39,
          top: 75,
          child: SizedBox(
            width: 53,
            height: 28,
            child: Text(
              '380p',
              textAlign: TextAlign.right,
              style: TextStyle(
                color: EcoeColors.white,
                fontSize: 20,
                fontFamily: 'Inter',
                fontWeight: FontWeight.w600,
                height: 0,
              ),
            ),
          ),
        ),
        Positioned(
          left: 38,
          top: 11.87,
          child: Container(
            width: 56.96,
            height: 56.96,
            child: Image.asset("assets/images/crown.png"),
          ),
        ),
      ],
    ),
  );
}
