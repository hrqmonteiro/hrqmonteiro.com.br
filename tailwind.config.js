/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./templates/**/*.{html,js,templ,go}",
        "./templates/components/**/*.{html,js,templ,go}",
        "./templates/pages/**/*.{html,js,templ,go}"
    ],
    theme: {
        extend: {
            colors: {
                primary: 'var(--primary)',
                'text-primary': 'var(--text-primary)',
                'text-secondary': 'var(--text-secondary)'
            },
            fontFamily: {
                sans: ['Roboto', 'ui-sans-serif', 'sans-serif']
            }
        }
    },
    plugins: []
}