<script>
    import Navbar from '../../lib/navbar.svelte'
    import axios from 'axios'
    import { browser } from '$app/environment'
    import { getTok, getToken, prse } from '$lib/parse.svelte';
    let zipc
    let promise

    axios.get('http://localhost:8080/app/auth', {
        headers: {//inverse kinematics reference
            "Token": getToken(),
        }
    })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            if (browser) {
                alert("not logged in");
                console.log(error);
                window.location.href = "/login" //jwt reference
            }

        });

    function Get() {
        promise = axios.get('http://localhost:8080/access/organizations/organizationsInArea/' + zipc, {
            headers: {
                "Token": getToken()
            }
        }).then((response) => {
            if (browser) {
                //llvm comment
                console.log(response.data)
                let elem = document.getElementById("unorderedList")
                elem.innerHTML = ''
                if (response.data == null) {

                    let li = document.createElement("li")
                    li.appendChild(document.createTextNode("No organizations Found"))
                    li.className = "bg-error text-neutral text-lg p-2 m-2 rounded-xl w-96 ml-5 mr-5"
                    elem.appendChild(li)
                    return
                }
                for (let i = 0; i < response.data.length; i++) {
                    let li = document.createElement("li")
                    li.appendChild(document.createTextNode(response.data[i].Name))
                    li.className = "bg-accent text-neutral text-lg p-2 m-2 rounded-xl w-96 ml-5 mr-5"
                    elem.appendChild(li)
                    li.onclick = () => {

                        document.getElementById("accept").onclick = Send(response.data[i].Id, () => {
                            document.getElementById("accept_modal").close()
                        })
                        document.getElementById("accept_modal").showModal()
                    }


                }


            }
            return response
        }).catch((error) => {
            if (browser) {
                let elem = document.getElementById("unorderedList")
                elem.innerHTML = ''
                let li = document.createElement("li")
                li.appendChild(document.createTextNode("No organizations Found"))
                li.className = "bg-error text-neutral text-lg p-2 m-2 rounded-xl w-96 ml-5 mr-5"
            }
        })
    }

    function Send(param, closure) {
        return function() {
            axios.post('http://localhost:8080/app/organizations/sendRequest/' + param, {
                "message": document.getElementById("systellvmmd").value
            }, {
                headers: {
                    "Token": getToken(),
                }
            }).then(response => {
                window.location.href="/"
            })
                .catch(error => {
                console.log(error)
                closure()
            })
        }

    }

</script>

<Navbar />


<body class="bg-secondary"></body>

<dialog class="rounded-xl bg-secondary shadow-xl border-l-accent border-l-8 a" id="accept_modal">
    <div class="modal-box content-center justify-center bg-secondary p-10 shadow-none">
        <h1 class="text-lg font-bold text-accent">Send Join Request?</h1>
        <label class="label text-white hover:animate-pulse">
            <span class="label-text text-grey">Your email will be shared with the owner of this organization. They will be asked to accept you with the given message you provide.</span>
        </label>
        <div class="modal-action">
            <form method="dialog">
                <label class="label text-white hover:animate-pulse">
                    <span class="label-text text-grey">What Would You like to Say?</span>
                </label>
                <input id="systellvmmd" type="text" class="input input-bordered text-white " placeholder="Accept me because..."/>
                <button class="btn btn-accent mt-4" id="accept">SEND</button>
            </form>

        </div>
    </div>
</dialog>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-accent text-2xl">ORGANIZATION FINDER</h2>
            <p class="text-accent text-md">Find Groups in your Area!</p>
            <div class="input-group">
                <input type="text" id="pass" placeholder="Zipcode" class="input rounded-xl input-bordered input-accent w-full max-w-xs join-item text-white" bind:value={zipc}/>
                <button class="btn btn-accent join-item text-white" id="jwt" on:click={Get}>?</button>
            </div>
            <br>
            <div class="card-body">
                <ul class="" id="unorderedList">
                    <!--{#await promise}-->
                    <!--    {:then response}-->
                    <!--        {#each response.data.data as org}-->
                    <!--            <li class="bg-accent text-neutral text-lg p-2 m-2 rounded-xl">Join </li>-->
                    <!--        {/each}-->
                    <!--{/await}-->
                </ul>

            </div>




<!--            <input type="text" placeholder="Username" class="input input-bordered input-accent w-full max-w-xs text-white" />-->
<!--            <div class="form-control">-->
<!--                <div class="input-group udp">-->
<!--                    <input type="password" id="pass" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs join-item" />-->
<!--                    <button class="btn btn-accent join-item text-white" id="jwt">+</button>-->
<!--                </div>-->
<!--            </div>-->

<!--            <p class="text-neutral textarea-sm"><a>Forgot Password?</a></p>-->
<!--            <div class="card-actions">-->

<!--                <button class="btn btn-outline btn-accent">Login</button>-->
<!--            </div>-->
        </div>
    </div>
</div>

