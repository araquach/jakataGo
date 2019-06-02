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
                                <p class="title is-4 has-text-primary">{{tm.Fname}} {{tm.Lname}}</p>
                                <p class="subtitle has-text-primary">{{tm.level}}</p>
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

        methods: {

            level(){
                switch(this.TeamMember.level) {
                    case 1:
                        "Director";
                        break;
                    case 2:
                        "Senior Stylist";
                        break;
                    case 3:
                        "Stylist";
                        break;
                    default:
                        "Stylist";
                }

                return;
            }
        },

        mounted() {
            axios.get('/api/team').then(response => this.TeamMembers = response.data);
        }
    }
</script>