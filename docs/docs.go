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
        "/actor/addActorInfo": {
            "post": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Add information about a new actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add actor information",
                "operationId": "addActorInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Actor information to add",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddActorInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Actor information added successfully",
                        "schema": {
                            "$ref": "#/definitions/model.ActorModel"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/actor/changeActorInfo": {
            "post": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Update information about an existing actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change actor information",
                "operationId": "changeActorInfo",
                "parameters": [
                    {
                        "description": "Actor information to change",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangeActorInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Actor information changed successfully",
                        "schema": {
                            "$ref": "#/definitions/model.ActorModel"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/actor/getList": {
            "get": {
                "security": [
                    {
                        "JWTRegularUserAuth": []
                    }
                ],
                "description": "Retrieve a list of actors along with the films they are associated with",
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of actors with their associated films",
                "operationId": "getActorsInfoListWithFilms",
                "responses": {
                    "200": {
                        "description": "List of actors with associated films",
                        "schema": {
                            "$ref": "#/definitions/model.GetActorsAndTeirFilmsService"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/actor/remove/{actorID}": {
            "delete": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Remove information about a specific actor by UUID",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove actor information",
                "operationId": "rmActorInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID of the actor to remove",
                        "name": "actorID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Actor information removed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/film/addFilmInfo": {
            "post": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Add information about a new film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add film information",
                "operationId": "addFilmInfo",
                "parameters": [
                    {
                        "description": "Film information to add",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddFilmInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Film information added successfully",
                        "schema": {
                            "$ref": "#/definitions/model.FilmModelResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/film/changeFilmInfo": {
            "post": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Update information about an existing film",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change film information",
                "operationId": "changeFilmInfo",
                "parameters": [
                    {
                        "description": "Film information to change",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangeFilmInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Film information changed successfully",
                        "schema": {
                            "$ref": "#/definitions/model.FilmModelResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/film/findFilm": {
            "get": {
                "security": [
                    {
                        "JWTRegularUserAuth": []
                    }
                ],
                "description": "Retrieve a list of films based on a search fragment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of films by fragment",
                "operationId": "getFilmsListByFragment",
                "parameters": [
                    {
                        "description": "Search fragment for films",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetFilmsListByFragmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of films matching the search fragment",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.FilmsListModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/film/getList": {
            "get": {
                "security": [
                    {
                        "JWTRegularUserAuth": []
                    }
                ],
                "description": "Retrieve a list of films based on specified criteria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of films",
                "operationId": "getFilmsList",
                "parameters": [
                    {
                        "description": "Criteria for retrieving films list",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetFilmsListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of films based on criteria",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.FilmsListModel"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/film/remove/{filmID}": {
            "delete": {
                "security": [
                    {
                        "JWTAdminAuth": []
                    }
                ],
                "description": "Remove information about a specific film by UUID",
                "produces": [
                    "application/json"
                ],
                "summary": "Remove film information",
                "operationId": "rmFilmInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID of the film to remove",
                        "name": "filmID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Film information removed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Log in with user credentials to get access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login as a user",
                "operationId": "loginUser",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User logged in successfully",
                        "schema": {
                            "$ref": "#/definitions/model.UserRequestModel"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to log in user",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register a new user",
                "operationId": "registerUser",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/model.UserRequestModel"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON or missing fields",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to register user",
                        "schema": {
                            "$ref": "#/definitions/model.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ActorModel": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "otherInfo": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.AddActorInfoRequest": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "otherInfo": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "model.AddFilmInfoRequest": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AddActorInfoRequest"
                    }
                },
                "description": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ChangeActorInfoRequest": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "otherInfo": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.ChangeFilmInfoRequest": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ChangeActorInfoRequest"
                    }
                },
                "description": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.ErrResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.FilmModelResponse": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ActorModel"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.FilmsListModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.GetActorsAndTeirFilmsService": {
            "type": "object",
            "properties": {
                "actorsList": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "model.GetFilmsListByFragmentRequest": {
            "type": "object",
            "properties": {
                "fragment": {
                    "type": "string"
                },
                "fragmentOf": {
                    "type": "string"
                }
            }
        },
        "model.GetFilmsListRequest": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "string"
                },
                "sortBy": {
                    "type": "string"
                }
            }
        },
        "model.Login": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Register": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserRequestModel": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "role_id": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Film Libriary RestAPI app",
	Description:      "Its the Film Libtiary app that uses pSQL, JWT netHttp lib",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
