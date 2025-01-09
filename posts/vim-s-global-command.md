Title: Vim's Global Command
Date: 09-Jan-2025

Vim's `:global` command, often abbreviated to `:g`, is a powerful tool for performing operations on lines matching a specific pattern.  It allows you to execute Ex commands on a subset of lines within a file, significantly enhancing your editing efficiency.  This blog post will delve into the mechanics and usage of `:global`.

## Basic Syntax

The fundamental syntax of the global command is as follows:

```
:[range]g/{pattern}/{command}
```

* `[range]` (Optional): Specifies the lines on which the command operates. If omitted, it defaults to the entire file.  You can use specific line numbers (e.g., `1,10`), relative line numbers (e.g., `.,+5`), or special symbols like `%` for the whole file.

* `g`:  The global command itself.

* `{pattern}`:  A regular expression that determines which lines are selected.  Lines matching this pattern will be subject to the specified command.

* `{command}`:  The Ex command to execute on the selected lines.  This could be anything from deleting lines (`d`) to substituting text (`s`) or even running external commands.


## Practical Examples

Let's illustrate the power of `:g` with some examples.

1. **Deleting lines containing a specific word:**

Suppose you want to remove all lines containing the word "error" in a log file. You can use the following command:

```vim
:g/error/d
```

This command searches for the pattern "error" and deletes all lines where it's found.


2. **Commenting out JavaScript code:**

Imagine you need to comment out multiple lines of JavaScript code. You can achieve this with `:g` combined with the `s` (substitute) command:

```vim
:g/^\s*console\.log/s/^/\/\/
```
This searches for lines that begin (using `^`) with optional whitespace (`\s*`) followed by `console.log` and prepends `//` to the start of those lines, effectively commenting them out.

3. **Numbering lines conditionally:**

You can number lines containing specific patterns. This is handy for debugging or analyzing sections of code.

```vim
:g/function/s/^/\=line('.') . ' '/
```

This finds lines containing "function" and substitutes the beginning of the line with the current line number (`line('.')`) followed by a space. The `\=` allows you to evaluate Vimscript expressions within the substitution.


4. **Executing external commands:**

For more complex tasks, you can even execute external commands on matched lines. For example, you could use this to perform operations like sorting or filtering:

```vim
:g/^TODO:/!sort  
```

This finds lines beginning with "TODO:" and runs the external `sort` command on *all other lines* because of the `!` operator which inverts the selection.


##  Conclusion

Vim's `:global` command is a remarkably versatile tool.  By mastering its usage, you can manipulate text and code with surgical precision and significantly boost your editing productivity.  Experiment with different patterns and commands to uncover the full potential of this essential Vim feature.
