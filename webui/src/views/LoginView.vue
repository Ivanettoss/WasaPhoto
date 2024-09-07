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
          errormsg= "empty username is not allowed"
        }
        else{

            try{
            let response = await this.$axios.post("/session", {username: this.username.trim()});
            idToken= response.data.success.split('id: ')[1]

            localStorage.setItem('username', this.username);
            localStorage.setItem('token', token);

            this.$router.replace("/home")

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
        <link href='log-instyle.css' rel='stylesheet' type='text/css'>
        <script src="https://code.jquery.com/jquery-3.5.0.js"></script>
       <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"> </script>
      
    </head>
    <video autoplay muted loop id="myVideo">
      <source src="backlogin.mp4" type="video/mp4">
    </video>
      <div id="login-button">
        <img src="WasaPhotoLogo.png"> </img>
      </div>


    <div id="container">
    <h1>Log-in </h1>
    <span class="close-btn">
      <img src="https://cdn4.iconfinder.com/data/icons/miu/22/circle_close_delete_-128.png"></img>
    </span>
  
     <script src="main.js"></script>
    <form action="logincontrol.php" method="POST">
      <input type="username" name="username" placeholder="username">
      <button id="bottone"  value="Login" onclick="login">Login</button>
     </form>
  </div>
</template>

<style>
 img{
    display: block;
    margin: auto;
    width: 100%;
    height: auto;
  }
  
  #login-button{
    cursor: pointer;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 30px;
    margin: auto;
    width: 200px;
    height: 200px;
    border-radius: 50%;
    background: rgba(3,3,3,.8);
    overflow: hidden;
    opacity: 70%;
    box-shadow: 10px 10px 30px #000;
}
  
  /* Login container */
  #container{
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin: auto;
    width: 260px;
    height: 260px;
    border-radius: 5px;
    background: rgba(3,3,3,0.25);
    box-shadow: 1px 1px 50px #000;
    display: none;
  }
  
  .close-btn{
    position: absolute;
    cursor: pointer;
    font-family: 'Open Sans Condensed', sans-serif;
    line-height: 18px;
    top: 3px;
    right: 3px;
    width: 20px;
    height: 20px;
    text-align: center;
    border-radius: 10px;
    opacity: .2;
    -webkit-transition: all 2s ease-in-out;
    -moz-transition: all 2s ease-in-out;
    -o-transition: all 2s ease-in-out;
    transition: all 0.2s ease-in-out;
  }
  
  .close-btn:hover{
    opacity: .5;
  }
  
  /* Heading */
  h1{
    font-family: 'Open Sans Condensed', sans-serif;
    position: relative;
    margin-top: 0px;
    text-align: center;
    font-size: 40px;
    color: #ddd;
    text-shadow: 3px 3px 10px #000;
  }
  
  /* Inputs */
  input{
    font-family: 'Open Sans Condensed', sans-serif;
    text-decoration: none;
    position: relative;
    width: 80%;
    display: block;
    margin: 9px auto;
    font-size: 17px;
    color: #fff;
    padding: 8px;
    border-radius: 6px;
    border: none;
    background: rgba(3,3,3,.1);
    -webkit-transition: all 2s ease-in-out;
    -moz-transition: all 2s ease-in-out;
    -o-transition: all 2s ease-in-out;
    transition: all 0.2s ease-in-out;
  }
  
  input:focus{
    outline: none;
    box-shadow: 3px 3px 10px #333;
    background: rgba(3,3,3,.18);
  }
  
  /* Placeholders */
  ::-webkit-input-placeholder {
     color: #ddd;  }
  :-moz-placeholder { /* Firefox 18- */
     color: red;  }
  ::-moz-placeholder {  /* Firefox 19+ */
     color: red;  }
  :-ms-input-placeholder {  
     color: #333;  }
  
  /* Link */
  #bottone{
    font-family: 'Open Sans Condensed', sans-serif;
    text-align: center;
    color: rgb(0,0,128);
    padding: 4px 8px;
    background: rgb(143,206,237);
  }
  
  a:hover{
    opacity: 0.7;
  }
  
  #remember-container{
    position: relative;
    margin: -5px 20px;
  }
  
  .checkbox {
    position: relative;
    cursor: pointer;
      -webkit-appearance: none;
      padding: 5px;
      border-radius: 4px;
    background: rgba(3,3,3,.2);
      display: inline-block;
    width: 16px;
    height: 15px;
  }
  
  .checkbox:checked:active {
      box-shadow: 0 1px 2px rgba(0,0,0,0.05), inset 0px 1px 3px rgba(0,0,0,0.1);
  }
  
  .checkbox:checked {
    background: rgba(3,3,3,.4);
      box-shadow: 0 1px 2px rgba(0,0,0,0.05), inset 0px -15px 10px -12px rgba(0,0,0,0.05), inset 15px 10px -12px rgba(255,255,255,0.5);
      color: #fff;
  }
  
  .checkbox:checked:after {
      content: '\2714';
      font-size: 10px;
      position: absolute;
      top: 0px;
      left: 4px;
      color: #fff;
  }
  
  #remember{
    position: absolute;
    font-size: 13px;
    font-family: 'Hind', sans-serif;
    color: rgba(255,255,255,.5);
    top: 7px;
    left: 20px;
  }
  
  #forgotten{
    position: absolute;
    font-size: 12px;
    font-family: 'Hind', sans-serif;
    color: rgba(255,255,255,.2);
    right: 0px;
    top: 8px;
    cursor: pointer;
    -webkit-transition: all 2s ease-in-out;
    -moz-transition: all 2s ease-in-out;
    -o-transition: all 2s ease-in-out;
    transition: all 0.2s ease-in-out;
  }
  
  #forgotten:hover{
    color: rgba(255,255,255,.6);
  }
  
  #forgotten-container{
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin: auto;
    width: 260px;
    height: 180px;
    border-radius: 10px;
    background: rgba(3,3,3,0.25);
    box-shadow: 1px 1px 50px #000;
    display: none;
  }
  
  .orange-btn{
    background: rgba(87,198,255,.5);
  }


  #myVideo {

    object-fit: cover;
    width: 100vw;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
  }
  
  /* Add some content at the bottom of the video/page */
  .content {
    position: fixed;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    color: #f1f1f1;
    width: 100%;
    padding: 20px;
  }
  /* se necessitiamo di aggiungere il pulsante per gestire il video di sfondo 
  /* Style the button used to pause/play the video */
  #myBtn {
    width: 200px;
    font-size: 18px;
    padding: 10px;
    border: none;
    background: #000;
    color: #fff;
    cursor: pointer;
  }
  
  #myBtn:hover {
    background: #ddd;
    color: black;
  }
</style>
    


  