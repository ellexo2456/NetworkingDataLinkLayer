// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/code": {
            "post": {
                "description": "Кодирует и декодирует полученный в виде байт сегмент, вносит ошибку, исправляет ее, так же с вероятностью возвращает сегмент на траспортный уровень.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "EncodeSegmentSimulate.",
                "parameters": [
                    {
                        "description": "Пользовательский объект в формате JSON",
                        "name": "segment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/segment.SegmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ"
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/swag.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/swag.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "segment.SegmentRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "что-то"
                },
                "segment_number": {
                    "type": "integer",
                    "example": 1
                },
                "sender": {
                    "type": "string",
                    "example": "Некто"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2024-03-09T12:04:08Z"
                },
                "total_segments": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "swag.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "DataLinkLayer API",
	Description:      "API server for DataLinkLayer application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
