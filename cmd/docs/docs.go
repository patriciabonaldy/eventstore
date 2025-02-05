// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/answers": {
            "post": {
                "description": "if ID params is not empty will update the record other cases will create a new answer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Create an event",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "201": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/answers/{id}": {
            "get": {
                "description": "get event by ID example:\"0bfce8da-bdc9-11ec-b9f3-acde48001122\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Show an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Update an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete event by ID example:\"0bfce8da-bdc9-11ec-b9f3-acde48001122\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Delete an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/answers/{id}/history": {
            "get": {
                "description": "get history of event by ID example:\"0bfce8da-bdc9-11ec-b9f3-acde48001122\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "answers"
                ],
                "summary": "Show a history event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CheckHandler"
                ],
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateRequest": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "data": {
                    "description": "data is answers list",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "the id for create a new event",
                    "type": "string",
                    "example": "0bfce8da-bdc9-11ec-b9f3-acde48001122"
                }
            }
        },
        "handler.Request": {
            "type": "object",
            "required": [
                "data",
                "event"
            ],
            "properties": {
                "data": {
                    "description": "data is answers list",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "en": "Map",
                        "kk": "Карталар",
                        "ru": "Карта"
                    }
                },
                "event": {
                    "description": "type of event",
                    "type": "string"
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "answer_id": {
                    "description": "event id",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "data": {
                    "description": "data is answers list",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "en": "Map",
                        "kk": "Карталар",
                        "ru": "Карта"
                    }
                },
                "event": {
                    "description": "type of event",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "version(1.0)",
	Host:             "0.0.0.0:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API document title",
	Description:      "Description of specifications",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
