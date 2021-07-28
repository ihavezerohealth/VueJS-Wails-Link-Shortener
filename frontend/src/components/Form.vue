<template>
    <form @submit.prevent="formSubmit">
        <div class="form-control">
            <label>Link to Shorten:</label>
            <input type="text" v-model="form.link" name="thisistheformlink" placeholder="Link">
        </div>
        <div class="form-control">
            <label>Service to use:</label>
            <select>
                <option value="bit.ly">Bit.ly</option>
            </select>
        </div>
        <div class="check-before-submit">
            <p>{{form.link}} will be shortened using the service Bit.ly</p>
            <label>Are you sure you want to carry out the operation?</label>
            <input type="checkbox" v-model="form.checked">
        </div>
        <input type="submit" value="Submit">
    </form>
</template>

<script>
    export default {
        name: "Form",
        data: () => ({
                form: {
                    link: "https://github.com/ihavezerohealth",
                    service: "bit.ly",
                    checked: false,
                }
        }),
        methods: {
            formSubmit() {
                if (!this.form.checked) {
                    return alert("Please confirm you want to carry out the operation!")
                }
                if (!this.form.link.length < 0) {
                    return alert("Please enter a link to shorten!")
                }
                this.$emit("submit", this.form);
            }
        }
    }
</script>

<style>
form {
    align-items: center;
}
.form-control {
  /*margin: 20px 0; */
  margin-right: 25px;
}
.form-control label {
  display: block;
}
.form-control input {
  width: 100%;
  height: 40px;
  margin: 5px;
  padding: 3px 7px;
  font-size: 17px;
}

</style>