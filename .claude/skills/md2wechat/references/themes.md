# 微信公众号排版风格提示词

## 风格选择指南

根据文章内容类型推荐风格：

| 文章类型 | 推荐风格 | 理由 |
|----------|----------|------|
| 技术教程、开发文档 | minimal | 清晰易读，专注内容 |
| 散文、随笔、故事 | elegant | 文艺感，情感表达 |
| 科技资讯、产品评测 | tech | 现代感，视觉冲击 |
| 商业报告、数据分析 | minimal | 专业简洁 |
| 生活方式、旅行 | elegant | 温暖舒适 |

---

## 风格 1: minimal (极简)

### 整体感觉
干净、简洁、专业，白色背景为主，黑色文字，无多余装饰。

### 颜色规范
```css
/* 背景色 */
--bg-primary: #FFFFFF
--bg-secondary: #F9F9F9
--bg-quote: #F5F5F5
--bg-code: #F0F0F0

/* 文字色 */
--text-primary: #333333
--text-secondary: #666666
--text-heading: #000000
--text-link: #007AFF
--text-code: #D14A28

/* 边框/分割线 */
--border-color: #EEEEEE
--border-quote: #DDDDDD
```

### 字体规范
```css
/* 使用系统默认无衬线字体栈 */
font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif

/* 字号 */
--h1-size: 24px
--h2-size: 20px
--h3-size: 18px
--h4-size: 16px
--body-size: 16px
--code-size: 14px

/* 行高/间距 */
--line-height: 1.8
--paragraph-spacing: 20px
--letter-spacing: normal
```

### 布局规范
```css
/* 容器 */
--max-width: 677px (微信标准宽度)
--padding-horizontal: 20px
--padding-vertical: 16px

/* 间距 */
--spacing-sm: 8px
--spacing-md: 16px
--spacing-lg: 24px
--spacing-xl: 32px
```

### 元素样式

#### 标题 (h1-h4)
```css
h1: font-size: 24px, font-weight: bold, color: #000000, margin-bottom: 20px, line-height: 1.4
h2: font-size: 20px, font-weight: bold, color: #000000, margin-top: 28px, margin-bottom: 16px
h3: font-size: 18px, font-weight: bold, color: #000000, margin-top: 20px, margin-bottom: 12px
h4: font-size: 16px, font-weight: bold, color: #333333, margin-top: 16px, margin-bottom: 10px
```

#### 段落 (p)
```css
font-size: 16px, line-height: 1.8, color: #333333, margin-bottom: 20px
```

#### 引用块 (blockquote)
```css
background-color: #F5F5F5
border-left: 4px solid #DDDDDD
padding: 16px
margin: 20px 0
color: #666666
font-style: italic
```

#### 代码块 (pre/code)
```css
pre: background-color: #F0F0F0, padding: 16px, overflow-x: auto, border-radius: 4px, margin: 20px 0
code: font-family: "SF Mono", Monaco, Consolas, monospace, font-size: 14px, color: #333333
inline-code: background-color: #F0F0F0, padding: 2px 6px, border-radius: 3px, font-size: 14px
```

#### 列表 (ul/ol)
```css
padding-left: 24px
margin-bottom: 20px
line-height: 1.8
li: margin-bottom: 8px
```

#### 链接 (a)
```css
color: #007AFF
text-decoration: none
border-bottom: 1px solid transparent
transition: border-color 0.2s
```

#### 分割线 (hr)
```css
border: none
border-top: 1px solid #EEEEEE
margin: 32px 0
```

#### 表格 (table)
```css
width: 100%
border-collapse: collapse
margin: 20px 0
th: background-color: #F9F9F9, padding: 12px, text-align: left, font-weight: bold, border-bottom: 2px solid #EEEEEE
td: padding: 12px, border-bottom: 1px solid #EEEEEE
```

#### 图片 (img)
```css
max-width: 100%
height: auto
display: block
margin: 20px auto
border-radius: 4px
```

---

## 风格 2: elegant (优雅)

### 整体感觉
温暖、文艺、有质感，米白色背景，深灰色文字，衬线字体，营造阅读氛围。

### 颜色规范
```css
/* 背景色 */
--bg-primary: #FAF9F6 (米白/奶油色)
--bg-secondary: #F5F3EF
--bg-quote: #FFF8DC (淡金/玉米丝色)
--bg-code: #EFEBE0

/* 文字色 */
--text-primary: #2C2C2C (深炭灰)
--text-secondary: #6B6B6B
--text-heading: #1A365D (深蓝)
--text-link: #8B4513 (鞍褐)
--text-code: #8B4513

/* 边框/分割线 */
--border-color: #E8E4D9
--border-quote: #D4C4A8
```

### 字体规范
```css
/* 优先使用衬线字体 */
font-family: Georgia, "Times New Roman", "Songti SC", "SimSun", serif

/* 字号 */
--h1-size: 26px
--h2-size: 22px
--h3-size: 19px
--h4-size: 17px
--body-size: 17px
--code-size: 14px

/* 行高/间距 */
--line-height: 2.0 (更宽松的呼吸感)
--paragraph-spacing: 24px
--letter-spacing: 0.02em (轻微字间距)
```

### 元素样式

#### 标题
```css
h1: font-size: 26px, font-weight: normal, color: #1A365D, margin-bottom: 24px, letter-spacing: 0.05em
h2: font-size: 22px, font-weight: normal, color: #1A365D, margin-top: 32px, margin-bottom: 20px, border-bottom: 1px solid #D4C4A8, padding-bottom: 8px
h3: font-size: 19px, font-weight: normal, color: #2C2C2C, margin-top: 24px, margin-bottom: 16px
h4: font-size: 17px, font-weight: normal, color: #6B6B6B, margin-top: 20px, margin-bottom: 12px, font-style: italic
```

#### 段落
```css
font-size: 17px, line-height: 2.0, color: #2C2C2C, margin-bottom: 24px, text-align: justify
```

#### 引用块
```css
background-color: #FFF8DC
border-left: 4px solid #D4C4A8
padding: 20px
margin: 24px 0
color: #6B6B6B
font-style: italic
border-radius: 0 4px 4px 0
```

#### 代码块
```css
pre: background-color: #EFEBE0, padding: 20px, border-radius: 4px, margin: 24px 0, border: 1px solid #E8E4D9
code: font-family: "SF Mono", Monaco, "Courier New", monospace, font-size: 14px, color: #2C2C2C
inline-code: background-color: #EFEBE0, padding: 3px 8px, border-radius: 3px, font-size: 14px, color: #8B4513
```

#### 分割线
```css
border: none
border-top: 1px solid #D4C4A8
margin: 40px 0
height: 1px
background: linear-gradient(to right, transparent, #D4C4A8, transparent)
```

#### 图片
```css
max-width: 100%
height: auto
display: block
margin: 24px auto
border-radius: 8px
box-shadow: 0 2px 8px rgba(0,0,0,0.1)
```

---

## 风格 3: tech (科技)

### 整体感觉
现代、科技感、暗色主题，深蓝黑背景，青色霓虹点缀，适合科技内容。

### 颜色规范
```css
/* 背景色 */
--bg-primary: #0D1117 (GitHub 深色背景)
--bg-secondary: #161B22
--bg-quote: #1F2937
--bg-code: #000000

/* 文字色 */
--text-primary: #E6EDF3
--text-secondary: #8B949E
--text-heading: #00D9FF (青色霓虹)
--text-link: #58A6FF
--text-code: #FF7B72

/* 强调色 */
--accent-primary: #00D9FF (青色)
--accent-secondary: #7EE787 (绿色)
--accent-tertiary: #FFA657 (橙色)
```

### 字体规范
```css
/* 使用现代无衬线字体 */
font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Noto Sans", "Roboto", sans-serif

/* 字号 */
--h1-size: 28px
--h2-size: 22px
--h3-size: 18px
--h4-size: 16px
--body-size: 15px
--code-size: 13px

/* 行高/间距 */
--line-height: 1.6
--paragraph-spacing: 16px
```

### 元素样式

#### 标题
```css
h1: font-size: 28px, font-weight: bold, color: #00D9FF, margin-bottom: 20px, text-shadow: 0 0 10px rgba(0,217,255,0.3)
h2: font-size: 22px, font-weight: bold, color: #00D9FF, margin-top: 24px, margin-bottom: 14px, border-left: 4px solid #00D9FF, padding-left: 12px
h3: font-size: 18px, font-weight: bold, color: #E6EDF3, margin-top: 20px, margin-bottom: 12px
h4: font-size: 16px, font-weight: bold, color: #8B949E, margin-top: 16px, margin-bottom: 10px
```

#### 段落
```css
font-size: 15px, line-height: 1.6, color: #E6EDF3, margin-bottom: 16px
```

#### 引用块
```css
background-color: #16213E
border-left: 4px solid #00D9FF
padding: 16px
margin: 20px 0
color: #8B949E
box-shadow: inset 0 0 20px rgba(0,217,255,0.05)
```

#### 代码块
```css
pre: background-color: #000000, padding: 16px, border-radius: 6px, margin: 20px 0, border: 1px solid #30363D
code: font-family: "SF Mono", "Fira Code", monospace, font-size: 13px, color: #E6EDF3
inline-code: background-color: #161B22, padding: 3px 6px, border-radius: 4px, font-size: 13px, color: #FF7B72, border: 1px solid #30363D
```

#### 列表
```css
padding-left: 24px
margin-bottom: 16px
line-height: 1.6
li: margin-bottom: 6px, color: #E6EDF3
li::marker: color: #00D9FF
```

#### 链接
```css
color: #58A6FF
text-decoration: none
border-bottom: 1px solid #58A6FF
```

#### 分割线
```css
border: none
border-top: 1px solid #30363D
margin: 32px 0
background: linear-gradient(to right, #0D1117, #00D9FF, #0D1117)
height: 2px
```

#### 表格
```css
width: 100%
border-collapse: collapse
margin: 20px 0
border: 1px solid #30363D
th: background-color: #161B22, padding: 12px, text-align: left, font-weight: bold, color: #00D9FF, border-bottom: 2px solid #30363D
td: padding: 12px, border-bottom: 1px solid #21262D, color: #E6EDF3
```

#### 图片
```css
max-width: 100%
height: auto
display: block
margin: 20px auto
border-radius: 6px
border: 1px solid #30363D
```

---

## 生成提示词模板

当用户选择某个风格时，向 Claude 提供以下指令：

### Minimal 风格生成指令

```
使用极简风格生成微信公众号 HTML：
- 白色背景 #FFFFFF，黑色文字 #333333
- 无衬线字体，行高 1.8
- 所有 CSS 必须内联（style 属性）
- 使用安全的 HTML 标签
- 图片使用占位符格式 <!-- IMG:index -->
```

### Elegant 风格生成指令

```
使用优雅风格生成微信公众号 HTML：
- 米白背景 #FAF9F6，深蓝标题 #1A365D
- 衬线字体，行高 2.0
- 所有 CSS 必须内联
- 引用块使用淡金背景 #FFF8DC
- 图片添加阴影效果
```

### Tech 风格生成指令

```
使用科技风格生成微信公众号 HTML：
- 深色背景 #0D1117，青色标题 #00D9FF
- 现代无衬线字体，行高 1.6
- 所有 CSS 必须内联
- 标题添加发光效果
- 代码块使用黑色背景
```
