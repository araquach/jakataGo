<template>
    <div id="feedback-feed" class="box">
        <p>&quot;{{review}}&quot;</p>
        <p>{{client}} - hair by {{stylist}}</p>
        <a href="/reviews">More reviews here</a>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                review: "Had a great experience - thanks very much, I will be coming back",
                stylist: "Nat",
                client: "Anna Alexander"
            }
        }
    }
</script>