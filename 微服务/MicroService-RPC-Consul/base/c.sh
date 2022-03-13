#!/bin/bash

protoc --proto_path=. --micro_out=. --go_out=. ./proto/user.proto