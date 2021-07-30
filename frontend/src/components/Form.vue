<template>
    <form @submit.prevent="formSubmit">
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
        <div class="form-control2">
            <input type="text" v-model="form.link" placeholder="Link">
            <div class="submit-button-wrapper">
                <input type="submit" value="Shorten!" id="submit-button">
            </div>
        </div>
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
                console.log(`In formSubmit(): ${this.form.link}`)
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
    display: block;
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

.form-control2 {
  /*margin: 20px 0; */
  margin-right: 0;
  display: flex;
}
.form-control2 input {
    width: 100%;
    height: 40px;
    margin: 5px;
    padding: 3px 7px;
    font-size: 17px;
    flex: 1;
    margin-right:0px;
    border: 1px solid steelblue;
    font-family: Poppins, sans-serif;
    border-radius: 12px;
    border-top-right-radius: 0px;
    border-bottom-right-radius: 0px;
}
.submit-button-wrapper input {
    width: 100%;
    height: 40px;
    margin: 5px;
    padding: 3px 7px;
    font-size: 17px;
    flex: 1;
    margin-right:0px;
    border: 1px solid steelblue;
    font-family: Poppins, sans-serif;
    border-radius: 8px;
}
#submit-button {
    border: 1px solid steelblue;
    border-left: 0;
    margin-left: -2px;
    height: 48px;
    padding: 3px 7px;
    font-size: 17px;
    background-color: steelblue;
    color: white;
    border-top-left-radius: 0px;
    border-bottom-left-radius: 0px;
}
</style>