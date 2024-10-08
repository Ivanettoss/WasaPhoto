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

      showComments: false, // Controlla se il box dei commenti è visibile
      comments: [], // Lista dei commenti
      newComment: "",

      showSettings: false,

      showPopup: false, // Controlla la visibilità del popup
      newUsername: '', // Nuovo nome utente

      searchQuery: '',
      users: [],

      myself: false
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
        this.myself = true;
      }
      return this.myself;
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
  
  }
};
</script>

<template>
  <header class="header">
    <a class="logo">WasaPhoto</a>
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
</template>
