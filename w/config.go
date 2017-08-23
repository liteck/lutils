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

var secretLst secretConfig

type secretConfig struct {
	Lst  map[string]Secret
	Lock sync.Mutex
}

func (s secretConfig) Get(k string) Secret {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	return s.Lst[k]
}

func (s secretConfig) Set(k string, v Secret) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Lst[k] = v
}

func (s secretConfig) Del(k string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	delete(s.Lst, k)
}

func RegisterSecret(s ...Secret) error {
	if len(s) == 0 {
		return errors.New("配置参数不能为空!")
	}

	for _, v := range s {
		if err := v.valid(); err != nil {
			return err
		}
		secretLst.Set(v.AppId, v)
	}

	return nil
}

func DeleteSecret(app_id string) {
	secretLst.Del(app_id)
}

func getSecret(appid string) Secret {
	return secretLst.Get(appid)
}

func init() {
	secretLst = secretConfig{}
	secretLst.Lst = map[string]Secret{}
}
