import 'package:flutter/cupertino.dart';

Widget buildDonationItem(String asset, String text) {
  return Container(
    width: 174,
    height: 200,
    margin: EdgeInsets.fromLTRB(10, 0, 10, 0),
    decoration: ShapeDecoration(
      shape: RoundedRectangleBorder(
        side: BorderSide(width: 1, color: Color(0xFF151A1E)),
        borderRadius: BorderRadius.circular(20),
      ),
    ),
    child: Container(
        decoration: ShapeDecoration(
          shape: RoundedRectangleBorder(),
        ),
        child: Container(
          child: Column(
            children: [
              const SizedBox(height: 18),
              Image.asset(asset),
              const SizedBox(height: 4),
              Row(
                mainAxisAlignment: MainAxisAlignment.end,
                children: [
                  Text(
                    textAlign: TextAlign.end,
                    text,
                    style: const TextStyle(
                        fontFamily: "PretendardSemiBold", fontSize: 16),
                  ),
                  const SizedBox(width: 10),
                ],
              )
            ],
          ),
        )),
  );
}

Widget buildDonationList() {
  return Container(
    height: 200, // 리스트뷰의 높이를 설정
    child: ListView(
      scrollDirection: Axis.horizontal, // 가로로 스크롤되도록 설정
      children: <Widget>[
        buildDonationItem("assets/images/dog.png", "Use of\nabandoned dog..."),
        buildDonationItem(
            "assets/images/chair.png", "plastic bench\nsenior citizen ce..."),
        buildDonationItem(
            "assets/images/blanket.png", "Blankets made\nfrom recycled..."),
        buildDonationItem(
            "assets/images/bag.png", "Sturdy diaper bag\nmade from e-pet"),
      ],
    ),
  );
}

Widget buildDonationListGrid() {
  return Container(
    child: Column(
      children: [
        Row(
          children: [
            buildDonationItem(
                "assets/images/dog.png", "Use of\nabandoned dog..."),
            buildDonationItem("assets/images/chair.png",
                "plastic bench\nsenior citizen ce..."),
          ],
        ),
        SizedBox(height: 14),
        Row(
          children: [
            buildDonationItem(
                "assets/images/blanket.png", "Blankets made\nfrom recycled..."),
            buildDonationItem(
                "assets/images/bag.png", "Sturdy diaper bag\nmade from e-pet"),
          ],
        ),
        SizedBox(height: 24)
      ],
    ),
  );
}
