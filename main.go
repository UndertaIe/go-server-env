package main

import (
	"flag"
	"log"
	"strings"

	"github.com/UndertaIe/passwd/cmd"
	"github.com/UndertaIe/passwd/config"
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/tracer"
)

var (
	port       int
	runMode    string
	configPath string
)

func init() { // 初始化工作
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func main() {
	cmd.Run()
}

func setupFlag() error {
	flag.IntVar(&port, "port", 7788, "启动端口7788")
	flag.StringVar(&runMode, "mode", "debug", "启动模式(debug,prod)")
	flag.StringVar(&configPath, "config", "./", "配置文件路径,当前路径下")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := config.NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.APPSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("MySQL", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("SQLITE3", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSettings)
	if err != nil {
		return err
	}
	return err
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = database.NewDBEngine(global.DatabaseSettings) // 在对全局变量赋值时不要使用 :=, 否则会导致左侧变量变为nil
	return err
}

func setupTracer() error {
	tracer, _, err := tracer.NewJaegerTracer("passwd-service", "127.0.0.1:6831")
	global.Tracer = tracer
	return err
}