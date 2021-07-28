<template>
    <div class=copyright-and-about>
        <p>Issued under the MIT License.</p>
        <button @click="this.$emit('learnMoreClicked')">Learn More</button>
    </div>
</template>

<script>
    export default {
        name: "CopyrightAndAbout",
        emits: ["learnMoreClicked"]
    }
</script>
<style scoped>
div {
    align-content: center;
}
button {
    align-self: center;
}
p {
    font-size: 12px;
    text-align: center;
}
</style>