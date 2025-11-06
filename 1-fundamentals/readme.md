# 🐹 Fundamentals

This directory contains the **fundamentals of Golang**, including syntax, types, loops, functions, structs, and more.  
It is organized into multiple directories, each covering a specific topic.

🔄 Each topic directory contains:  
- One or more `main.go` files with lesson examples  
- An `exercises/` directory with **each exercise in its own directory**, each as a **standalone `package main`**  
- An `exercises.md` file inside `exercises/` listing **all exercise descriptions**, for easy navigation and access to the desired exercise  

📌 Current Topics (examples):  
- Variables and Types  
- Loops and Conditionals  
- Functions and Methods  
- Structs and Interfaces  
- Pointers  
- And more…

⚡ How to Use:  
1. Navigate to the topic or exercise directory you want to run.
   
2. For lesson examples:  
   ```bash
   go run main.go

3. For exercises: enter the exercise directory and run:
   ```
   go run main.go

4. For exercises with tests:
   ```
   go test ./...

💡 **Tip:** if you are using an editor like VS Code with *Code Runner*, you can run the active file directly using `Ctrl + Alt + N`, without navigating to the directory.
