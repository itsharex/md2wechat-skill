---
name: md2wechat
description: Convert Markdown articles to WeChat Official Account formatted HTML with styled CSS and optionally upload to draft box. Use when user wants to convert markdown to WeChat article, publish to WeChat, or format articles for WeChat Official Account.
---

# MD to WeChat

Converts Markdown articles to WeChat Official Account formatted HTML with inline CSS and optionally uploads to draft box.

## Quick Start

```bash
# Show HTML preview
bash scripts/run.sh preview article.md

# Upload to WeChat draft box
bash scripts/run.sh draft article.md
```

## Workflow

Copy this checklist to track progress:

```
Progress:
- [ ] Step 1: Analyze Markdown structure and images
- [ ] Step 2: Confirm style theme with user
- [ ] Step 3: Generate HTML with inline CSS
- [ ] Step 4: Process images (upload to WeChat)
- [ ] Step 5: Replace image URLs in HTML
- [ ] Step 6: Preview or upload to draft
```

---

## Step 1: Analyze Markdown

Read the markdown file and extract:

| Element | How to Extract |
|---------|----------------|
| **Title** | First `# heading` or filename |
| **Author** | Look for `Author:` or `作者:` in frontmatter |
| **Digest** | First paragraph or generate from content (max 120 chars) |
| **Images** | Collect all `![alt](src)` references |
| **Structure** | Headings, lists, code blocks, quotes, tables |

**Image Reference Types**:

| Type | Syntax | Processing |
|------|--------|------------|
| Local | `![alt](./path/image.png)` | Upload to WeChat |
| Online | `![alt](https://example.com/image.png)` | Download then upload |
| AI Generate | `![alt](__generate:prompt__)` | Generate via AI then upload |

---

## Step 2: Confirm Style Theme

Ask user which style to use:

| Style | Description | Best For |
|-------|-------------|----------|
| **minimal** | Clean, white background, black text | Technical tutorials |
| **elegant** | Warm tones, serif fonts | Essays, literature |
| **tech** | Dark theme, neon accents | Tech news, reviews |

Read detailed style prompts from [references/themes.md](references/themes.md)

**Default**: Use `minimal` if user doesn't specify.

---

## Step 3: Generate HTML

Read the selected style prompt from `references/themes.md` and generate HTML with **inline CSS**.

**Important Rules**:

1. All CSS must be **inline** (in `style` attributes)
2. No external stylesheets or scripts
3. Use WeChat-safe HTML tags only
4. Image placeholder format: `<!-- IMG:0 -->`, `<!-- IMG:1 -->`, etc.

**Safe HTML Tags**:
- `<p>`, `<br>`, `<strong>`, `<em>`, `<u>`, `<a>`
- `<h1>` to `<h6>`
- `<ul>`, `<ol>`, `<li>`
- `<blockquote>`, `<pre>`, `<code>`
- `<table>`, `<thead>`, `<tbody>`, `<tr>`, `<th>`, `<td>`
- `<section>` (with inline styles)

**Avoid**:
- `<script>`, `<iframe>`, `<form>`
- External CSS/JS references
- Complex positioning (fixed, absolute)

---

## Step 4: Process Images

For each image reference in order:

### Local Image

```bash
bash scripts/run.sh upload_image "/path/to/image.png"
```

Response:
```json
{"success": true, "wechat_url": "https://mmbiz.qpic.cn/...", "media_id": "xxx"}
```

### Online Image

```bash
bash scripts/run.sh download_and_upload "https://example.com/image.png"
```

### AI Generated Image

```bash
bash scripts/run.sh generate_image "A cute cat sitting on a windowsill"
```

**Note**: Requires `IMAGE_API_KEY` environment variable.

**Image Processing Pipeline**:
1. If AI generation: Call image API → get URL
2. If online: Download image to temp
3. If local: Read file
4. Compress if width > 1920px (configurable)
5. Upload to WeChat material API
6. Return `wechat_url` and `media_id`
7. Store result for HTML replacement

---

## Step 5: Replace Image URLs

Replace placeholders in HTML:

```html
<!-- Before -->
<!-- IMG:0 -->
<!-- IMG:1 -->

<!-- After -->
<img src="https://mmbiz.qpic.cn/..." />
<img src="https://mmbiz.qpic.cn/..." />
```

Use the WeChat URLs returned from image processing.

---

## Step 6: Preview or Upload

Ask user:

1. **Preview only** - Show HTML for review
2. **Upload to draft** - Create WeChat draft article

### Preview Mode

Display HTML in markdown code block for user to copy.

### Upload Mode

Create draft JSON and run:

```bash
bash scripts/run.sh create_draft "/path/to/draft.json"
```

**Draft JSON Format**:
```json
{
  "articles": [
    {
      "title": "Article Title",
      "author": "Author Name",
      "digest": "Article summary (120 chars max)",
      "content": "<!DOCTYPE html><html>...</html>",
      "thumb_media_id": "cover_image_media_id",
      "show_cover_pic": 1,
      "content_source_url": "https://original.url"
    }
  ]
}
```

Response:
```json
{"success": true, "media_id": "draft_media_id", "draft_url": "https://mp.weixin.qq.com/..."}
```

---

## Configuration

Required environment variables (user must set):

| Variable | Description | Required |
|----------|-------------|----------|
| `WECHAT_APPID` | WeChat Official Account AppID | Yes |
| `WECHAT_SECRET` | WeChat API Secret | Yes |
| `IMAGE_API_KEY` | Image generation API key | For AI images |
| `IMAGE_API_BASE` | Image API base URL | For AI images |
| `COMPRESS_IMAGES` | Compress images > 1920px (true/false) | No, default true |
| `MAX_IMAGE_WIDTH` | Max width in pixels | No, default 1920 |

---

## Error Handling

| Error | Action |
|-------|--------|
| Missing config | Ask user to set environment variable |
| Image upload fails | Log error, continue with placeholder |
| WeChat API fails | Show error message, return HTML for manual upload |
| Markdown parse error | Ask user to check file format |

---

## Complete Examples

### Example 1: Simple Article (No Images)

**Input**: `simple.md`
```markdown
# My First Article

This is a simple article with no images.
```

**Process**:
1. Generate HTML with `minimal` style
2. Skip image processing
3. Ask: preview or upload?
4. If upload → create draft

### Example 2: Article with Local Images

**Input**: `with-images.md`
```markdown
# Travel Diary

Day 1 in Paris:

![Eiffel Tower](./photos/eiffel.jpg)
```

**Process**:
1. Analyze: 1 local image
2. Generate HTML with `<!-- IMG:0 -->` placeholder
3. Run: `upload_image "./photos/eiffel.jpg"`
4. Replace placeholder with WeChat URL
5. Preview or upload

### Example 3: Article with AI Generated Images

**Input**: `ai-images.md`
```markdown
# Future Cities

![Futuristic skyline](__generate:A futuristic city skyline at sunset with flying cars and neon lights__)
```

**Process**:
1. Analyze: 1 AI image request
2. Generate HTML with placeholder
3. Run: `generate_image "A futuristic city..."`
4. Replace placeholder
5. Preview or upload

### Example 4: Mixed Image Types

**Input**: `mixed.md`
```markdown
# Tech Review

![Product Photo](./product.jpg)

![Comparison Chart](https://example.com/chart.png)

![Concept Art](__generate:Futuristic gadget design__)
```

**Process**:
1. Process 3 images in order
2. Each returns WeChat URL
3. Replace all placeholders
4. Final HTML with all WeChat-hosted images

---

## References

- [Style Themes](references/themes.md) - Detailed style prompts
- [HTML Guide](references/html-guide.md) - WeChat HTML constraints
- [Image Syntax](references/image-syntax.md) - Image reference syntax
- [WeChat API](references/wechat-api.md) - API reference

---

## Troubleshooting

**Q: Image upload fails with "invalid filetype"**
A: WeChat supports JPG, PNG, GIF. Ensure image is in correct format.

**Q: Draft created but images not showing**
A: Images must use WeChat-hosted URLs (`mmbiz.qpic.cn`), not external URLs.

**Q: "AppID not configured" error**
A: Set `WECHAT_APPID` and `WECHAT_SECRET` environment variables.

**Q: AI image generation fails**
A: Check `IMAGE_API_KEY` is set and API base URL is correct.
