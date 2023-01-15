Vue.createApp({
    data() {
        return {
            input: "",
            output: "",
            urlPrefix: "/api/v1.0/"
        }
    },
    methods: {
        getFunc() {
            // console.log(this.input)
            // this.output = this.input
            // console.log(this.output)
            axios.get(this.urlPrefix + "session")
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        postFunc() {
            let data = {
                mobile: JSON.stringify(this.input),
                password: JSON.stringify(this.input)
            }
            axios.post(this.urlPrefix + "sessions", data)
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        putFunc() {
            let data = {
                name: JSON.stringify(this.input)
            }
            axios.put(this.urlPrefix + "user/name", data)
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        deleteFunc() {
            let data = {}
            axios.delete(this.urlPrefix + "session", { data: data })
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
    }
}).mount("#app");