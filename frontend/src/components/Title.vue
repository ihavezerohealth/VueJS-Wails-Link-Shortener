<template>
    <h1>{{title}}</h1>
</template>

<script>
    export default {
        name: "Title",
        props: ["title"]
    }
</script>

<style scoped>
h1 {
    color: #333;
    display: flex;
}
</style>