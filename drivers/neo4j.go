package drivers

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/viper"
	"sync"
)

// 单例实例的全局变量，使用结构体指针
var instance *Neo4jSessionManager

// 锁，用于初始化时避免并发问题
var once sync.Once

// Neo4jSessionManager 管理 Neo4j 会话
type Neo4jSessionManager struct {
	driver neo4j.DriverWithContext
}

// NewNeo4jSessionManager 创建一个新的会话管理器
func NewNeo4jSessionManager(dbUri, dbUser, dbPassword string) (*Neo4jSessionManager, error) {
	auth := neo4j.BasicAuth(dbUser, dbPassword, "")
	driver, err := neo4j.NewDriverWithContext(dbUri, auth)
	if err != nil {
		return nil, err
	}
	return &Neo4jSessionManager{
		driver: driver,
	}, nil
}

// Close 关闭驱动器
func (m *Neo4jSessionManager) Close(ctx context.Context) {
	m.driver.Close(ctx)
}

// ExecuteWrite 执行写入操作
func (m *Neo4jSessionManager) ExecuteWrite(ctx context.Context, query string, params map[string]interface{}) error {
	session := m.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, query, params)
	return err
}

// ExecuteRead 执行读取操作
func (m *Neo4jSessionManager) ExecuteRead(ctx context.Context, query string, params map[string]interface{}) (neo4j.ResultWithContext, error) {
	session := m.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	return session.Run(ctx, query, params)
}

// GetInstance 获取单例实例
func GetInstance() *Neo4jSessionManager {
	once.Do(func() {
		var err error
		dbUri := viper.GetString("database.uri")
		dbUser := viper.GetString("database.user")
		dbPassword := viper.GetString("database.password")
		instance, err = NewNeo4jSessionManager(dbUri, dbUser, dbPassword)
		if err != nil {
			panic(err)
		}
	})
	return instance
}
