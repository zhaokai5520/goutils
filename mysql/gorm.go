package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(dbType string, dsn string, maxIdleConns int, maxOpenConns int) *gorm.DB {
	switch dbType {
	case "mysql":
		return GormMysql(dsn, maxIdleConns, maxOpenConns)
	default:
		return GormMysql(dsn, maxIdleConns, maxOpenConns)
	}

}

func GormMysql(dsn string, maxIdleConns int, maxOpenConns int) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: &LogxGorm,
	})
	// db.Logger = utils.GetGormLogger
	if err != nil {

		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		return db
	}
}
