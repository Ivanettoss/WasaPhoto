<script>
    export default{
        data: function(){
            return {
                errormsg:null,
                username:"",
                token: localStorage.getItem("token"),
                localUser: localStorage.getItem("username"),
                nFollowers:0,
                nFollowing:0,
                followState:false, 
                banState: false,
                photos:[],
                nPost:0,
                selectedFile:null

            }

        },
        methods:{
        async buildProfile(){
            try{
                console.log(this.$route.params.username)
                let response = await this.$axios.get("/searchuser/" + this.$route.params.username)
                this.username=response.data.username
                this.photos=response.data.photos
                this.nFollowers=response.data.nfollower
                this.nFollowing=response.data.nfollowed
                this.nPost=response.data.npost

                if (this.username != localStorage.getItem("username"))
                {
                    this.followState=response.data.followstate
                    this.banState=response.data.banstate
                }


            } 
            catch(e){
                 this.errormsg = e.toString();
            }
        },
        async Logout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({path: '/'})
		},

        async mySelf(){
            let myself=false
            if (this.username == localStorage.getItem('username')){
               this.myself=true
               console.log(this.username)
               console.log(localStorage.getItem('username'))
            }
            return this.myself

        },

        async doFollow(){
             let response = await this.$axios.put("/user/" +this.localUser +"/followed/"+this.$route.params.username)
             this.followState=True
             this.nFollowers+=1
        },
        async unFollow(){
             let response = await this.$axios.delete("/user/"+this.localUser +"/followed/"+this.$route.params.username)
             this.followState=false
             this.nFollowers-=1
        },

        handleButtonClick() {
      if (!this.selectedFile) {
        // Se non c'è nessun file selezionato, apri il file picker
        this.$refs.fileInput.click();
        console.log("ho pickato la foto")
      } else {
        // Se c'è un file selezionato, caricalo
        this.uploadPic();
      }
    },
       handleFileSelect(event) {
      const file = event.target.files[0];
      console.log("sto nell'handler")
      if (file) {
        this.selectedFile = file;
        console.log("selected file:", file.name);
      }
    },
    
    async uploadPic() {
      if (!this.selectedFile) {
        console.error("Nessun file selezioato.");
        this.errormsg = "Please select a file to upload."
        return;
      }

      const formData = new FormData();
      formData.append("file", this.selectedFile);

      try {
        let response = await this.$axios.post( "/user/"+this.localUser+"/upload",formData);
        this.nPost += 1;
        console.log("Upload completato");
        this.selectedFile = null; // Resetta lo stato dopo l'upload
      } catch (error) {
        console.error("Errore durante l'upload:", error);
      }
    },
  },
     mounted(){
            this.buildProfile()
        }
    
    }
</script>

<template>

<header class="header">
    <a class="logo">WasaPhoto</a>
    <nav class="navbar">
        <input id="searchForm" type="text"
        placeholder="Search users..."
        v-model="searchQuery" 
        @keyup.enter="searchUsers"  
      />
        <a id="menubu">Home</a>
        <a id="menubu">Profile</a>
        <a id="menubu">Settings</a>
    </nav>
</header>
 <div class="profile-container">
        <div class="profile-picture">
            <img src="https://www.shareicon.net/data/512x512/2016/09/15/829452_user_512x512.png" alt="Foto profilo">
        </div>
        <div class="profile-info">
            <div class="username">{{ username }}
            <input type="file" accept="image/*"  style="display: none" ref="fileInput" @change="handleFileSelect">
            <button v-if="mySelf()"  class="buttonsDynamic" @click="handleButtonClick">  {{ selectedFile ? 'Upload' : 'Choose an image' }} </button>
            <button v-else-if="followState" class="buttonsDynamic" @click="unFollow()"> unfollow</button>
            <button v-else-if="!followState" class="buttonsDynamic" @click="doFollow()"> Follow</button>
            </div>
            
            
            <div class="stats">
                <div class="stat-item">
                    <span>{{ nPost }}</span>
                    Post
                </div>
                <div class="stat-item">
                    <span>{{ nFollowing }}</span>
                    Seguiti
                </div>
                <div class="stat-item">
                    <span>{{ nFollowers }}</span>
                    Seguaci
                </div>
            </div>
        </div>
    </div>
    <ul class="card-list">
	
	<li class="card">
		<div class="card-image"  >
			<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/82/London_Big_Ben_Phone_box.jpg/220px-London_Big_Ben_Phone_box.jpg" >
		</div>
		<div class="card-description"  target="_blank">
			
			<button class="custom-button"> </button>
            <button class="custom-button2"> </button>
               
		</div>
	</li>
    </ul>

</template>

<style>
.view {
    height:100%;
}

.header{
    position:fixed;
    top:0;
    left:0;
    width:100%;
    padding: 20px 100px;
    background: cadetblue;
    display: flex;
    justify-content: space-between;
    align-items: center;
    z-index: 100;
}

.logo{
    font-size: 22px;
    color:whitesmoke;
    text-decoration: none;
    font-weight:500;
}

.navbar a,p{
    font-size :15px;
    font-weight: 300;
    margin-left: 50px;
    margin-right:30px;
}
#searchForm{
    background-color:transparent;
    border: none;
    border-bottom: 2px solid whitesmoke;
    outline: none;
    color:whitesmoke
}
#menubu{
    color:whitesmoke
}

        .profile-container {
            background-color : rgba(255, 255, 255, 0.5);
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
            margin:auto 
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
            margin:10px;
        }
        .username {
            font-size: 36px;
            font-weight: bold;
            margin-bottom: 20px;
            color:cadetblue
        }
        .stats {
            display: flex;
            justify-content: space-between;
            width: 400px;

        }
        .stat-item {
            text-align: center;
            color:cadetblue
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
	background-color : rgba(255, 255, 255, 0.5);
	background-size: cover;
    display: block;
	width: 100%;
}



.card-image.is-loaded {
	filter: none; /* remove the blur on fullres image */
	transition: filter 1s;
}




/* Layout Styles */


.card-list {
	display: block;
	margin: 1rem auto;
	padding: 0;
	font-size: 0;
	text-align: center;
	list-style: none;
}

.card {
    background-color : rgba(255, 255, 255, 0.5);
	display: inline-block;
	width: 90%;
	max-width: 20rem;
	margin: 1rem;
	font-size: 1rem;
	text-decoration: none;
	overflow: hidden;
	box-shadow: 0 0 3rem -1rem rgba(0,0,0,0.5);
	transition: transform 0.1s ease-in-out, box-shadow 0.1s;
}

.card:hover {
	transform: translateY(-0.5rem) scale(1.0125);
	box-shadow: 0 0.5em 3rem -1rem rgba(0,0,0,0.5);
}

.card-description {
	display: block;
	padding: 1em 0.5em;
	
}


.custom-button{
    background-image: url("heart.png"); /* URL dell'immagine */
    background-size: cover; /* L'immagine copre l'intero pulsante */
    width: 40px; /* Imposta la larghezza */
    height: 40px; /* Imposta l'altezza */
    border: none; /* Rimuove il bordo */
    cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
    background-color: rgba(255, 255, 255, 0);
}

.custom-button2{
    background-image: url("speech-bubble.png"); /* URL dell'immagine */
    background-size: cover; /* L'immagine copre l'intero pulsante */
    width: 30px; /* Imposta la larghezza */
    height: 30px; /* Imposta l'altezza */
    border: none; /* Rimuove il bordo */
    cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
    background-color: rgba(255, 255, 255, 0);
}
.custom-button:hover {
    background-color: rgba(255, 255, 255, 0.5)
	
}

.buttonsDynamic{
            display: inline-block;
            padding: 15px 30px;
            background-color: cadetblue;
            color:white;
            text-align: center;
            text-decoration: none;
            border: none;
            border-radius: 5px;
            font-size: 16px;
            transition: background-color 0.3s, transform 0.3s;
        }

.buttonsDynamic:hover {
            background-color: #0056b3;
            transform: scale(1.05);
        };


</style>