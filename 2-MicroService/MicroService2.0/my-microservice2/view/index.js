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
            let data = {
                aid: aid,
                sd: sd,
                ed: ed,
                sk: sk,
            }
            axios.get(this.urlPrefix + "houses", { params: data })
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        postFunc() {
            let data = {
                "acreage": JSON.stringify(this.input),
                "address": JSON.stringify(this.input),
                "area_id": JSON.stringify(this.input),
                "beds": JSON.stringify(this.input),
                "capacity": JSON.stringify(this.input),
                "deposit": JSON.stringify(this.input),
                "facility": JSON.stringify(this.input),
                "max_days": JSON.stringify(this.input),
                "min_days": JSON.stringify(this.input),
                "price": JSON.stringify(this.input),
                "room_count": JSON.stringify(this.input),
                "title": JSON.stringify(this.input),
                "unit": JSON.stringify(this.input),
            }
            axios.post(this.urlPrefix + "houses", data)
                .then((response) => {
                    this.output = response.data;
                    console.log(this.output);
                }, (err) => {
                    console.log(err);
                })
        },
        putFunc() {
            let data = {
                action: JSON.stringify(this.input),
                reason: JSON.stringify(this.input)
            }
            axios.put(this.urlPrefix + `/orders/${id}/status`, data)
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