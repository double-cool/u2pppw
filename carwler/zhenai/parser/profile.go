package main

import (
	"encoding/json"
	"fmt"
)

const profileJsonRe = `window.__INITIAL_STATE__=(.*?)`

type BasicInfo struct {
	Mari      string `json:"0"`
	Age       string `json:"1"`
	Xingzuo   string `json:"2"`
	Height    string `json:"3"`
	Weight    string `json:"4"`
	JobAddr   string `json:"5"`
	Salary    string `json:"6"`
	Work      string `json:"7"`
	Education string `json:"8"`
}

//type ObjectInfo struct {
//	BasicInfo BasicInfo `json:"basicInfo"`
//}

type ObjectInfo struct {
	BasicInfo BasicInfo `json:"basicInfo"`
}

type profile struct {
	ObjectInfo map[string]ObjectInfo `json:"objectInfo"`
}

func main() {
	profileJson := []byte(`{
    "objectInfo":{
        "age":39,
        "avatarPhotoID":259890821,
        "avatarPraiseCount":2,
        "avatarPraised":false,
        "avatarURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F263471\u002F1053881159\u002F23262817746149784.png",
        "basicInfo":[
            "离异",
            "39岁",
            "魔羯座(12.22-01.19)",
            "178cm",
            "64kg",
            "工作地:大理大理市",
            "月收入:5-8千",
            "服务业",
            "中专"
        ],
        "deliverLove":{
            "show":false
        },
        "detailInfo":[
            "汉族",
            "籍贯:四川南充",
            "体型:瘦长",
            "稍微抽一点烟",
            "不喝酒",
            "租房",
            "未买车",
            "有孩子但不在身边",
            "是否想要孩子:想要孩子",
            "何时结婚:一年内"
        ],
        "educationString":"中专",
        "emotionStatus":0,
        "gender":0,
        "genderString":"男士",
        "hasIntroduce":true,
        "hasSendMail":false,
        "heightString":"178cm",
        "hideVerifyModule":false,
        "introduceContent":"以诚相待，珍惜，责任。彼此包容，人生最最大的幸福就是不离不弃。",
        "introducePraiseCount":0,
        "isActive":true,
        "isFollowing":false,
        "isInBlackList":false,
        "isOfflineSuperRecUser":false,
        "isRecentlyActive":false,
        "isStar":false,
        "isSuperVip":false,
        "isZhenaiMail":true,
        "lastLoginTimeString":"",
        "liveAudienceCount":0,
        "liveType":0,
        "marriageString":"离异",
        "memberID":1053881159,
        "momentCount":0,
        "nickname":"小何",
        "objectAgeString":"18-41岁",
        "objectBodyString":"未填写",
        "objectChildrenString":"未填写",
        "objectEducationString":"未填写",
        "objectHeightString":"155cm以上",
        "objectInfo":[
            "18-41岁",
            "155cm以上",
            "工作地:云南大理",
            "是否想要孩子:想要孩子"
        ],
        "objectMarriageString":"未填写",
        "objectSalaryString":"未填写",
        "objectWantChildrenString":"想要孩子",
        "objectWorkCityString":"云南大理",
        "onlive":0,
        "photoCount":3,
        "photos":[
            {
                "createTime":"2019-07-01 15:20:08",
                "isAvatar":true,
                "photoID":259890821,
                "photoType":1,
                "photoURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F263471\u002F1053881159\u002F23262817746149784.png",
                "praiseCount":2,
                "praised":false,
                "verified":true
            },
            {
                "createTime":"2019-07-10 02:17:46",
                "isAvatar":false,
                "photoID":262614063,
                "photoType":1,
                "photoURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F263471\u002F1053881159\u002F37409560969612135.png",
                "praiseCount":2,
                "praised":false,
                "verified":true
            },
            {
                "createTime":"2020-05-15 00:02:43",
                "isAvatar":false,
                "photoID":337908905,
                "photoType":1,
                "photoURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F263471\u002F1053881159\u002F231655674300555826.png",
                "praiseCount":1,
                "praised":false,
                "verified":true
            }
        ],
        "praisedIntroduce":false,
        "previewPhotoURL":"",
        "pycreditCertify":false,
        "recommendUpgrade2":false,
        "recommendUpgrade3":false,
        "salaryString":"5001-8000元",
        "showHighVipPic":false,
        "showValidateIDCardFlag":false,
        "superRecClickTip":"将在24小时内推荐给大量心仪异性",
        "superRecGuideTip":"",
        "superRecommend":false,
        "totalPhotoCount":3,
        "validateEducation":false,
        "validateFace":false,
        "validateIDCard":true,
        "videoCount":0,
        "videoID":0,
        "workCity":10130150,
        "workCityString":"大理",
        "workProvinceCityString":"大理大理市"
    },
    "interest":[
        {
            "answerContent":"默认选项",
            "answerContentDetail":"",
            "answerContentDetailStatus":1,
            "answerGuideWord":"一个菜名，说一道最符合你口味的菜？",
            "answerID":1061052,
            "answerOrder":0,
            "answerRecordID":959625,
            "answerWriteRule":0,
            "iconURL":"https:\u002F\u002Fimages.zastatic.com\u002Fapp\u002Fmarriageview\u002Finterest_1060889_p.png",
            "newInterest":false,
            "praiseCount":0,
            "questionGuideWord":"",
            "questionID":1060889,
            "questionName":"喜欢的一道菜",
            "questionType":0,
            "verifyStatus":4
        },
        {
            "answerContent":"默认选项",
            "answerContentDetail":"",
            "answerContentDetailStatus":1,
            "answerGuideWord":"一个名字，最让你敬佩的名人是谁？",
            "answerID":1061054,
            "answerOrder":0,
            "answerRecordID":0,
            "answerWriteRule":0,
            "iconURL":"https:\u002F\u002Fimages.zastatic.com\u002Fapp\u002Fmarriageview\u002Finterest_1060890_p.png",
            "newInterest":false,
            "praiseCount":0,
            "questionGuideWord":"",
            "questionID":1060890,
            "questionName":"欣赏的一个名人",
            "questionType":0,
            "verifyStatus":5
        },
        {
            "answerContent":"默认选项",
            "answerContentDetail":"",
            "answerContentDetailStatus":1,
            "answerGuideWord":"一首歌名，哪首歌最让你感动？",
            "answerID":1061056,
            "answerOrder":0,
            "answerRecordID":0,
            "answerWriteRule":0,
            "iconURL":"https:\u002F\u002Fimages.zastatic.com\u002Fapp\u002Fmarriageview\u002Finterest_1060891_p.png",
            "newInterest":false,
            "praiseCount":0,
            "questionGuideWord":"",
            "questionID":1060891,
            "questionName":"喜欢的一首歌",
            "questionType":0,
            "verifyStatus":5
        },
        {
            "answerContent":"默认选项",
            "answerContentDetail":"",
            "answerContentDetailStatus":1,
            "answerGuideWord":"一本书名，让你印象最深的是哪一本？",
            "answerID":1061058,
            "answerOrder":0,
            "answerRecordID":0,
            "answerWriteRule":0,
            "iconURL":"https:\u002F\u002Fimages.zastatic.com\u002Fapp\u002Fmarriageview\u002Finterest_1060892_p.png",
            "newInterest":false,
            "praiseCount":0,
            "questionGuideWord":"",
            "questionID":1060892,
            "questionName":"喜欢的一本书",
            "questionType":0,
            "verifyStatus":5
        },
        {
            "answerContent":"默认选项",
            "answerContentDetail":"",
            "answerContentDetailRecords":[

            ],
            "answerContentDetailStatus":1,
            "answerGuideWord":"还有什么是你喜欢的？",
            "answerID":1061060,
            "answerOrder":0,
            "answerRecordID":0,
            "answerWriteRule":0,
            "iconURL":"https:\u002F\u002Fimages.zastatic.com\u002Fapp\u002Fmarriageview\u002Finterest_1060893_p.png",
            "newInterest":true,
            "praiseCount":0,
            "questionGuideWord":"",
            "questionID":1060893,
            "questionName":"喜欢做的事",
            "questionType":0,
            "verifyStatus":5
        }
    ],
    "memberList":[

    ],
    "seoInfo":{
        "location":[
            {
                "currLocation":false,
                "linkContent":"珍爱首页",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com"
            },
            {
                "currLocation":false,
                "linkContent":"大理征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fdali1"
            },
            {
                "currLocation":false,
                "linkContent":"大理男士征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fdali1\u002Fnan"
            }
        ],
        "nearbyCitys":[
            {
                "linkContent":"昆明征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fkunming"
            },
            {
                "linkContent":"曲靖征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fqujing"
            },
            {
                "linkContent":"昭通征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fzhaotong"
            },
            {
                "linkContent":"文山征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fwenshan"
            },
            {
                "linkContent":"楚雄征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fchuxiong"
            },
            {
                "linkContent":"临沧征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Flincang"
            },
            {
                "linkContent":"保山征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fbaoshan1"
            },
            {
                "linkContent":"玉溪征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fyuxi"
            },
            {
                "linkContent":"丽江征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Flijiang"
            },
            {
                "linkContent":"普洱征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fpuer"
            },
            {
                "linkContent":"红河征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fhonghe1"
            },
            {
                "linkContent":"西双版纳征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fxishuangbanna"
            },
            {
                "linkContent":"德宏征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fdehong"
            },
            {
                "linkContent":"怒江征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fnujiang"
            },
            {
                "linkContent":"迪庆征婚",
                "linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Fdiqing"
            }
        ]
    },
    "isSearchEntryFlag":false
}`)
	//re := regexp.MustCompile(profileJsonRe)
	//match := re.Education(contents)

	//	result := engine.ParseResult{}
	//	if match != nil {
	//	profileJson := match[1]
	var profileArr profile
	//var baseInfo BasicInfo

	err := json.Unmarshal(profileJson, &profileArr)
	if err != nil {
		fmt.Println("aaaaaaaaaaaa")
	}
	//if err == nil {
	//	test := profileArr["objectInfo"].(map[string]interface{})["basicInfo"].([]interface{})
	//
	//	mari := fmt.Sprintf("%v", test[0])
	//	baseInfo.Mari = mari
	//	Age := fmt.Sprintf("%v", test[1])
	//	baseInfo.Age = Age
	//	Xingzuo := fmt.Sprintf("%v", test[2])
	//	baseInfo.Xingzuo = Xingzuo
	//	Height := fmt.Sprintf("%v", test[3])
	//	baseInfo.Height = Height
	//	Weight := fmt.Sprintf("%v", test[4])
	//	baseInfo.Weight = Weight
	//	JobAddr := fmt.Sprintf("%v", test[5])
	//	baseInfo.JobAddr = JobAddr
	//	Salary := fmt.Sprintf("%v", test[6])
	//	baseInfo.Salary = Salary
	//	Work := fmt.Sprintf("%v", test[7])
	//	baseInfo.Work = Work
	//	Education := fmt.Sprintf("%v", test[8])
	//	baseInfo.Education = Education
	//
	//}
	fmt.Println("aaaa", profileArr)

	return
}
