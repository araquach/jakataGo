<template>
    <div class="columns box has-margin-3">
        <a id="slug" class="anchor"></a>
        <div class="column is-5">
            <figure class="image">
                <img src="http://via.placeholder.com/1000x1000" alt="">
            </figure>
        </div>
        <div class="column">
            <h2 class="title is-4">{{title}}</h2>
            <p>{{para}}</p>
            <p><a href="link-to-blog" class="button is-primary">Read more</a></p>
            <small class="is-size-7">Published by {{author}}</small><br>
            <time class="is-size-7">Created at {{published}}</time>
        </div>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                title: "This is a top Blog Post",
                para: "Welcome to my Vue blog",
                author: "Adam",
                published: "15th June"
            }
        }
    }
</script>