module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      backgroundImage: {
        "cone-gradient": "conic-gradient(from 90deg, green, orange, red)",
      },
    },
    colors: {
      "dark-green": "#265642",
      "light-green": "#34783B",
      "gold-yellow": "#F8D648",
    },
  },
  plugins: [],
};
