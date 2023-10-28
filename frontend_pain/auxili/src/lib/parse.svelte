<script context="module" lang="js">
    import axios from 'axios';
    import {browser} from "$app/environment";
    export function prse(dat) {
        let jso = JSON.parse(JSON.stringify(dat));
        return jso.data.data.Data
    }
    export function prsse(dat) {
        let jso = JSON.parse(JSON.stringify(dat));
        return jso.data.data
    }

    export function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    //sig syscall

    export function checkErr(err) {
        if (err.response.status == 401 && browser) {
            alert("You cannot invite anyone. You don't own this organization!")
            window.location.href = "http://localhost:5173/login";
        }
    }



    export function auth() {
        axios.get('http://localhost:8080/app/au th')
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            if (error.response.status == 401) {
                window.location.href = "localhost:8080/app/login";
            }
        });
    }

    function aaa() {
        return axios.get('http://localhost:8080/access/users/token', {
            headers: {
                "Token": getToken()
            }
        })
    }

    export function getTok() {
        return axios.get('http://localhost:8080/access/users/token', {
            headers: {
                "Token": getToken(),
                "Content-Type": "application/json",
            }
        })


    }
</script>