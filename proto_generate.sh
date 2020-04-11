#!/bin/bash

protoc ./procpb/proc_1.proto --go_out=plugins=grpc:.