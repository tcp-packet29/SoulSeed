<script>
    import { page } from '$app/stores'
    import axios from 'axios';
    import {browser} from "$app/environment";
    import Navbar from "../../../lib/navbar.svelte";
    let x = $page.params.latlong.split('~')
    console.log(x)
    let zipcode = ""
    let lat = x[0]
    let liblong = x[2]
    let amqp = ""
    let tcp = ""
    let providing = false

    function test(callback) {
        let among = "";
        axios.get(' https://api.mapbox.com/geocoding/v5/mapbox.places/' + liblong + ',' + lat + '.json?types=postcode&limit=1&access_token=pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA')
            .then(function (response) {
                if (response.data.features.length > 0) {
                    console.log(response.data.features[0].text);
                    //synchroous
                    callback(response.data.features[0].text)
                } else {
                    alert("That is not a nearby area!")
                    if(browser) {
                        window.location.href = "http://localhost:5173/mainPg"
                    }
                }

            })
            .catch(function (error) {
                console.log(error);
            })
    }
    try {
        axios.get(' https://api.mapbox.com/geocoding/v5/mapbox.places/' + liblong + ',' + lat + '.json?types=postcode&limit=1&access_token=pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA')
            .then(function (response) {
                console.log(response);
                if (response.data.features.length > 0) {
                    console.log(response.data.features[0].text);
                    zipcode = response.data.features[0].text //synchroous
                }
                if (browser) {
                    let x = window.localStorage.getItem("org");
                    if (x == "global") {
                        setLatLong(lat, liblong)
                    } else {
                        axios.get('http://localhost:8080/app/organizations/loc/' + x+ '/' + zipcode, {
                            headers: {
                                "Token": getToken(), //this isnt returning null or underfined
                                "Content-Type": "application/json"
                            }
                        })

                            .then(function (response) {
                                console.log(response)

                                // if (response.status != 201 && response.status != 200) {
                                //     console.log(response.status)
                                //     alert("You are not logged in")//will it show
                                //
                                // }

                            }).catch((error) => {
                            console.log(error)
                            alert("That is not a nearby area!")

                            })
                    }
                    // { AMOUNT: 0999990329053295258934888888
                    }


            })
            .catch(function (error) {
                console.log(error);
            })

    } catch(e) {
        alert("aa")
        if(browser) {
            window.location.href = "http://localhost:5173/mainPg"
        }

    }




    function getToken() {
        if (browser) {
            return window.localStorage.getItem("token")
        }
    }
    function setLatLong(lat, long)
    {
        test(function(val) {
            if (browser) {
                axios.get('http://localhost:8080/access/geo/' + val, {
                    headers: {
                        "Token": getToken(), //this isnt returning null or underfined
                        "Content-Type": "application/json"
                    }
                })

                    .then(function (response) {
                        console.log(response)
                        // if (response.status != 201 && response.status != 200) {
                        //     console.log(response.status)
                        //     alert("You are not logged in")//will it show
                        //
                        // }


                    }).catch(function (error) {
                    console.log(error)
                    alert("That is not a nearby area!")
                    //window.location.href = "http://localhost:5173/mainPg"

                })
            }
        })

    }


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

    let jwtamqp = 0;
    let a = []
    let item = ""
    function work() {
        if (browser && item.trim() != "" && jwtamqp < 5 && a.includes(item) == false) {
            let ele = document.getElementById("jwt")
            let elem = document.createElement("li")
            elem.className = "badge badge-accent badge-lg badge-outline"
            if (kv[item] != undefined) {
                elem.innerText = kv[item] + "  " + item
            } else {
                elem.innerText = item
            }
            elem.addEventListener("click", function () {
                ele.removeChild(elem)
                a.splice(a.indexOf(item), 1)
                jwtamqp--
            })
            ele.appendChild(elem)


            jwtamqp++
            a.push(item)
            item = ""
            console.log(a)
        }
    }

    function jwtcudp() {
        if (browser) {
            let valu = window.localStorage.getItem("org")


            axios.get('http://localhost:8080/access/users/token', {
                headers: {
                    "Token": getToken(), //this isnt returning null or underfined
                    "Content-Type": "application/json"
                }
            })
                .then((response) => {
                    console.log(response)
                    console.log(response.data.data.Data.Id)
                    axios.post('http://localhost:8080/app/trades/create/' + zipcode, {
                        "makerid": response.data.data.Data.Id,
                        "name": amqp,
                        "items": a,
                        "description": tcp,
                        "open": true,
                        "orgid": valu,
                        "latlong": [lat, liblong],
                        "offers": [],
                        "providing": providing,
                    }, {
                        headers: {
                            "Token": getToken(), //this isnt returning null or underfined
                            "Content-Type": "application/json"
                        }
                    }).then((response) => {
                        console.log(response)
                        if (response.status == 201) {
                            alert("Trade Created!")
                            window.location.href = "http://localhost:5173/organization/" + valu + "/trade/" + response.data.data.data.Id
                        }

                    }).catch((response) => {
                        console.log(response)
                    })
                })
                .catch((response) => {
                    console.log(response)
                })
        }
    }

</script>

<body class="bg-primary"></body>
<Navbar />
<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral tooltip" data-tip={lat + ", " + liblong}>Create Trade</h2>
            <div class="form-control">
                <label class="label cursor-pointer">
                    <span class="label-text tooltip" data-tip="Are you offering material or asking for some?">Providing Resources? </span>
                    <span class="text-secondary">-</span>
                    <input type="checkbox" class="checkbox checkbox-accent" bind:checked={providing} />
                </label>
            </div>

            <input type="text" placeholder="Name" class="input input-bordered w-64 text-white" bind:value={amqp} />
            <textarea placeholder="Description" class="textarea textarea-bordered textarea-md w-64 text-white" bind:value={tcp}></textarea>
            <div class="form-control">
                <div class="input-group">
                    <input placeholder="Items" class="input input-bordered w-48 text-white" bind:value={item}/>
                    <button class="btn btn-square btn-accent" on:click={work} accesskey="enter">+</button>
                </div>

            </div>
            <button class="btn btn-accent" on:click={jwtcudp}>Create</button>


            <ul class="flex justify-center items-center grid" id="jwt">
            </ul>

        </div>
    </div>
</div>