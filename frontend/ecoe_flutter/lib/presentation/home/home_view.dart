import 'dart:developer';

import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:ecoe_flutter/presentation/donation.dart';
import 'package:ecoe_flutter/presentation/my/widget/my_point.dart';
import 'package:ecoe_flutter/presentation/my/widget/point_history.dart';
import 'package:ecoe_flutter/presentation/rank.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../button.dart';
import '../font.dart';
import '../main_bottom.dart';
import '../main_bottom_circle.dart';
import '../point.dart';
import '../title.dart';

class Home extends StatelessWidget {
  // static const route = '/main/MyPageBloc';

  const Home({super.key});

  @override
  Widget build(BuildContext context) {
    return const HomeView();
  }
}

class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  @override
  void initState() {
    // context.read<MyPageBloc>().add(MyPageBlocMainComponentsRequested(isInit: true));
    _init();
    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  void _init() {
    log('MyPageBloc init');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: EcoeColors.white,
      body: SafeArea(
        child: SingleChildScrollView(
          child: Column(
            children: [
              _buildHomeTop(),
              SizedBox(height: 10),
              buildTitle("Co cycle"),
              SizedBox(height: 10),
              RankBoardHome(),
              SizedBox(height: 28),
              buildTitle("New cycle"),
              _buildMiddleTitle(),
              SizedBox(height: 10),
              buildDonationList(),
              SizedBox(height: 20),
              buildTitle("Rank cycle"),
              buildBottomCircle(),
              SizedBox(height: 10),
              buildMainBottom(),
              SizedBox(height: 20),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildMiddleTitle() {
    return Row(
      children: [
        SizedBox(
          width: 30,
        ),
        Text(
          "Upload New donation",
          style: TextStyle(
              fontSize: 20,
              fontFamily: Font.PretendardSemiBold,
              color: EcoeColors.textGrey),
        ),
      ],
    );
  }

  Widget _buildHomeTop() {
    return Container(
        decoration: BoxDecoration(
          image: DecorationImage(
            image: AssetImage('assets/images/main_bg.png'),
            fit: BoxFit.cover,
          ),
          borderRadius: BorderRadius.circular(12), // 모서리를 둥글게 처리 (옵션)
        ),
        child: Center(
            child: Column(
          children: [
            SizedBox(height: 22),
            const Row(
              children: [
                SizedBox(width: 32),
                Text(
                  "Cheer up~! Jinnie",
                  style: TextStyle(
                    fontSize: 18,
                    fontFamily: 'PretendardBold',
                  ),
                ),
                SizedBox(height: 25),
              ],
            ),
            SizedBox(height: 25),
            Image.asset("assets/images/level1.png"),
            SizedBox(height: 20),
            buildPoint("10,200", MainAxisAlignment.center),
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                buildEcoeButton(
                    150, 40, "Get point >", EcoeColors.black, EcoeColors.white),
                const SizedBox(width: 17)
              ],
            ),
            const SizedBox(height: 18)
          ],
        )));
  }
}
