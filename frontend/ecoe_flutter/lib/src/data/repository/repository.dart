import 'dart:async';

class Repository {
  final String _domain;

  Repository(this._domain);

  // final Client client = Client();

  ///Used from server
  // late final BannerApi banner = BannerApi(client, _domain);
  // late final ArApi ar = ArApi(client, _domain);
  // late final CommonApi common = CommonApi(client, _domain);
  // late final RoomApi room = RoomApi(client, _domain);
  // late final MemberApi member = MemberApi(client, _domain, cachedMember, auth);
  // late final NoticeApi notice = NoticeApi(client, _domain);
  // late final AlarmApi alarm = AlarmApi(client, _domain);
  // late final ScheduleApi schedule = ScheduleApi(client, _domain);
  // late final KeywordApi keyword = KeywordApi(client, _domain);
  // late final ChatApi chat = ChatApi(client, _domain);
  // late final StoveApi stove = StoveApi(client);

  ///Used for local
  // final CachedMember cachedMember = CachedMember();
  // final AuthRepository auth = AuthRepository();
  // final VersionRepository version = VersionRepository();
  // late final AppSettingRepository appSetting = AppSettingRepository(auth);
  // final ContactsRepository contacts = ContactsRepository();
  // final ConstantsRepository constants = ConstantsRepository();
  // final CachedDeepLink cachedDeepLink = CachedDeepLink();
  // final WaitingRoomAccessStatus waitingRoomAccessStatus = WaitingRoomAccessStatus();
  // final CachedNotificationCount cachedNotificationCount = CachedNotificationCount();
  // final Operation operation = Operation();
  // final PageRefreshController pageRefreshController = PageRefreshController();
  // Timer? _timer;
  //
  // void setAccessToken(AccessToken accessToken, {String? providerCode}) {
  //   client.type = 'Bearer';
  //   client.token = accessToken.token;
  //   auth.setAccessToken(accessToken);
  //   if (providerCode != null) {
  //     appSetting.setLastLoginProviderCode(providerCode);
  //   }
  //   _startAccessTokenTimer();
  // }
  //
  // void refreshAccessToken() async {
  //   final accessToken = await Auth.accessToken();
  //   if (accessToken == null) return;
  //   final result = await accessToken.refresh();
  //   if (!result.isSuccess()) stopAccessTokenTimer();
  // }
  //
  // void _startAccessTokenTimer() {
  //   stopAccessTokenTimer();
  //   _timer = Timer(const Duration(hours: 3), () => refreshAccessToken());
  // }
  //
  // void stopAccessTokenTimer() {
  //   _timer?.cancel();
  //   _timer = null;
  // }
  //
  // void setVisitor(Visitor visitor) {
  //   client.type = 'jwt';
  //   client.token = visitor.token;
  //   auth.setVisitor(visitor);
  // }
  //
  // void logout() {
  //   stopAccessTokenTimer();
  //   client.type = null;
  //   client.token = null;
  //   if (auth.visitor != null) {
  //     appSetting.clearVisitorKeys();
  //   }
  //   auth.logout();
  //   cachedMember.clear();
  //   contacts.clear();
  // }
}
