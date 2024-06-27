# MyShell

MyShell is a simple shell program written in Go. It provides basic shell functionalities such as navigating directories, printing the current directory, and executing commands.

## Features

- **exit**: Exit the shell.
- **pwd**: Print the current working directory.
- **cd [directory]**: Change the current directory to the specified directory.
- **echo [text]**: Print the specified text to the standard output.
- **type [command]**: Display whether the specified command is a shell builtin or an external command.

## Usage

1. **exit**: 
   ```
   $ exit
   ```

2. **pwd**: 
   ```
   $ pwd
   /path/to/current/directory
   ```

3. **cd [directory]**: 
   ```
   $ cd /path/to/directory
   ```

4. **echo [text]**: 
   ```
   $ echo Hello, World!
   Hello, World!
   ```

5. **type [command]**: 
   ```
   $ type echo
   echo is a shell builtin
   ```

