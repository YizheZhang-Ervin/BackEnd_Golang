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
            let data = {}
            axios.get(this.urlPrefix + "session", { params: data })
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        postFunc() {
            let data = {
                "mobile": JSON.stringify(this.input),
                "password": JSON.stringify(this.input),
                "sms_code": JSON.stringify(this.input)
            }
            axios.post(this.urlPrefix + "users", data)
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        putFunc() {
            let data = {}
            axios.put(this.urlPrefix + ``, data)
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        deleteFunc() {
            let data = {}
            axios.delete(this.urlPrefix + "", { data: data })
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
    }
}).mount("#app");