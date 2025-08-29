
# 初始化项目目录变量
HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output

# 应用名称/二进制文件名称
APPNAME = ddd-demo-go

GOPKGS  := $$(go list ./...| grep -vE "vendor")


#GOROOT  := $(GO_1_16_HOME)
GOROOT  := $(GO_1_19_HOME)
GOOS:=linux
GO      := $(GOROOT)/bin/go
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")
# 将 go 命令加入到 PATH，这样依赖 go 命令的其他编译脚本可以正常运行
PATH    := $(GOROOT)/bin:/opt/compiler/gcc-8.2/bin/:$(PATH)

# 设置编译时所需要的 Go 环境
export GOENV = $(HOMEDIR)/go.env

#执行编译，可使用命令 make 或 make all 执行， 顺序执行 prepare -> compile -> test -> package 几个阶段
all: prepare compile test package

# prepare阶段， 使用 bcloud 下载非 Go 依赖，可单独执行命令: make prepare
prepare: prepare-dep
prepare-dep:
	#bcloud local -U # 下载非 Go 依赖，依赖 BCLOUD 文件
	git version     # 低于 2.17.1 可能不能正常工作
	go env          # 打印出 go 环境信息，可用于排查问题

set-env:
	go mod download -x || go mod download -x # 下载 Go 依赖

# compile 阶段，执行编译命令，可单独执行命令: make compile
compile:build
build: set-env
	go build -o $(HOMEDIR)/bin/$(APPNAME)

# test 阶段，进行单元测试， 可单独执行命令: make test
# cover 平台会优先执行此命令
test: test-case
test-case: set-env
	go test -race -timeout=300s -v -cover $(GOPKGS) -coverprofile=coverage.out | tee unittest.txt

# package 阶段，对编译产出进行打包，输出到 output 目录， 可单独执行命令: make package
package: package-bin
package-bin:
	$(shell rm -rf $(OUTDIR))
	$(shell mkdir -p $(OUTDIR))
	$(shell mkdir -p $(OUTDIR)/var/)
	$(shell cp -a bin $(OUTDIR)/bin)
	$(shell cp -a conf $(OUTDIR)/conf)
	$(shell cp -a hestia_online $(OUTDIR)/hestia)
	$(shell date +"%F %T" >> $(OUTDIR)/var/app_version.txt) # 更新应用版本信息
	$(shell if [ -d "data_online" ]; then cp -r data_online $(OUTDIR)/data; fi)
	$(shell if [ -d "script" ]; then cp -r script $(OUTDIR)/script; fi)
	$(shell if [ -d "noahdes" ]; then cp -r noahdes $(OUTDIR)/; fi)
	$(shell if [ -d "webroot" ]; then cp -r webroot $(OUTDIR)/; fi)
	tree $(OUTDIR)

# clean 阶段，清除过程中的输出， 可单独执行命令: make clean
clean:
	rm -rf $(OUTDIR)

# avoid filename conflict and speed up build
.PHONY: all prepare compile test package  clean build
