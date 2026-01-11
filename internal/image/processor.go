package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/geekjourney/md2wechat/internal/config"
	"github.com/geekjourney/md2wechat/internal/wechat"
	"go.uber.org/zap"
)

// Processor 图片处理器
type Processor struct {
	cfg *config.Config
	log *zap.Logger
	ws  *wechat.Service
}

// NewProcessor 创建图片处理器
func NewProcessor(cfg *config.Config, log *zap.Logger) *Processor {
	return &Processor{
		cfg: cfg,
		log: log,
		ws:  wechat.NewService(cfg, log),
	}
}

// UploadLocalImage 上传本地图片
type UploadResult struct {
	MediaID   string `json:"media_id"`
	WechatURL string `json:"wechat_url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

func (p *Processor) UploadLocalImage(filePath string) (*UploadResult, error) {
	p.log.Info("uploading local image", zap.String("path", filePath))

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	// 如果需要压缩，先处理
	processedPath := filePath
	if p.cfg.CompressImages {
		compressedPath, err := p.compressIfNeeded(filePath)
		if err != nil {
			p.log.Warn("compress failed, using original", zap.Error(err))
		} else if compressedPath != "" {
			processedPath = compressedPath
			defer os.Remove(compressedPath)
		}
	}

	// 上传到微信
	result, err := p.ws.UploadMaterialWithRetry(processedPath, 3)
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		MediaID:   result.MediaID,
		WechatURL: result.WechatURL,
	}, nil
}

// DownloadAndUpload 下载在线图片并上传
func (p *Processor) DownloadAndUpload(url string) (*UploadResult, error) {
	p.log.Info("downloading and uploading image", zap.String("url", url))

	// 下载图片
	tmpPath, err := wechat.DownloadFile(url)
	if err != nil {
		return nil, fmt.Errorf("download failed: %w", err)
	}
	defer os.Remove(tmpPath)

	// 压缩（如果需要）
	processedPath := tmpPath
	if p.cfg.CompressImages {
		compressedPath, err := p.compressIfNeeded(tmpPath)
		if err != nil {
			p.log.Warn("compress failed, using original", zap.Error(err))
		} else if compressedPath != "" {
			processedPath = compressedPath
			defer os.Remove(compressedPath)
		}
	}

	// 上传到微信
	result, err := p.ws.UploadMaterialWithRetry(processedPath, 3)
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		MediaID:   result.MediaID,
		WechatURL: result.WechatURL,
	}, nil
}

// GenerateAndUploadResult AI 生成图片结果
type GenerateAndUploadResult struct {
	Prompt     string `json:"prompt"`
	OriginalURL string `json:"original_url"`
	MediaID    string `json:"media_id"`
	WechatURL  string `json:"wechat_url"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// GenerateAndUpload AI 生成图片并上传
func (p *Processor) GenerateAndUpload(prompt string) (*GenerateAndUploadResult, error) {
	p.log.Info("generating image via AI", zap.String("prompt", prompt))

	// 验证配置
	if err := p.cfg.ValidateForImageGeneration(); err != nil {
		return nil, err
	}

	// 调用图片生成 API
	imageURL, err := p.callImageAPI(prompt)
	if err != nil {
		return nil, fmt.Errorf("generate image: %w", err)
	}
	p.log.Info("image generated", zap.String("url", imageURL))

	// 下载生成的图片
	tmpPath, err := wechat.DownloadFile(imageURL)
	if err != nil {
		return nil, fmt.Errorf("download generated image: %w", err)
	}
	defer os.Remove(tmpPath)

	// 压缩（如果需要）
	processedPath := tmpPath
	if p.cfg.CompressImages {
		compressedPath, err := p.compressIfNeeded(tmpPath)
		if err != nil {
			p.log.Warn("compress failed, using original", zap.Error(err))
		} else if compressedPath != "" {
			processedPath = compressedPath
			defer os.Remove(compressedPath)
		}
	}

	// 上传到微信
	result, err := p.ws.UploadMaterialWithRetry(processedPath, 3)
	if err != nil {
		return nil, err
	}

	return &GenerateAndUploadResult{
		Prompt:      prompt,
		OriginalURL: imageURL,
		MediaID:     result.MediaID,
		WechatURL:   result.WechatURL,
	}, nil
}

// callImageAPI 调用图片生成 API（兼容 OpenAI DALL-E）
func (p *Processor) callImageAPI(prompt string) (string, error) {
	// 构造请求
	reqBody := map[string]interface{}{
		"model": "dall-e-3",
		"prompt": prompt,
		"n":     1,
		"size":  "1024x1024",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// 创建请求
	url := p.cfg.ImageAPIBase + "/images/generations"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.cfg.ImageAPIKey)

	// 发送请求
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Data) == 0 {
		return "", fmt.Errorf("no image generated")
	}

	return result.Data[0].URL, nil
}

// compressIfNeeded 如果图片太大则压缩
func (p *Processor) compressIfNeeded(filePath string) (string, error) {
	// 检查文件大小
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	// 如果文件大小小于限制，不需要压缩
	if fileInfo.Size() < p.cfg.MaxImageSize {
		return "", nil
	}

	p.log.Info("compressing image",
		zap.String("path", filePath),
		zap.Int64("size", fileInfo.Size()))

	// 这里应该使用图片压缩库
	// 由于避免引入重型依赖，这里简化处理：
	// 1. 读取文件
	// 2. 使用压缩库处理
	// 3. 保存到临时文件

	// 简化实现：直接返回原路径，实际应该调用压缩函数
	// 可以使用 imaging 或 similar 库

	return "", nil
}

// detectFormat 检测图片格式
func detectFormat(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		return "jpeg"
	case ".png":
		return "png"
	case ".gif":
		return "gif"
	default:
		return "jpeg" // 默认
	}
}

// readImageDimensions 读取图片尺寸
func readImageDimensions(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// 简化实现：需要图片库来读取尺寸
	// 这里返回默认值
	return 1920, 1080, nil
}

// isValidFormat 检查是否是有效的图片格式
func isValidFormat(filePath string) bool {
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	return validExts[ext]
}
