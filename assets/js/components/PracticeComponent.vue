<template>
    <div class="columns is-multiline section">
        <div class="column is-4" v-for="person in people">
            <div class="box">
                <div>
                    <p class="is-size-4">{{person.name}}</p>
                </div>
                <div>
                    <p>
                        <em>gender:</em> {{person.gender}}
                    </p>
                    <p>
                        <em>height:</em> {{person.height}}
                    </p>
                    <p>
                        <em>hair color:</em> {{person.hair_color}}
                    </p>
                    <p>
                        <em>eye color:</em> {{person.eye_color}}
                    </p>
                    <p>
                        <em>birth year:</em> {{person.birth_year}}
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    const url = 'https://swapi.co/api/people/?format=json';
    export default {

        data() {
            return {
                people: []
            }
        },
        methods: {
            loadPeople() {
                axios.get(url)
                    .then((response) => {
                        this.people = response.data.results;
                    });
            }
        },
        mounted() {
            // When the Component is ready fetch the JSON from the Server Backend
            this.loadPeople();
        }
    }
</script>