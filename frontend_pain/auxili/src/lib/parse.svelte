<script context="module" lang="js">
    import axios from 'axios';
    import {browser} from "$app/environment";
    export function prse(dat) {
        let jso = JSON.parse(JSON.stringify(dat));
        return jso.data.data.Data
    }

    export function getToken() {
        if (browser) {
            return window.localStorage.getItem("token");

        }
    }

    export function auth() {
        axios.get('http://localhost:8080/app/auth')
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
                "Token": getToken()
            }
        })


    }
</script>