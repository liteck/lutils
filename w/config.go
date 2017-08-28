package w

import (
	"errors"
	"sync"
)

var (
	ErrMethodNotSupport  = errors.New("METHOD NOT SUPPORT")
	ErrMethodNameNil     = errors.New("METHOD NAME NIL")
	ErrBizContentNameNil = errors.New("BIZ CONTENT NIL")
	ErrAppIdNil          = errors.New("APPID NIL")
	ErrSecretNil         = errors.New("SECRET NIL")
	ErrSign              = errors.New("SIGN ERROR")
	ErrVerifySign        = errors.New("VERIFY SIGN ERROR")
)

const (
	CAN_NOT_NIL  = "不能为空"
	FORAMT_ERROR = "格式错误"
)

type Secret struct {
	AppId     string
	AppSecret string
}

func (s *Secret) valid() error {
	if len(s.AppId) == 0 {
		return errors.New("appid 不能为空")
	}

	if len(s.AppSecret) == 0 {
		return errors.New("appsecret 不能为空")
	}

	return nil
}

var secretLst sync.Map

func RegisterSecret(s ...Secret) error {
	if len(s) == 0 {
		return errors.New("配置参数不能为空!")
	}

	for _, v := range s {
		if err := v.valid(); err != nil {
			return err
		}
		secretLst.Store(v.AppId, v)
	}

	return nil
}

func DeleteSecret(app_id string) {
	secretLst.Delete(app_id)
}

func getSecret(appid string) Secret {
	if v, ok := secretLst.Load(appid); !ok || v == nil {
		return Secret{}
	} else {
		return v.(Secret)
	}
}

func init() {
	secretLst = sync.Map{}
}
