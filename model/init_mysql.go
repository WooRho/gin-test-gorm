package model

import (
	"fmt"
	"gin-test-gorm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var userDbs sync.Map

// GetDb 获取指定数据库
func GetDb(mysql ...*config.Mysql) *gorm.DB {

	data, ok := userDbs.Load(config.Config.Mysql.DbName)
	if !ok || data.(*gorm.DB) == nil {
		singleDb := NewDatabase(&config.Config.Mysql)
		userDbs.Store(config.Config.Mysql.DbName, NewDatabase(&config.Config.Mysql))
		return singleDb
	}
	return data.(*gorm.DB)
}

func NewDatabase(config *config.Mysql) *gorm.DB {

	dblink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%v&collation=%v&loc=Local&parseTime=true",
		config.UserName,
		config.PassWord,
		config.Host,
		config.Port,
		config.DbName,
		config.Charset,
		config.Collation,
	)
	//加载日志
	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为3s
			SlowThreshold: 1 * time.Second,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(mysql.Open(dblink), &gorm.Config{
		SkipDefaultTransaction: false, //跳过事务执行
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "k_", //设置表前缀
			SingularTable: true, //在创建表名时采用单复数  也就是比如用结构体去创建表时 不加s
		},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键
		Logger:                                   slowLogger,
	})
	if err != nil {
		panic(err)
	}

	SetDB(config, db)

	return db
}

func SetDB(config *config.Mysql, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * 3)

	log.Println(`正在连接"` + config.DbName + `"数据库`)
	err = sqlDB.Ping()
	if err != nil {
		log.Println(config.DbName + `"数据库,连接失败,err:` + err.Error())
	} else {
		log.Println(`"` + config.DbName + `"数据库,连接成功`)
	}
}

const (
	SqlPathLocal = "docs/sql"
	SqlPathTest  = "sql/"
)

var sqlStr = ""

func SqlInit() {
	var (
		tplFileList []string
		err         error
	)
	tplFileList, err = GetAllTplFile(SqlPathLocal, nil)
	if err != nil {
		return
	}

	fileCount := len(tplFileList)
	if fileCount == 0 {
		return
	}
	for _, path := range tplFileList {
		c, ioErr := ioutil.ReadFile(path)
		if ioErr != nil {
			return
		}
		sqlStr = string(c)
		if sqlStr == "" {
			return
		}
		splitSQLs := strings.Split(sqlStr, ";")
		for _, splitSQL := range splitSQLs {
			if splitSQL == "" || strings.Trim(splitSQL, "\r\n") == "" {
				continue
			}
			err := GetDb().Exec(strings.Trim(splitSQL, "\n")).Error
			if err != nil {
				continue
			}
		}
	}
	return
}

// 获取 pathName 文件夹下所有 sql 文件
func GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".sql") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}
