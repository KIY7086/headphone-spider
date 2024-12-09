#!/bin/bash

# 创建输出目录
mkdir -p build

# 编译 Windows 版本
echo "正在编译 Windows 版本..."
GOOS=windows GOARCH=amd64 go build -o build/Windows-amd64-headphone-spider.exe
GOOS=windows GOARCH=386 go build -o build/Windows-386-headphone-spider.exe

# 编译 Linux 版本
echo "正在编译 Linux 版本..."
GOOS=linux GOARCH=amd64 go build -o build/Linux-amd64-headphone-spider
GOOS=linux GOARCH=386 go build -o build/Linux-386-headphone-spider
GOOS=linux GOARCH=arm64 go build -o build/Linux-arm64-headphone-spider
GOOS=linux GOARCH=arm go build -o build/Linux-arm-headphone-spider

# 编译 macOS 版本
echo "正在编译 macOS 版本..."
GOOS=darwin GOARCH=amd64 go build -o build/macOS-amd64-headphone-spider
GOOS=darwin GOARCH=arm64 go build -o build/macOS-arm64-headphone-spider

# 编译 FreeBSD 版本
echo "正在编译 FreeBSD 版本..."
GOOS=freebsd GOARCH=amd64 go build -o build/FreeBSD-amd64-headphone-spider
GOOS=freebsd GOARCH=386 go build -o build/FreeBSD-386-headphone-spider

# 编译 NetBSD 版本
echo "正在编译 NetBSD 版本..."
GOOS=netbsd GOARCH=amd64 go build -o build/NetBSD-amd64-headphone-spider
GOOS=netbsd GOARCH=386 go build -o build/NetBSD-386-headphone-spider

# 编译 OpenBSD 版本
echo "正在编译 OpenBSD 版本..."
GOOS=openbsd GOARCH=amd64 go build -o build/OpenBSD-amd64-headphone-spider
GOOS=openbsd GOARCH=386 go build -o build/OpenBSD-386-headphone-spider

# 编译 Solaris 版本
echo "正在编译 Solaris 版本..."
GOOS=solaris GOARCH=amd64 go build -o build/Solaris-amd64-headphone-spider

echo "编译完成！所有文件已保存在 build 目录中" 
gh release create ./build/*