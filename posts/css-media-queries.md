Title: CSS Media Queries
Date: 07-Feb-2025

Media queries are a powerful CSS feature that allows developers to apply different styles based on the characteristics of the device accessing the website.  This is the cornerstone of responsive web design, letting you create layouts that adapt seamlessly to various screen sizes, orientations, resolutions, and even input types.

Essentially, a media query acts like a conditional statement for CSS.  It checks if certain conditions about the user's device are met, and if so, applies the associated styles.

```css
/* Basic media query structure */
@media (condition) {
  /* CSS rules to apply if the condition is true */
}
```

Here's how you can use media queries to create a responsive layout:

```css
/* Default styles for larger screens */
body {
  font-size: 16px;
  line-height: 1.5;
}

.container {
  width: 960px;
  margin: 0 auto;
}

/* Media query for smaller screens (e.g., tablets) */
@media (max-width: 768px) {
  .container {
    width: 720px;
  }
}

/* Media query for even smaller screens (e.g., phones) */
@media (max-width: 480px) {
  body {
    font-size: 14px;
  }

  .container {
    width: 90%; /* Fluid width */
    padding: 0 15px; /* Add some padding */
  }
}
```

In this example:

1. **Default styles:**  The initial CSS applies to larger screens (typically desktops).

2. **`@media (max-width: 768px)`:** This media query targets devices with a maximum width of 768 pixels (common for tablets). It adjusts the width of the `.container`.

3. **`@media (max-width: 480px)`:** This query targets even smaller devices like phones.  It further reduces the font size and makes the container fluid, adjusting to the screen width using percentages. It also adds padding for better readability.

**Common Media Features:**

Besides `max-width` and `min-width`, you can use many other media features:

* `orientation`: `portrait` or `landscape`.
* `aspect-ratio`: The ratio of width to height.
* `resolution`:  Pixels per inch (dpi).
* `hover`: Whether the primary input mechanism can hover (useful for distinguishing touch devices).

**Combining Media Features:**

You can combine multiple media features using `and`:

```css
@media (min-width: 768px) and (max-width: 1024px) {
  /* Styles for devices between 768px and 1024px wide */
}

@media (orientation: landscape) and (min-resolution: 2dppx) {
  /* Styles for landscape orientation on high-resolution screens */
}
```


By strategically using media queries, you can ensure your website provides an optimal viewing experience across a wide range of devices, enhancing user accessibility and satisfaction.
