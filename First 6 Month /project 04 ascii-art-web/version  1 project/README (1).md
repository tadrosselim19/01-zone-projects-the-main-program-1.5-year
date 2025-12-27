# 🎨 ASCII Art Encoding Web App

**Author:** Tadros Selim  
**Academy Project:** 01 Academy – ASCII Art Generator  

A web application written in **Go** that converts text into ASCII art, supports multiple banner styles, allows coloring of full text or specific substrings, previews results in the browser, and enables downloading the output as a `.txt` file.

---

## 📁 Project Structure

```
.
├── banners/               # ASCII banner templates (.txt)
│   ├── standard.txt
│   ├── thinkertoy.txt
│   └── shadow.txt
├── static/                # Static assets
│   ├── styles.css
│   └── script.js
├── templates/             # HTML templates
│   ├── layout.html        # Main UI
│   └── failure.html       # Error page
├── main.go                # HTTP server & routing
├── ascii-art.go           # ASCII generation logic
├── go.mod                 # Go module
└── Dockerfile             # Docker configuration
```

---

## 🚀 Features

- Generate ASCII art from text
- Three banner styles:
  - `standard`
  - `thinkertoy`
  - `shadow`
- Color the whole text or a specific substring
- Live preview in the browser
- Download result as `.txt`
- Horizontal scrolling for wide ASCII output

---

## 🖥 How It Works

1. User enters text and selects:
   - Banner style
   - Color
   - Optional substring to color
2. Request is sent to `/submit`
3. Server:
   - Loads banner file from `banners/`
   - Generates ASCII art
   - Wraps colored parts using `<span>`
4. Result is rendered safely using `template.HTML`
5. User can download the output via `/downlood_as_txt`

---

## 🛠 Installation

### Requirements
- Go 1.20+
- (Optional) Docker

### Run Locally

```bash
git clone https://github.com/your-username/ascii-art-encoding.git
cd ascii-art-encoding
go run main.go ascii-art.go
```

Open your browser at:
```
http://localhost:8080
```

---

## 🐳 Run with Docker

```bash
docker build -t ascii-art-app .
docker run -p 8080:8080 ascii-art-app
```

---

## ⚙ Usage

1. Enter text
2. Choose banner style
3. Select color
4. (Optional) Enter substring to color
5. Click **Generate**
6. Preview ASCII art
7. Download as `.txt`

---

## 💡 Notes

- `.txt` downloads do **not** preserve colors (terminal colors are not supported in plain text files)
- Browser coloring is done using CSS
- ASCII output uses monospace fonts for alignment

---

## 📜 License

This project is part of **01 Academy** training.

MIT License.
