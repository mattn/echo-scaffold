# Echo Scaffold

`Echo Scaffold` is CLI to generate scaffolds for the `echo` framework.
For now the project only supports `mongodb` and `mgo` as database.

## Installation

	go get github.com/mattn/echo-scaffold

## Initializing a project

	echo-scaffold init <project path>

## Creating a model

	echo-scaffold model <model name> <field name>:<field type>

## Creating a controller

	echo-scaffold controller <controller name>

## Creating a scaffold

	echo-scaffold scaffold <controller name> <field name>:<field type>

## Running

	go run <project name>.go

## Accessing

	Open browser, and access to http://localhost:4000. (Default port:4000)

## Thanks

This is based on [gin-scaffold](https://github.com/dcu/gin-scaffold)
