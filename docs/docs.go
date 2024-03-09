// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/balances": {
            "get": {
                "description": "get balances for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balances"
                ],
                "summary": "Get Balances",
                "parameters": [
                    {
                        "type": "string",
                        "name": "begin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "command",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns user balances",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ledger.Balance"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    }
                }
            }
        },
        "/queries": {
            "get": {
                "description": "get queries for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queries"
                ],
                "summary": "Get Queries",
                "responses": {
                    "200": {
                        "description": "Returns user queries",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Query"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    }
                }
            }
        },
        "/templates": {
            "get": {
                "description": "get templates for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "templates"
                ],
                "summary": "Get Templates",
                "responses": {
                    "200": {
                        "description": "Returns user templates",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    }
                }
            }
        },
        "/txs": {
            "get": {
                "description": "get transactions for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get Transactions",
                "parameters": [
                    {
                        "type": "string",
                        "name": "begin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "command",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns user transactions",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "New Transaction",
                "parameters": [
                    {
                        "description": "Transaction Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns new transaction",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ledger.Balance": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                }
            }
        },
        "main.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.Query": {
            "type": "object",
            "required": [
                "name",
                "query"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "query": {
                    "type": "string"
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "commodity": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Transaction": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Account"
                    }
                },
                "date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
