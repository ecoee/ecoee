import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: PieChartExample(),
    );
  }
}

class PieChartExample extends StatelessWidget {
  final List<double> data = [70, 30]; // 두 개의 값

  @override
  Widget build(BuildContext context) {
    double total = data.reduce((a, b) => a + b); // 전체 합 계산

    return Scaffold(
      appBar: AppBar(
        title: Text('Pie Chart with Two Values'),
      ),
      body: Center(
        child: Stack(
          alignment: Alignment.center,
          children: [
            Container(
              width: 200, // 차트의 너비
              height: 200, // 차트의 높이
              child: PieChart(
                PieChartData(
                  sections: [
                    PieChartSectionData(
                      color: Colors.blue,
                      value: data[0],
                      title: '${(data[0] / total * 100).toStringAsFixed(1)}%', // 첫 번째 섹션의 퍼센트
                      radius: 50,
                      titleStyle: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.bold,
                        color: Colors.white,
                      ),
                    ),
                    PieChartSectionData(
                      color: Colors.red,
                      value: data[1],
                      title: '${(data[1] / total * 100).toStringAsFixed(1)}%', // 두 번째 섹션의 퍼센트
                      radius: 50,
                      titleStyle: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.bold,
                        color: Colors.white,
                      ),
                    ),
                  ],
                  centerSpaceRadius: 60, // 중앙의 빈 공간 크기
                  sectionsSpace: 0, // 섹션 간 간격을 없앰
                  // 애니메이션 제거 (기본값으로 애니메이션 없음)
                ),
              ),
            ),
            Positioned(
              child: Text(
                '${total.toStringAsFixed(0)}', // 중앙에 표시할 텍스트 (총합)
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
