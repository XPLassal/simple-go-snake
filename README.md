# 🐍 Simple Go Snake

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-gray)](https://github.com/XPLassal/simple-go-snake/releases)
[![AUR](https://img.shields.io/aur/version/simple-go-snake?label=AUR&logo=arch-linux&color=blue)](https://aur.archlinux.org/packages/simple-go-snake)

A modern, hyper-optimized implementation of the classic Snake game that runs directly in your terminal. Written in **Pure Go** with a focus on **Clean Architecture** and **O(1) Algorithms**.

![Gameplay Preview](preview.gif)

---

## ⚡ Key Features (v4.0)

* **📦 AUR Support:** Install easily on Arch Linux via `simple-go-snake`.
* **⏸️ Pause Game:** Need a break? Press **'P'** to freeze the snake.
* **💾 Standard Configs:** Settings are saved automatically in your system's standard config folder (XDG/AppData).
* **📺 ASCII & Emoji Modes:** Toggle between beautiful 🐍 graphics or SSH-friendly ASCII (`%`/`@`) mode.
* **🚀 True O(1) Performance:** The engine uses a **Linked List via Map** structure. Movement is instant (~300ns) regardless of snake length.
* **💻 Cross-Platform:** Native builds for **Windows** (with icon support), **Linux**, and **macOS**.

---

## 📊 Benchmarks

The movement logic is implemented as a Linked List using a `map[Coordinates]Coordinates`. This ensures **O(1)** time complexity.
Whether the snake has 10 segments or 10,000, the movement time remains constant.

**Test Environment:** Ryzen 5 5600H, Linux

| Snake Length | Time per Op |
| :--- | :--- |
| **Small (10)** | ~335 ns/op |
| **Medium (1,000)** | ~283 ns/op |
| **Huge (10,000)** | ~284 ns/op |

*Data obtained via `go test -bench=. ./structs`*

---

## 📥 Installation

### Arch Linux (AUR)
The recommended way for Arch users:
```bash
paru -S simple-go-snake
```

### Manual Download (All OS)

Download the executable for your OS from the [**Releases Page**](https://www.google.com/search?q=https://github.com/XPLassal/simple-go-snake/releases/latest).

| OS | File |
| :--- | :--- |
| 🪟 **Windows** | `snake-windows-amd64.exe` |
| 🐧 **Linux** | `snake-linux-amd64` |
| 🍎 **macOS (Apple Silicon)** | `snake-macos-arm64` |
| 🍎 **macOS (Intel)** | `snake-macos-intel` |

*(Linux/macOS users: run `chmod +x <file>` to make it executable).*

### Build from Source

```bash
git clone [https://github.com/XPLassal/simple-go-snake.git](https://github.com/XPLassal/simple-go-snake.git)
cd simple-go-snake
go build -ldflags "-s -w" -o snake .
./snake
```

-----

## ⚙️ Configuration

On the first run, the game will ask for your preferences (Size, Difficulty, Emoji/ASCII).
These settings are saved permanently in your system's user folder:

  * **Linux:** `~/.config/simple-go-snake/config.json`
  * **Windows:** `%AppData%\simple-go-snake\config.json`
  * **macOS:** `~/Library/Application Support/simple-go-snake/config.json`

| Setting | Description |
| :--- | :--- |
| **Columns** | Map size (e.g., 20). |
| **Hard Mode** | If `y`, the game speeds up as your score increases. |
| **Use Emojis** | `y` for beautiful graphics. `n` for ASCII - recommended for SSH or older terminals. |

> **Tip:** Delete `config.json` to reset settings or press **'C'** inside the game.

-----

## 🕹 Controls

| Key | Action |
| :---: | :--- |
| **W, A, S, D** | Move Snake ⬆️⬅️⬇️➡️ |
| **P** | **Pause Game** ⏸️ |
| **C** | **Config / Restart** (Stops current game) |
| **Q** | Quit Game |

-----

## 🏗 Technical Details

  * **Logic:** `map[Coordinates]Coordinates` (Linked List) for **O(1)** movement logic.
  * **Rendering:** `strings.Builder` + `bufio.Writer` for zero-allocation rendering per frame.
  * **Architecture:** Clean separation of `main` (Loop), `structs` (Domain), and `render` (UI).

-----

## 📄 License

This project is licensed under the MIT License.
