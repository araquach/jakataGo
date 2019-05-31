<template>
    <a href="/blog">
        <div class="box columns is-paddingless">
            <div class="column is-4 is-paddingless">
                <figure class="image">
                    <img src="http://via.placeholder.com/1000x1000" alt="Image 1">
                </figure>
            </div>
            <div class="column">
                <h3 class="title is-4 has-text-primary">{{title}}</h3>
                <p class="is-size-6">{{para}}</p>
            </div>
        </div>
    </a>
</template>

<script>
    export default {

        data() {
            return {
                title: "Blog Post Title is cool",
                para: "This is the cool blog posts paragraph"
            }
        }
    }
</script>