package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/geekjourney/md2wechat/internal/config"
	"github.com/geekjourney/md2wechat/internal/draft"
	"github.com/geekjourney/md2wechat/internal/image"
	"go.uber.org/zap"
)

var (
	cfg *config.Config
	log *zap.Logger
)

func init() {
	var err error
	cfg, err = config.Load()
	if err != nil {
		responseError(err)
		os.Exit(1)
	}

	log, err = zap.NewProduction()
	if err != nil {
		responseError(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "upload_image":
		cmdUploadImage(args)
	case "download_and_upload":
		cmdDownloadAndUpload(args)
	case "generate_image":
		cmdGenerateImage(args)
	case "create_draft":
		cmdCreateDraft(args)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`md2wechat - Markdown to WeChat Official Account converter

Commands:
  upload_image <file_path>       Upload local image to WeChat material library
  download_and_upload <url>      Download online image and upload to WeChat
  generate_image <prompt>        Generate image via AI and upload to WeChat
  create_draft <json_file>       Create WeChat draft article from JSON file

Environment Variables:
  WECHAT_APPID                   WeChat Official Account AppID (required)
  WECHAT_SECRET                  WeChat API Secret (required)
  IMAGE_API_KEY                  Image generation API key (for AI images)
  IMAGE_API_BASE                 Image API base URL (default: https://api.openai.com/v1)
  COMPRESS_IMAGES                Compress images > 1920px (default: true)
  MAX_IMAGE_WIDTH                Max image width in pixels (default: 1920)

Examples:
  # Upload local image
  md2wechat upload_image ./photo.jpg

  # Download and upload online image
  md2wechat download_and_upload https://example.com/image.jpg

  # Generate AI image
  md2wechat generate_image "A cute cat"

  # Create draft
  md2wechat create_draft draft.json`)
}

func cmdUploadImage(args []string) {
	if len(args) < 1 {
		responseError(fmt.Errorf("file_path is required"))
		return
	}

	filePath := args[0]
	processor := image.NewProcessor(cfg, log)
	result, err := processor.UploadLocalImage(filePath)
	if err != nil {
		responseError(err)
		return
	}

	responseSuccess(result)
}

func cmdDownloadAndUpload(args []string) {
	if len(args) < 1 {
		responseError(fmt.Errorf("url is required"))
		return
	}

	url := args[0]
	processor := image.NewProcessor(cfg, log)
	result, err := processor.DownloadAndUpload(url)
	if err != nil {
		responseError(err)
		return
	}

	responseSuccess(result)
}

func cmdGenerateImage(args []string) {
	if len(args) < 1 {
		responseError(fmt.Errorf("prompt is required"))
		return
	}

	prompt := args[0]
	processor := image.NewProcessor(cfg, log)
	result, err := processor.GenerateAndUpload(prompt)
	if err != nil {
		responseError(err)
		return
	}

	responseSuccess(result)
}

func cmdCreateDraft(args []string) {
	if len(args) < 1 {
		responseError(fmt.Errorf("json_file is required"))
		return
	}

	jsonFile := args[0]
	svc := draft.NewService(cfg, log)
	result, err := svc.CreateDraftFromFile(jsonFile)
	if err != nil {
		responseError(err)
		return
	}

	responseSuccess(result)
}

func responseSuccess(data interface{}) {
	response := map[string]interface{}{
		"success": true,
		"data":    data,
	}
	printJSON(response)
}

func responseError(err error) {
	response := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}
	printJSON(response)
	os.Exit(1)
}

func printJSON(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "JSON encode error: %v\n", err)
		os.Exit(1)
	}
}
