<script>
    import axios from 'axios';
    import { browser } from '$app/environment';
    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    function logOut() {
        if (browser) {
            window.localStorage.removeItem("token");
            window.location.href = "/login";
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
                if (response.status != 201 && response.status != 200) {
                    console.log(response.status)
                    alert("You are not logged in")//will it show

                }
                console.log(response);


            }).catch(function (error) {
            console.log(error);
            if (browser) {
                alert("You are not logged in")//will it show
                window.location.href = "/login";
            }

        });


</script>

<div class="drawer drawer-end">
    <input id="opensettings" type="checkbox" class="drawer-toggle">
    <div class="drawer-content">
        <div class="navbar bg-base-200 rounded-b-box">
            <div class="flex-1">
                <a class="btn btn-ghost btn-sm rounded-btn btn btn-md border-accent">restapi</a>
            </div>

            <div class="flex-none">
                <label for="opensettings" class="btn btn-square btn-ghost border-accent">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-8 h-8 stroke-current"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
                </label>
            </div>
        </div>
    </div>
    <div class="drawer-side">
        <label for="opensettings" class="drawer-overlay"></label>
        <ul class="menu p-4 w-80 bg-base-100 text-base-content text-accent bg-secondary rounded-l-box border-accent">
            <li>
                <a href="/mainPg">Home</a>
            </li>
            <li>
                <a href="/login">Login</a>
            </li>
            <li>
                <a href="/register">Register</a>
            </li>
            <li>
                <button on:click={logOut}><a>Logout</a></button>
            </li>
        </ul>
    </div>
</div>