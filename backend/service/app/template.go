package app

const ShareEmailTemplate = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>卡片切换效果</title>\n    <style>\n        .swiper-container {\n            box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;\n            min-height: 450px;\n            padding: 5px 20px;\n            border-radius: 10px;\n        }\n\n        .swiper-wrapper {\n            font-size: 14px;\n            opacity: 0.9;\n        }\n\n        .link {\n            color: #067bef;\n        }\n    </style>\n</head>\n<body>\n<div class=\"swiper-container\">\n    <div class=\"swiper-wrapper\">\n        <p>Hi <a href=\"\" class=\"link\">%s</a> :</p>\n        <p>&nbsp;&nbsp;&nbsp;&nbsp;您的好友 '%s' 给您分享了一个文件链接【%s】邀请您一起进行协作，请在电脑端打开以下的分享链接:</p>\n        <p>\n            &nbsp;&nbsp;&nbsp;&nbsp;<a href=\"%s\" target=\"_blank\" class=\"link\">%s</a>\n        </p>\n    </div>\n</div>\n</body>\n</html>"

const WechatReplyTemplate = "验证码：%s，5分钟内有效"

const PerfectUserBaseInfoMessageTemplate = "恭喜🎉! 您已完善用户基本消息, 奖励5个链接数已下发"
const OpenMembershipMessageTemplate = "您已成为我们的会员，感谢您的信任和支持！我们将为您提供更多的优质服务，帮助您更好地进行资源管理和收藏。"
