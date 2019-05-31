<template>
    <div>
        <h1 v-for="(review, index) in reviews">{{review.Review}}</h1>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                reviews: []
            }
        },

        mounted() {
            axios.get('/api/reviews').then(response => this.reviews = response.data);
        }
    }
</script>