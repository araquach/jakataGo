<template>
    <div>
        <div class="columns box has-margin-3">
            <a :id="blog.Slug" class="anchor"></a>
            <div class="column is-5">
                <figure class="image">
                    <img :src="blog.Img" :alt="blog.Alt">
                </figure>
            </div>
            <div class="column">
                <h2 class="title is-4">{{blog.Title}}</h2>
                <p v-for="(Para, index) in blog.Paras">{{Para}}</p>
                <p><a :href="blog.Link" class="button is-primary">Read more</a></p>
                <small class="is-size-7">Published by {{blog.Author}}</small><br>
                <time class="is-size-7">Created at {{blog.CreatedAt}}</time>
            </div>
        </div>
    </div>
</template>

<script>

    export default {

        props:

        data() {
            return {
                blogPosts: []
            }
        },

        mounted() {
            axios.get('/api/blogs').then(response => this.blogs = response.data);
        }
    }
</script>