/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/routes/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],

  daisyui: {
    themes: [
      {
        "mycodecompiled": {

          "primary": "#3C787E",
          "secondary": "#17191E",
          "accent": "#9D6A9D",
          "neutral": "#ffffff",
          "base-100": "#17191E"
          

        }
      }
    ]
  }
}
