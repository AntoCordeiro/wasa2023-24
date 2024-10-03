<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: "",
			userID: null,
		}
	},
	methods: {
		async doLogin() {
			try {
				let response = await this.$axios.post("/session", {username: this.username});
				this.userID = response.data;
				localStorage.setItem("userID", this.userID)
				localStorage.setItem("username", this.username)
				this.$router.push({ path: "/myStream"})
			} catch (e) {
				if (e.response && e.response.status === 400) {
					this.errormsg = "Status Bad Request: try again"
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Status Internal Server Error"
				}
				else {
					this.errormsg = e.toString();
				}
			}
		},
	},
	mounted() {
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Sign in</h1>
			<div class="mx-auto">
				<input type="text" id="username" v-model="username" required @keyup.enter="doLogin"/>
				<button  type="submit" @click="doLogin">Login</button>
			</div>
		</div>

		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
