/** @type {import('tailwindcss').Config} */

const dracula = require('tailwind-dracula/colors');

module.exports = {
  mode: 'jit',
  content: ['./public/templates/*.html'],
  purge: ['./public/templates/*.html'],
  theme: {
    extend: {
      backgroundImage: {
        'dracula-roll': "url('./bg.webm')",
      },
      colors: {
          ...dracula,
          primary: dracula.purple,
          gray: dracula.darker,
          blue: dracula.blue,
          cyan: dracula.cyan,
          green: dracula.green,
          orange: dracula.orange,
          pink: dracula.pink,
          purple: dracula.purple,
          red: dracula.red,
          yellow: dracula.yellow,
          dark: dracula.dark,
          darker: dracula.darker,
          light: dracula.light,
      }
    },
  },
  plugins: [
    require('tailwind-dracula')('dracula'),
  ],
}

