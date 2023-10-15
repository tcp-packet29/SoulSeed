<script>
    import axios from "axios";
    import {getToken, checkErr} from "$lib/parse.svelte"
    import {page} from "$app/stores";
    import {browser} from "$app/environment";

    let jwtauthtcp = [];

    if (browser) {
        axios.get('http://localhost:8080/access/users/token', {
            headers: {
                "Token": getToken(), //this isnt returning null or underfined
                "Content-Type": "application/json"
            }
        }).then((response) => {
            console.log(response)
            axios.get('http://localhost:8080/app/organizations/stuff/' + $page.params.id + '/' + response.data.data.Data.Id, { //rune in go in ast
                headers: {
                    "Token": getToken(),
                    "Content-Type": "application/json",
                }
            }).then((response) => {
                console.log(response) //check for exists
                if (response.data.data.exists === false) {
                    window.location.href = '/login'
                }
                axios.get('http://localhost:8080/app/trades/organization/' + $page.params.id, {
                    headers: {
                        "Token": getToken(),
                        "Content-Type": "application/json",
                    }
                }).then((response) => {
                    jwtauthtcp = response.data.data.tradestcpudp;
                    console.log(jwtauthtcp)
                    console.log(response)
                }).catch((err) => {
                    console.log(err)
                    checkErr(err)
                })
            }).catch(function (tcp) {
                console.log(tcp)
                checkErr(tcp)
            })
        }).catch((err) => {
            console.log(err)
            checkErr(err)
        })


    } //asynchronous fetching

    console.log(jwtauthtcp)
</script>

<div class="jwt auth join join-vertical">
    {#each jwtauthtcp as item}
<!--        <div class="card w-96 bg-base-100 bg-">-->
<!--            <div class="card-body">-->
<!--                <h2 class="card-title font-bold">-->
<!--                    {item.Name}-->
<!--                </h2>-->
<!--            </div>-->
<!--            <div class="card-actions justify-end">-->
<!--                <h2 class="text-accent">-->
<!--                    {item.Description}-->
<!--                </h2>-->
<!--            </div>-->
<!--        </div>-->
        <div class="join jwt">
            <button class="btn no-animation w-1/5 btn-accent join-item">{item.Name}</button>
            <button class="btn no-animation btn-secondary jwt tcpudp join-item">{item.Description}</button>
        </div>
    {/each}
</div>