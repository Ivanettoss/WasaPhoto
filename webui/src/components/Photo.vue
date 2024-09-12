<script>

export default{
    data: function(){
        return{
        bytes:"",
        nLikes:0 ,
        nComments:0,
        dataUpload:"",
        errormsg:""
        }
    },

    props: ['owner','nLikes','nComments',"photo_id","isOwner", "dataUpload",], 

    methods:{
        async loadPhoto(){
            try{
            let photo=await this.$axios.get("/photo/" + this.photoId)
            console.log(this.$route.params.photo_id)
            this.bytes=photo.data.photobytes
            this.nLikes=photo.data.nlikes
            this.nComments=photo.data.ncomments
            this.dataUpload=photo.data.uploaddatatime
            } catch(e){
                 this.errormsg = e.toString()
            }

        }

    }
    
     
    }







</script>