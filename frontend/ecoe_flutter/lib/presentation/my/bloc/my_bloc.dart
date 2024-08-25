import 'dart:async';

import 'package:flutter_bloc/flutter_bloc.dart';

import 'my_event.dart';
import 'my_state.dart';

class MyPageBloc extends Bloc<MyPageEvent, MyPageState> {
  MyPageBloc(super.initialState);

  @override
  Future<void> close() async {
    return super.close();
  }

  FutureOr<void> _onMyPageConnect(
      MyPageEvent event, Emitter<MyPageState> emit) async {}
}
