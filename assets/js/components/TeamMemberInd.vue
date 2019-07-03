<template>
    <div class="columns section">
        <div class="column is-6">
            <figure class="image">
                <img :src="tm[0].Image" :alt="tm[0].Alt">
            </figure>
        </div>
        <div class="column">
            <h2 class="title is-3">{{tm[0].Name}}</h2>
            <p><strong>{{tm[0].Para1}}</strong></p>
            <p>{{tm[0].Para2}}<</p>
            <p>{{tm[0].Para3}}<</p>
            <p><strong>Favourite Style: </strong>{{tm[0].FavStyle}}</p>
            <p><strong>Favourite Product: </strong>{{tm[0].FavProd}}</p>
            <p>Average Cut &amp; Colour with Team Member name: &pound;{{tm[0].Price}}</p>

            <a :href="'/team#' + tm[0].Class">Back to the full team</a>

            <div id="fb-like">
                <div class="fb-like" data-href="http://www.jakatasalon.co.uk/team/team-member-name" data-width="250" data-layout="button" data-action="like" data-show-faces="true" data-share="true"></div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                tm : []
            }
        },

        mounted() {
            axios.get('/api/team').then(response => this.tm = response.data);
        }
    }
</script>