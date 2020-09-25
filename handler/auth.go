package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": gin.H{
			"id":            "124",
			"name":          "zhangsan",
			"username":      "admin",
			"password":      "",
			"avatar":        "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
			"status":        1,
			"telephone":     "",
			"lastLoginIp":   "27.154.74.117",
			"lastLoginTime": 1534837621348,
			"creatorid":     "admin",
			"createTime":    1497160610259,
			"deleted":       0,
			"roleid":        "admin",
			"lang":          "zh-CN",
			"token":         "4291d7da9005377ec9aec4a71ea837f",
		},
	})
}

const info = `{
    "id": "4291d7da9005377ec9aec4a71ea837f",
    "name": "子不语",
    "username": "admin",
 	"password": "",
	"avatar": "https://tva2.sinaimg.cn/crop.0.0.1006.1006.180/a6a53118jw8fc26p9zxeuj20ry0ryq4l.jpg",
    "status": 1,
    "telephone": "",
    "lastLoginIp": "27.154.74.117",
    "lastLoginTime": 1534837621348,
    "creatorid": "admin",
    "createTime": 1497160610259,
    "merchantCode": "TLif2btpzg079h15bk",
    "deleted": 0,
    "roleid": "admin",
	"role": {
		"id": "admin",
		"name": "管理员",
		"describe": "拥有所有权限",
		"status": 1,
		"creatorid": "system",
		"createTime": 1497160610259,
		"deleted": 0,
		"permissions": [{
			"roleid": "admin",
			"permissionid": "dashboard",
			"permissionName": "仪表盘",
			"actions": [{
				"action": "add",
				"defaultCheck": false,
				"describe": "新增"
			}, {
				"action": "query",
				"defaultCheck": false,
				"describe": "查询"
			}, {
				"action": "get",
				"defaultCheck": false,
				"describe": "详情"
			}, {
				"action": "update",
				"defaultCheck": false,
				"describe": "修改"
			}, {
				"action": "delete",
				"defaultCheck": false,
				"describe": "删除"
			}],
			"actionEntitySet": [{
					"action": "add",
					"describe": "新增",
					"defaultCheck": false
				},
				{
					"action": "query",
					"describe": "查询",
					"defaultCheck": false
				},
				{
					"action": "get",
					"describe": "详情",
					"defaultCheck": false
				},
				{
					"action": "update",
					"describe": "修改",
					"defaultCheck": false
				},
				{
					"action": "delete",
					"describe": "删除",
					"defaultCheck": false
				}
			],
			"actionList": [],
			"dataAccess": []
		}]
	}
}`

func UserInfo(c *gin.Context) {
	map2 := make(map[string]interface{})
	err := json.Unmarshal([]byte(info), &map2)
	if err != nil {
		fmt.Print(err)
	}
	c.JSON(200, gin.H{
		"result": map2,
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": gin.H{},
	})

}

// http://localhost:8000/api/user/nav
const nav = `
[{
	"parentid": 0,
	"id": 11,
	"path": "/dashboard",
	"name": "dashboard",
	"redirect": "/dashboard/analysis",
	"component": "RouteLayout",
	"meta": {
		"title": "工作台",
		"keepAlive": true,
		"icon": "dashboard",
		"permission": ["dashboard"]
	}
}, {
	"parentid": 11,
	"id": 111,
	"path": "/dashboard/analysis",
	"name": "analysis",
	"component": "Index",
	"meta": {
		"title": "hello",
		"keepAlive": true,
		"permission": ["dashboard"]
	}
}, {
	"parentid": 0,
	"id": 12,
	"path": "/example",
	"name": "example",
	"redirect": "/example/table",
	"component": "RouteLayout",
	"meta": {
		"title": "示例页面",
		"keepAlive": true,
		"icon": "thunderbolt",
		"permission": ["dashboard"]
	}
}, {
	"parentid": 12,
	"id": 121,
	"path": "/example/table",
	"name": "exampleTableList",
	"component": "TableList",
	"meta": {
		"title": "table",
		"keepAlive": true,
		"permission": ["dashboard"]
	}
}, {
	"parentid": 12,
	"id": 122,
	"path": "/example/test",
	"name": "exampleTest",
	"component": "TableTest",
	"meta": {
		"title": "test",
		"keepAlive": true,
		"permission": ["dashboard"]
	}
}, {
	"parentid": 0,
	"id": 13,
	"path": "https://pro.loacg.com/docs/getting-started",
	"name": "docs",
	"meta": {
		"title": "在线文档",
		"icon": "select",
		"target": "_blank"
	}
}, {
	"parentid": 0,
	"id": 10024,
	"path": "/403",
	"name": "403",
	"component": "403",
	"meta": {
		"title": "403",
		"show": false
	}
}, {
	"parentid": 0,
	"id": 10024,
	"path": "/404",
	"name": "404",
	"component": "404",
	"meta": {
		"title": "404",
		"show": false
	}
}, {
	"parentid": 0,
	"id": 10024,
	"path": "/500",
	"name": "500",
	"component": "500",
	"meta": {
		"title": "500",
		"show": false
	}
}]
`

func UserNav(c *gin.Context) error {
	var map2 []map[string]interface{}
	err := json.Unmarshal([]byte(nav), &map2)
	if err != nil {
		return err
	}
	c.Set("data", map2)
	return nil
}
