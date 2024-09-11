<script>
	export default {
		data: function() {
			return {
				errormsg: null,
				username: localStorage.getItem("username"),
				userID: localStorage.getItem("userID"),
				stream: [],
				comments: [],
				showComments: false,
				currentPhotoId: null,
				newComment: "",
				banList: [],
			}
		},
		methods: {
			async refresh() {
				try {
					let response = await this.$axios.get("/users/" + this.username + "/stream", {
					headers: {Authorization: "Bearer " + this.userID }
					});
					this.stream = response.data
					if (this.stream) {
						for (let i = 0; i < this.stream.length; i++) {
      						const photo = this.stream[i];
     						photo.photoData = `data:image/octet-stream;base64,${photo.photoData}`; 
    					}	
					}
					let banResponse = await this.$axios.get("/users/" + this.username + "/bans", {
					headers: {Authorization: "Bearer " + this.userID }
					});
					this.banList = banResponse.data
					console.log(banResponse.data)
				} catch (e) {
					if (e.response && e.response.status === 401) {
					this.errormsg = "Status Unauthorized"
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Status Internal Server Error"
					}
					else {
						this.errormsg = e.toString();
					}
				}
			},
			async getComments(photoID) {
				try {
					let response = await this.$axios.get("/users/" + this.username + "/photos/" + photoID + "/comments", {
						headers: {Authorization: "Bearer " + this.userID }
					});
					this.comments = response.data
					this.refresh()
				} catch(e) {
					if (e.response && e.response.status === 401) {
						this.errormsg = "Status Unauthorized"
					} else if (e.response && e.response.status === 400) {
						this.errormsg = "Status Bad Request"
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Status Internal Server Error"
					} else {
						this.errormsg = e.toString();
					}
				}
      			this.showComments = true;
      			this.currentPhotoId = photoID;
    		},
    		async postComment(photoID, content) {
				try {
					let response = await this.$axios.post("/users/" + this.username + "/photos/" + photoID + "/comments", { content }, {
					headers: {Authorization: "Bearer " + this.userID, "Content-Type": "application/json"}
					});
					this.newComment = ""
					this.getComments(photoID)
				} catch(e) {
					if (e.response && e.response.status === 401) {
						this.errormsg = "Status Unauthorized"
					} else if (e.response && e.response.status === 400) {
						this.errormsg = "Status Bad Request"
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Status Internal Server Error"
					} else {
						this.errormsg = e.toString();
					}
				}
    		},
			async deleteComment(commentID, photoID) {
				try {
					let response = await this.$axios.delete("/users/" + this.username + "/photos/" + photoID + "/comments/" + commentID, {
					headers: {Authorization: "Bearer " + this.userID }
					});
					this.getComments(photoID)
				} catch(e) {
					if (e.response && e.response.status === 401) {
						this.errormsg = "Status Unauthorized"
					} else if (e.response && e.response.status === 400) {
						this.errormsg = "Status Bad Request"
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "Status Internal Server Error"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async goToSearch() {
				this.$router.push({ path: "/profile"})
			},
			async unbanUser(username) {
				try {
				let response = await this.$axios.delete("/users/" + this.username + "/bans/" + username, {
					headers: {Authorization: "Bearer " + this.userID }
				});
				this.refresh()
				}
				catch(e) {
					if (e.response && e.response.status === 401) {
					this.errormsg = "Status Unauthorized"
				} else if (e.response && e.response.status === 400) {
					this.errormsg = "Status Bad Request"
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Status Internal Server Error"
				} else {
					this.errormsg = e.toString();
				}
				}
			},
		},
		mounted() {
			this.refresh()
		}
	}
	</script>
	
	<template>
		<div>
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">Home page</h1>
				<button type="submit" class="btn btn-sm btn-primary" @click="goToSearch()">Search profile</button>
			</div>
			<div class="col-md-4" v-for="photo in stream" :key="photo.id">
				<div class="card mb-4 shadow-sm">
					<img class="card-img-top" :src=photo.photoData alt="Card image cap">
					<div class="card-body">
						<p class="card-text">Uploaded on: {{ photo.uploadDate }}</p>
						<button v-if="!photo.isLiked" type="button" class="btn btn-sm btn-outline-primary">Like</button>
						<button v-if="photo.isLiked" type="button" class="btn btn-sm btn-outline-primary">Unlike</button>
						<button type="button" class="btn btn-sm btn-outline-primary" @click=getComments(photo.id)>Comments</button>
						<span>Likes: {{ photo.likesCount }}</span><br>
						<span>Comments: {{ photo.commentsCount }}</span><br>
					</div>
	
					<div class="card mb-4 shadow-sm" v-if="showComments && currentPhotoId === photo.id">
						<div class="card-header">
						  Comments
						</div>
						<div class="card-body">
							<ul>
								<li v-for="comment in comments" :key="comment.id">
								{{ comment.userID }}: {{ comment.content }}
								<a href="javascript:" @click="deleteComment(comment.id)">[Delete]</a>
								</li>
							</ul>
							<form @submit.prevent="postComment(photo.id, newComment)">
							<input type="text" v-model="newComment" placeholder="Add a comment">
							<button type="submit" class="btn btn-sm btn-primary">Add Comment</button>
							</form>
						</div>
					</div>
				</div>
			</div>
			<ul>
				<p>List of banned users</p>
				<li v-for="ban in banList" :key="ban.banID">
				{{ ban.banID }}: {{ ban.username }}
				<a href="javascript:" @click="unbanUser(ban.username)">[Delete]</a>
				</li>
			</ul>
			
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</template>
	
	<style>
	</style>