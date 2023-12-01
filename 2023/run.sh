#!/bin/bash

export LUA_PATH="$(pwd)/?.lua;${LUA_PATH}"

cat $1.data | lua $1.lua
