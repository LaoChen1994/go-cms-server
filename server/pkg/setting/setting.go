package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type SqlConf struct {
	DatabaseType string
	User         string
	Password     string
	Host         string
	Name         string
	Prefix       string
}

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize     int
	JwtSecret    string
	DatabaseConf *SqlConf
)

func init() {
	t, err := ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse app.ini", err)
		return
	}

	Cfg = t

	loadBase()
	loadApp()
	loadServer()
	loadDatabase()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadApp() {
	app, err := Cfg.GetSection("app")

	if err != nil {
		log.Fatalf("load App Fail, ", err)
		return
	}

	PageSize = app.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = app.Key("JWT_SECRET").MustString("")

	if JwtSecret == "" {
		panic("jwt secret cannot be empty")
	}
}

func loadServer() {
	app, err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("load server =>", err)
		return
	}

	HttpPort = app.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(app.Key("READ_TIMEOUT").MustInt(8000)) * time.Second
	WriteTimeout = time.Duration(app.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadDatabase() {
	sec, err := Cfg.GetSection("database")

	if err != nil {
		log.Fatalf("load database =>", err)
		return
	}

	DatabaseConf = &SqlConf{
		DatabaseType: sec.Key("TYPE").MustString("mysql"),
		User:         sec.Key("USER").MustString(""),
		Password:     sec.Key("PASSWORD").MustString(""),
		Host:         sec.Key("HOST").MustString(""),
		Name:         sec.Key("NAME").MustString(""),
		Prefix:       sec.Key("TABLE_PREFIX").MustString(""),
	}
}
