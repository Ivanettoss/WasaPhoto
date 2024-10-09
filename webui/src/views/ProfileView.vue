<script>
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

      showComments: false,      // Controlla se il box dei commenti è visibile
      comments: [],             // Lista dei commenti
      newComment: "",

      showSettings: false,

      showPopup: false, // Controlla la visibilità del popup
      newUsername: '',  // Nuovo nome utente

      searchQuery: '',
      users: [],

      myself: false


    }

  },
  methods: {
    async buildProfile() {
      this.myself = false
      try {
        let response = await this.$axios.get("/searchuser/" + this.$route.params.username, {
          headers: {
            Authorization: this.token
          }
        })

        this.username = response.data.username
        this.nFollowers = response.data.nfollower
        this.nFollowing = response.data.nfollowed
        this.nPost = response.data.npost
        this.photos = []
        if (response.data.photos) {
          for (let i = 0; i < response.data.photos.length; i++) {
            response.data.photos[i].commentsOpen = false
            this.photos.push(response.data.photos[i])
          }
        }

        if (this.username != localStorage.getItem("username")){
          this.followState = response.data.followstate
          this.banState = response.data.banstate
        }


      }
      catch (e) {
        this.errormsg = e.toString();
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

    async searchUsers(searchQuery) {
      this.users = []; // Resetta la lista se il campo è vuoto

      try {
        let response = await this.$axios.get("/user/" + this.localUser + "/searchusers/" + this.searchQuery)
        for (let i = 0; i < response.data.userlist.length; i++) {
          this.users.push(response.data.userlist[i])
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    async doLogOut() {
      localStorage.removeItem("token")
      localStorage.removeItem("username")
      this.$router.push({ path: '/' })
    },

    mySelf() {
      if (this.username == this.localUser) {
        console.log(this.username)
        console.log(this.localUser)
        this.myself = true
      }
      console.log(this.myself)
      return this.myself

    },

    async doFollow() {

      try {
        let response = await this.$axios.put("/user/" + this.localUser + "/followed/" + this.$route.params.username, {},
          {
            headers: {
              "Authorization": this.token
            }
          })
        console.log("follows pre", this.followState)
        this.followState = true
        this.nFollowers += 1
        console.log("follows post", this.followState)

      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async unFollow() {
      try {
        let response = await this.$axios.delete("/user/" + this.localUser + "/followed/" + this.$route.params.username)
        console.log("follows pre", this.followState)
        this.followState = false
        this.nFollowers -= 1
        console.log("follows post", this.followState)
      } catch (e) {
        this.errormsg = e.toString();
      }
    },


    toggleComments(photo) {
      photo.commentsOpen = !photo.commentsOpen;
    },

    settingsDropdown() {
      this.showSettings = !this.showSettings;
    },
    async changeUsername() {
      if (this.newUsername) {
        try {
          // Aggiungi la logica per cambiare il nome utente, ad esempio:
          let response = await this.$axios.put("/user/" + this.localUser + "/set_username", { "username": this.newUsername },
            {
              headers: {
                "Authorization": this.token
              }
            })
          console.log(response)
          this.username = response.data.username
          this.showPopup = false; // Chiude il popup dopo il salvataggio
          this.newUsername = ''; // Resetta il campo di input
          localStorage.setItem('username', user.username);
          this.$router.push({ path: '/profile/' + this.username })
        } catch (e) {
          this.errormsg = e.toString();
        }
      }
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

    handleButtonClick() {
      if (!this.selectedFile) {
        this.$refs.fileInput.click();
        console.log("ho pickato la foto")
      } else {
        // Se c'è un file selezionato, caricalo
        console.log("else fatto")
        this.uploadPic();
      }
    },
    handleFileSelect(event) {
      const file = this.$refs.fileInput.files[0];
      if (file) {
        console.log("selected file:", file.name);
        const reader = new FileReader();
        reader.onload = (e) => {
          this.selectedFile = e.target.result;
        };
        reader.readAsDataURL(file);
        console.log("reader eseguito")
      }
    },

    async uploadPic() {
      if (!this.selectedFile) {
        console.error("No file selected");
        this.errormsg = "Please select a file to upload."
        return;
      }

      try {
        console.log(this.selectedFile)
        let response = await this.$axios.post("/user/" + this.localUser + "/upload", { "photobytes": this.selectedFile }, {
          headers: {
            Authorization: this.token
          }
        });
        console.log("post fatto")
        this.nPost += 1;
        console.log("Upload completato");
        this.selectedFile = null; // Resetta lo stato dopo l'upload
        this.buildProfile()
      } catch (error) {
        console.error("Errore durante l'upload:", error);
      }
    },
    async deletePic(idPhotoToDelete) {
      try {
        let response = await this.$axios.delete("/user/" + this.localUser + "/photo/" + idPhotoToDelete, {
          headers: {
            "Authorization": this.token
          }
        });
        console.log("delete eseguito")
        this.buildProfile()
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    goToProfile(profileUsername) {
      console.log("swag goto")
      this.$router.push({ path: '/profile/' + profileUsername })
      this.buildProfile()
      this.mySelf()

    },

    async doBan() {
      try {
        let response = await this.$axios.put("/user/" + this.localUser + "/ban_user/" + this.$route.params.username, {},
          {
            headers: {
              "Authorization": this.token
            }
          })
        this.banState = true
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    async dounBan() {
      try {
        let response = await this.$axios.delete("/user/" + this.localUser + "/ban_user/" + this.$route.params.username, {},
          {
            headers: {
              "Authorization": this.token
            }
          })
        this.banState = false
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    goToHome() {
      this.$router.push({ path: '/home' })
    }


  },
  mounted() {
    this.buildProfile()
  }

}
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

  <div class="profile-container">
    <div class="profile-picture">
      <img
        src="https://www.shareicon.net/data/512x512/2016/09/15/829452_user_512x512.png"
        alt="Foto profilo"
      />
    </div>
    <div class="profile-info">
      <div class="username">
        {{ username }}
        <input
          type="file"
          accept="image/*"
          style="display: none"
          ref="fileInput"
          @change="handleFileSelect"
        />
        <button
          v-if="mySelf()"
          class="buttonsDynamic"
          @click="handleButtonClick"
        >{{ selectedFile ? 'Upload' : 'Choose an image' }}</button>
        <button v-else-if="followState" class="buttonsDynamic" @click="unFollow()">unfollow</button>
        <button v-else-if="!followState" class="buttonsDynamic" @click="doFollow()">Follow</button>

        <button v-if="!mySelf() && !banState" class="buttonsDynamicBan" @click="doBan()">Ban</button>
        <button v-if="!mySelf() && banState" class="buttonsDynamicBan" @click="dounBan()">unban</button>
      </div>

      <div class="stats">
        <div class="stat-item">
          <span>{{ nPost }}</span>
          Post
        </div>
        <div class="stat-item">
          <span>{{ nFollowing }}</span>
          Followed
        </div>
        <div class="stat-item">
          <span>{{ nFollowers }}</span>
          Followers
        </div>
      </div>
    </div>
  </div>

  <ul class="card-list">
    <li class="card" v-for="photo in this.photos" :key="photo.idphoto">
      <div class="card-image">
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

        <button v-if="photo.username==localUser" class="custom-button3" @click="deletePic(photo.idphoto)"></button>
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

<style>
.view {
  height: 100%;
}

.header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  padding: 20px 100px;
  background: cadetblue;
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 100;
}

.logo {
  font-size: 20px;
  color: whitesmoke;
  text-decoration:none;
  font-weight: 500;
  
}

.logos{
  display: flex; /* Dispone gli elementi al suo interno (l'immagine e la scritta) in una riga */
  align-items: center;
}

#logopic{
  width:50px;
  height: auto; /* Mantiene le proporzioni dell'immagine */
  margin-right: 10px; 
}

.navbar a,
p {
  font-size: 15px;
  font-weight: 300;
  margin-left: 50px;
  margin-right: 30px;
}
.navbar a:hover {
  transform: scale(1.05);
}
#searchForm {
  background-color: transparent;
  border: none;
  border-bottom: 2px solid whitesmoke;
  outline: none;
  color: whitesmoke;
}
#menubu {
  color: whitesmoke;
}

.profile-container {
  background-color: rgba(255, 255, 255, 0.5);
  border: 1px solid #ddd;
  padding: 40px;
  display: flex;
  align-items: center;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  border-radius: 15px;
  width: 100%;
  max-width: 800px;
  position: relative;
  align-self: center;
  margin: auto;
}
.profile-picture {
  overflow: hidden;
  margin-right: 30px;
}
.profile-picture img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.profile-info {
  display: flex;
  flex-direction: column;
  margin: 10px;
}
.username {
  font-size: 36px;
  font-weight: bold;
  margin-bottom: 20px;
  color: cadetblue;
  display: flex;
  justify-content: space-between;
}
.stats {
  display: flex;
  justify-content: space-between;
  width: 400px;
}
.stat-item {
  text-align: center;
  color: cadetblue;
}
.stat-item span {
  display: block;
  font-size: 24px;
  font-weight: bold;
}
.stat-item {
  font-size: 20px;
}

.card-image {
  display: block;
  min-height: 20rem; /* layout hack */
  background-color: rgba(255, 255, 255, 0.5);
  background-size: cover;
  display: block;
  width: 100%;
  width: 100%;
  height: 200px; /* Imposta un'altezza fissa per la carta */
  
  overflow: hidden; /* Nasconde le parti in eccesso dell'immagine */
  position: relative;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover; /* L'immagine si adatta al contenitore mantenendo le proporzioni */
}

/* Layout Styles */

.card-list {
  margin: 1rem auto;
  padding: 0;
  font-size: 0;
  text-align: center;
  list-style: none;
}

.card {
  background-color: rgba(255, 255, 255, 0.5);
  display: inline-block;
  width: 90%;
  
  max-width: 20rem;
  margin: 1rem;
  font-size: 1rem;
  text-decoration: none;
  overflow: hidden;
  box-shadow: 0 0 3rem -1rem rgba(0, 0, 0, 0.5);
  transition: transform 0.1s ease-in-out, box-shadow 0.1s;
}

.card:hover {
  transform: translateY(-0.5rem) scale(1.0125);
  box-shadow: 0 0.5em 3rem -1rem rgba(0, 0, 0, 0.5);
}

.card-description {
  display: flex;
  justify-content: space-between; /* Distribuisce i pulsanti su entrambi i lati */
  align-items: center;
  padding: 15px 20px 5px;
}

.lNc {
  display: flex; /* Mantiene i pulsanti su una linea orizzontale */
  gap: 10px; /* Spazio tra il cuore e il commento */
  justify-content: flex-start; /* Sposta i pulsanti cuore e commento verso sinistra */
}

.card-image.is-loaded {
  filter: none; /* remove the blur on fullres image */
  transition: filter 1s;
}
.custom-button {
  background-image: url("heart.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 30px; /* Imposta la larghezza */
  height: 30px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  background-color: rgba(255, 255, 255, 0);
  margin-right: auto;
}
.custom-buttonLiked {
  background-image: url("heart_filled.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 30px; /* Imposta la larghezza */
  height: 30px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  background-color: rgba(255, 255, 255, 0);
  margin-right: auto;
}

.custom-button2 {
  background-image: url("speech-bubble.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 25px; /* Imposta la larghezza */
  height: 25px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  background-color: rgba(255, 255, 255, 0);
  margin-right: auto;
}

.custom-button3 {
  background-image: url("bin.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 25px; /* Imposta la larghezza */
  height: 25px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  margin-left: auto;
  background-color: rgba(255, 255, 255, 0);
}
.search-button {
  background-image: url("search.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 15px; /* Imposta la larghezza */
  height: 15px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  margin-left: auto;
  background-color: rgba(255, 255, 255, 0);
}
.search-button:hover {
  transform: scale(1.5);
}

.custom-buttonDelC {
  background-image: url("bin.png"); /* URL dell'immagine */
  background-size: cover; /* L'immagine copre l'intero pulsante */
  width: 10px; /* Imposta la larghezza */
  height: 10px; /* Imposta l'altezza */
  border: none; /* Rimuove il bordo */
  cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
  margin-left: auto;
}

.buttonsDynamic {
  display: inline-block;
  padding: 15px 30px;
  background-color: cadetblue;
  color: white;
  text-align: center;
  text-decoration: none;
  border: none;
  border-radius: 5px;
  font-size: 16px;
  transition: background-color 0.3s, transform 0.3s;
}

.buttonsDynamicBan {
  display: inline-block;
  padding: 15px 30px;
  background-color: rgb(231, 97, 73);
  color: white;
  width: 20%;
  text-align: center;
  text-decoration: none;
  border: none;
  border-radius: 5px;
  font-size: 12px;
  transition: background-color 0.3s, transform 0.3s;
}

.buttonsDynamicBan:hover {
  transform: scale(1.05);
}

.buttonsDynamic:hover {
  background-color: #0056b3;
  transform: scale(1.05);
}

.comment-box {
  margin: 10px auto; /* Imposta il margine automatico per centrare orizzontalmente */
  padding: 10px;
  background-color: rgba(255, 255, 255, 0.5);
  border: 1px solid #ddd;
  border-radius: 5px;
  width: 80%; 
}

.comment-box ul {
  list-style-type: none;
  padding: 0;
}

.comment-box li {
  margin: 0 auto 5px; /* Imposta il margine su auto per centrare orizzontalmente */
  padding: 5px;
  background-color: rgba(255, 255, 255, 0.5);
  border: 1px solid #ddd;
  border-radius: 5px;
  width: 80%; /* La larghezza è del 80% del genitore */
  display: block;
}

.comment-box textarea {
  width: 75%;
  height: 75px;
  margin-top: 10px;
  padding: 5px;
  border-radius: 5px;
  background-color: rgba(255, 255, 255, 0.5);
  border: 1px solid #ddd;
}

.custom-buttonC {
  margin-top: 5px;
  margin-bottom: 5px;
  padding: 10px;
  background-color: #007bff;
  color: whitesmoke;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.custom-buttonC:hover {
  background-color: #0056b3;
}

.dropdown {
  position: relative;
  display: inline-block;
  background-color: cadetblue;
  top: 100%; /* Questo allinea perfettamente il dropdown sotto il bottone "Settings" */
  left: 0;
}

.dropdown-content {
  display: block;
  position: relative; /* Posiziona il contenuto del dropdown rispetto a "Settings" */
  top: 100%; /* Questo allinea perfettamente il dropdown sotto il bottone "Settings" */
  left: 0;
  background: cadetblue;
  box-shadow: 0px 8px 16px rgba(0, 0, 0, 0.2);
}

.dropdown-content a {
  color: whitesmoke;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
}

.dropdown-content a:hover {
  transform: scale(1.05);
}

/* Stile per il popup */
.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Sfondo semi-trasparente */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.popup-box {
  background-color: whitesmoke;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  text-align: center;
  width: 300px;
}

.popup-box h1 {
  margin-bottom: 20px;
  color: cadetblue;
}

.popup-box input {
  width: 80%;
  padding: 10px;
  margin-bottom: 20px;
  border-radius: 5px;
  border: 1px solid #ddd;
  color: #0056b3;
}

.popup-box button {
  padding: 10px 20px;
  background-color: cadetblue;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin: 5px;
}

.popup-box button:hover {
  background-color: #0056b3;
}

.search-container {
  position: relative; /* Ensures dropdown is positioned relative to this container */
}
.search-dropdown {
  position: relative;
  background-color: whitesmoke;
  width: 200px;
  max-height: 200px;
  overflow-y: auto;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  z-index: 100;
  top: 100%; /* Position it right below the input */
  left: 0;
}

.search-dropdown ul {
  list-style: none;
  padding: 0;
  margin: 0;
  color: cadetblue;
}

.search-dropdown li {
  padding: 10px;
  cursor: pointer;
}

.search-dropdown li:hover {
  background-color: #f0f0f0;
}
</style>