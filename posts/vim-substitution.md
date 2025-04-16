Title: Vim Substitution
Date: 16-Apr-2025

Substituting text within Vim is a powerful editing feature, allowing for complex replacements and manipulations.  This goes beyond simple find-and-replace functionality, offering pattern matching and special characters for fine-grained control.

The basic substitution command in Vim follows this structure:

```
:s/{pattern}/{replacement}/{flags}
```

Let's break down each component:

* **`:`**:  Enters command mode.
* **`s`**:  Indicates the substitution command.
* **`{pattern}`**: The text you want to search for. This can be a literal string or a regular expression.
* **`{replacement}`**: The text you want to replace the pattern with.  This can also include special characters.
* **`{flags}`**:  Optional flags that modify the substitution's behavior.


**Examples:**

* **Simple Replacement:**
   Replace "apple" with "orange" on the current line:

   ```vim
   :s/apple/orange
   ```

* **Global Replacement:**
    Replace all occurrences of "apple" with "orange" in the entire file:

    ```vim
    :%s/apple/orange/g
   ```

   Here, `%` represents all lines, and `g` stands for global.

* **Case-Insensitive Replacement:**
    Replace "apple" with "orange", ignoring case:

    ```vim
    :s/apple/orange/i
    ```

* **Confirming Replacements:**
    Replace "apple" with "orange", but ask for confirmation before each replacement:

   ```vim
   :%s/apple/orange/gc
   ```

   The `c` flag prompts for confirmation.

* **Using Regular Expressions:**
    Replace all occurrences of a digit followed by the letter "a" with "X":

   ```vim
   :%s/\d\+a/X/g
   ```

   Here, `\d\+` represents one or more digits.

* **Using Capture Groups:**
   Swap the order of two words separated by a space:

   ```vim
   :s/\(\w\+\) \(\w\+\)/\2 \1/
   ```
   The `\(...\)` define capture groups, and `\1` and `\2` refer to the captured text.

* **Special Characters in Replacement:**
    Replace "apple" with "apple (replaced)":


   ```vim
    :s/apple/apple (&)/
    ```

    `&` represents the entire matched pattern.  You can also use `\0`.  Similarly, `\1`, `\2`, etc., refer to captured groups if regular expressions are used.

**Beyond the Basics:**

The substitution command offers much more. Experiment with different regular expressions and flags to harness its full potential.  The `:help :s` command within Vim provides comprehensive documentation on all the options available.  Mastering substitution will significantly enhance your text editing efficiency in Vim.
