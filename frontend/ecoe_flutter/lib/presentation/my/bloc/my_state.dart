import 'package:equatable/equatable.dart';

class MyPageState extends Equatable {
  /// Used for prevent chat room leave.
  final int? point;

  const MyPageState({
    this.point,
  });

  @override
  List<Object?> get props => [];
}
