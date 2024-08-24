// import 'dart:core';
// import 'dart:io';
// import 'dart:ui';
//
// import 'package:flutter/animation.dart';
// import 'package:stove_sdk/base/constants.dart';
//
// import '../data/model/national_info.dart';
//
// enum Environment {
//   live('live'),
//   sandbox('sandbox'),
//   sandbox2('sandbox2'),
//   qa('qa');
//
//   const Environment(this.value);
//
//   final String value;
// }
//
// class Config {
//   static Environment environment = Environment.live;
//
//   static String get domain {
//     switch (environment) {
//       case Environment.live:
//         return 'https://tv-api.ppool.us';
//       case Environment.sandbox:
//         return 'https://tv-api-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://tv-api-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://tv-api-qa.ppool.us';
//     }
//   }
//
//   static String get stoveMembershipDomain {
//     switch (environment) {
//       case Environment.live:
//         return 'https://s-api.onstove.com';
//       case Environment.sandbox:
//         return 'https://s-api.gate8.com';
//       case Environment.sandbox2:
//         return 'https://s-api.gate8.com';
//       case Environment.qa:
//         return 'https://s-api-qa.onstove.com';
//     }
//   }
//
//   static String get stoveDomain {
//     switch (environment) {
//       case Environment.live:
//         return 'https://papi.ppool.us';
//       case Environment.sandbox:
//         return 'https://api.gate8.com';
//       case Environment.sandbox2:
//         return 'https://papi-sb.ppool.us';
//       case Environment.qa:
//         return 'https://papi-qa.ppool.us';
//     }
//   }
//
//   static String get webviewDomain {
//     switch (environment) {
//       case Environment.live:
//         return 'https://m.ppool.us';
//       case Environment.sandbox:
//         return 'https://m-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://m-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://m-qa.ppool.us';
//     }
//   }
//
//   static String get imageDomain {
//     switch (environment) {
//       case Environment.live:
//         return 'https://tv-contents.ppool.us';
//       case Environment.sandbox:
//         return 'https://tv-contents-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://tv-contents-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://tv-contents-qa.ppool.us';
//     }
//   }
//
//   static String get originUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://tv-origin.ppool.us';
//       case Environment.sandbox:
//         return 'https://tv-origin-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://tv-origin-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://tv-origin-qa.ppool.us';
//     }
//   }
//
//   static String get edgeUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://tv-edge.ppool.us';
//       case Environment.sandbox:
//         return 'https://tv-edge-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://tv-edge-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://tv-edge-qa.ppool.us';
//     }
//   }
//
//   static String get messageUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://tv-msg.ppool.us';
//       case Environment.sandbox:
//         return 'https://tv-msg-sb.ppool.us';
//       case Environment.sandbox2:
//         return 'https://tv-msg-sb2.ppool.us';
//       case Environment.qa:
//         return 'https://tv-msg-qa.ppool.us';
//     }
//   }
//
//   static String get guideUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://support.ppool.us/guide';
//       case Environment.sandbox:
//       case Environment.sandbox2:
//         return 'https://support.ppool.us/guide-sb';
//       case Environment.qa:
//         return 'https://support.ppool.us/guide-qa';
//     }
//   }
//
//   static String get noticeUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://support.ppool.us/notice';
//       case Environment.sandbox:
//       case Environment.sandbox2:
//         return 'https://support.ppool.us/notice-sb';
//       case Environment.qa:
//         return 'https://support.ppool.us/notice-qa';
//     }
//   }
//
//   static String get customerSupportUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://help.ppool.us/mobile/faq/serviceMain/page';
//       case Environment.sandbox:
//       case Environment.sandbox2:
//         return 'https://help-sb.ppool.us/mobile/faq/serviceMain/page';
//       case Environment.qa:
//         return 'https://help-qa.ppool.us/mobile/faq/serviceMain/page';
//     }
//   }
//
//   static String get termsOfUseUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://m.ppool.us/$urlLanguageCode/terms?category=service';
//       case Environment.sandbox:
//         return 'https://m-sb.ppool.us/$urlLanguageCode/terms?category=service';
//       case Environment.sandbox2:
//         return 'https://m-sb2.ppool.us/$urlLanguageCode/terms?category=service';
//       case Environment.qa:
//         return 'https://m-qa.ppool.us/$urlLanguageCode/terms?category=service';
//     }
//   }
//
//   static String get operationalPolicyUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://m.ppool.us/$urlLanguageCode/terms?category=oppolicy';
//       case Environment.sandbox:
//         return 'https://m-sb.ppool.us/$urlLanguageCode/terms?category=oppolicy';
//       case Environment.sandbox2:
//         return 'https://m-sb2.ppool.us/$urlLanguageCode/terms?category=oppolicy';
//       case Environment.qa:
//         return 'https://m-qa.ppool.us/$urlLanguageCode/terms?category=oppolicy';
//     }
//   }
//
//   static String get youthProtectionPolicyUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://m.ppool.us/$urlLanguageCode/terms?category=safeguard';
//       case Environment.sandbox:
//         return 'https://m-sb.ppool.us/$urlLanguageCode/terms?category=safeguard';
//       case Environment.sandbox2:
//         return 'https://m-sb2.ppool.us/$urlLanguageCode/terms?category=safeguard';
//       case Environment.qa:
//         return 'https://m-qa.ppool.us/$urlLanguageCode/terms?category=safeguard';
//     }
//   }
//
//   static String get privacyPolicyUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://m.ppool.us/$urlLanguageCode/terms?category=privacyauth';
//       case Environment.sandbox:
//         return 'https://m-sb.ppool.us/$urlLanguageCode/terms?category=privacyauth';
//       case Environment.sandbox2:
//         return 'https://m-sb2.ppool.us/$urlLanguageCode/terms?category=privacyauth';
//       case Environment.qa:
//         return 'https://m-qa.ppool.us/$urlLanguageCode/terms?category=privacyauth';
//     }
//   }
//
//   static String get loginPageUrl {
//     switch (environment) {
//       case Environment.live:
//         return 'https://pmember.ppool.us/auth/login';
//       case Environment.sandbox:
//       case Environment.sandbox2:
//         return 'https://pmember-sb.ppool.us/auth/login';
//       case Environment.qa:
//         return 'https://pmember-qa.ppool.us/auth/login';
//     }
//   }
//
//   static const Duration connectTimeout = Duration(seconds: 10);
//   static const Duration receiveTimeout = Duration(seconds: 10);
//   static const Duration sendTimeout = Duration(seconds: 10);
//
//   static final List<String> defaultDefaultImageUrls = [
//     '$imageDomain/stovetv/profile/profile_img_1.png',
//     '$imageDomain/stovetv/profile/profile_img_2.png',
//     '$imageDomain/stovetv/profile/profile_img_3.png',
//     '$imageDomain/stovetv/profile/profile_img_4.png',
//     '$imageDomain/stovetv/profile/profile_img_5.png',
//   ];
//
//   static const List<Color> backgroundProfileColors = [
//     Color(0xFF4848FF),
//     Color(0xFF0089EC),
//     Color(0xFFFF4848),
//     Color(0xFFFF6948),
//     Color(0xFFFBB80D),
//     Color(0xFF00BD62),
//     Color(0xFF8D23FF),
//   ];
//
//   static const Duration animationDefault = Duration(milliseconds: 300);
//   static const Duration animationFast = Duration(milliseconds: 150);
//   static const Duration animationLong = Duration(milliseconds: 900);
//   static const Curve animationDefaultCurve = Curves.easeInOut;
//
//   static const Duration bottomSheetDuration = Duration(milliseconds: 300);
//   static const Curve bottomSheetCurve = ElasticOutCurve(2.5);
//
//   static const Duration toastDuration = Duration(seconds: 2);
//   static const Duration actionToastDuration = Duration(minutes: 60);
//
//   static const Duration networkTimeout = Duration(seconds: 10);
//   static const Duration networkDownloadTimeout = Duration(seconds: 30);
//   static const Duration networkUploadTimeout = Duration(seconds: 300);
//
//   static const Duration debounce = Duration(milliseconds: 300);
//
//   static const int videoBitrate = 200;
//
//   static const Locale korea = Locale('ko', 'KR');
//   static const Locale english = Locale('en', 'US');
//   static const Locale japan = Locale('ja', 'JP');
//   static const Locale china = Locale('zh', 'CN');
//   static const Locale taiwan = Locale('zh', 'TW');
//
//   static List<Locale> locales = [
//     korea,
//     english,
//     japan,
//     china,
//     taiwan,
//   ];
//
//   static Locale getPlatformLocale() {
//     final platformLocale = Platform.localeName;
//     if (platformLocale.contains('zh_Hant') || platformLocale.contains('zh_TW')) {
//       return taiwan;
//     }
//
//     final languageCode = Platform.localeName.split('_').first;
//     return locales.firstWhere((e) => e.languageCode == languageCode, orElse: () => english);
//   }
//
//   static Locale locale = english;
//
//   static const String tag = '@';
//
//   static const String nicknameRegex = r'^[\W\w]{2,20}$';
//   static const String emailRegex = r'^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$';
//   static const String tagErrorStartRegex = r'^[.]$';
//   static const String tagRegex = r'^[a-zA-Z0-9._]{1,20}$';
//   static const String passwordRegex = r'^.*(?=.*\d)(?=.*[a-zA-Z])(?=.*[!@#$%^&+*=()/_]).{1,20}$';
//   static const String youtubeRegex =
//       'http(?:s)?://(?:m.)?(?:www\\.)?youtu(?:\\.be\\/|be\\.com\\/(?:watch\\?(?:feature=youtu.be\\&)?v=|v\\/|embed\\/|shorts\\/|user\\/(?:[\\w#]+\\/)+))([^&#?\\n]+)(?:\\&[\\S]+)?';
//
//   static const double navigationHeight = 70;
//
//   static const double chatMoreThreshold = 15;
//
//   static const double refreshPercent = 0.6;
//   static const Duration refreshWaitDuration = Duration(seconds: 3);
//   static const Duration refreshedWaitDuration = Duration(seconds: 1);
//
//   static const cryptoKey = 'smilegateppool56';
//
//   static const int maxGuestCount = 25;
//
//   static const Duration emojiDuration = Duration(milliseconds: 3000);
//   static const Duration sendEmojiWaitDuration = Duration(milliseconds: 500);
//   static const Duration cameraOrientationAnimationDuration = Duration(milliseconds: 2000);
//   static const Duration cameraOrientationAnimationMinDuration = Duration(milliseconds: 500);
//
//   static const int maxEmojiCount = 10;
//   static const int maxRecentlyUsedEmojiCount = 12;
//
//   static const allowUpdateProfileIdDays = 30;
//
//   static const termsPushKey = 'STOVETV_MOBILEPUSH';
//   static const termsPpoolRequired = 'STOVETV_MOBILESERVICE';
//
//   static NationalInfo defaultNation = NationalInfo(nationalCode: 'us', nationalImageUrl: '', nationalName: '', nationalCallNumber: '+1');
//
//   static String get urlLanguageCode => switch (Constants.gds.countryCode) {
//         'KR' => 'ko',
//         'JP' => 'ja',
//         'CN' => 'zh-cn',
//         'TW' => 'zh-tw',
//         _ => 'en',
//       };
// }
