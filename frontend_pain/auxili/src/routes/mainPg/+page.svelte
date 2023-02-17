<script>
    import axios from 'axios';
    import { browser } from '$app/environment';
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    //cpidl add of sttet,emt nroser but then would ahe to add if not browser then rhrow erero


        axios.get('http://localhost:8080/app/auth', {

            headers: {
                "Token": getToken(), //this isnt returning null or underfined
                "Content-Type": "application/json"
            }
        })
            .then(function (response) {
                if (response.status != 201 || response.status != 200) {

                    alert("You are not logged in")//will it show

                }
                console.log(response);


            }).catch(function (error) {
            console.log(error);
        });


</script>
