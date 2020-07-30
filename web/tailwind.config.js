module.exports = {
    purge: [],
    theme: {
        extend: {
            boxShadow: {
                t: '0px -4px 5px -2px rgba(0, 0, 0, .3)',
                b: '0px 4px 5px -2px rgba(0, 0, 0, .3)',
            },
            maxHeight: {
                '0': '0',
                '1/4': '25%',
                '1/2': '50%',
                '3/4': '75%',
                'full': '100%',
            },
            maxWidth: {
                'xxs': '14.8rem',
            }
        },
    },
    variants: {},
    plugins: [],
}
