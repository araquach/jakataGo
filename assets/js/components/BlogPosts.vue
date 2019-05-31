<template>
    <div>
        <blog-post
                v-for="blogPost in blogPosts"
                v-bind:key="blogPost.id"
                v-bind:blogPost="blogPost"
        ></blog-post>
    </div>
</template>
