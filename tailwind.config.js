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
                secondary: 'var(--secondary)',
            },
            textColor: {
                primary: 'var(--text-primary)',
                secondary: 'var(--text-secondary)',
                mono: 'var(--mono)'
            },
            fontFamily: {
                sans: ['Inter', 'ui-sans-serif', 'sans-serif']
            },
            grayscale: ['hover', 'group-hover']
        }
    },
    plugins: []
}