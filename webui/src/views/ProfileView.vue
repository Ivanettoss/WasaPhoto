<script>
    export default{
        data: function(){
            return {
                errormsg:null,
                username:"",
                token: localStorage.getItem("token"),
                nFollowers:0,
                nFollowing:0,
                followState:false, 
                banState: false,
                photos:[],
                nPost:0

            }

        },
        async buildProfile(){
            try{
                let response = await this.$axios.get("/searchuser/" + this.$route.params.uname)
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
        }
    
    }
</script>

<template>
<!--Navbar-->
<nav class="navbar navbar-fixed navbar-dark info-color">

    <!-- Collapse button-->
    <button class="navbar-toggler hidden-sm-up" type="button" data-toggle="collapse" data-target="#collapseEx2">
        <i class="fa fa-bars"></i>
    </button>

    <div class="container">

        <!--Collapse content-->
        <div class="collapse navbar-toggleable-xs" id="collapseEx2">
            <!--Navbar Brand-->
            <a class="navbar-brand">WasaPhoto</a>
            <!--Links-->
            <ul class="nav navbar-nav">
                <li class="nav-item active">
                    <a class="nav-link">Home <span class="sr-only">(current)</span></a>
                </li>
                        
                <li class="nav-item btn-group">
                    <a class="nav-link dropdown-toggle" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Settings</a>
                    <div class="dropdown-menu" aria-labelledby="dropdownMenu1">
                        <a class="dropdown-item">LogOut</a>
                    </div>
                </li>
            </ul>
            
            <ul class="nav navbar-nav nav-flex-icons">
                <li class="nav-item"><a href="#!" class="nav-link"><i class="fa fa-search" aria-hidden="true"></i></a></li>
                <li class="nav-item"><a href="#!" class="nav-link"><i class="fa fa-ellipsis-v" aria-hidden="true"></i></a></li>
            </ul>
        </div>
        <!--/.Collapse content-->

    </div>

</nav>
<!--/.Navbar-->

<main>
    <div class="container">
        <div class="row m-b-r m-t-3">
            <div class="col-md-2 offset-md-1">
                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQb-NGEQDekk2BwsllLjk4tcIM_BPIzXECdsg&s" alt="" class="img-circle img-fluid">
            </div>
            <div class="col-md-9 p-t-2">
                <h2 class="h2-responsive">{{ username }}

                    <button v-if="this.username != localStorage.getItem('username')" type="button" class="btn btn-info-outline waves-effect">Follow</button>
                    <button v-else type="button" class="btn btn-info-outline waves-effect">upload</button>

                </h2>

                <ul class="flex-menu">
                    <li><strong>{{ nPost }}</strong> Posts</li>
                    <li><strong>{{ nFollowers }}</strong> Followers</li>
                    <li><strong>{{ nFollowing }}</strong> Following</li>
                </ul>
            </div>
        </div>
   </div> 
        <div class="row">
            <div class="col-sm-12 col-md-6 col-lg-4">
                <div class="view overlay hm-black-light m-b-2">
                    <img src="https://mdbootstrap.com/images/regular/nature/img%20(1).jpg" class="img-fluid " alt="">
                    <div class="mask flex-center">
                        <ul class="flex-menu">
                            <li><p class="white-text"><i class="fa fa-comment" aria-hidden="true"></i> 32</p></li>
                            <li><p class="white-text"><i class="fa fa-heart" aria-hidden="true"></i> 1.2K</p></li>
                        </ul>
                    </div>
                </div>
            </div>
          </div> 
</main>
</template>

<style>
.view {
    height:100%;
}

.navbar-brand {
    color:#fff !important;
}

.profile-photo {
    width:48px;
    height:auto;
}

.flex-menu {
    display:flex;
}

.flex-menu li:not(:last-child) {
    margin-right:40px;
}
</style>
