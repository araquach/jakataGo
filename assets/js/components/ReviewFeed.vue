<template>
    <div id="feedback-feed" class="box">
        <ul>
            <li>&quot;{{review}}&quot;<br><span>{{client}} - hair by {{stylist}}</span></li>
        </ul>
        <a href="/reviews">More reviews here</a>
    </div>
</template>

<script>
    export default {

        data() {
            return {
                review: "Had a great experience - thenks",
                stylist: "Nat",
                client: "Anna Alexander"
            }
        }
    }
</script>