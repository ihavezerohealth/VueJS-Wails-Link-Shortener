<template>
    <div id="app">
        <div class="container">
            <Title title="Link Shortener"/>
            <Form @submit="passFormToBackend"/>
            <h1>{{link}}</h1>
        </div>
        <button @click="link = ''">Clear Link</button>
        <!--
        <LearnMore v-show="learningMore"/>
        <CopyrightAndAbout @learnMoreClicked="toggleLearnMore()"/>
        -->
    </div>
</template>

<script>
import Title from './components/Title.vue'
import Form from './components/Form.vue'
//import CopyrightAndAbout from './components/CopyrightAndAbout.vue'
//import LearnMore from './components/LearnMore.vue'

    export default {
        name: 'App',
        components: {
            Title,
            Form,
            //CopyrightAndAbout,
            //LearnMore,
        },
        methods: {
            //toggleLearnMore () {
            //    this.learningMore = !(this.learningMore);
            //},
            async passFormToBackend(form) {
                /* eslint-disable no-unused-vars */
                var link = "";
                this.link = "";
                console.log(form);
                let response = await fetch('https://api-ssl.bitly.com/v4/shorten', {
                    method: 'POST',
                    headers: {
                        'Authorization': 'dc6bd94d1d19f0e027c070a0fd0eb47238e72ff3',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ "long_url": form.link, "domain": "bit.ly" })
                })
                link = await response.json();
                console.log(link);
                link = link["link"];
                console.log(link);
                /*.then(response => response.json(), link = response).then(data => console.log(data),
                    link = data,
                    console.log("link is" + link),
                )
                */
                /* eslint-enable no-unused-vars */
                /*
                window.backend.shortenLink(form.link).then(result => {
                self.message = result;
                console.log(result);
                })
                */
                this.link = link;
            },
        },
        data: () => ({
            showAbout: false,
            link: '',
            form: {
                link: null,
                service: 'bitly',
            }
        })
    }
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap');
body {
    font-family: Poppins, sans-serif;
}
.container {
    max-width: 500px;
    margin: 30px auto;
    overflow: auto;
    min-height: 300px;
    border: 2px solid steelblue;
    padding: 30px;
    border-radius: 12px;
}
#app {
    justify-content: center;
    text-align: center;
}
button {
    font-family: Poppins, sans-serif;
}
</style>