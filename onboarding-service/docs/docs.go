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
        "/onboarding": {
            "post": {
                "description": "Create onboarding",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "onboarding"
                ],
                "summary": "Create onboarding",
                "parameters": [
                    {
                        "description": "Onboarding request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.OnboardingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.OnboardingStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/onboarding/{id}": {
            "get": {
                "description": "Get onboarding",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "onboarding"
                ],
                "summary": "Get onboarding",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Onboarding ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.OnboardingStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/onboarding/{id}/signature": {
            "post": {
                "description": "Sign agreement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "onboarding"
                ],
                "summary": "Sign agreement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Onboarding ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Signature request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.SignatureRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.OnboardingStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "http.OnboardingRequest": {
            "type": "object",
            "properties": {
                "account_type": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "http.OnboardingStatus": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "id": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/state.OnboardingState"
                }
            }
        },
        "http.SignatureRequest": {
            "type": "object",
            "properties": {
                "signature": {
                    "type": "string"
                }
            }
        },
        "state.OnboardingState": {
            "type": "string",
            "enum": [
                "processing",
                "failed",
                "fraud_not_passed",
                "signature_not_valid",
                "waiting_for_agreement_signature",
                "completed"
            ],
            "x-enum-varnames": [
                "ProcessingState",
                "FailedState",
                "FraudNotPassedState",
                "SignatureNotValidSate",
                "WaitingForAgreementSign",
                "CompletedState"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Onboarding Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
