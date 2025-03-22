Title: Vim Visual Mode
Date: 22-Mar-2025

Visual mode is a powerful feature in Vim that allows you to select text for operations like copying, deleting, changing, or indenting.  It offers several variations, each with its own advantages.

Entering Visual Mode:

You can enter visual mode using the following keys:

* `v`: Character-wise visual mode. This is the most common mode and selects characters individually.
* `V`: Line-wise visual mode. This selects entire lines at a time.
* `Ctrl-v` (or `Ctrl-q` in some terminals): Block-wise visual mode. This selects rectangular blocks of text.


Using Visual Mode:

Once in visual mode, you can navigate using the standard Vim movement keys (e.g., `h`, `j`, `k`, `l`, `w`, `b`, `e`, `0`, `$`, `gg`, `G`).  As you move, the selection will expand or contract.

Performing Operations:

After selecting the desired text, you can perform an operation. Some common examples:

* `y`: Yank (copy) the selected text.
* `d`: Delete the selected text.
* `c`: Change (delete and enter insert mode) the selected text.
* `>`: Indent the selected text.
* `<`: Unindent the selected text.
* `u`: Convert the selected text to lowercase.
* `U`: Convert the selected text to uppercase.

Example:

Let's say you have the following text:

```
This is line one.
This is line two.
This is line three.
```

1. **Character-wise visual mode:**  Place your cursor on the 'o' in "one" and press `v`. Move your cursor to the 't' in "two" using `w` (move to next word). The text "one.This is line " will be highlighted. Pressing `d` will delete the selection.

2. **Line-wise visual mode:** Place your cursor anywhere on line two and press `V`.  Press `j` to select line three as well. Pressing `y` will copy both lines.

3. **Block-wise visual mode:** Place your cursor on the 'T' of the first line. Press `Ctrl-v`.  Press `2j` to move down two lines. Press `$` to extend the selection to the end of each line. Pressing `I` (insert at beginning of selection) followed by "// " and then `Esc` will add "// " to the beginning of each selected line, effectively commenting them out.  This results in:


```
// This is line one.
// This is line two.
// This is line three.
```

Exiting Visual Mode:

Press `Esc` or any non-movement key that doesn't initiate a Vim command (e.g., a letter key if you aren't in insert mode) to exit visual mode.


Visual mode, combined with Vim's movement keys and operators, offers a highly efficient way to manipulate text. Mastering it significantly enhances your editing speed and productivity.
