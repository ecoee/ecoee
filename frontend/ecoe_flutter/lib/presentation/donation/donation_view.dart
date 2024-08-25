import 'dart:developer';

import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:ecoe_flutter/presentation/donation.dart';
import 'package:ecoe_flutter/presentation/my/widget/my_point.dart';
import 'package:ecoe_flutter/presentation/my/widget/point_history.dart';
import 'package:ecoe_flutter/presentation/rank.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../button.dart';
import '../font.dart';
import '../main_bottom.dart';
import '../main_bottom_circle.dart';
import '../point.dart';
import '../title.dart';

class Donation extends StatelessWidget {
  // static const route = '/main/MyPageBloc';

  const Donation({super.key});

  @override
  Widget build(BuildContext context) {
    return const DonationView();
  }
}

class DonationView extends StatefulWidget {
  const DonationView({super.key});

  @override
  State<DonationView> createState() => _DonationViewState();
}

class _DonationViewState extends State<DonationView> {
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
              SizedBox(height: 10),
              RankBoardDonation(),
              SizedBox(height: 18),
              buildVoteButton(),
              SizedBox(height: 27),
              buildDonationListGrid(),
            ],
          ),
        ),
      ),
    );
  }

  Widget buildVoteButton() {
    bool isShort = false;
    return isShort ? buildShortPoint(10) : buildFullPoint();
  }

  Widget buildShortPoint(int shortPoint) {
    return Container(
      padding: EdgeInsets.fromLTRB(20, 10, 10, 10),
      width: 362,
      height: 78,
      clipBehavior: Clip.antiAlias,
      decoration: ShapeDecoration(
        color: Color(0xFFF3F3F3),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(20),
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "You can't vote",
            style: TextStyle(fontSize: 18, fontFamily: Font.PretendardBold),
          ),
          Text(
            "Our team is $shortPoint points short.",
            style: TextStyle(fontSize: 18, fontFamily: Font.PretendardBold),
          ),
        ],
      ),
    );
  }

  Widget buildFullPoint() {
    return Container(
      width: 362,
      height: 78,
      padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 9),
      clipBehavior: Clip.antiAlias,
      decoration: ShapeDecoration(
        color: Color(0xFF151A1E),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(20),
        ),
      ),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Container(
            height: double.infinity,
            child: Row(
              mainAxisSize: MainAxisSize.min,
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                SizedBox(
                  width: 302,
                  child: Text(
                    'Please vote for your favorite\ndonation bin.',
                    style: TextStyle(
                      color: Color(0xFFF3F3F3),
                      fontSize: 18,
                      fontFamily: Font.PretendardSemiBold,
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
