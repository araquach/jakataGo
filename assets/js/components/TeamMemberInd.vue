<template>
    <div>

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