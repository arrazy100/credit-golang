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
        "/auth/login": {
            "post": {
                "description": "Login to get token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register Debtor User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    }
                }
            }
        },
        "/debtor/detail": {
            "get": {
                "description": "Detail Debtor of logged in User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debtor"
                ],
                "summary": "Detail Debtor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DebtorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    }
                }
            }
        },
        "/debtor/register": {
            "post": {
                "description": "Register Debtor to get Tenor Limits",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "debtor"
                ],
                "summary": "Register Debtor",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterDebtorPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RegisterDebtorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    }
                }
            }
        },
        "/list/debtor": {
            "get": {
                "description": "List all registered Debtor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List Debtor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ListDebtorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custom_errors.ErrorValidation"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "custom_errors.ErrorField": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "custom_errors.ErrorValidation": {
            "type": "object",
            "properties": {
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/custom_errors.ErrorField"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "enums.TenorLimitType": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "Monthly",
                "Yearly"
            ]
        },
        "enums.UserRole": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "Admin",
                "Debtor"
            ]
        },
        "request.LoginPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 5
                }
            }
        },
        "request.RegisterDebtorPayload": {
            "type": "object",
            "required": [
                "date_of_birth",
                "full_name",
                "identity_picture_url",
                "legal_name",
                "nik",
                "place_of_birth",
                "salary",
                "selfie_picture_url"
            ],
            "properties": {
                "date_of_birth": {
                    "type": "string",
                    "minLength": 1
                },
                "full_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "identity_picture_url": {
                    "type": "string",
                    "maxLength": 2048,
                    "minLength": 1
                },
                "legal_name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "nik": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 16
                },
                "place_of_birth": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "salary": {
                    "type": "string"
                },
                "selfie_picture_url": {
                    "type": "string",
                    "maxLength": 2048,
                    "minLength": 1
                }
            }
        },
        "request.RegisterPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 5
                }
            }
        },
        "response.DebtorResponse": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "identity_picture_url": {
                    "type": "string"
                },
                "legal_name": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "place_of_birth": {
                    "type": "string"
                },
                "salary": {
                    "type": "string"
                },
                "selfie_picture_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "response.DebtorTenorLimitResponse": {
            "type": "object",
            "properties": {
                "limit_amount": {
                    "type": "string"
                },
                "tenor_duration": {
                    "type": "integer"
                },
                "tenor_limit_type": {
                    "$ref": "#/definitions/enums.TenorLimitType"
                }
            }
        },
        "response.ListDebtorResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.DebtorResponse"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/utils.SimpleAuth"
                }
            }
        },
        "response.RegisterDebtorResponse": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "debtor_tenor_limits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.DebtorTenorLimitResponse"
                    }
                },
                "full_name": {
                    "type": "string"
                },
                "identity_picture_url": {
                    "type": "string"
                },
                "legal_name": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "place_of_birth": {
                    "type": "string"
                },
                "salary": {
                    "type": "string"
                },
                "selfie_picture_url": {
                    "type": "string"
                }
            }
        },
        "response.RegisterResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enums.UserRole"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "utils.SimpleAuth": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enums.UserRole"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Debtor API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}