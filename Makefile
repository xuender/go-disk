# 获取最后的tag
TAG = $(shell git tag -l --sort=v:refname | tail -n 1)

all: www build pack changelog

www:
	@echo 编译 www
	ng build --prod

build:
	@echo 编译
	go build -ldflags="-s -w" -o go-disk main.go

pack:
	@echo 打包 $(TAG)
	zip go-disk-$(TAG)-linux_amd64.zip go-disk www/* data/*
	rm go-disk

changelog:
	@echo 修改记录
	conventional-changelog -p angular -i CHANGELOG.md -s -r 0

todo:
	@echo 寻找未完成代码
	grep -rnw "TODO" gds src client

test:
	@echo 测试
	go test ./gds
