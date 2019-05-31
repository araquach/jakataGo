<template>
    <h1>{{ message }}</h1>
</template>

<script>
    export default {

        data() {
            return {
                message: "Hello from Vue"
            }
        }
    }

    // mounted() {
    //     axios.get('/staffall').then(response => this.employees = response.data);
    // }
</script>