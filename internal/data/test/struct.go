package test

import "time"

// CrUpfileavatar 用户头像
type CrUpfileavatar struct {
	ID         uint32    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"` // id
	UserID     string    `gorm:"column:user_id;type:varchar(10);not null;default:''"` // 用户id
	Imgurl     string    `gorm:"column:imgurl;type:varchar(200);not null;default:''"` // 图片地址
	Imgname    string    `gorm:"column:imgname;type:varchar(50);not null;default:''"`
	Ratio      int       `gorm:"column:ratio;type:int(1);not null"`                                    // 图片的尺寸比例(1为原始180-180,2是80-80,3是48-48)
	IsDelete   int       `gorm:"column:is_delete;type:int(1);not null;default:0"`                      // 是否删除(0为默认，9为删除，8是机器人)
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 修改时间
}

// CrUpfileavatarColumns get sql column name.获取数据库列名
var CrUpfileavatarColumns = struct {
	ID         string
	UserID     string
	Imgurl     string
	Imgname    string
	Ratio      string
	IsDelete   string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	UserID:     "user_id",
	Imgurl:     "imgurl",
	Imgname:    "imgname",
	Ratio:      "ratio",
	IsDelete:   "is_delete",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// CrUser crawler用户表
type CrUser struct {
	UserID          uint32    `gorm:"primaryKey;index:index_name;column:user_id;type:int(10) unsigned;not null"` // 用户id
	Card            string    `gorm:"column:card;type:char(32);not null;default:''"`                             // 用户唯一身份标识
	Nickname        string    `gorm:"column:nickname;type:varchar(30);not null;default:''"`                      // 用户昵称
	Phone           string    `gorm:"unique;index:index_name;column:phone;type:char(11);not null;default:''"`    // 用户名
	Password        string    `gorm:"column:password;type:char(32);not null;default:''"`                         // 用户密码
	Token           string    `gorm:"column:token;type:varchar(255);not null;default:''"`
	UserStatus      int       `gorm:"column:user_status;type:int(10);not null;default:0"`                   // 用户认证类型(1是机构2是个人3是传媒公司，9是认证中)
	IsTeam          int       `gorm:"column:is_team;type:int(1);not null;default:0"`                        // 是否是团队
	Status          bool      `gorm:"column:status;type:tinyint(1);not null;default:0"`                     // 用户账号状态（0是默认，2是禁用，3是软删除，9是硬删）
	CreateTime      time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 修改时间
	TokenExpireTime int       `gorm:"column:token_expire_time;type:int(10);not null;default:0"`             // token过期时间
	Fans            int       `gorm:"column:fans;type:int(10);not null;default:0"`                          // 粉丝数
	Likes           int       `gorm:"column:likes;type:int(10);not null;default:0"`                         // 关注数
	IsRobot         bool      `gorm:"column:is_robot;type:tinyint(1);not null;default:0"`                   // 是否是机器人（0是默认正常账号，1是机器人，2是自有平台账号，3是adsaas等特殊账号）
	Integral        int       `gorm:"column:integral;type:int(10);not null;default:0"`                      // 用户积分
	Credit          int       `gorm:"column:credit;type:int(10);not null;default:0"`                        // 用户信用分
}

// CrUserColumns get sql column name.获取数据库列名
var CrUserColumns = struct {
	UserID          string
	Card            string
	Nickname        string
	Phone           string
	Password        string
	Token           string
	UserStatus      string
	IsTeam          string
	Status          string
	CreateTime      string
	UpdateTime      string
	TokenExpireTime string
	Fans            string
	Likes           string
	IsRobot         string
	Integral        string
	Credit          string
}{
	UserID:          "user_id",
	Card:            "card",
	Nickname:        "nickname",
	Phone:           "phone",
	Password:        "password",
	Token:           "token",
	UserStatus:      "user_status",
	IsTeam:          "is_team",
	Status:          "status",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	TokenExpireTime: "token_expire_time",
	Fans:            "fans",
	Likes:           "likes",
	IsRobot:         "is_robot",
	Integral:        "integral",
	Credit:          "credit",
}

// CrUserMes crawler用户简介表
type CrUserMes struct {
	ID         uint32    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`                  // id
	UserID     int       `gorm:"unique;column:user_id;type:int(10);not null;default:0"`                // 用户id
	Sex        int       `gorm:"column:sex;type:int(1);not null;default:1"`                            // 性别(1是男2是女3是未知,默认是0)
	Birthday   time.Time `gorm:"column:birthday;type:datetime;not null;default:1970-01-01 00:00:00"`   // 生日
	Province   string    `gorm:"column:province;type:varchar(50);not null;default:''"`                 // 省
	City       string    `gorm:"column:city;type:varchar(50);not null;default:''"`                     // 市
	Address    string    `gorm:"column:address;type:varchar(200);not null;default:''"`                 // 地址
	Des        string    `gorm:"column:des;type:text;not null"`                                        // 个人简介
	School     string    `gorm:"column:school;type:varchar(50);not null;default:''"`                   // 毕业院校
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 修改时间
	Card       string    `gorm:"column:card;type:char(32);not null;default:''"`                        // 用户唯一身份标识
	UserStatus int       `gorm:"column:user_status;type:int(11);not null;default:0"`                   // 用户认证类型(1是机构2是个人3是传媒公司，9是认证中)
	IsTeam     int       `gorm:"column:is_team;type:int(11);not null;default:0"`                       // 是否是团队
	Fans       int       `gorm:"column:fans;type:int(11);not null;default:0"`                          // 粉丝数
	Likes      int       `gorm:"column:likes;type:int(11);not null;default:0"`                         // 关注数
	IsRobot    int8      `gorm:"column:is_robot;type:tinyint(4);not null;default:0"`                   // 是否是机器人（0是默认正常账号，1是机器人，2是自有平台账号，3是adsaas等特殊账号）
	Integral   int       `gorm:"column:integral;type:int(11);not null;default:0"`                      // 用户积分
	Credit     int       `gorm:"column:credit;type:int(11);not null;default:0"`                        // 用户信用分
}

// CrUserMesColumns get sql column name.获取数据库列名
var CrUserMesColumns = struct {
	ID         string
	UserID     string
	Sex        string
	Birthday   string
	Province   string
	City       string
	Address    string
	Des        string
	School     string
	CreateTime string
	UpdateTime string
	Card       string
	UserStatus string
	IsTeam     string
	Fans       string
	Likes      string
	IsRobot    string
	Integral   string
	Credit     string
}{
	ID:         "id",
	UserID:     "user_id",
	Sex:        "sex",
	Birthday:   "birthday",
	Province:   "province",
	City:       "city",
	Address:    "address",
	Des:        "des",
	School:     "school",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Card:       "card",
	UserStatus: "user_status",
	IsTeam:     "is_team",
	Fans:       "fans",
	Likes:      "likes",
	IsRobot:    "is_robot",
	Integral:   "integral",
	Credit:     "credit",
}

// UcDeptMid uc用户部门信息表
type UcDeptMid struct {
	ID         uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`           // id
	UId        uint32 `gorm:"index:uid;column:uid;type:int(10) unsigned;not null;default:0"` // 用户id
	DepID      int    `gorm:"column:dep_id;type:int(11);not null;default:0"`                 // 部门id
	CompanyID  int    `gorm:"column:company_id;type:int(11);not null;default:0"`             // 公司id
	Status     int16  `gorm:"column:status;type:smallint(6);default:0"`                      // 1正常,9删除
	CreateTime int    `gorm:"column:create_time;type:int(11);not null;default:0"`            // 创建时间
	ModifyTime int    `gorm:"column:modify_time;type:int(11);not null;default:0"`            // 更新时间
}

// UcDeptMidColumns get sql column name.获取数据库列名
var UcDeptMidColumns = struct {
	ID         string
	UId        string
	DepID      string
	CompanyID  string
	Status     string
	CreateTime string
	ModifyTime string
}{
	ID:         "id",
	UId:        "uid",
	DepID:      "dep_id",
	CompanyID:  "company_id",
	Status:     "status",
	CreateTime: "create_time",
	ModifyTime: "modify_time",
}

// UcInshoot 文件追加附件表
type UcInshoot struct {
	ID   uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`
	UcID int    `gorm:"unique;column:uc_id;type:int(11);not null;default:0"`   // uc的用户id
	InID uint32 `gorm:"column:in_id;type:int(10) unsigned;not null;default:0"` // inshoot的用户id
}

// UcInshootColumns get sql column name.获取数据库列名
var UcInshootColumns = struct {
	ID   string
	UcID string
	InID string
}{
	ID:   "id",
	UcID: "uc_id",
	InID: "in_id",
}

// UcUsers 用户信息表
type UcUsers struct {
	ID            uint64 `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	UserEmail     string `gorm:"column:user_email;type:varchar(50);not null;default:''"`           // 邮箱
	UserPhone     string `gorm:"column:user_phone;type:char(11);not null"`                         // 手机号
	UserPassword  string `gorm:"column:user_password;type:varchar(255);not null"`                  // 密码（未确定加密方式）
	UserNickname  string `gorm:"column:user_nickname;type:varchar(50);not null;default:''"`        // 昵称
	UserName      string `gorm:"column:user_name;type:varchar(50);not null;default:''"`            // 真实姓名
	UserSex       int8   `gorm:"column:user_sex;type:tinyint(4);not null;default:0"`               // 性别0|未知 1|男 2|女
	UserLevel     uint8  `gorm:"column:user_level;type:tinyint(3) unsigned;not null;default:0"`    // 等级
	UserFace      string `gorm:"column:user_face;type:varchar(150);not null;default:''"`           // 头像
	UserLasttime  int    `gorm:"column:user_lasttime;type:int(11);not null;default:0"`             // 上次登录时间
	UserTopsign   uint16 `gorm:"column:user_topsign;type:smallint(5) unsigned;not null;default:0"` // 企业资质关联的ID 企业管理员和当前企业的用户用的是同一个资质id 默认0
	UserSource    uint8  `gorm:"column:user_source;type:tinyint(3) unsigned;not null;default:0"`   // 来源  1:企业微信 2:QQ 3:adsaas
	IsDimission   int8   `gorm:"column:is_dimission;type:tinyint(4);not null;default:0"`           // 是否离职  1|离职 0|在职
	DimissionTime int    `gorm:"column:dimission_time;type:int(11);not null;default:0"`            // 离职时间
	UserType      int8   `gorm:"column:user_type;type:tinyint(4);not null;default:0"`              // 标识adsaas的客户，1为adsaas客户
	UserEnable    uint8  `gorm:"column:user_enable;type:tinyint(3) unsigned;not null;default:1"`   // 状态 1开启 0停用
	CreateTime    uint32 `gorm:"column:create_time;type:int(10) unsigned;not null;default:0"`
	UpdateTime    uint32 `gorm:"column:update_time;type:int(10) unsigned;not null;default:0"`
	DeleteTime    uint32 `gorm:"column:delete_time;type:int(10) unsigned;not null;default:0"`
}

// UcUsersColumns get sql column name.获取数据库列名
var UcUsersColumns = struct {
	ID            string
	UserEmail     string
	UserPhone     string
	UserPassword  string
	UserNickname  string
	UserName      string
	UserSex       string
	UserLevel     string
	UserFace      string
	UserLasttime  string
	UserTopsign   string
	UserSource    string
	IsDimission   string
	DimissionTime string
	UserType      string
	UserEnable    string
	CreateTime    string
	UpdateTime    string
	DeleteTime    string
}{
	ID:            "id",
	UserEmail:     "user_email",
	UserPhone:     "user_phone",
	UserPassword:  "user_password",
	UserNickname:  "user_nickname",
	UserName:      "user_name",
	UserSex:       "user_sex",
	UserLevel:     "user_level",
	UserFace:      "user_face",
	UserLasttime:  "user_lasttime",
	UserTopsign:   "user_topsign",
	UserSource:    "user_source",
	IsDimission:   "is_dimission",
	DimissionTime: "dimission_time",
	UserType:      "user_type",
	UserEnable:    "user_enable",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	DeleteTime:    "delete_time",
}
