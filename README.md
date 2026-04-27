# ASCII Art Generator (Go)

## Overview

+ This project is a command-line ASCII Art generator written in Go.
It takes a string as input and displays it in a graphical representation using ASCII characters.

+ The program supports:

+ Letters, numbers, and symbols
Multi-line input (\n)
Multiple banner styles (standard, shadow, thinkertoy)
🚀 Features
✅ Convert text into ASCII art
✅ Support for multiple banner fonts
✅ Handle new lines (\n)
✅ Clean and modular Go code structure

```
🧱 Project Structure
ascii-art/
│
├── main.go          # Entry point
├── banner.go        # Banner selection & file loading
├── render.go        # ASCII rendering logic
│
├── standard.txt     # Default font
├── shadow.txt       # Shadow style font
├── thinkertoy.txt   # Thinkertoy style font
```
## Requirements

+ Go installed (version 1.18+ recommended)
▶️ Usage
Basic usage
go run . "Hello"
With banner style
go run . "Hello" shadow
go run . "Hello" thinkertoy
🧪 Examples
go run . "Hello"
go run . "Hello\nWorld"
go run . "ASCII ART" shadow
⚠️ Rules & Behavior
Default banner is standard
Invalid banner name → program exits
Only ASCII characters (32–126) are rendered
Empty input produces no output
\n creates a new ASCII-art line
🧠 How It Works

+ The program reads a banner file (ASCII font)

+ Each character is mapped using ASCII index:

index = (char - 32) * 9
Each character is printed line-by-line (8 rows)
Output is assembled into ASCII-art text
🛠️ Installation

+ Clone the repository:

+ git clone https://github.com/Onoja217/ascii-art.git
cd ascii-art

+ Run the program:

go run . "Your Text"
🔄 Development Workflow
git add .
git commit -m "your message"
git push
📌 Notes
Banner files must not be empty
File names are case-sensitive
Ensure .txt files are in the project root
✨ Future Improvements
Add color support
Support custom banner files
Add CLI flags (--font)
Export output to file

👨‍💻 Author

+ Onoja Monday Ojonugba

📄 License

+ This project is for educational purposes.