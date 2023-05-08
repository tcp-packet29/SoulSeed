<script>
    import { page } from '$app/stores'
    import axios from 'axios';
    import {browser} from "$app/environment";
    let x = $page.params.latlong.split('~')
    console.log(x)
    let zipcode = ""
    let lat = x[0]
    let liblong = x[2]

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
                if (response.data.features.length > 0) {
                    console.log(response.data.features[0].text);
                    zipcode = response.data.features[0].text //synchroous
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

    setLatLong(lat, liblong)

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




</script>

<body class="bg-primary"></body>
<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>

</div>
<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral tooltip" data-tip={lat + ", " + liblong}>Create Trade</h2>

            <input type="text" placeholder="Name" class="input input-bordered w-64" />
            <textarea placeholder="Description" class="textarea textarea-bordered textarea-md w-64"></textarea>
            <div class="form-control">
                <div class="input-group">
                    <input placeholder="Items" class="input input-bordered w-48 " bind:value={item}/>
                    <button class="btn btn-square btn-accent" on:click={work} accesskey="enter">+</button>
                </div>

            </div>
            <button class="btn btn-accent">Create</button>


            <ul class="flex justify-center items-center grid" id="jwt">
            </ul>

        </div>
    </div>
</div>