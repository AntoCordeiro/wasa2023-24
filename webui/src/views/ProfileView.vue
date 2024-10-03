<script>
export default {
	data: function() {
		return {
			errormsg: null,
			isMyProfile: false,
			searchedUsername: "",
			username: localStorage.getItem("username"),
			userID: localStorage.getItem("userID"),
			userProfile: {
        		user: {},
        		photos: [],
        		follows: [],
        		followers: [],
      		},
			newPhoto: {},
			decodedImage: null,
			showComments: null,
			currentPhotoID: null,
			newComment: "",
			comments: [],
			searchQuery: "",
		}
	},
	methods: {
		async refresh() {
			try {
				if (this.searchedUsername === this.username) {
					this.isMyProfile = true
				} else {
					this.isMyProfile = false
				}
				let response = await this.$axios.get("/users/" + this.username + "/profiles/" + this.searchedUsername, {
					headers: {Authorization: "Bearer " + this.userID }
				});
				this.userProfile = response.data;
				console.log(this.userProfile);
				if (this.userProfile.photos) {
					for (let i = 0; i < this.userProfile.photos.length; i++) {
      					const photo = this.userProfile.photos[i];
     					photo.photoData = `data:image/octet-stream;base64,${photo.photoData}`; 
    				}
				}
				console.log("changed phtos:", this.userProfile);
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
		async uploadPhoto(event) {
			let file = event.target.files[0];
			if (file) {
				let formData = new FormData();
				formData.append('file', file);

				try {
					let response = await this.$axios.post("/users/" + this.username + "/photos", formData, {
					headers: {Authorization: "Bearer " + this.userID, 
							  'Content-Type': 'multipart/form-data'}
					});
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
				}}
			}
		},
		async deletePhoto(photoID) {
			try {
				let response = await this.$axios.delete("/users/" + this.username + "/photos/" + photoID, {
					headers: {Authorization: "Bearer " + this.userID }
				});
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
				}}
		},
		async followUser(username) {
			try {
				let response = await this.$axios.post("/users/" + this.username + "/follows", { username }, {
					headers: {Authorization: "Bearer " + this.userID, "Content-Type": "application/json" }
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
		async unfollowUser(username) {
			try {
				let response = await this.$axios.delete("/users/" + this.username + "/follows/" + username, {
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
		async getComments(photoID) {
			try {
				let response = await this.$axios.get("/users/" + this.username + "/photos/" + photoID + "/comments", {
					headers: {Authorization: "Bearer " + this.userID }
				});
				this.comments = response.data
				console.log("Comments data:", this.comments);
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
		async searchUserProfile(searchQuery) {
			this.searchedUsername = searchQuery;
			this.searchQuery = "";
			this.refresh()
		},
		async banUser() {
			try {
				let response = await this.$axios.post("/users/" + this.username + "/bans", { username: this.searchedUsername }, {
				headers: {Authorization: "Bearer " + this.userID, "Content-Type": "application/json"}
				});
				this.$router.push({ path: "/myStream"})
			} catch(e) {
				if (e.response && e.response.status === 401) {
					this.errormsg = "Status Unauthorized"
				} else if (e.response && e.response.status === 400) {
					this.errormsg = "Status Bad Request"
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Status Internal Server Error"
				} else {
					this.errormsg = e.toString();					}
			}
		},
		async LikePhoto(photoID) {
			try {
				const authorizationHeader = `Bearer ${this.userID}`; // Assuming your backend expects "Bearer " format

    			const response = await this.$axios.post(`/users/${this.username}/photos/${photoID}/likes`, {}, {headers: { Authorization: authorizationHeader },});

			    this.refresh();
			} catch(e) {
					if (e.response && e.response.status === 401) {
					this.errormsg = "Status Unauthorized"
				} else if (e.response && e.response.status === 400) {
					this.errormsg = "Status Bad Request"
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "Status Internal Server Error"
				} else {
					this.errormsg = e.toString();
				}}
		},
	mounted() {
		this.refresh()
	},
}}


</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<div v-if="userProfile.user">
			  <h1 class="h2">{{ userProfile.user.username }}</h1>
			  <span class="post-count">Post Count: {{ userProfile.user.postCount }}</span>
			</div>
			<p v-else>Please refresh</p>
			<input v-if="isMyProfile" type="file" @change="uploadPhoto">
			<input type="text" v-model="searchQuery" placeholder="Search for a user">
			<button type="button" class="btn btn-sm btn-primary" @click="searchUserProfile(searchQuery)">Search</button>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button v-if="!isMyProfile" type="button" class="btn btn-sm btn-outline-secondary" @click="followUser(searchedUsername)">
						Follow
					</button>
					<button v-if="!isMyProfile" type="button" class="btn btn-sm btn-outline-secondary" @click="unfollowUser(searchedUsername)">
						Unfollow
					</button>
				</div>
				<div class="btn-group me-2">
					<button v-if="!isMyProfile" type="button" class="btn btn-sm btn-outline-primary" @click="banUser()">
						Ban
					</button>
				</div>
			</div>
		</div>
		<div class="Row">
			<h2>Follows</h2>
    		<ul>
      			<li v-for="follow in userProfile.follows" :key="follow">
        		{{ follow }}
      			</li>
    		</ul>
			<h2>Followers</h2>
    		<ul>
      			<li v-for="follower in userProfile.followers" :key="follower">
        		{{ follower }}
      			</li>
    		</ul>
		</div>
		<div class="col-md-4" v-for="photo in userProfile.photos" :key="photo.id">
            <div class="card mb-4 shadow-sm">
                <img class="card-img-top" :src=photo.photoData alt="Card image cap">
				<div class="card-body">
					<p class="card-text">Uploaded on: {{ photo.uploadDate }}</p>
					<button v-if="!photo.isLiked" type="button" class="btn btn-sm btn-outline-primary" @click=LikePhoto(photo.id)>Like</button>
      				<button v-if="photo.isLiked" type="button" class="btn btn-sm btn-outline-primary">Unlike</button>
      				<button type="button" class="btn btn-sm btn-outline-primary" @click=getComments(photo.id)>Comments</button><br>
					<span>Likes: {{ photo.likesCount }}</span><br>
					<span>Comments: {{ photo.commentsCount }}</span><br>
					<button v-if="isMyProfile" type="button" class="btn btn-sm btn-danger" @click="deletePhoto(photo.id)">Delete Photo</button>
				</div>

				<div class="card mb-4 shadow-sm" v-if="showComments && currentPhotoId === photo.id">
					<div class="card-header">
					  Comments
					</div>
					<div class="card-body">
					  	<ul>
							<li v-for="comment in comments" :key="comment.commentData.id">
							{{ comment.commentData.userID }}: {{ comment.commentData.content }}
							<a href="javascript:" @click="deleteComment(comment.commentData.id, photo.id)">[Delete]</a>
							</li>
					  	</ul>
					  	<form @submit.prevent="postComment(photo.id, newComment)">
							<input
							  type="text"
							  v-model="newComment"
							  placeholder="Add a comment"
							/>
							<button type="submit" class="btn btn-sm btn-primary">Add Comment</button>
						  </form>						  
					</div>
				</div>
			</div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
    button { 
    	margin-right: 10px; 
  	}
</style>
