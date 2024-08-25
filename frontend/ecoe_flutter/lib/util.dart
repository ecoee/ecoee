import 'package:intl/intl.dart';

class Util {
  static String formatNumberWithComma(int number) {
    final NumberFormat formatter = NumberFormat('#,###');
    return formatter.format(number);
  }
}
