# 🛠️ System Tools in Go

A collection of **system tools** written in **Go**, designed for **Linux** 🐧 and tailored for use in **i3blocks**. These tools provide real-time system statistics with minimal overhead.

## 🚀 Features
- **Filesystem Info**: Displays filesystem usage 📂
- **CPU Usage**: Monitors CPU usage 💻
- **System Time**: Shows formatted system time 🕒

## 📁 Project Structure
```
📦 project-root
├── README.md         # Project documentation
├── go.mod            # Go module file
├── dist/             # Compiled binaries
└── src/              # Source code
```

## ⚙️ Usage
1. **Clone the repo**:
   ```bash
   git clone [https://github.com/jean0t/your-repo.git](https://github.com/jean0t/system_tools.git)
   cd system_tools
   ```

2. **Run the tools**:
   - For **CPU usage**:
     ```bash
     ./dist/cpu_usage
     ```
   - For **Filesystem info**:
     ```bash
     ./dist/filesystem_info
     ```
   - For **System time**:
     ```bash
     ./dist/time_pt
     ```

## 🖥️ i3blocks Integration
Add the following to your **i3blocks** configuration to display the system statistics:

```ini
[cpu_usage]
command=/path/to/cpu_usage
interval=1

[filesystem_info]
command=/path/to/filesystem_info
interval=60

[time_pt]
command=/path/to/time_pt
interval=5
```

## 🔜 Coming Soon
More system tools are on the way! Stay tuned for additional utilities to enhance your **Linux** experience.

## 📃 License
This project is licensed under the **GPL-3 License**. See the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.html)) file for details.

---

Made with ❤️ by [Jean0t](https://github.com/jean0t)
