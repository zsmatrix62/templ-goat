/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./templates/**/*.templ'],
    darkMode: 'class',
    theme: {},
    plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
    daisyui: {
        darkTheme: 'light', // name of one of the included themes for dark mode
        themeRoot: ':root', // The element that receives theme color CSS variables
    },
};
