package mq

var cfg *MqConfig

var alipay_mq_cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIDPDCCAqWgAwIBAgIJAMRsb0DLM1fsMA0GCSqGSIb3DQEBBQUAMHIxCzAJBgNV
BAYTAkNOMQswCQYDVQQIEwJIWjELMAkGA1UEBxMCSFoxCzAJBgNVBAoTAkFCMRAw
DgYDVQQDEwdLYWZrYUNBMSowKAYJKoZIhvcNAQkBFht6aGVuZG9uZ2xpdS5semRA
YWxpYmFiYS5jb20wIBcNMTcwMzA5MTI1MDUyWhgPMjEwMTAyMTcxMjUwNTJaMHIx
CzAJBgNVBAYTAkNOMQswCQYDVQQIEwJIWjELMAkGA1UEBxMCSFoxCzAJBgNVBAoT
AkFCMRAwDgYDVQQDEwdLYWZrYUNBMSowKAYJKoZIhvcNAQkBFht6aGVuZG9uZ2xp
dS5semRAYWxpYmFiYS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBALZV
bbIO1ULQQN853BTBgRfPiRJaAOWf38u8GC0TNp/E9qtI88A+79ywAP17k5WYJ7XS
wXMOJ3h1qkQT2TYJVetZ6E69CUJq4BsOvNlNRvmnW6eFymh5QZsEz2MTooxJjVjC
JQPlI2XRDjIrTVYEQWUDxj2JhB8VVPEed+6u4KQVAgMBAAGjgdcwgdQwHQYDVR0O
BBYEFHFlOoiqQxXanVi2GUoDiKDD33ujMIGkBgNVHSMEgZwwgZmAFHFlOoiqQxXa
nVi2GUoDiKDD33ujoXakdDByMQswCQYDVQQGEwJDTjELMAkGA1UECBMCSFoxCzAJ
BgNVBAcTAkhaMQswCQYDVQQKEwJBQjEQMA4GA1UEAxMHS2Fma2FDQTEqMCgGCSqG
SIb3DQEJARYbemhlbmRvbmdsaXUubHpkQGFsaWJhYmEuY29tggkAxGxvQMszV+ww
DAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOBgQBTSz04p0AJXKl30sHw+UM/
/k1jGFJzI5p0Z6l2JzKQYPP3PfE/biE8/rmiGYEenNqWNy1ZSniEHwa8L/Ux98ci
4H0ZSpUrMo2+6bfuNW9X35CFPp5vYYJqftilJBKIJX3C3J1ruOuBR28UxE42xx4K
pQ70wChNi914c4B+SxkGUg==
-----END CERTIFICATE-----
`)

type MqConfig struct {
	Topics     []string
	Servers    []string
	Ak         string
	Password   string
	ConsumerId string
	CertBytes  []byte
}

func SetConfig(config MqConfig) {
	cfg = &config
	cfg.CertBytes = alipay_mq_cert
}

func (m *MqConfig) SetTopics(topics ...string) {
	for _, topic := range topics {
		same := false
		for _, t := range m.Topics {
			if t == topic {
				same = true
				break
			}
		}
		if same {
			continue
		}
		m.Topics = append(m.Topics, topic)
	}
}

func (m *MqConfig) SetServers(services ...string) {
	for _, service := range services {
		same := false
		for _, s := range m.Servers {
			if s == service {
				same = true
				break
			}
		}
		if same {
			continue
		}
		m.Servers = append(m.Servers, service)
	}
}

func (m *MqConfig) SetAk(ak string) {
	m.Ak = ak
}

func (m *MqConfig) SetPassword(password string) {
	m.Password = password
}

func (m *MqConfig) SetConsumerId(consumerId string) {
	m.ConsumerId = consumerId
}
