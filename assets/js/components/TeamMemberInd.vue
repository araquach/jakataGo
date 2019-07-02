<template>
    <div class="columns section">
        <div class="column is-6">
            <figure class="image">
                <img src="http://via.placeholder.com/800x800" alt="team member">
            </figure>
        </div>
        <div class="column">
            <h2>Team Member Name</h2>
            HEllo
            <p><strong>Para 1</strong></p>
            <p>Para 2</p>
            <p>Para 3</p>
            <p><strong>Favourite Style: </strong>fav_style</p>
            <p><strong>Favourite Product: </strong>fav_prod</p>
            <p>Average Cut &amp; Colour with Team Member name: &pound;team_member-price</p>

            <a href="/team #teammember-id">Back to the full team</a>

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
                TeamMembers: []
            }
        },

        mounted() {
            axios.get('/api/team').then(response => this.TeamMembers = response.data);
        }
    }
</script>