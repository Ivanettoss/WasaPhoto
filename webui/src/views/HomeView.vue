<script>import axios from "axios";

export default {
  data: function () {
    return {
      errormsg: null,
      username: "",
      token: localStorage.getItem("token"),
      localUser: localStorage.getItem("username"),
      nFollowers: 0,
      nFollowing: 0,
      followState: false,
      banState: false,
      photos: [],
      nPost: 0,
      selectedFile: null,

      showComments: false, // Controlla se il box dei commenti è visibile
      comments: [], // Lista dei commenti
      newComment: "",

      showSettings: false,

      showPopup: false, // Controlla la visibilità del popup
      newUsername: '', // Nuovo nome utente

      searchQuery: '',
      users: [],

      myself: false,
      photoStream:[]
    };
  },
  methods: {
    // Navbar Methods
    async searchUsers(searchQuery) {
      this.users = []; // Resetta la lista se il campo è vuoto

      try {
        let response = await this.$axios.get("/user/" + this.localUser + "/searchusers/" + this.searchQuery);
        for (let i = 0; i < response.data.userlist.length; i++) {
          this.users.push(response.data.userlist[i]);
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    goToProfile(profileUsername) {
      this.$router.push({ path: '/profile/' + profileUsername });
      this.buildProfile();
      this.mySelf();
    },

    goToHome() {
      this.$router.push({ path: '/home' });
    },

    settingsDropdown() {
      this.showSettings = !this.showSettings;
    },

    async doLogOut() {
      localStorage.removeItem("token");
      localStorage.removeItem("username");
      this.$router.push({ path: '/' });
    },

    async buildStream(){
      try{
        let response=await this.$axios.get("/user/"+this.localUser+"/stream", 
        {headers: {
            Authorization: this.token
          }
        }
        );
        this.photoStream=[]
        if (response.data.photos) {
          for (let i = 0; i < response.data.photos.length; i++) {
            response.data.photos[i].commentsOpen = false;
            this.photoStream.push(response.data.photos[i]);
          }
        }
      }catch(e){

      }
    },
     async doLike(photo) {

      try {

        let response = await this.$axios.put("/user/" + photo.username + "/photo/" + photo.idphoto + "/like/" + this.localUser, {}, {
          headers: {
            "Authorization": this.token
          }
        })

      }
      catch (e) {
        this.errormsg = e.toString();
        console.log(e)
      }
      photo.isliked = true
      photo.nlikes += 1

    },
    async doUnLike(photo) {

      try {
        let response = await this.$axios.delete("/user/" + photo.username + "/photo/" + photo.idphoto + "/like/" + this.localUser, {
          headers: {
            "Authorization": this.token
          }
        })
      }
      catch (e) {
        this.errormsg = e.toString();
      }
      photo.isliked = false
      photo.nlikes -= 1
    },
    toggleComments(photo) {
      photo.commentsOpen = !photo.commentsOpen;
    },
     async addComment(photo) {
      if (this.newComment.trim()) {
        try {
          console.log(this.newComment)
          let idUser = parseInt(this.token)
          let response = await this.$axios.post("/user/" + photo.username + "/photo/" + photo.idphoto + "/comments",
            { "text": this.newComment, "user": { "username": this.localUser, "id": idUser } },
            {
              headers: {
                "Authorization": this.token
              }
            })
          photo.ncomments += 1
          photo.comments.push(response.data)
        } catch (e) {
          this.errormsg = e.toString();
        }
        this.newComment = ""; // Resetta il campo dopo l'invio del commento
      }
    },

    async deleteComment(photo, idcomment) {
      try {
        let response = await this.$axios.delete("/user/" + photo.username + "/photo/" + photo.idphoto + "/comments/" + idcomment, {
          headers: {
            "Authorization": this.token
          }
        })

        for (let i = 0; i < photo.comments.length; i++) {
          if (photo.comments[i].idcomment == response.data.idcomment)
            photo.comments.splice(i, 1)

        }
        photo.comments = photo.comments
        photo.ncomments -= 1
      } catch (e) {
        this.errormsg = e.toString();
      }
    },


    // Profile Building and User Actions
    async buildProfile() {
      this.myself = false;
      try {
        let response = await this.$axios.get("/searchuser/" + this.$route.params.username, {
          headers: {
            Authorization: this.token
          }
        });

        this.username = response.data.username;
        this.nFollowers = response.data.nfollower;
        this.nFollowing = response.data.nfollowed;
        this.nPost = response.data.npost;
        this.photos = [];
        if (response.data.photos) {
          for (let i = 0; i < response.data.photos.length; i++) {
            response.data.photos[i].commentsOpen = false;
            this.photos.push(response.data.photos[i]);
          }
        }

        if (this.username != localStorage.getItem("username")) {
          this.followState = response.data.followstate;
          this.banState = response.data.banstate;
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

     mySelf() {
      if (this.username == this.localUser) {
        console.log(this.username)
        console.log(this.localUser)
        this.myself = true
      }
      else{
        this.myself = false
      }
      return this.myself

    },

    async changeUsername() {
      if (this.newUsername) {
        try {
          let response = await this.$axios.put("/user/" + this.localUser + "/set_username", { "username": this.newUsername }, {
            headers: {
              "Authorization": this.token
            }
          });
          this.username = response.data.username;
          this.showPopup = false; // Chiude il popup dopo il salvataggio
          this.newUsername = ''; // Resetta il campo di input
          localStorage.setItem('username', this.username);
          this.$router.push({ path: '/profile/' + this.username });
        } catch (e) {
          this.errormsg = e.toString();
        }
      }
    },
  },
  mounted() {
    this.buildStream()
  }
};
</script>

<template>

  <header class="header">
    <div class="logos">
    <img id="logopic" src="./wasacircle.png">
    <a class="logo">WasaPhoto</a>
    </div>
    <nav class="navbar">
      <div class="search-container">
        <input
          id="searchForm"
          type="text"
          placeholder="Search users..."
          v-model="searchQuery"
          @keyup.enter="searchUsers(searchQuery)"
        />
        <button class="search-button" @click="searchUsers(searchQuery)"></button>

        <div v-if="users.length > 0" class="search-dropdown">
          <ul>
            <li
              v-for="user in users"
              :key="user.username"
              @click="goToProfile(user.username)"
            >{{ user.username }}</li>
          </ul>
        </div>
      </div>

      <a id="menubu" @click="goToHome()">Home</a>
      <a id="menubu" @click="goToProfile(localUser)">Profile</a>
      <a id="menubu" @click="settingsDropdown">Settings</a>

      <div class="dropdown">
        <div v-if="showSettings" class="dropdown-content">
          <a id="menubu" @click="showPopup = true">Change Username</a>
          <a id="menubu" @click="doLogOut()">Logout</a>
        </div>
      </div>
    </nav>
  </header>

  <div v-if="showPopup" class="popup-overlay">
    <div class="popup-box">
      <h1>Change Username</h1>
      <input type="text" v-model="newUsername" placeholder="Enter new username" />
      <button @click="changeUsername()">Save</button>
      <button @click="showPopup = false">Cancel</button>
    </div>
  </div>







  <ul class="card-list">
    <li class="card" v-for="photo in this.photoStream" :key="photo.idphoto">
      <div class="card-image">
        {{ photo.username }}
        <img :src="photo.photobytes" />
      </div>
      <div class="card-description" target="_blank">
        {{ photo.nlikes }}
        <button
          v-if="!photo.isliked"
          class="custom-button"
          @click="doLike(photo)"
        ></button>
        <button v-if="photo.isliked" class="custom-buttonLiked" @click="doUnLike(photo)"></button>
        {{ photo.ncomments }}
        <button class="custom-button2" @click="toggleComments(photo)"></button>
      </div>

      <!-- Box dei commenti nascosto sotto la foto -->
      <div v-if="photo.commentsOpen" class="comment-box">
        <ul>
          <li v-for="comment in photo.comments" :key="comment.idcomment">
            {{ comment.user.username }}: {{ comment.text }}
            <button
              v-if="comment.user.id == this.token"
              class="custom-buttonDelC"
              @click="deleteComment(photo, comment.idcomment)"
            ></button>
          </li>
        </ul>
        <textarea v-model="newComment" placeholder="write a comment..."></textarea>
        <button class="custom-buttonC" @click="addComment(photo)">Comment</button>
      </div>
    </li>
  </ul>
</template>
<style scoped>
.card-list{
    margin-top: 100px; 
}
body {
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center; /* Allinea il contenuto al centro orizzontalmente */
  min-height: 100vh; /* Imposta l'altezza minima del body su 100% dello schermo */
  box-sizing: border-box;
}
</style>
