package rabbitmq

import "go_bbs/setting"

func init() {
	dbCfg := setting.RabbitConfig{
		Host:     "127.0.0.1",
		User:     "guest",
		Password: "guest",
		Port:     5672,
	}

	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}
