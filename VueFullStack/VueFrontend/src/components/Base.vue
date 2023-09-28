<script>
import axios from 'axios'

export default {
    name: 'base',
    data() {
        return {
            msg: ''
        }
    },
    methods: {
        getmessage() {
            const path = 'https://localhost:5000/';
            axios.get(path)
                .then((res) => {
                    this.msg = res.data;
                })
                .catch((error) => {
                    console.error(error);
                })
        }
    },
    created() {
        this.getmessage();
    }
}
</script>

<template>
    <div>
        <p>{{ msg.messge }}</p>
    </div>
</template>