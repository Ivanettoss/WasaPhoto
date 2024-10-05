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
                selectedFile:null,

                showComments: false,      // Controlla se il box dei commenti è visibile
                comments: [],             // Lista dei commenti
                newComment: "",

                showSettings: false,

                showPopup: false, // Controlla la visibilità del popup
                newUsername: ''   // Nuovo nome utente

            }

        },
        methods:{
        async buildProfile(){
            try{
                let response = await this.$axios.get("/searchuser/" + this.$route.params.username)
                console.log(response)
                console.log("ciao, guarda photos", response.data.photos)
                this.username=response.data.username
                this.photos=[]
                for (let i = 0; i < response.data.photos.length; i++){
                    console.log(response.data.photos[i])
                    this.photos.push(response.data.photos[i])   
                }
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


        async doLike(){
            try{

            }
            catch(e){
                this.errormsg = e.toString();
            }


        },
        async doUnLike(){
            try{

            }
            catch(e){
                this.errormsg = e.toString();
            }


        },async searchUsers() {
          

        },



        async doLogOut() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({path: '/'})
		},

        async mySelf(){
            let myself=false
            if (this.username == localStorage.getItem('username')){
               this.myself=true
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


        toggleComments() {
      this.showComments = !this.showComments;
    },

    settingsDropdown() {
      this.showSettings = !this.showSettings;
    },
    async changeUsername() {
                if (this.newUsername) {
                  try{
                    // Aggiungi la logica per cambiare il nome utente, ad esempio:
                    let response= await this.$axios.put("/user/" +this.localUser +"/set_username",{ "username": this.newUsername},
                    { headers: {
            Authorization: this.token
          }})
                    console.log(response)
                    this.username=response.data.username
                    this.showPopup = false; // Chiude il popup dopo il salvataggio
                    this.newUsername = ''; // Resetta il campo di input
                    localStorage.setItem('username', user.username);
                    this.$router.push({path: '/profile/'+this.username})
                } catch (e)
                {
                   this.errormsg = e.toString();
                }
                }},

    addComment() {
      if (this.newComment.trim()) {
        this.comments.push({
          id: Date.now(),   // ID univoco per ogni commento
          text: this.newComment,
        });
        this.newComment = ""; // Resetta il campo dopo l'invio del commento
      }},

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
        let response = await this.$axios.post( "/user/"+this.localUser+"/upload",{ "photobytes": this.selectedFile },{
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
    async deletePic(idPhotoToDelete){
    try{
    console.log(idPhotoToDelete)
    let responde=await this.$axios.delete( "/user/"+this.localUser+"/photo/"+ idPhotoToDelete,{
          headers: {
            Authorization: this.token
          }
        });
        console.log("delete eseguito")
        this.buildProfile()
    }catch(e){
    this.errormsg = e.toString();
  }},

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
      <button class="search-button" @click="searchusers()"></button>
        <a  id="menubu">Home</a>
        <a  id="menubu">Profile</a>
        <a  id="menubu" @click="settingsDropdown">Settings</a>

        <div class="dropdown">
         <div v-if="showSettings" class="dropdown-content">

        <a  id="menubu" @click="showPopup=true">Change Username</a>
        <a   id="menubu" @click="doLogOut()">Logout</a>
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
	
	<li class="card" v-for="photo in this.photos" :key="photo.idphoto">
		<div class="card-image"  >
			 <img :src="photo.photobytes">
		</div>
		<div class="card-description"  target="_blank">
      {{ photo.nlikes }}
			<button class="custom-button"></button>
      {{ photo.ncomments }}
      <button class="custom-button2" @click="toggleComments()"></button>

      <button class="custom-button3" @click="deletePic( photo.idphoto)"></button>
               
		</div>

        <!-- Box dei commenti nascosto sotto la foto -->
      <div v-if="showComments" class="comment-box">
        <ul>
          <li v-for="comment in comments" :key="comment.id">
           autore: {{ comment.text }}
          </li>
        </ul>
        <textarea v-model="newComment" placeholder="write a comment..."></textarea>
        <button  @click="addComment">Comment</button>
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
.navbar a:hover {
  transform: scale(1.05);
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
  width: 100%;
  height: 300px; /* Imposta un'altezza fissa per la carta */
  overflow: hidden; /* Nasconde le parti in eccesso dell'immagine */
  position: relative;
}

  .card-image img {
    width: 100%;
    height: 100%;
    object-fit: cover; /* L'immagine si adatta al contenitore mantenendo le proporzioni */
}

.lNc{
  display: flex; /* Mantiene i pulsanti su una linea orizzontale */
    gap: 10px; /* Spazio tra il cuore e il commento */
    justify-content: flex-start; /* Sposta i pulsanti cuore e commento verso sinistra */
}






.card-image.is-loaded {
	filter: none; /* remove the blur on fullres image */
	transition: filter 1s;
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
 display: flex;
  justify-content: space-between; /* Distribuisce i pulsanti su entrambi i lati */
  align-items: center;
    padding: 15px 20px 5px;

	
}


.custom-button{
    background-image: url("heart.png"); /* URL dell'immagine */
    background-size: cover; /* L'immagine copre l'intero pulsante */
    width: 30px; /* Imposta la larghezza */
    height: 30px; /* Imposta l'altezza */
    border: none; /* Rimuove il bordo */
    cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
    background-color: rgba(255, 255, 255, 0);
    margin-right: auto;
}

.custom-button2{
    background-image: url("speech-bubble.png"); /* URL dell'immagine */
    background-size: cover; /* L'immagine copre l'intero pulsante */
    width: 25px; /* Imposta la larghezza */
    height: 25px; /* Imposta l'altezza */
    border: none; /* Rimuove il bordo */
    cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
    background-color: rgba(255, 255, 255, 0);
    margin-right: auto;
}

.custom-button3{
    background-image: url("bin.png"); /* URL dell'immagine */
    background-size: cover; /* L'immagine copre l'intero pulsante */
    width: 25px; /* Imposta la larghezza */
    height: 25px; /* Imposta l'altezza */
    border: none; /* Rimuove il bordo */
    cursor: pointer; /* Cambia il cursore quando passa sopra il pulsante */
    margin-left: auto;
    background-color: rgba(255, 255, 255, 0);
}
.search-button{
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
     transform: scale(1.5)
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

.comment-box {
  margin-top: 10px;
  padding: 10px;
  background-color: #f1f1f1;
  border: 1px solid #ddd;
  border-radius: 5px;
  width: 60%;
  height: 60%;
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

.comment-box button {
  margin-top: 5px;
  margin-bottom: 5px;
  padding: 10px;
  background-color: #007bff;
  color: whitesmoke;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.comment-box button:hover {
  background-color: #0056b3;
}


.dropdown {
  position: relative;
  display: inline-block;
  background-color:cadetblue ;
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
  z-index: 1;
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
    color:   #0056b3;
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



</style>