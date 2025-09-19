@import('tailwindcss').Config;

module.exports = {
    darkMode: 'class', // Enables dark mode via class
    theme: {
      extend: {
        colors: {
          background: "#1a1a1a", // Dark background
          foreground: "#2e2e2e", // Slightly lighter gray
          accent: "#d76d77", // Pastel red
          muted: "#8a8a8a", // Greyish text
          text: "#e0e0e0", // Light text
        },
      },
    },
    plugins: [],
};
