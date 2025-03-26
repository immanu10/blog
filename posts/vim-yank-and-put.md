Title: Vim Yank and Put
Date: 26-Mar-2025

Yanking and putting in Vim are analogous to copying and pasting in other text editors.  They provide a powerful and efficient way to move text around within your file or even between different files.

**Yanking (Copying)**

The `y` operator is used for yanking (copying) text.  It's often combined with a *motion* to specify what text to yank.  Here are a few examples:

* `yw`: Yank the current word.
* `y2w`: Yank the next two words.
* `yl`: Yank the character under the cursor.
* `y$`: Yank to the end of the line.
* `y0`: Yank to the beginning of the line.
* `yy`: Yank the entire current line.  Equivalent to `Y`.
* `y3j`: Yank the current line and the three lines below it.
* `yip`: Yank inside the current paragraph.
* `ya'`: Yank inside single quotes, including the quotes.
* `yi(`: Yank inside the parentheses, excluding the parentheses.


**Putting (Pasting)**

The `p` operator puts (pastes) the yanked text *after* the cursor or the current line. `P` (uppercase) puts the yanked text *before* the cursor or the current line.


**Examples:**

* Yank the current word and paste it after the next word: `ywep`
* Yank the current line and paste it below: `yyp`
* Yank the current line and paste it above: `yyP`
* Yank inside the parentheses and paste it at the end of the line: `yi(A)` (Note: `A` appends to the end of the line).

**Using Registers**

Vim has named registers that allow you to store multiple yanked items.  You can specify a register by using `"register` before the yank or put command.

* `"ayiw`: Yank the current word into register `a`.
* `"ap`: Put the contents of register `a`.
* `"0p`: Put the yanked text from the default register (the last yank without a specified register).
* `"+y`: Yank the current line into the system clipboard (often used for pasting outside of Vim).
* `"+p`: Paste the contents of the system clipboard.


**Example with Registers:**

Let's say you want to swap two words.  You can use registers to achieve this:

1. Position the cursor on the first word.
2. `"ayiw`: Yank the first word into register `a`.
3. `dw`: Delete the first word.
4. Move the cursor to the second word.
5. `"byiw`: Yank the second word into register `b`.
6. `dw`: Delete the second word.
7. `"ap`: Put the first word (from register `a`).
8. `"bp`: Put the second word (from register `b`).

By mastering yanking, putting, and registers, you'll significantly enhance your text editing efficiency in Vim.
