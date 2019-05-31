<template>
    <div>
        <div v-for="(review, index) in reviews" id="feedback-feed" class="box">
            <p>&quot;{{review.Review}}&quot;</p>
            <p>{{review.Client}} - hair by {{review.Stylist}}</p>
            <a href="/reviews">More reviews here</a>
        </div>
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