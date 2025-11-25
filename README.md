# 🐍 High-Performance Snake in Go

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-gray)](https://github.com/XPLassal/simple-snake-on-go/releases)

A modern, hyper-optimized implementation of the classic Snake game that runs directly in your terminal. Written in **Pure Go** with a focus on **Clean Architecture** and **O(1) Algorithms**.

![Gameplay Preview](preview.gif)

---

## ⚡ Key Features

* **🚀 True O(1) Performance:** The game engine relies on a custom **Linked List via Map** data structure. Movement and collision checks are instant, regardless of whether the snake has 10 or 10,000 segments.
* **🎨 Zero-Allocation Rendering:** (v2.1) The rendering engine uses a pre-allocated buffer and `bufio.Writer` to eliminate GC pressure and system call overhead.
* **📺 Flicker-Free:** Uses ANSI cursor management (`\033[H`) instead of screen clearing for smooth 60 FPS visuals without flashing.
* **💻 Cross-Platform:** Runs natively on **Windows**, **Linux**, and **macOS** (Intel & Apple Silicon).
* **⚙️ Dynamic Gameplay:**
    * Customizable map size.
    * **"Hard Mode"**: Game speed increases automatically as you score points.

---

## 🎮 How to Play

### Option 1: Download Binary (Recommended)
You don't need Go installed. Just grab the executable for your OS from the [**Releases Page**](https://github.com/XPLassal/simple-snake-on-go/releases/latest).

| OS | File |
| :--- | :--- |
| 🪟 **Windows** | `snake-windows-amd64.exe` |
| 🐧 **Linux** | `snake-linux-amd64` |
| 🍎 **macOS (M1/M2)** | `snake-macos-arm64` |
| 🍎 **macOS (Intel)** | `snake-macos-intel` |
## 🕹 Controls
| Key | Action |
| :---: | :--- |
| **W** | Move Up ⬆️ |
| **S** | Move Down ⬇️ |
| **A** | Move Left ⬅️ |
| **D** | Move Right ➡️ |
| **Q** | Quit Game |

**Linux/macOS Note:**
If the file doesn't run, give it permission:
```bash
chmod +x snake-linux-amd64
./snake-linux-amd64
```
