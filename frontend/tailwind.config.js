/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        "box-color": "#171717",
        blue: "#7380F4",
        white: "#e3e2e1",
        yellow: "#7a5b04",
      },
      spacing: {
        one: "0.5rem",
        two: "1rem",
        three: "8rem",
        four: "4rem",
      },
      textColor: "#e3e2e1",
      fontSize: {
        title1: "2rem",
        smm: "0.7rem",
      },
    },
  },
  plugins: [],
};
