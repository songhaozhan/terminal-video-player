# 终端视频播放器

🎬 一个高性能的终端视频播放器，使用先进的象限块渲染技术将视频转换为超高分辨率的彩色ASCII艺术。

[English](README.md) | **中文**

## ✨ 功能特性

- **超高分辨率ASCII艺术**: 使用2x2像素采样和16个Unicode象限块字符，实现最大细节呈现
- **真彩色支持**: 24位ANSI颜色渲染，呈现生动的视频显示效果
- **性能优化**: 预分配缓冲区、原子输出和精确时序控制，确保流畅播放
- **模块化架构**: 遵循Go最佳实践的清晰内部包结构
- **跨平台**: 支持macOS、Linux和Windows，需要ffmpeg支持
- **自定义输出**: 可调节终端尺寸和帧率
- **实时处理**: 实时帧提取和转换，延迟最小

## 🎥 演示

### 原始视频与终端输出对比

| 原始视频 | 彩色模式 | 黑白模式 |
|---------|---------|---------|
| ![原始视频](video/sample.gif) | ![终端视频播放器演示](video/result.gif) | ![黑白模式演示](video/mono.gif) |

*左：原始视频输入 | 中：彩色ASCII艺术输出 | 右：黑白ASCII艺术输出*

## 🚀 快速开始

### 环境要求

- **Go 1.24+** - [下载Go](https://golang.org/dl/)
- **ffmpeg** - 视频帧提取必需

### 安装ffmpeg:
```bash
# macOS
brew install ffmpeg

# Ubuntu/Debian
sudo apt install ffmpeg

# CentOS
sudo yum install ffmpeg

# Windows
# 从 https://ffmpeg.org/download.html 下载并添加到PATH
```

### 安装

```bash
# 克隆仓库
git clone https://github.com/songhaozhan/terminal-video-player.git
cd terminal-video-player

# 构建播放器
go build -o video_player main.go
```

### 使用方法

```bash
# 基本用法
go run main.go video/sample.mp4

# 自定义尺寸和帧率
go run main.go video/sample.mp4 120 30 24

# 或使用构建的二进制文件
./video_player video/sample.mp4
```

## 📋 命令行选项

```
go run main.go <视频文件> [宽度] [高度] [帧率] [模式]
```

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `视频文件` | 视频文件路径（必需） | - |
| `宽度` | 终端宽度（字符数） | 80 |
| `高度` | 终端高度（字符数） | 24 |
| `帧率` | 播放帧率 | 15.0 |
| `模式` | 显示模式："color"(彩色)、"mono"/"black"/"bw"(黑白) | color |

### 使用示例

```bash
# 使用默认设置播放视频
go run main.go video/sample.mp4

# 为大型终端自定义分辨率
go run main.go video/sample.mp4 160 40

# 高帧率播放
go run main.go video/sample.mp4 100 30 30

# 宽屏格式（彩色模式）
go run main.go video/sample.mp4 200 25 20 color

# 黑白模式播放
go run main.go video/sample.mp4 120 30 24 mono
```

## 🏗️ 架构设计

项目采用模块化的内部包结构：

```
├── internal/
│   └── modules/
│       ├── converter/     # ASCII艺术转换算法
│       ├── frames/        # 视频帧提取和加载
│       ├── player/        # 核心播放引擎
│       └── terminal/      # 终端控制功能
├── video/                 # 示例视频资源
├── main.go               # 应用程序入口点
└── go.mod               # Go模块定义
```

### 核心组件

- **帧提取**: 使用ffmpeg将视频帧提取为PNG图像
- **ASCII转换**: 高级象限块字符映射与亮度计算
- **播放引擎**: 优化的渲染，包含帧时序和终端控制
- **终端控制**: 用于流畅显示的ANSI转义序列

## 🎨 技术细节

### 象限块渲染

ASCII转换使用复杂的2x2像素采样技术，配合16个Unicode象限块字符：

```
▘ ▝ ▀ ▖ ▌ ▞ ▛ ▗ ▚ ▐ ▜ ▄ ▙ ▟ █
```

每个字符代表亮/暗象限的特定模式，相比传统ASCII艺术分辨率提高4倍。

### 颜色处理

- **亮度计算**: 使用ITU-R BT.601公式 (0.299R + 0.587G + 0.114B)
- **自适应阈值**: 动态亮度阈值以获得最佳对比度
- **真彩色输出**: 24位RGB ANSI转义序列实现准确色彩重现

### 性能优化

- **预分配缓冲区**: 播放期间最小化内存分配
- **原子输出**: 单次写入渲染减少终端闪烁
- **精确时序**: 帧精确播放和时序补偿
- **高效缩放**: 基于FFmpeg的帧预处理

## 🛠️ 开发

### 从源码构建

```bash
# 克隆并构建
git clone https://github.com/songhaozhan/terminal-video-player.git
cd terminal-video-player
go mod tidy
go build -o video_player main.go
```

### 运行测试

```bash
go test ./...
```

### 代码格式化

```bash
go fmt ./...
```

## 📺 支持格式

通过ffmpeg集成，播放器支持多种视频格式：

- **MP4** (H.264, H.265)
- **AVI**
- **MOV**
- **MKV**
- **WebM**
- **FLV**
- 以及更多...

## ⚡ 性能建议

1. **终端尺寸**: 更大的尺寸需要更多处理能力
2. **帧率**: 更高的FPS会增加CPU使用率
3. **视频分辨率**: 考虑对极高分辨率视频进行降采样
4. **终端**: 使用硬件加速终端以获得最佳性能
5. **SSD存储**: 更快的存储可提高帧加载速度

## ⚠️ 重要提示

- **终端窗口大小**: 如果播放过程中出现闪烁，请尝试放大终端窗口。更大的终端窗口能提供更好的渲染稳定性，减少视觉伪影。
- **终端兼容性**: 为获得最佳效果，请使用支持24位颜色的现代终端模拟器（iTerm2、Windows Terminal、现代版GNOME Terminal等）

## 🤝 贡献

欢迎贡献！请随时提交Pull Request。对于重大更改，请先开启issue讨论您想要更改的内容。

### 开发设置

1. Fork仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m '添加某个惊人功能'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启Pull Request

## 📄 许可证

本项目采用MIT许可证 - 查看 [LICENSE](LICENSE) 文件了解详细信息。

## 🙏 致谢

- [ffmpeg](https://ffmpeg.org/) - 强大的多媒体框架
- Unicode联盟提供的象限块字符
- Go社区提供的优秀工具和库

## 📞 支持

如果您遇到任何问题或有疑问：

1. 查看 [Issues](https://github.com/songhz/ai_tool2/issues) 页面
2. 创建新issue并提供详细信息
3. 包含您的操作系统、Go版本和ffmpeg版本

---

**用 ❤️ 和 Go 制作**