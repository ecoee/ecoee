import 'dart:developer';

import 'package:ecoe_flutter/model/ecoe_colors.dart';
import 'package:ecoe_flutter/presentation/my/widget/my_point.dart';
import 'package:ecoe_flutter/presentation/my/widget/point_history.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../button.dart';
import '../donation.dart';
import '../title.dart';
import 'bloc/my_bloc.dart';
import 'bloc/my_state.dart';

class MyPage extends StatelessWidget {
  // static const route = '/main/MyPageBloc';

  const MyPage({super.key});

  @override
  Widget build(BuildContext context) {
    return const MyPageView();
  }
}

class MyPageView extends StatefulWidget {
  const MyPageView({super.key});

  @override
  State<MyPageView> createState() => _MyPageViewState();
}

class _MyPageViewState extends State<MyPageView> {
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
      body: SafeArea(
        child: SingleChildScrollView(
          child: Column(
            children: [
              buildMyPoint(),
              _buildHistoryList(),
              buildTitle("My donation history"),
              SizedBox(height: 11),
              buildDonationList(),
              SizedBox(height: 18),
              buildTitle("Point store"),
              SizedBox(height: 11),
              buildDonationList(),
              SizedBox(height: 18),
              // Expanded(child: _buildBody(context)),
              // _buildExit(context),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildHistoryList() {
    return Container(
      color: EcoeColors.upperGrey,
      child: Column(
        children: [
          buildHistoryItem("2024.21.22", "asdasd", "0202"),
          buildHistoryItem("2024.21.22", "asdasd", "0202"),
          buildHistoryItem("2024.21.22", "asdasd", "0202"),
          SizedBox(height: 10),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              buildEcoeButton(
                  105, 40, "more >", EcoeColors.black, Colors.white),
              const SizedBox(
                width: 24,
              )
            ],
          ),
          SizedBox(height: 18)
        ],
      ),
    );
  }
}
