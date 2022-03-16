// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "xiayoushuang",
            "email": "york-xia@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/reserve/project": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建储备库项目",
                "tags": [
                    "储备库 - 项目"
                ],
                "summary": "创建储备库项目",
                "parameters": [
                    {
                        "description": "ReserveReq",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ReserveReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "创建储备库项目成功"
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/reserve/project/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取储备库项目",
                "tags": [
                    "储备库 - 项目"
                ],
                "summary": "获取储备库项目",
                "parameters": [
                    {
                        "type": "string",
                        "description": "储备库项目id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询储备库项目成功",
                        "schema": {
                            "$ref": "#/definitions/vo.ReserveResp"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改储备库项目",
                "tags": [
                    "储备库 - 项目"
                ],
                "summary": "修改储备库项目",
                "parameters": [
                    {
                        "type": "string",
                        "description": "储备库项目id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改储备库项目成功"
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除储备库项目",
                "tags": [
                    "储备库 - 项目"
                ],
                "summary": "删除储备库项目",
                "parameters": [
                    {
                        "type": "string",
                        "description": "储备库项目id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除储备库项目成功"
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/reserve/projects": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取储备库项目列表",
                "tags": [
                    "储备库 - 项目"
                ],
                "summary": "获取储备库项目列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请求页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "description": "ReserveFilterParam",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ReserveFilterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询储备库项目列表成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/vo.DataPagination"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/vo.ListReserveProResp"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "登录"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "LoginReq",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应成功",
                        "schema": {
                            "$ref": "#/definitions/vo.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "vo.DataPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "数据"
                },
                "pagination": {
                    "description": "分页信息",
                    "$ref": "#/definitions/vo.Pagination"
                }
            }
        },
        "vo.Error": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "vo.InvestDetail": {
            "type": "object",
            "properties": {
                "comment": {
                    "description": "备注",
                    "type": "string"
                },
                "total": {
                    "description": "资金类别 总投资",
                    "type": "number"
                },
                "value": {
                    "description": "投资数额",
                    "type": "number"
                },
                "year": {
                    "description": "年份",
                    "type": "string"
                }
            }
        },
        "vo.InvestmentDetail": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "投资情况数额细节",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.InvestDetail"
                    }
                },
                "total": {
                    "description": "总投资",
                    "type": "number"
                },
                "type": {
                    "description": "投资类型",
                    "type": "integer"
                }
            }
        },
        "vo.ListReserveProResp": {
            "type": "object",
            "properties": {
                "construct_subject": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "project_type": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "vo.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "vo.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "token",
                    "type": "string"
                },
                "expiry": {
                    "description": "token 到期时间 默认两小时",
                    "type": "integer"
                },
                "token_type": {
                    "description": "认证类型",
                    "type": "string"
                }
            }
        },
        "vo.Pagination": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "请求页",
                    "type": "integer"
                },
                "page_size": {
                    "description": "页大小",
                    "type": "integer"
                },
                "total_count": {
                    "description": "数据总条数",
                    "type": "integer"
                }
            }
        },
        "vo.ReserveFilterParam": {
            "type": "object",
            "properties": {
                "construct_subject": {
                    "description": "建设主体",
                    "type": "string"
                },
                "level": {
                    "description": "项目级别",
                    "type": "integer"
                },
                "name": {
                    "description": "以下所有参数，有就传，无则不传\n项目名称",
                    "type": "string"
                },
                "period": {
                    "description": "计划周期(根据起止时间计算相差的月数)",
                    "type": "integer"
                },
                "plan_begin": {
                    "description": "计划开始时间",
                    "type": "string"
                },
                "project_type": {
                    "description": "项目类型",
                    "type": "integer"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                }
            }
        },
        "vo.ReserveReq": {
            "type": "object",
            "properties": {
                "add": {
                    "description": "新增建设用地",
                    "type": "number"
                },
                "company_business": {
                    "description": "企/事业单位(家)",
                    "type": "integer"
                },
                "construct_basis_necessity": {
                    "description": "建设依据及必要性",
                    "type": "string"
                },
                "construct_content_scope": {
                    "description": "建设内容及规模",
                    "type": "string"
                },
                "construct_site": {
                    "description": "建设地点",
                    "type": "string"
                },
                "construct_subject": {
                    "description": "建设主体",
                    "type": "string"
                },
                "contract": {
                    "description": "前期工作联系人",
                    "type": "string"
                },
                "enter_db_type": {
                    "description": "入库类别 0:A类,1:B类;2:C类",
                    "type": "integer"
                },
                "implement_type": {
                    "description": "实施类型 0:新开工,1:续建",
                    "type": "integer"
                },
                "investment_detail": {
                    "description": "资金详情 eg:\n\"[{\\\\\"type\\\\\":0, \\\\\"total\\\\\":100, \\\\\"detail\\\\\":[{\\\\\"total\\\\\": 100,\\\\\"year\\\\\": \\\\\"2022\\\\\",\\\\\"value\\\\\":20,\\\\\"comment\\\\\":\\\\\"xxx\\\\\"}, {\\\\\"total\\\\\": 100,\\\\\"year\\\\\": \\\\\"2023\\\\\",\\\\\"value\\\\\":30,\\\\\"comment\\\\\":\\\\\"xxx\\\\\"}, ...]}, {}...]\"\ntype说明： 0:区财政;1:自筹;2:其他",
                    "type": "string"
                },
                "is_land_use": {
                    "description": "是否有用地情况",
                    "type": "boolean"
                },
                "level": {
                    "description": "项目级别 0:区级,1:街镇级",
                    "type": "integer"
                },
                "move_land_comsumption": {
                    "description": "征迁/土地费用",
                    "type": "number"
                },
                "name": {
                    "description": "项目名称",
                    "type": "string"
                },
                "need_collect": {
                    "description": "需征地面积",
                    "type": "number"
                },
                "need_people_move": {
                    "description": "需拆迁农户/居民数(人)",
                    "type": "integer"
                },
                "no_conform_use_plan": {
                    "description": "不符合土地利用规划面积",
                    "type": "number"
                },
                "period": {
                    "description": "建设周期",
                    "type": "integer"
                },
                "phone": {
                    "description": "联系人手机号",
                    "type": "string"
                },
                "plan_begin": {
                    "description": "计划开工时间",
                    "type": "string"
                },
                "point_type": {
                    "description": "重点类型; 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型",
                    "type": "integer"
                },
                "project_consumption": {
                    "description": "工程费用",
                    "type": "number"
                },
                "project_type": {
                    "description": "项目类型; 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他",
                    "type": "integer"
                },
                "site_photo": {
                    "description": "无拆迁照片",
                    "type": "string"
                },
                "site_red": {
                    "description": "选址红线 0:有拆迁,1:无拆迁",
                    "type": "integer"
                },
                "total": {
                    "description": "总用亩",
                    "type": "number"
                },
                "total_investment": {
                    "description": "总投资",
                    "type": "number"
                },
                "upload_cad_id": {
                    "description": "CAD文件ID(上传文件接口返回的ID)",
                    "type": "string"
                }
            }
        },
        "vo.ReserveResp": {
            "type": "object",
            "properties": {
                "add": {
                    "description": "新增建设用地",
                    "type": "number"
                },
                "company_business": {
                    "description": "企/事业单位(家)",
                    "type": "integer"
                },
                "construct_basis_necessity": {
                    "description": "建设依据及必要性",
                    "type": "string"
                },
                "construct_content_scope": {
                    "description": "建设内容及规模",
                    "type": "string"
                },
                "construct_site": {
                    "description": "建设地点",
                    "type": "string"
                },
                "construct_subject": {
                    "description": "建设主体",
                    "type": "string"
                },
                "contract": {
                    "description": "前期工作联系人",
                    "type": "string"
                },
                "create_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "enter_db_type": {
                    "description": "入库类别 0:A类,1:B类;2:C类",
                    "type": "integer"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "implement_type": {
                    "description": "实施类型 0:新开工,1:续建",
                    "type": "integer"
                },
                "investment_detail": {
                    "description": "资金详情",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.InvestmentDetail"
                    }
                },
                "is_land_use": {
                    "description": "是否有用地情况",
                    "type": "boolean"
                },
                "level": {
                    "description": "项目级别 0:区级,1:街镇级",
                    "type": "integer"
                },
                "move_land_comsumption": {
                    "description": "征迁/土地费用",
                    "type": "number"
                },
                "name": {
                    "description": "项目名称",
                    "type": "string"
                },
                "need_collect": {
                    "description": "需征地面积",
                    "type": "number"
                },
                "need_people_move": {
                    "description": "需拆迁农户/居民数(人)",
                    "type": "integer"
                },
                "no_conform_use_plan": {
                    "description": "不符合土地利用规划面积",
                    "type": "number"
                },
                "period": {
                    "description": "建设周期",
                    "type": "integer"
                },
                "phone": {
                    "description": "联系人手机号",
                    "type": "string"
                },
                "plan_begin": {
                    "description": "计划开工时间",
                    "type": "string"
                },
                "point_type": {
                    "description": "重点类型; 0:省重点实施项目,1:省重点预备项目,2:省重大产业项目;3:省4+1项目;4:省6千亿项目;5:市重点实施项目;6:市重点预备项目;7:无重点类型",
                    "type": "integer"
                },
                "project_consumption": {
                    "description": "工程费用",
                    "type": "number"
                },
                "project_type": {
                    "description": "项目类型; 0:安置房,1:道路交通,2:市政设施;3:提升整治;4:卫生;5:五水共治;6:学校;7:其他",
                    "type": "integer"
                },
                "site_photo": {
                    "description": "无拆迁照片",
                    "type": "string"
                },
                "site_red": {
                    "description": "选址红线 0:有拆迁,1:无拆迁",
                    "type": "integer"
                },
                "status": {
                    "description": "项目状态 0:草稿,1:已入库,2:前期计划;3:已发文\"",
                    "type": "integer"
                },
                "total": {
                    "description": "总用亩",
                    "type": "number"
                },
                "total_investment": {
                    "description": "总投资",
                    "type": "number"
                },
                "upload_cad_id": {
                    "description": "CAD文件ID(上传文件接口返回的ID)",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "临安区政府投资项目管理后台API",
	Description:      "临安区政府投资项目管理后台API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
