package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"
	httpclient "youdao/util"
)

const (
	signType = "v3"   //签名版本
	to       = "auto" //目标语言 auto为自动
	from     = "auto" //源语言 auto为自动
)

type youDaoConf struct {
	YouDaoUrl    string `json:"YouDaoUrl"`    //有道的翻译地址
	YouDaoAppID  string `json:"YouDaoAppID"`  //应用ID
	YouDaoAppKey string `json:"YouDaoAppKey"` //应用密钥
}

//后期可能会添加配置实时更新功能
func getYDConf() (youDaoConf, error) {
	//获取有道配置
	ydConf := youDaoConf{}
	err := viper.UnmarshalKey("app.youDao", &ydConf)
	if err != nil {
		return ydConf, errors.WithStack(err)
	}
	if ydConf.YouDaoUrl == "" || ydConf.YouDaoAppID == "" || ydConf.YouDaoAppKey == "" {
		return ydConf, errors.Errorf("youdao conf has a error, struct is : %+v", ydConf)
	}
	return ydConf, nil
}

type YDRequest struct {
	CurTime  string `json:"curtime"`  //当前UTC时间戳(秒)
	SignType string `json:"signType"` //签名类型
	Sign     string `json:"sign"`     //签名
	Salt     string `json:"salt"`     //UUID
	AppKey   string `json:"appKey"`   //应用ID
	To       string `json:"to"`       //目标语言
	From     string `json:"from"`     //源语言
	Q        string `json:"q"`        //待翻译文本
}

func (r *FYReq) YouDaoServer() (res *AutoGenerated, err error) {

	ydConf, err := getYDConf()
	if err != nil {
		return nil, err
	}

	res = &AutoGenerated{}
	req := YDRequest{
		CurTime:  strconv.FormatInt(time.Now().Unix(), 10),
		SignType: signType,
		Salt:     r.Ip, //把客户端的ip当作用户id
		AppKey:   ydConf.YouDaoAppID,
		To:       to,
		From:     from,
		Q:        r.Content,
	}
	//生成签名
	req.buildSign(ydConf.YouDaoAppKey)
	raw, err := json.Marshal(req)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	//切片通过json转成map
	var reqMap map[string]string
	err = json.Unmarshal(raw, &reqMap)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	//请求有道词典接口
	status, err := httpclient.PostForm(ydConf.YouDaoUrl, reqMap, nil, res)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	if status != http.StatusOK {
		err = errors.Errorf("失败，%+v", status)
		return nil, err
	}
	if res.ErrorCode != "0" {
		return nil, nil
	}
	return
}

func (S *YDRequest) buildSign(key string) {
	//生成input
	input := S.buildInputString()
	//sign=sha256(应用ID+input+salt+curtime+应用密钥)；
	sha := sha256.Sum256([]byte(S.AppKey + input + S.Salt + S.CurTime + key))
	S.Sign = hex.EncodeToString(sha[:])
}
func (S *YDRequest) buildInputString() (input string) {
	//input的计算方式为：input=q前10个字符 + q长度 + q后10个字符（当q长度大于20）或 input=q字符串（当q长度小于等于20）
	qLength := utf8.RuneCountInString(S.Q)
	if qLength <= 20 {
		input = S.Q
		return
	}

	input = string(exutf8.RuneSub([]byte(S.Q), 0, 10)) //性能比string([]rune(S.Q)[:10])快10倍+
	input += strconv.Itoa(qLength)
	input += string(exutf8.RuneSub([]byte(S.Q), -10, 10))
	return
}
