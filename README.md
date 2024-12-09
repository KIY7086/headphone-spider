# Headphone Spider

Headphone Spider 是一个用于从线上耳机评测网站抓取耳机评测数据的工具。它支持从 [huihifi.com](https://huihifi.com) 和 [rtings.com](https://rtings.com) 抓取数据，并将结果保存为 CSV 文件。

## 前置要求

本工具依赖于 Chrome 或 Chromium 浏览器进行数据抓取。请确保你的系统上已安装以下任一浏览器：

- [Google Chrome](https://www.google.cn/chrome/)
- [Chromium](https://www.chromium.org/getting-involved/download-chromium)

### 各系统安装 Chrome 的方法

#### Windows
1. 访问 [Chrome 官网](https://www.google.cn/chrome/)
2. 下载并运行安装程序

#### macOS
1. 访问 [Chrome 官网](https://www.google.cn/chrome/)
2. 下载并安装 .dmg 文件

#### Linux
Ubuntu/Debian:
```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

CentOS/RHEL:
```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
sudo dnf localinstall google-chrome-stable_current_x86_64.rpm
```

Arch Linux:（需要先安装[paru](https://github.com/Morganamilo/paru)）
```bash
paru -S google-chrome
```

## 下载和使用

### 安装方法

#### 方法一：直接下载编译好的文件（推荐）

1. 访问 [GitHub Release 页面](https://github.com/kiy7086/headphone-spider/releases)
2. 找到最新的 Release 版本
3. 根据你的操作系统下载对应的编译文件：
   - Windows 用户下载 `Windows-x86_64-headphone-spider.exe`
   - Windows 32位用户下载 `Windows-x86-headphone-spider.exe`
   - macOS 用户下载 `macOS-x86_64-headphone-spider`
   - Apple Silicon用户下载 `macOS-arm64-headphone-spider`
   - Linux 用户下载 `Linux-x86_64-headphone-spider`
   - Linux Arm（Android Termux或树莓派等）用户下载 `Linux-arm64-headphone-spider`

#### 方法二：从源码编译

1. 克隆仓库
```bash
git clone https://github.com/kiy7086/headphone-spider.git
cd headphone-spider
```

2. 安装依赖
```bash
go mod download
```

3. 编译（需要先安装[Go](https://golang.org/dl/)）
```bash
go build
```

编译完成后会在当前目录生成可执行文件：
- Windows系统下生成 `headphone-spider.exe`
- macOS/Linux系统下生成 `headphone-spider`

### 使用方法

1. 打开终端（Windows 用户可以使用命令提示符或 PowerShell）
2. 导航到下载的文件所在的目录
3. 运行以下命令来抓取数据：

   ```bash
   ./headphone-spider <URL>
   ```

   例如：

   ```bash
   ./headphone-spider https://huihifi.com/evaluation/5e14542b-be71-49e8-add2-d6177bf900dc
   ```

### 支持的 URL 类型

- `https://huihifi.com/evaluation/xxx`
- `https://rtings.com/headphones/reviews/xxx`

### 注意事项

- 如果程序报错找不到 Chrome，请确保已正确安装 Chrome 或 Chromium 浏览器
- 程序运行时会自动启动 Chrome 浏览器（默认为无头模式，不会显示界面）
- 如果遇到任何问题，请在 [GitHub Issues](https://github.com/kiy7086/headphone-spider/issues) 页面报告

## 许可证

该项目使用 GPL-v3.0 许可证。详细信息请参阅 [LICENSE](LICENSE) 文件。
