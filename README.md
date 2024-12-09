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
   - Windows 64位用户下载 `Windows-amd64-headphone-spider.exe`
   - Windows 32位用户下载 `Windows-386-headphone-spider.exe`
   - macOS 64位用户下载 `macOS-amd64-headphone-spider`
   - Apple Silicon用户下载 `macOS-arm64-headphone-spider`
   - Linux 64位用户下载 `Linux-amd64-headphone-spider`
   - Linux 32位用户下载 `Linux-386-headphone-spider`
   - Linux ARM 64位用户下载 `Linux-arm64-headphone-spider`
   - Linux ARM 32位用户下载 `Linux-arm-headphone-spider`

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

1. 打开终端（Windows 用户可以使��命令提示符或 PowerShell）
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
- `https://squig.link/?share=xxx`
- `https://pw.squig.link/?share=xxx`
- `https://graph.hangout.audio/iem/xxx/?share=xxx`
- `https://earphonesarchive.squig.link/headphones/?share=xxx`
- 以及其他使用相同数据格式的网站

### 使用说明

#### Squig.Link 及相关网站

当使用 Squig.Link 或其相关网站（如 hangout.audio）时：

1. URL 必须包含 `share` 参数
2. 如果 URL 中包含多个耳机型号（以逗号分隔），程序会提示选择要下载的型号
3. 程序会自动过滤掉 "Custom_Tilt" 选项
4. 支持不同的数据格式（制表符分隔或 CSV）

示例：
```bash
# 单个耳机型号
./headphone-spider "https://pw.squig.link/?share=KZ_PRX"

# 多个耳机型号
./headphone-spider "https://earphonesarchive.squig.link/headphones/?share=Custom_Tilt,Apple_AirPods_Max"

# Hangout Audio
./headphone-spider "https://graph.hangout.audio/iem/5128/?share=AirPods_Pro_2_(ANC_on)"
```

注意：
- 对于包含多个耳机型号的 URL，程序会提示用户选择要下载的型号
- 程序会自动将下载的数据保存为 CSV 格式
- 输出文件名将使用选择的耳机型号，空格会被保留

### 注意事项

- 如果程序报错找不到 Chrome，请确保已正确安装 Chrome 或 Chromium 浏览器
- 程序运行时会自动启动 Chrome 浏览器（默认为无头模式，不会显示界面）
- 如果遇到任何问题，请在 [GitHub Issues](https://github.com/kiy7086/headphone-spider/issues) 页面报告

## 许可证

该项目使用 GPL-v3.0 许可证。详细信息请参阅 [LICENSE](LICENSE) 文件。
