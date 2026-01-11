package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config 应用配置
type Config struct {
	// 微信公众号配置
	WechatAppID   string
	WechatSecret string

	// 图片生成 API 配置
	ImageAPIKey  string
	ImageAPIBase string

	// 图片处理配置
	CompressImages bool
	MaxImageWidth  int
	MaxImageSize   int64 // bytes

	// 超时配置
	HTTPTimeout int // seconds
}

// Load 从环境变量加载配置
func Load() (*Config, error) {
	cfg := &Config{
		WechatAppID:    os.Getenv("WECHAT_APPID"),
		WechatSecret:  os.Getenv("WECHAT_SECRET"),
		ImageAPIKey:   os.Getenv("IMAGE_API_KEY"),
		ImageAPIBase:  os.Getenv("IMAGE_API_BASE"),
		CompressImages: getEnvBool("COMPRESS_IMAGES", true),
		MaxImageWidth:  getEnvInt("MAX_IMAGE_WIDTH", 1920),
		MaxImageSize:   int64(getEnvInt("MAX_IMAGE_SIZE", 5*1024*1024)), // 5MB
		HTTPTimeout:    getEnvInt("HTTP_TIMEOUT", 30),
	}

	// 验证必需配置
	if cfg.WechatAppID == "" {
		return nil, fmt.Errorf("WECHAT_APPID is required")
	}
	if cfg.WechatSecret == "" {
		return nil, fmt.Errorf("WECHAT_SECRET is required")
	}

	// 设置默认值
	if cfg.ImageAPIBase == "" {
		cfg.ImageAPIBase = "https://api.openai.com/v1"
	}

	return cfg, nil
}

// ValidateForImageGeneration 验证图片生成配置
func (c *Config) ValidateForImageGeneration() error {
	if c.ImageAPIKey == "" {
		return fmt.Errorf("IMAGE_API_KEY is required for image generation")
	}
	return nil
}

// getEnvBool 获取布尔型环境变量
func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val == "true" || val == "1"
}

// getEnvInt 获取整型环境变量
func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}
