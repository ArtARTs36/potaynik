var app = new Vue({
    el: '#app',
    data: {
        secret_create_options: [
            {
                "key": "password",
                "label": "Пароль",
                "valueKey": "password",
                "example": "cdJ$FDA_1",
            },
            {
                "key": "ip",
                "label": "IP или подсеть",
                "valueKey": "ip",
                "example": "192.168.0.1"
            }
        ],
        show_options_enabled: false,
        new_secret: {
            value: null,
            ttl: 60,
            auth_factors: new Map(),
        },
    },
    methods: {
        arrayChunk(array, size) {
            let chunks = []

            for (let i = 0; i < array.length; i += size) {
                chunks.push(array.slice(i, i + size))
            }

            return chunks
        },
        toggleShowOptions() {
            this.show_options_enabled = !this.show_options_enabled
        },
        createSecret() {
            console.log(`try to create secret, uses ${this.new_secret.auth_factors.length} auth factors`)

            let newSecret = {
                "value": this.new_secret.value,
                "ttl": this.new_secret.ttl,
            }

            if (this.new_secret.auth_factors.length > 0) {
                newSecret.auth_factors = this.new_secret.auth_factors;
            }

            axios.post('http://api.potaynik.ru/api/secrets/add', newSecret)
        },
    }
});

