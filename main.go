package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"terminal-video-player/internal/modules/player"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run main.go <视频文件路径>")
		fmt.Println("示例: go run main.go video/aaa.mp4")
		return
	}

	videoPath := os.Args[1]

	// Check if file exists
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		fmt.Printf("视频文件不存在: %s\n", videoPath)
		return
	}

	// Check if ffmpeg is available
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		fmt.Println("错误: 需要安装 ffmpeg")
		fmt.Println("安装方法:")
		fmt.Println("  macOS: brew install ffmpeg")
		fmt.Println("  Ubuntu: sudo apt install ffmpeg")
		fmt.Println("  Windows: 下载 ffmpeg 并添加到 PATH")
		return
	}

	videoPlayer := player.New(videoPath)

	// Allow user to set custom dimensions
	if len(os.Args) >= 3 {
		if width, err := strconv.Atoi(os.Args[2]); err == nil {
			if len(os.Args) >= 4 {
				if height, err := strconv.Atoi(os.Args[3]); err == nil {
					videoPlayer.SetDimensions(width, height)
				}
			}
		}
	}

	if len(os.Args) >= 5 {
		if fps, err := strconv.ParseFloat(os.Args[4], 64); err == nil {
			videoPlayer.SetFPS(fps)
		}
	}

	fmt.Printf("终端视频播放器\n")
	fmt.Printf("视频文件: %s\n", videoPath)
	fmt.Printf("输出尺寸: %dx%d\n", videoPlayer.Width, videoPlayer.Height)
	fmt.Printf("播放帧率: %.1f FPS\n", videoPlayer.FPS)
	fmt.Println()

	if err := videoPlayer.Play(); err != nil {
		fmt.Printf("播放错误: %v\n", err)
	}
}
