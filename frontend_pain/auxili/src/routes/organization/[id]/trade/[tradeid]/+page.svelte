<script>
    import { page } from '$app/stores'
    import  axios  from 'axios'
    import {getTok, prse, getToken} from "$lib/parse.svelte"
    import {browser} from "$app/environment";
    import Parsetwo from "../../../../../lib/parsetwo.svelte";
    import Navbar from "../../../../../lib/navbar.svelte";
    let val = "rsi"
    let udp = ""
    let llvm = " "
    let promisea
    let promise
    let editing = "Creating a new offer."
    let thispromise

    if (browser) {

        getTok()
            .then(function (response) {
                let jsotn = prse(response)
                console.log("a", response)
                console.log("a", response.data.data.Data)
                val = jsotn
                console.log("test", val)
                axios.get("http://localhost:8080/app/organizations/stuff/" + $page.params.id + "/" + val.Id, {
                    headers: {
                        "Token": getToken()
                    }
                }).then(response => {
                    if (!response.data.data.exists) {
                        //lex parse inrto ast
                        window.location.href="/"
                    }
                }).catch(error => {
                    window.location.href="/"
                })
            })
            .catch(function (error) {
                console.log(error)
                console.log("error internal")
                if (browser) {
                    window.location.href = "http://localhost:5173/login";
                }
            })

        axios.get("http://localhost:8080/app/trades/offer/" + $page.params.tradeid + "/amnt", {
            headers: {
                "Token": getToken()
            }
        }).then(response => {
            console.log(response)
            document.getElementById("works").innerText = "There are currently " + response.data.data.data + " offers for this trade."
        }).catch(error => {
            console.log(error)
        })


    }
     promise = axios.get('http://localhost:8080/app/trades/tcpreference/' + $page.params.tradeid + '/' + $page.params.id, {
            headers: {
                "Token": getToken()
            }
        })
        //.then((response) => {
        //console.log(response)
        //should i use prisma if im using backend in trypsecript and sql db
        // udp = response.data.data.data

        //}).catch(function(error) {
        //    console.log(error)
        // window.location.href = '/'
        //})


        function tcp() {
            if (browser) {
                let elem = new HTMLDivElement()
                let inside = new HTMLSpanElement().innerText = "ORM"
                elem.setAttribute("class", "alert alert-primary")


            }
        }


            promisea = axios.get('http://localhost:8080/app/trades/checkOwner/' + $page.params.tradeid, {
                headers: {
                    "Token": getToken()
                }
            }).then((response) => {
                // console.log(response)
                // //should i use prisma if im using backend in trypsecript and sql db
                // // udp = response.data.data.data
                // llvm = response.data.data.data
                // comment node inthe ast
                console.log(response)
                console.log('isowner');
                if (browser) {
                    document.getElementById("tcp").innerText = "Close Trade"
                    document.getElementById('by').innerText = "By you"
                    let x = document.getElementById("tcpstream")
                    x.addEventListener("click", () => {
                        x.style.display = "none"
                    })
                }

            }).catch(function (error) {
                if (browser) {
                    document.getElementById("tcp").innerText = "Edit Offer"
                    let x = document.getElementById("tcpstream")
                    x.addEventListener("click", () => {
                        x.style.display = "none"
                    })

                    let vale = getToken()
                    document.getElementById("tcp").addEventListener("click", () => {
                        document.getElementById("confirmation_modal").showModal()
                    })
                    document.getElementById("submit").addEventListener("click", () => {
                        axios.post('http://localhost:8080/app/trades/offer/' + $page.params.tradeid, {
                            "Usr": val,
                            "Address": document.getElementById("systemd").value,
                            "Description": document.getElementById("tcppacket").value,
                            "Approved": false,
                            "UserId": val.Id,
                        },{
                            headers: {
                                "Token": vale
                            },
                        }).then(response => {
                            console.log(response)
                            location.reload(true)
                        }).catch(error => {
                            console.log(error)
                        })
                    })
                    axios.get("http://localhost:8080/app/trades/offer/" + $page.params.tradeid + "/" + val.Id, {
                        headers: {
                            "Token": getToken()
                        }
                    }).then(response => {
                        console.log(response)
                        if (browser) {
                            document.getElementById("systemd").value = response.data.data.data.Address
                            document.getElementById("tcppacket").value = response.data.data.data.Description
                            editing = "Editing an offer."
                            document.getElementById("submit").innerText = "Edit Offer"
                        }
                    }).catch(error => {
                        document.getElementById("tcp").innerText = "Make Offer"
                    })
                }

                // window.location.href = '/'
            })
        // thispromise = axios.get('http://localhost:8080/app/trades/offersFetch/' + $page.params.tradeid, {
        //     headers: {
        //         "Token": getToken()
        //     }
        // })
    //tcp




//ssh
    //tcp realtime str
    //wjt jwt
    //amqp
    //mqtt
    //xmpp
    //astr
    ///macro rpeprocessor
    //tcp

    let kv = {

        "apple" : "🍎",
        "banana": "🍌",
        "orange": "🍊",
        "pineapple": "🍍",
        "grapes": "🍇",
        "watermelon": "🍉",
        "pear": "🍐",
        "peach": "🍑",
        "cherries": "🍒",
        "strawberry": "🍓",
        "kiwi": "🥝",
        "tomato": "🍅",
        "coconut": "🥥",
        "avocado": "🥑",
        "eggplant": "🍆",
        "potato": "🥔",
        "carrot": "🥕",
        "corn": "🌽",
        "pepper": "🌶",
        "cucumber": "🥒",
        "broccoli": "🥦",
        "mushroom": "🍄",
        "peanuts": "🥜",
        "chestnut": "🌰",




        //make a key value pair of more fruits matched to their emojis

    }
</script>


<!--<div class="hero min-screen bg-base-200 bg-primary">-->
<!--    <div class="hero-content flex-col">-->
<!--        <div class="text-center lg:text-right">-->
<!--            <h1 class="text-accent text-neutral text-9xl">-->
<!--                {val}-->
<!--            </h1>-->
<!--        </div>-->

<!--    </div>-->
<!--</div>-->
<div class="alert alert-accent bg-black z-60 absolute m-5 w-1/4 text-center justify-center  left-0 bottom-0 " id="tcpstream">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
    <span id="works" class="text-accent"></span>
</div>


{#await promise}
    {:then response}
    {#if response.data.data.data.Maker.Id == val.Id}
        <Parsetwo id="{$page.params.id}"/>
    {:else}
        <Navbar id="{$page.params.id}"/>
    {/if}

    <dialog class="rounded-xl bg-secondary shadow-xl border-l-accent border-l-8 a" id="confirmation_modal">
        <div class="modal-box content-center justify-center bg-secondary p-10 shadow-none">
            <h1 class="text-lg font-bold text-accent">Submit Offer?</h1>
                <label class="label text-white hover:animate-pulse">
                    <span class="label-text text-grey">{editing}</span>
                </label>
               <div class="modal-action">
                   <form method="dialog">
                       {#if !response.data.data.Providing}
                           <label class="label text-white hover:animate-pulse">
                               <span class="label-text text-grey">Address</span>
                           </label>
                          <input id="systemd" type="text" class="input input-bordered text-white" placeholder="What's your address?"/>
                           <label class="label text-white hover:animate-pulse">
                               <span class="label-text text-grey">Memo</span>
                           </label>
                           <input id="tcppacket" type="text" class="input input-bordered text-white" placeholder="Why you?"/>
                           <label class="label text-white hover:animate-pulse">
                               <span class="label-text text-grey">If you are chosen you may receive an email or text by SMS.</span>
                           </label>



                       {/if}
                       <button class="btn btn-accent mt-4" id="submit">SUBMIT</button>

                   </form>

               </div>
        </div>
    </dialog>
    <dialog class="rounded-xl bg-secondary shadow-xl border-l-accent border-l-8 a" id="accept_modal">
        <div class="modal-box content-center justify-center bg-secondary p-10 shadow-none">
            <h1 class="text-lg font-bold text-accent">Accept Offer?</h1>
            <label class="label text-white hover:animate-pulse">
                <span class="label-text text-grey">By accepting the offer you take responsibility for all consequential actions. This will also close the trade.</span>
            </label>
            <div class="modal-action">
                <form method="dialog">
                    <h1 class="text-md text-white" id="Username"></h1>
                    <h1 class="text-md text-white" id="Addr"></h1>
                    <h1 class="text-md text-white" id="Mem"></h1>
                    <label class="label text-white hover:animate-pulse">
                        <span class="label-text text-grey">What Would You like to Say?</span>
                    </label>
                    <input id="systellvmmd" type="text" class="input input-bordered text-white" placeholder="What's your address?"/>
                    <button class="btn btn-accent mt-4" id="accept">ACCEPT</button>
                </form>

            </div>
        </div>
    </dialog>
    <div class="hero min-h-screen bg-base-200 bg-secondary">
        <div class="hero-content flex-col lg:flex-row-reverse border-2 rounded-lg ">


            <div class="text-center lg:text-right">
                <div class="grid grid-cols-1 grid-rows-2 flex">
                    <div class="border-l-2 border-r-2 border-t-2 border-accent rounded-t-md border-b-2 p-10">
                    <h1 class="mb-5 font-bold text-neutral lg:text-5xl md:text-5xl sm:text-md uppercase text-center">{response.data.data.data.Name}</h1>
                        <hr>
                        <br>
                        <h1 class="text-accent text-neutral text-xl text-left font-bold">
                            {response.data.data.data.Description}
                        </h1>
                    </div>
                    <div class="grid grid-cols-2 flex">
                        <div class="border-l-2 border-b-2 border-accent rounded-b-md p-5 text-center content-center">
                            <p class="text-4xl font-bold text-neutral text-left">ITEMS</p>
                            <br>
                            <ul>
                                {#each response.data.data.data.Items as udp}

                                    <li class="mb-5 text-xl text-neutral text-left">{kv[udp] == undefined ? udp : kv[udp] + " " + udp}</li>
                                {/each}
                            </ul>
                        </div>
                        <div class="border-l-2 border-r-2 border-b-2 border-accent rounded-b-md p-5 text-center content-center">
                            <p class="text-4xl font-bold text-neutral ">INFO</p>
                            <br>
                            <a class="mb-5 text-xl text-accent" id="by" href="/usr/{response.data.data.data.Maker.Username}">By {response.data.data.data.Maker.Username}</a>
                            <br>
                            <br>
                            {#if response.data.data.data.Open}
                                <p class="mb-5 text-xl text-accent font-bold">Trade is Open</p>
                                <button class="btn btn-accent btn-lg w-4/5 hover:animate-pulse" id="tcp"><span class="loading loading-dots loading-md"></span></button>
                            {:else}
                                <p class="mb-5 text-xl text-accent font-bold">Closed Trade</p>
                            {/if}

                            <!--{#await promisea}-->
                            <!--    {:then response}-->
                            <!--        <button class="btn btn-accent btn-lg w-4/5">Check Offers</button>-->
                            <!--        <h3 class="text-sm text-accent mb-5">You made this post.</h3>-->
                            <!--    {:catch error}-->

                            <!--{/await}-->
                        </div>
                    </div>
                    {#if response.data.data.data.Maker.Id === val.Id && response.data.data.data.Open} <!-- in outb register stack -->
                        <div class="divider"><h1 class="text-neutral font-bold text-2xl text-center p-2">OFFERS</h1></div>
                        <div class="overflow-x-auto overflow-y-auto h-60 content-center text-center justify-center bg-secondary p-4 border-l-2  border-t-2  rounded-lg border-r-2 border-b-2 border-accent">
                            <table class="table justify-center">
                                <thead>
                                <tr>
                                    <th>Name</th>
                                    <th>Address</th>
                                    <th>Memo</th>
                                    <th>Approve?</th>
                                </tr>
                                </thead>
                                <tbody>
                                {#each response.data.data.data.Offers as udp}
                                    <tr>
                                        <td>{udp.Usr.Username}</td>
                                        <td>{udp.Address}</td>
                                        <td>{udp.Description}</td>
                                        <td><button class="btn btn-accent text-white" on:click={()=> {
                                            document.getElementById("Username").innerHTML = udp.Usr.Username;
                                            document.getElementById("Addr").innerHTML = udp.Address;
                                            document.getElementById("Mem").innerHTML = udp.Description;
                                            document.getElementById("accept").onclick = function() {
                                                console.log(udp)
                                                axios.post("http://localhost:8080/app/trades/acceptOffer/" + $page.params.tradeid + "/" + udp.UserId, {
                                                    "message": document.getElementById("systellvmmd").value
                                                }, {
                                                    headers: {
                                                        "Token": getToken()
                                                    }
                                                }).then(response => {
                                                    console.log(response)
                                                }).catch(error => {
                                                    document.getElementById("accept_modal").showModal();
                                                    console.log(error)
                                                    //use error message from response
                                                })
                                            };
                                            document.getElementById("accept_modal").showModal();
                                        }}>Approve</button></td>
                                    </tr>
                                {/each}
                                <!-- //commentnode-->
                                <!--{#await thispromise}-->
                                <!--    {:then response}-->
                                <!--    {#each response.data.Data.data as udp}-->
                                <!--        <tr>-->
                                <!--            <td>{udp.Usr.Username}</td>-->
                                <!--            <td>{udp.Address}</td>-->
                                <!--            <td>{udp.Description}</td>-->
                                <!--            <td>{udp.Approved}</td>-->
                                <!--        </tr>-->
                                <!--    {/each}-->
                                <!--{/await}-->
                                </tbody>
                            </table>
                        </div>
                    {/if}


                </div>

            </div>



        </div>



    </div>

{/await}

