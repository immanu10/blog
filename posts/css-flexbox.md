Title: CSS Flexbox
Date: 25-Feb-2025

Flexbox is a powerful layout module in CSS that simplifies arranging items within a container. It offers flexibility in controlling alignment, direction, and order, making it ideal for both small-scale component layouts and complex application designs.

Here's a breakdown of key Flexbox concepts:

**1. Flex Container:**

The parent element containing the flex items. You enable flexbox by setting the `display` property to `flex` or `inline-flex`.

```css
.container {
  display: flex; /* or inline-flex */
}
```

**2. Flex Items:**

The direct children of the flex container.  Their layout is controlled by the flex container's properties.

**3. Main Axis & Cross Axis:**

Flexbox lays out items along two axes:

- **Main Axis:**  The primary direction items are placed (default is row, horizontally).
- **Cross Axis:** The perpendicular direction to the main axis.


**4. Key Flex Container Properties:**

- `flex-direction`:  Determines the main axis direction.  Common values:
    - `row` (default): left to right
    - `row-reverse`: right to left
    - `column`: top to bottom
    - `column-reverse`: bottom to top

- `justify-content`:  Aligns items along the main axis. Common values:
    - `flex-start` (default): items at the start of the main axis
    - `flex-end`: items at the end of the main axis
    - `center`: items at the center of the main axis
    - `space-between`: items evenly distributed, space between them
    - `space-around`: items evenly distributed, space around each item

- `align-items`: Aligns items along the cross axis. Common values:
    - `flex-start`:  items at the start of the cross axis
    - `flex-end`: items at the end of the cross axis
    - `center`: items at the center of the cross axis
    - `stretch` (default): items stretched to fill the cross axis

- `flex-wrap`:  Controls whether items wrap onto multiple lines.
    - `nowrap` (default): items stay on one line, may overflow
    - `wrap`: items wrap onto multiple lines
    - `wrap-reverse`:  items wrap onto multiple lines in reverse order


**5. Key Flex Item Properties:**

- `flex-grow`: How much an item can grow relative to other items.
- `flex-shrink`: How much an item can shrink relative to other items.
- `flex-basis`: The initial size of the item.
- `order`:  Controls the order in which items appear (lower numbers appear first).
- `align-self`: Allows overriding the `align-items` property for individual items.


**Example:**

```html
<div class="container">
  <div class="item">Item 1</div>
  <div class="item">Item 2</div>
  <div class="item">Item 3</div>
</div>
```

```css
.container {
  display: flex;
  justify-content: space-between; /* Space between items */
  align-items: center; /* Center vertically */
}

.item {
  background-color: lightblue;
  padding: 10px;
  border: 1px solid blue;
}
```

This example creates a flex container with three items spaced evenly apart and centered vertically.  Experimenting with the different properties will give you a better understanding of how Flexbox works. Flexbox's intuitive syntax and powerful features make it a must-know for modern web development.
