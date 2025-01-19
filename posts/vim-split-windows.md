Title: Vim Split Windows
Date: 19-Jan-2025

Splitting windows in Vim allows you to view and edit multiple files simultaneously or different parts of the same file. This is incredibly useful for cross-referencing code, comparing changes, or simply keeping an eye on related information. Here's a breakdown of how to manage split windows effectively in Vim:

**Creating Splits:**

* **Vertical Split:**  `:vsplit [filename]` (or `:vs`) splits the window vertically. If you provide a filename, that file opens in the new split. Otherwise, the current file is displayed in both splits.
* **Horizontal Split:** `:split [filename]` (or `:sp`) splits the window horizontally.  Similar to vertical split, providing a filename opens that file, otherwise, the current file is duplicated.
* **New Tab:** `:tabnew [filename]` (or `:tabnew`) opens a new tab. This is like having a separate Vim instance, offering more isolation between your workspaces.

**Navigating Splits:**

* **Ctrl+w + Direction:** Use `Ctrl+w` followed by a directional key (`h` (left), `j` (down), `k` (up), `l` (right)) to move the cursor between splits.
* **Ctrl+w + w:** Quickly cycle through open splits.

**Resizing Splits:**

* **Ctrl+w + +/-:** Increase or decrease the size of the active split.  `Ctrl+w + +` makes the split taller (horizontal) or wider (vertical), while `Ctrl+w + -` does the opposite.
* **Ctrl+w + =:**  Equalize the size of all splits.

**Closing Splits:**

* **Ctrl+w + q:**  (or `:q`) Close the active split. If it's the last split, it closes the entire Vim window (or tab if in a tab).
* **Ctrl+w + o:** Close all other splits, leaving only the active split open.


**Example Workflow:**

Let's say you're working on a JavaScript project and want to compare two files, `index.js` and `utils.js`:

1. Open `index.js`: `vim index.js`
2. Create a vertical split containing `utils.js`: `:vs utils.js`
3. Navigate between the splits using `Ctrl+w + h` and `Ctrl+w + l`.
4. Resize the splits for a comfortable view:  `Ctrl+w + +` or `Ctrl+w + -`
5. Close a split when finished: `Ctrl+w + q`


**Advanced Tip:**

You can combine split commands with file wildcards. For example, `:vsplit *.js` will open all JavaScript files in the current directory in vertical splits.  This is incredibly powerful for quickly browsing related files.

By mastering split windows, you'll significantly enhance your efficiency and workflow within Vim.  These commands enable a flexible and customizable editing environment, facilitating everything from small code edits to large project navigation.
