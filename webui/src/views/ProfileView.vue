<script>
    export default{
        data: function(){
            return {
                errormsg:"",
                username: localStorage.getItem("token"),
                token: localStorage.getItem("token"),
                nFollowers:""
                nFollowing:""
                followState:false 
                banState: false 

            }

        }
        async buildProfile(){
            try{
                let response = await this.$axios.get("/user/" + this.$route.params.uname, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});
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
                        <a class="dropdown-item">Change password</a>
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
                <h2 class="h2-responsive">Ivanettos <button type="button" class="btn btn-info-outline waves-effect">Follow</button></h2>

                <ul class="flex-menu">
                    <li><strong>41</strong> posts</li>
                    <li><strong>47k</strong> followers</li>
                    <li><strong>208</strong> following</li>
                </ul>
            </div>
        </div>

        <div class="row">
            <div class="col-sm-12 col-md-6 col-lg-4">
                <div class="view overlay hm-black-light m-b-2">
                    <img src="https://mdbootstrap.com/images/regular/nature/img%20(1).jpg" class="img-fluid " alt="">
                    <div class="mask flex-center">
                        <ul class="flex-menu">
                            <li><p class="white-text"><i class="fa fa-comment" aria-hidden="true"></i> 32</p></li>
                            <li><p class="white-text"><i class="fa fa-heart" aria-hidden="true"></i> 1.2K</i></p></li>
                        </ul>
                    </div>
                </div>
            </div>
          </div> 
</main>
</template>

<style>
body {
    background:#eee;
}

html,
body,
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
