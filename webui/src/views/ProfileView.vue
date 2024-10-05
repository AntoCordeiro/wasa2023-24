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
			fileToUpload: null,
			successfulMsg: "",
			emptySearch: true,
			isFollowed: false,
		}
	},
	methods: {
		async refresh() {
			this.errormsg = null
			this.successfulMsg = ""
			try {
				if (this.searchedUsername != "") {
					this.emptySearch = false
				}
				if (this.searchedUsername === this.username) {
					this.isMyProfile = true
				} else {
					this.isMyProfile = false
				}
				let response = await this.$axios.get("/users/" + this.username + "/profiles/" + this.searchedUsername, {
					headers: {Authorization: "Bearer " + this.userID }
				});
				this.userProfile = response.data;
				if (this.userProfile.photos) {
					for (let i = 0; i < this.userProfile.photos.length; i++) {
      					const photo = this.userProfile.photos[i];
     					photo.photoData = `data:image/octet-stream;base64,${photo.photoData}`; 
    				}
				}
				if (this.userProfile.followers && this.userProfile.followers.includes(this.username)) {
					this.isFollowed = true	
				}
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
		async selectFile(event) {
			this.fileToUpload = event.target.files[0];
			this.successfulMsg = ""
		},
		async uploadPhoto() {
			if (this.fileToUpload) {
				let formData = new FormData();
				formData.append('file', this.fileToUpload);

				try {
					let response = await this.$axios.post("/users/" + this.username + "/photos", formData, {
					headers: {Authorization: "Bearer " + this.userID, 
							  'Content-Type': 'multipart/form-data'}
					});
					this.selectedFile = null;  
      				this.$refs.fileInput.value = ''; 
					this.successfulMsg = "Photo uploaded successfully!"
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
				console.log(this.userProfile.followers)
				let response = await this.$axios.post("/users/" + this.username + "/follows", { username }, {
					headers: {Authorization: "Bearer " + this.userID, "Content-Type": "application/json" }
				});
				this.isFollowed = true
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
				console.log(this.userProfile.followers)
				let response = await this.$axios.delete("/users/" + this.username + "/follows/" + username, {
					headers: {Authorization: "Bearer " + this.userID }
				});
				this.isFollowed = false
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
				let response = await this.$axios.post("/users/" + this.username + "/photos/" + photoID + "/likes", {}, {
					headers: {Authorization: "Bearer " + this.userID}
				});
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
		async UnLikePhoto(photoID) {
			try {
				let response = await this.$axios.delete("/users/" + this.username + "/photos/" + photoID + "/likes", {
					headers: {Authorization: "Bearer " + this.userID }
					});
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
		async goToMyStream() {
			this.$router.push({ path: "/myStream"})
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
			  <span class="post-count" v-if="!emptySearch">Post Count: {{ userProfile.user.postCount }}</span>
			</div>
			<p v-else>Please refresh</p>
			<div>
				<input type="text" v-model="searchQuery" placeholder="Search for a user" @keyup.enter="searchUserProfile(searchQuery)">
				<button type="button" class="btn btn-sm btn-primary" @click="searchUserProfile(searchQuery)">Search</button>
			</div>
			<div>
				<button type="button" class="btn btn-sm btn-primary" @click="goToMyStream">My Stream</button>
			</div>
		</div>
		<div class="row" style="margin-bottom: 20px;">
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button v-if="!isMyProfile && !emptySearch && !isFollowed" type="button" class="btn btn-sm btn-outline-secondary" @click="followUser(searchedUsername)">
						Follow
					</button>
					<button v-if="!isMyProfile && !emptySearch && isFollowed" type="button" class="btn btn-sm btn-outline-secondary" @click="unfollowUser(searchedUsername)">
						Unfollow
					</button>
				</div>
				<div class="btn-group me-2">
					<button v-if="!isMyProfile && !emptySearch" type="button" class="btn btn-sm btn-outline-primary" @click="banUser()">
						Ban
					</button>
				</div>
			</div>
			<div>
				<h5 v-show="isMyProfile">Upload Photo: </h5>
				<input v-show="isMyProfile" type="file" @change="selectFile" ref="fileInput">
				<button v-if="fileToUpload" @click="uploadPhoto">Upload Photo</button>
				<p v-if="successfulMsg" style="color: green;">{{ successfulMsg }}</p>
			  </div>
		</div>
		<div class="col-md-4" v-for="photo in userProfile.photos" :key="photo.id">
            <div class="card mb-4 shadow-sm">
                <img class="card-img-top" :src=photo.photoData alt="Card image cap">
				<div class="card-body">
					<p class="card-text">Uploaded on: {{ photo.uploadDate }}</p>
					<button v-if="!photo.isLiked" type="button" class="btn btn-sm btn-outline-primary" @click=LikePhoto(photo.id)>Like</button>
      				<button v-if="photo.isLiked" type="button" class="btn btn-sm btn-outline-primary" @click="UnLikePhoto(photo.id)">Unlike</button>
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
							{{ comment.commentData.username }}: {{ comment.commentData.content }}
							<a v-if="comment.commentData.userID == userID" href="javascript:" @click="deleteComment(comment.commentData.id, photo.id)">[Delete]</a>
							</li>
					  	</ul>
					  	<form @submit.prevent="postComment(photo.id, newComment)">
							<input
							  type="text"
							  v-model="newComment"
							  placeholder="Add a comment"
							  />
							<button type="submit" class="btn btn-sm btn-primary">Post</button>
						  </form>						  
					</div>
				</div>
			</div>
		</div>
		<div class="Row">
			<h2 v-if="!emptySearch">Follows</h2>
    		<ul>
      			<li v-for="follow in userProfile.follows" :key="follow">
        		{{ follow }}
      			</li>
    		</ul>
			<h2 v-if="!emptySearch">Followers</h2>
    		<ul>
      			<li v-for="follower in userProfile.followers" :key="follower">
        		{{ follower }}
      			</li>
    		</ul>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
    button { 
    	margin-right: 10px; 
  	}
</style>
