package frames

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/exec"
)

// ExtractFrames extracts video frames using ffmpeg
func ExtractFrames(videoPath string, width, height int, fps float64) error {
	tempDir := "/tmp/video_frames"
	os.RemoveAll(tempDir)
	os.MkdirAll(tempDir, 0755)

	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-vf", fmt.Sprintf("fps=%f,scale=%d:%d", fps, width*2, height*2),
		"-y",
		tempDir+"/frame_%04d.png")

	return cmd.Run()
}

// LoadFrame loads and returns a frame as image.Image
func LoadFrame(frameNum int) (image.Image, error) {
	framePath := fmt.Sprintf("/tmp/video_frames/frame_%04d.png", frameNum)

	file, err := os.Open(framePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// CleanupFrames removes temporary frame files
func CleanupFrames() {
	os.RemoveAll("/tmp/video_frames")
}

// GetVideoInfo gets video information using ffprobe
func GetVideoInfo(videoPath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", videoPath)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("无法获取视频信息: %v", err)
	}

	return string(output), nil
}
