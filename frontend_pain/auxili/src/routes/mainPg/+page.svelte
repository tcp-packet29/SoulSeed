<script>
    import axios from 'axios';
    import { browser } from '$app/environment';
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    //cpidl add of sttet,emt nroser but then would ahe to add if not browser then rhrow erero
    if (browser) {

        axios.get('http://localhost:8080/app/auth', {
            headers: {
                "Token": getToken()
            }
        })
            .then(function (response) {
                if (response.status != 201 || response.status != 200) {
                    window.location.href = "http://localhost:8080/login"
                    alert("You are not logged in")//will it show

                }
                console.log(response);


            }).catch(function (error) {
            console.log(error);
        });
    }

</script>
