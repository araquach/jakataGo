<template>
    <div class="section columns is-3 is-multiline">
        <div v-for="(tm, index) in TeamMembers" class="column is-half">
            <a :href="tm.ReviewLink">
                <div class="card">
                    <figure class="image">
                        <img :src="tm.Image" :alt="tm.Fname">
                    </figure>
                    <div class="card-content">
                        <div class="media">
                            <div class="media-content">
                                <p class="title is-4 has-text-primary">{{tm.Name}}</p>
                                <p class="subtitle has-text-primary">{{tm.Level}}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </a>
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
            axios.get('/api/team').then(response => this.TeamMembers = response.data)
                .catch(error => {
                    console.log(error)
                })
        }
    }
</script>