/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "sans-serif"], // ฟอนต์ Inter
        kanit: ["Kanit", "sans-serif"], // ฟอนต์ Kanit
        montserrat: ["Montserrat", "sans-serif"], // ฟอนต์ Montserrat
      },
      boxShadow: {
        'text-glow': '0 0 10px rgba(0, 255, 0, 0.7), 0 0 20px rgba(0, 255, 0, 0.7)', // สีเขียวเรืองแสง
      }
    },
  },
  plugins: [function({ addUtilities }) {
    addUtilities({
      '.scrollbar-hide': {
        '-ms-overflow-style': 'none', /* IE and Edge */
        'scrollbar-width': 'none', /* Firefox */
        '&::-webkit-scrollbar': {
          display: 'none', /* Chrome, Safari, Edge */
        },
      },
    });
  },],
}