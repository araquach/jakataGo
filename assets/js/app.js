import Vue from "vue";
// import NewsItem from './components/NewsItem.vue';
import ReviewFeed from './components/ReviewFeed.vue';
// import BlogPost from './components/BlogPost.vue';
// import TeamMember from './components/TeamMember.vue';
// import TeamMemberInd from './components/TeamMemberInd';

import PracticeComponent from './components/PracticeComponent';

window.axios = require('axios');
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

// Vue.component('news-item', NewsItem);
Vue.component('review-feed', ReviewFeed);
// Vue.component('blog-post', BlogPost);
// Vue.component('team-member', TeamMember);
// Vue.component('team-ind', TeamMemberInd);
//
// Vue.component('practice-component', PracticeComponent);



const app = new Vue({
    el: '#app'
});

// Bulma NavBar Burger Script
document.addEventListener('DOMContentLoaded', function () {
    // Get all "navbar-burger" elements
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

    // Check if there are any navbar burgers
    if ($navbarBurgers.length > 0) {

        // Add a click event on each of them
        $navbarBurgers.forEach(function ($el) {
            $el.addEventListener('click', function () {

                // Get the target from the "data-target" attribute
                let target = $el.dataset.target;
                let $target = document.getElementById(target);

                // Toggle the class on both the "navbar-burger" and the "navbar-menu"
                $el.classList.toggle('is-active');
                $target.classList.toggle('is-active');

            });
        });
    }

});
