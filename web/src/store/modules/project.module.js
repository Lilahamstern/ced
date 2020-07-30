const state = {
    projects: [{
            id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
            orderId: 1026304,
            title: 'SU, Ombyggnad av OP1-OP7 for nya sterilhantering',
            sector: 'Skola och larande',
            manager: 'Lena Karlsson',
            client: 'Vastfastigheter',
            co2: 210,
            updatedAt: 1 * (Math.random() * 10).toFixed()
        },
        {
            id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
            orderId: 10215136,
            title: 'Nya KS, ombyggnad soder del 2',
            sector: 'Halso och Sjukvard',
            manager: 'Karl Gustav Olof',
            client: 'Goteborgs Stad',
            co2: 430,
            updatedAt: 2 * (Math.random() * 10).toFixed()
        },
        {
            id: 'b7e414b8-f7c6-413a-a248-f1c04fbe85c0',
            orderId: 10274356,
            title: 'Evakueringsbyggnad',
            sector: 'Skyddsombud',
            manager: 'Lars Erik Glensson',
            client: 'Vastra Gotaland Regionen',
            co2: 630,
            updatedAt: 3 * (Math.random() * 10).toFixed()
        }
    ],
    search: '',
}

const mutations = {
    search(state, payload) {
        state.search = payload
    }
}

const actions = {
    search({commit}, payload) {
        console.log(payload)
        commit('search', payload)
    }
}

const getters = {}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}