<script>
  export default {
    data: function () {
      return {
        errormsg: null,
        username: "",
      }
    },
    methods: {
      async login(){
        
        if (this.username==""){
          this.errormsg= "empty username is not allowed"
          
        }
        else{
 
            try{
            let response = await this.$axios.post("/session", {username: this.username.trim()});
            let user = response.data 

            localStorage.setItem('username', user.username);
            localStorage.setItem('token', user.id);

            this.$router.push({path:"/profile/"+user.username})
            

          }
        catch (e) {
            this.errormsg = e.toString();
          }
        }
		           },
             }
            }

</script>

<template>
    <head> 
        <title>Log in </title>
    </head>

    <div class="blurred-box">
       <div class="user-login-box">
        <span class="user-icon"></span>
        <div class="user-name">WasaPhoto</div>
    
    <input class="user-password" v-model="username" name="username" placeholder="username">
   <button  class="btn-hover color-9" id="bottone"  value="Login" @click="login">Login</button>
  </div>
</div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

</template>

<style>

/* CSS which you need for blurred box */
body {
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: top;
  background-image: url("wasaback.jpeg");
  font-family: Arial, Helvetica;
  letter-spacing: 0.02em;
  font-weight: 400;
  -webkit-font-smoothing: antialiased;
  
  display: flex; /* Aggiungi questa proprietà */
  justify-content: center; /* Centra orizzontalmente */
  align-items: center; /* Centra verticalmente */
  min-height: 100vh; /* Altezza minima pari alla viewport */
}

.blurred-box {
  position: relative;
  width: 250px;
  height: 350px;
  background: inherit;
  border-radius: 2px;
  overflow: hidden;
}

.blurred-box:after {
  content: '';
  width: 300px;
  height: 300px;
  background: inherit;
  position: absolute;
  left: -25px;
  right: 0;
  top: -25px;
  bottom: 0;
  box-shadow: inset 0 0 0 200px rgba(255, 255, 255, 0.05);
  filter: blur(10px);
}

/* Altri stili rimangono invariati */



/* Form which you dont need */
.user-login-box{
  position: relative;
  margin-top: 50px;
  text-align: center;
  z-index: 1;
}
.user-login-box > *{
  display: inline-block;
  width: 200px;
}

.user-icon{
  width: 100px;
  height: 100px;
  position: relative;
  border-radius: 50%;
  background-size: contain;
  background-image:url("wasalogo.jpeg");
}

.user-name{
  margin-top: 15px;
  margin-bottom: 15px;
  color: white;
}

input.user-password{
  width: 120px;
  height: 18px;
  opacity: 0.4;
  border-radius: 2px;
  padding: 5px 15px;
  border: 0;
}

.btn-hover.color-9 {
    background-image: linear-gradient(to right, #25aae1, #4481eb, #04befe, #3f86ed);
    box-shadow: 0 4px 15px 0 rgba(65, 132, 234, 0.75);
}

.buttons {
    margin: 5%;
    text-align: center;
    width: 2px
      
}

.btn-hover {
    width: 200px;
    font-size: 16px;
    font-weight: 600;
    color: #fff;
    cursor: pointer;
    margin: 20px;
    height: 55px;
    text-align:center;
    border: none;
    background-size: 300% 100%;

    border-radius: 50px;
    moz-transition: all .4s ease-in-out;
    -o-transition: all .4s ease-in-out;
    -webkit-transition: all .4s ease-in-out;
    transition: all .4s ease-in-out;
}

.btn-hover:hover {
    background-position: 100% 0;
    moz-transition: all .4s ease-in-out;
    -o-transition: all .4s ease-in-out;
    -webkit-transition: all .4s ease-in-out;
    transition: all .4s ease-in-out;
}

.btn-hover:focus {
    outline: none;
}
</style>
    


  