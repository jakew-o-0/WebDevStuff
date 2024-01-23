/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    extend: {
      colors: {
        custOrange: "#db9c4f",
        custBlue: "#5c8b9c",
        custLightBlue: "#4bb5dc",
        custBeige: "#e1d2be",
        custGrey: "#4d585c",
      },
      fontFamily: {
        "sans": ['Rubik', "sans-serif"],
      }
    },
  },
  plugins: [],
}

