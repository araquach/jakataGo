<template>
    <div>
        <div v-for="(newsItem, index) in newsItems">
            <a href="{{newsItem.Link}}">
                <div class="box columns is-paddingless">
                    <div class="column is-4 is-paddingless">
                        <figure class="image">
                            <img src="{{ newsItem.Img}}0" alt="{{newsItem.Alt}}">
                        </figure>
                    </div>
                    <div class="column">
                        <h3 class="title is-4 has-text-primary">{{newsItem.Title}}</h3>
                        <p>{{newsItem.Para}}</p>
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
                newsItems: []
            }
        },

        mounted() {
            axios.get('/api/newsItems').then(response => this.newsItems = response.data);
        }
    }
</script>