/** @type {import('tailwindcss').Config} */
/* eslint-env node */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    fontFamily: {
      sans: ["Roboto"],
      title: ["Roboto Condensed"]
    },
    extend: {
      gridTemplateColumns: {
        "fill-10": "repeat(auto-fill, minmax(10rem, 1fr))",
        "fill-12": "repeat(auto-fill, minmax(12rem, 1fr))",
        "fill-15": "repeat(auto-fill, minmax(15rem, 1fr))",
        "fill-20": "repeat(auto-fill, minmax(20rem, 1fr))",
      },
    },
  },
  plugins: [],
};
