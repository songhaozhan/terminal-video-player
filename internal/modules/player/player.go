package player

import (
	"fmt"
	"strings"
	"time"

	"terminal-video-player/internal/modules/converter"
	"terminal-video-player/internal/modules/frames"
	"terminal-video-player/internal/modules/terminal"
)

// VideoPlayer encapsulates video playback functionality
type VideoPlayer struct {
	VideoPath string
	Width     int
	Height    int
	FPS       float64
	ColorMode bool // true for color, false for monochrome
}

// New creates a new VideoPlayer with default settings
func New(videoPath string) *VideoPlayer {
	return &VideoPlayer{
		VideoPath: videoPath,
		Width:     80,   // Terminal width
		Height:    24,   // Terminal height
		FPS:       15.0, // Frames per second for ASCII playback
		ColorMode: true, // Default to color mode
	}
}

// SetDimensions sets custom terminal dimensions
func (vp *VideoPlayer) SetDimensions(width, height int) {
	vp.Width = width
	vp.Height = height
}

// SetFPS sets custom playback frame rate
func (vp *VideoPlayer) SetFPS(fps float64) {
	vp.FPS = fps
}

// SetColorMode sets color mode (true for color, false for monochrome)
func (vp *VideoPlayer) SetColorMode(colorMode bool) {
	vp.ColorMode = colorMode
}

// Play starts video playback in terminal with optimized smooth display
func (vp *VideoPlayer) Play() error {
	fmt.Println("正在提取视频帧...")
	if err := frames.ExtractFrames(vp.VideoPath, vp.Width, vp.Height, vp.FPS); err != nil {
		return fmt.Errorf("failed to extract frames: %v", err)
	}

	fmt.Println("开始播放视频 (按 Ctrl+C 停止)...")
	time.Sleep(1 * time.Second)

	// Prepare terminal for smooth playbook
	terminal.ClearScreen()
	terminal.HideCursor()
	defer terminal.ShowCursor() // Ensure cursor is restored on exit

	frameDelay := time.Duration(1000/vp.FPS) * time.Millisecond
	frameNum := 1

	// Pre-allocate buffer with estimated size
	var frameBuffer strings.Builder
	frameBuffer.Grow(vp.Width * vp.Height * 25) // Estimate buffer size

	for {
		// Load frame with timing
		start := time.Now()
		img, err := frames.LoadFrame(frameNum)
		if err != nil {
			// End of video or error
			break
		}

		// Convert to ASCII
		asciiFrame := converter.ImageToASCII(img, vp.Width, vp.Height, vp.ColorMode)
		loadTime := time.Since(start)

		// Build complete output in one go
		frameBuffer.Reset()
		frameBuffer.WriteString(asciiFrame)
		frameBuffer.WriteString(fmt.Sprintf("\033[0m帧: %d | FPS: %.1f | 加载: %v\033[K", frameNum, vp.FPS, loadTime))

		// Atomic output to minimize flicker
		terminal.MoveCursorHome()
		fmt.Print(frameBuffer.String())

		// Precise timing adjustment
		elapsed := time.Since(start)
		if elapsed < frameDelay {
			time.Sleep(frameDelay - elapsed)
		}
		frameNum++
	}

	terminal.MoveCursorHome()
	terminal.ClearScreen()
	fmt.Println("视频播放完毕!")

	// Cleanup temporary frames
	frames.CleanupFrames()
	return nil
}

// GetVideoInfo displays video information
func (vp *VideoPlayer) GetVideoInfo() error {
	info, err := frames.GetVideoInfo(vp.VideoPath)
	if err != nil {
		return err
	}

	fmt.Printf("视频信息:\n%s\n", info)
	return nil
}
