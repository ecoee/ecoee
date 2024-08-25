// import 'dart:developer';
//
// import 'package:design/design.dart';
// import 'package:flutter/cupertino.dart';
// import 'package:flutter/material.dart';
// import 'package:flutter_bloc/flutter_bloc.dart';
// import 'package:stove_sdk/log/action_type.dart';
//
// import '../../constants/messages.dart';
// import '../../data/event/room/room_bloc.dart';
// import '../../utils/log.dart';
// import '../create/model/create_room_schedule_config.dart';
// import '../../constants/assets.gen.dart';
// import '../../constants/config.dart';
// import '../../data/model/room.dart';
// import '../../data/repository/repository.dart';
// import '../info/model/schedule_info_result.dart';
// import '../room/widgets/host_appoint_view.dart';
// import '../routes.dart';
// import 'bloc/main_bloc.dart';
// import 'model/navigation_item.dart';
// import 'profile/bloc/profile_bloc.dart';
//
// final List<NavigationItemBase> mainTabs = [
//   NavigationItem(mainTab: const MainHomeTab(), icon: Assets.images.navigationHome, iconSelected: Assets.images.navigationHomeSelected, title: Strings.mainHomeIcLabel.tr),
//   NavigationItem(mainTab: const MainFriendTab(), icon: Assets.images.navigationFriend, iconSelected: Assets.images.navigationFriendSelected, title: Strings.mainFriendIcLabel.tr),
//   NavigationItemBase(title: Strings.mainCreateRoomIcLabel.tr, iconSize: 40),
//   NavigationItem(
//       mainTab: const MainNotificationTab(), icon: Assets.images.navigationNotification, iconSelected: Assets.images.navigationNotificationSelected, title: Strings.mainNotificationIcLabel.tr),
//   NavigationItem(mainTab: const MainProfileTab(), icon: Assets.images.navigationProfile, iconSelected: Assets.images.navigationProfileSelected, title: Strings.mainProfileIcLabel.tr),
// ];
//
// class BottomNavigationView extends StatelessWidget {
//   final Color _backgroundColor = PpoolColors.bgDepth1;
//
//   const BottomNavigationView({super.key});
//
//   @override
//   Widget build(BuildContext context) {
//     return Container(
//       decoration: _decoration(),
//       child: Column(
//         mainAxisSize: MainAxisSize.min,
//         children: [
//           CustomPaint(
//             painter: NavigationPainter(_backgroundColor),
//             child: Container(
//               height: Config.navigationHeight,
//               padding: const EdgeInsets.symmetric(horizontal: 16),
//               child: Row(
//                 mainAxisAlignment: MainAxisAlignment.spaceAround,
//                 children: _getItemView(context),
//               ),
//             ),
//           ),
//           Container(
//             height: MediaQuery.of(context).padding.bottom,
//             color: _backgroundColor,
//           ),
//         ],
//       ),
//     );
//   }
//
//   BoxDecoration _decoration() {
//     return const BoxDecoration(
//       boxShadow: [
//         BoxShadow(
//           color: Color(0x0C000000),
//           blurRadius: 10,
//           offset: Offset(0, -4),
//           spreadRadius: 0,
//         ),
//       ],
//     );
//   }
//
//   List<Widget> _getItemView(BuildContext context) {
//     return mainTabs.map((item) {
//       if (item is NavigationItem) {
//         return _buildItemView(context, item);
//       } else {
//         return _buildBigItemView(context, item);
//       }
//     }).toList();
//   }
//
//   Widget _buildBigItemView(BuildContext context, NavigationItemBase item) {
//     return AnimatedButton(
//       onTap: () async {
//         Log.sendData(ActionType.clickMenuBarMakeRoom);
//         if (context.read<Repository>().auth.isNotLogin()) {
//           Routes.toLogin();
//           return;
//         }
//
//         final roomBloc = context.read<RoomBloc>();
//
//         if (roomBloc.state.room != null) {
//           final result = await Routes.showAlert(
//             description: Strings.vcPipConfirmMsgNewVc.tr,
//             positive: Strings.vcPipConfirmNewVc2btn.tr,
//             negative: Strings.vcPipConfirmNewVc1btn.tr,
//           );
//
//           if (!(result == DialogViewResult.positive) || !context.mounted) {
//             return;
//           } else {
//             if (roomBloc.state.hostCount == 1 && roomBloc.state.isHost) {
//               // 호스트 지정 필요
//               final speakerList = roomBloc.state.speakerList ?? [];
//               if (speakerList.isEmpty || (speakerList.length == 1 && speakerList.firstOrNull?.length == 1 && speakerList.firstOrNull?.firstOrNull?.id == roomBloc.state.memberId)) {
//                 // 바로 퇴장
//               } else {
//                 final result = await Routes.showBottomSheet(const HostAppointView());
//                 if (result != true) {
//                   return;
//                 }
//               }
//             }
//           }
//
//           if (!context.mounted) return;
//
//           context.read<RoomBloc>().add(LeaveRoom());
//         }
//
//         //TODO : show create room popup
//         final schedule = await context.read<MainBloc>().hasReservationSchedule();
//         if (!context.mounted) return;
//
//         final id = schedule?.id;
//
//         if (id != null) {
//           final result = await Routes.toScheduleGettingStarted(id);
//
//           if (result == null) {
//             return;
//           }
//
//           if (result.action == ScheduleInfoResultAction.start) {
//             final sequence = result.liveRoomSequence;
//             if (sequence != null && context.mounted) {
//               Routes.toRoom(Room(sequence, isStartedByMe: false));
//             }
//             return;
//           }
//         }
//
//         if (!context.mounted) return;
//
//         final menu = await Routes.toCreateMenu();
//         if (menu == null || !context.mounted) return;
//
//         if (menu == CreateRoomScheduleType.room) {
//           final room = await Routes.toCreateRoom();
//           if (room != null && context.mounted) {
//             Routes.toRoom(Room(room.id, isStartedByMe: true));
//           }
//         } else if (menu == CreateRoomScheduleType.schedule) {
//           final schedule = await Routes.toCreateSchedule();
//           if (schedule != null && context.mounted) {
//             final mainBloc = context.read<MainBloc>();
//             mainBloc.add(MainTabRefreshed(refreshTab: RefreshTab(mainTab: const MainHomeTab(), refreshTime: DateTime.now())));
//             mainBloc.add(MainTabClicked(const MainProfileTab(profileTab: ProfileTab.reserved)));
//             mainBloc.add(MainTabRefreshed(refreshTab: RefreshTab(mainTab: const MainProfileTab(), refreshTime: DateTime.now())));
//
//             final result = await Routes.toScheduleInfo(schedule.id);
//             if (result != null) {
//               if (!context.mounted) return;
//               _handleScheduleInfoResult(context, result);
//             }
//           }
//         }
//       },
//       child: NavigationBigItemView(item: item),
//     );
//   }
//
//   void _handleScheduleInfoResult(BuildContext context, ScheduleInfoResult result) async {
//     log('_handleScheduleInfoResult action2 => ${result.action}');
//     final mainBloc = context.read<MainBloc>();
//     switch (result.action) {
//       case ScheduleInfoResultAction.delete:
//       case ScheduleInfoResultAction.close:
//         {
//           mainBloc.add(MainTabRefreshed(refreshTab: RefreshTab(mainTab: const MainHomeTab(), refreshTime: DateTime.now())));
//           mainBloc.add(MainTabRefreshed(refreshTab: RefreshTab(mainTab: const MainProfileTab(), refreshTime: DateTime.now())));
//         }
//       case ScheduleInfoResultAction.start:
//         {
//           final sequence = result.liveRoomSequence;
//           if (sequence != null) {
//             Routes.toStartedRoom(
//               sequence,
//               result.isStartedByMe,
//               isWaiting: result.isWaiting,
//             );
//           }
//         }
//       case _:
//         {}
//     }
//   }
//
//   Widget _buildItemView(BuildContext context, NavigationItem item) {
//     return BlocBuilder<MainBloc, MainState>(
//       builder: (context, state) {
//         final mainTab = state.mainTab;
//         final selected = item.mainTab == mainTab;
//         return AnimatedButton(
//           onTap: () {
//             if (item.mainTab.isCheckLogin() && context.read<Repository>().auth.isNotLogin()) {
//               Routes.toLogin();
//               return;
//             }
//
//             context.read<MainBloc>().add(MainTabClicked(item.mainTab));
//             Log.sendData(getMainTabActionType(item.mainTab));
//           },
//           child: NavigationItemView(
//             item: item,
//             selected: selected,
//           ),
//         );
//       },
//     );
//   }
// }
//
// class NavigationItemView extends StatelessWidget {
//   final NavigationItem item;
//   final bool selected;
//   final bool hasNew;
//
//   const NavigationItemView({super.key, required this.item, required this.selected, this.hasNew = false});
//
//   @override
//   Widget build(BuildContext context) {
//     return AnimatedContainer(
//       width: 56,
//       color: Colors.transparent,
//       duration: Config.animationDefault,
//       child: Column(
//         mainAxisSize: MainAxisSize.min,
//         children: [
//           const Spacer(flex: 20),
//           _buildIcon(),
//           const SizedBox(height: 4),
//           Text(
//             item.title,
//             style: Style.medium(fontSize: 9, color: selected ? PpoolColors.textPpool : PpoolColors.text03),
//           ),
//           const Spacer(flex: 4),
//         ],
//       ),
//     );
//   }
//
//   Widget _buildIcon() {
//     final icon = selected ? item.iconSelected.svg(width: item.iconSize, height: item.iconSize) : item.icon.svg(width: item.iconSize, height: item.iconSize);
//     if (item.mainTab.isProfile()) {
//       return BlocSelector<MainBloc, MainState, String?>(
//         selector: (state) => state.member?.profileImageUrl,
//         builder: (context, profileImageUrl) {
//           if (profileImageUrl != null) {
//             return ProfileImageView(size: Size.square(item.iconSize), url: profileImageUrl);
//           }
//           return icon;
//         },
//       );
//     }
//     if (item.mainTab.isNotification()) {
//       return BlocSelector<MainBloc, MainState, bool>(
//         selector: (state) => state.hasNewNotification,
//         builder: (context, hasNewNotification) {
//           return Stack(
//             alignment: Alignment.topCenter,
//             children: [
//               icon,
//               if (hasNewNotification)
//                 Align(
//                     alignment: Alignment.topRight,
//                     child: Padding(
//                       padding: const EdgeInsets.only(right: 12),
//                       child: Assets.images.dotFill.svg(width: 6, height: 6),
//                     )),
//             ],
//           );
//         },
//       );
//     }
//     return icon;
//   }
// }
//
// class NavigationBigItemView extends StatelessWidget {
//   final NavigationItemBase item;
//
//   const NavigationBigItemView({super.key, required this.item});
//
//   @override
//   Widget build(BuildContext context) {
//     return Container(
//       color: Colors.transparent,
//       child: Column(
//         mainAxisSize: MainAxisSize.min,
//         children: [
//           const Spacer(flex: 8),
//           Container(
//             width: 40,
//             height: 40,
//             decoration: const ShapeDecoration(
//               gradient: LinearGradient(
//                 begin: Alignment(0, -1),
//                 end: Alignment(0, 1),
//                 colors: PpoolColors.navigationCreate,
//               ),
//               shape: CircleBorder(),
//             ),
//             alignment: Alignment.center,
//             child: Assets.images.navigationCreateRoomPlus.image(width: 28, height: 28),
//           ),
//           const SizedBox(height: 4),
//           Text(
//             item.title,
//             style: Style.medium(fontSize: 9, color: PpoolColors.text03),
//             maxLines: 1,
//             overflow: TextOverflow.ellipsis,
//           ),
//           const Spacer(flex: 4),
//         ],
//       ),
//     );
//   }
// }
//
// class NavigationPainter extends CustomPainter {
//   final _topWidthSize = 213.63 - 161.37;
//
//   final Color color;
//
//   NavigationPainter(this.color);
//
//   double _width(Size size, double width) => size.width * (width / 375);
//
//   double _height(Size size, double height) => size.height * (height / 70);
//
//   @override
//   void paint(Canvas canvas, Size size) {
//     Path path = Path();
//     path.moveTo(0, size.height);
//     path.lineTo(0, _height(size, 50));
//     path.cubicTo(0, _height(size, 31.09), 0, _height(size, 23.79), _width(size, 5.04), _height(size, 17.29));
//     path.cubicTo(_width(size, 6.27), _height(size, 15.7), _width(size, 7.7), _height(size, 14.27), _width(size, 9.29), _height(size, 13.04));
//     path.cubicTo(_width(size, 15.78), 8, _width(size, 25.24), 8, _width(size, 44), 8);
//     path.cubicTo(_width(size, 15.78), 8, _width(size, 25.24), 8, _width(size, 44), 8);
//     final topLeftStartX = (size.width - _topWidthSize) / 2;
//     path.lineTo(topLeftStartX, 8);
//     path.cubicTo(topLeftStartX + (165.56 - 161.37), 8, topLeftStartX + (169.47 - 161.37), 6.17, topLeftStartX + (173.06 - 161.37), 4);
//     path.cubicTo(topLeftStartX + (177.42 - 161.37), 1.37, topLeftStartX + (182.41 - 161.37), -0.01, topLeftStartX + (187.5 - 161.37), 0);
//     path.cubicTo(topLeftStartX + (192.78 - 161.37), 0, topLeftStartX + (197.73 - 161.37), 1.46, topLeftStartX + (201.94 - 161.37), 4);
//     path.cubicTo(topLeftStartX + (205.53 - 161.37), 6.17, topLeftStartX + (209.44 - 161.37), 8, topLeftStartX + (213.63 - 161.37), 8);
//     path.lineTo(topLeftStartX + _topWidthSize, 8);
//     path.cubicTo(_width(size, 349.76), 8, _width(size, 359.21), 8, _width(size, 365.71), _height(size, 13.04));
//     path.cubicTo(_width(size, 367.3), _height(size, 14.27), _width(size, 368.73), _height(size, 15.7), _width(size, 369.96), _height(size, 17.29));
//     path.cubicTo(size.width, _height(size, 23.78), size.width, _height(size, 31.08), size.width, _height(size, 50));
//     path.lineTo(size.width, size.height);
//     canvas.drawPath(path, Paint()..color = color);
//   }
//
//   @override
//   bool shouldRepaint(CustomPainter oldDelegate) {
//     NavigationPainter old = oldDelegate as NavigationPainter;
//     return color != old.color;
//   }
// }
