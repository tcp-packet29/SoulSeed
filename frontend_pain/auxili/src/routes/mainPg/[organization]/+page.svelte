<link href='https://api.mapbox.com/mapbox-gl-js/v2.8.1/mapbox-gl.css' rel='stylesheet' />
<style>

    #map {
        width: 100%;
        position: center;
        height: 100%;





    }
</style>
<script>
    import Navbar from "../../../lib/navbar.svelte";
    import axios from 'axios';
    import { onMount } from 'svelte';
    import mapboxgl from 'mapbox-gl';
    import {page} from '$app/stores'



if (browser && ($page.params.organization != "global"))  { //jwtauthiasndmasjhaseprskosrjejnder
    axios.get('http://localhost:8080/access/users/token', {
        headers: {
            "token": getToken(),
        }
    })
        .then(function (response) {
            let usId = response.data.data.Data.Id
            console.log(usId)
            axios.get('http://localhost:8080/app/organizations/stuff/' + $page.params.organization + '/' + usId, {
                headers: {
                    "token": getToken(),
                }
            })//webrtc refere since server
                .then(function (response) {
                    console.log(response)

                })
                .catch(function (response) {
                    console.log(response)
                    alert("no org yes")
                    // window.location.href = "/login";
                })
        })
        .catch(function (error) {
            console.log(error);
            // window.location.href = "/login";
            console.log("this wont come")
        })
}
    let zipc = ""

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

    function setLatLong(e) {
        if (browser) {
            if ($page.params.organization == 'global') {
                axios.get('http://localhost:8080/access/geo/' + zipc, {
                    headers: {
                        "Token": getToken(), //this isnt returning null or underfined
                        "Content-Type": "application/json"
                    }
                })

                    .then(function (response) {
                        console.log(response)
                        window.localStorage.setItem("org", $page.params.organization);
                        // if (response.status != 201 && response.status != 200) {
                        //     console.log(response.status)
                        //     alert("You are not logged in")//will it show
                        //
                        // }

                        window.location.href = "http://localhost:5173/tradeCreate/" + e.lngLat.wrap().lat + "~amqp~" + e.lngLat.wrap().lng


                    }).catch(function (error) {
                    console.log(error)
                    alert("That is not a nearby area!")

                })
            } else {
                axios.get('http://localhost:8080/app/organizations/loc/' + $page.params.organization + '/' + zipc, {
                    headers: {
                        "Token": getToken(), //this isnt returning null or underfined
                        "Content-Type": "application/json"
                    }
                })

                    .then(function (response) {
                        console.log(response)
                        window.localStorage.setItem("org", $page.params.organization);
                        // if (response.status != 201 && response.status != 200) {
                        //     console.log(response.status)
                        //     alert("You are not logged in")//will it show
                        //
                        // }

                        window.location.href = "http://localhost:5173/tradeCreate/" + e.lngLat.wrap().lat + "~amqp~" + e.lngLat.wrap().lng
                }).catch((error) => {
                    console.log(error)
                    alert("That is not a nearby area!")
                })
            }
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
    let lati = 0;
    let longi = 0;


    onMount(async () => {
        let x;
        let tcp = []
        navigator.geolocation.getCurrentPosition((position) => {
            window.localStorage.setItem('lat', position.coords.latitude);
            window.localStorage.setItem('long', position.coords.longitude)//redis kv referlibtoeitisokcasifjauexteideidjsadocksajriemaldaloaslcloa
            x = position
            console.log("first")
            // map.flyTo(
            //     {
            //         center: [position.coords.longitude, position.coords.latitude],
            //         zoom: 10
            //     }
            // )
        })


        mapboxgl.accessToken = "pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA" //should probably env this

        const map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [window.localStorage.getItem('long'), window.localStorage.getItem('lat')],
            zoom: 11.5,
        })


        map.on("click", (e) => {
            axios.get(' https://api.mapbox.com/geocoding/v5/mapbox.places/' + e.lngLat.wrap().lng + ',' + e.lngLat.wrap().lat + '.json?types=postcode&limit=1&access_token=pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA')
                .then(function (response) {
                    console.log(response.data.features[0].text);
                    zipc = response.data.features[0].text

                })
                .catch(function (error) {
                    console.log(error);
                })
            console.log(e.lngLat.wrap())
            let pop = new mapboxgl.Popup()
                .setLngLat(e.lngLat.wrap())
                .setHTML("<button class='text-xl text-secondary btn-accent btn' href=''>Create</button>")
                .addTo(map)

            pop.getElement().addEventListener('click', () => {
                setLatLong(e)
            })

            map.on('closepopups', () => {
                pop.remove()
            }
            )
        })
        axios.get('http://localhost:8080/app/trades/organization/' + $page.params.organization, {
            headers: {
                "Token": getToken(),
                "Content-Type": "application/json",
            }
        }).then((response) => {
            tcp = response.data.data.tradestcpudp;
            console.log(tcp)
            console.log(response)
            for (let iteramtcp = 0; iteramtcp < tcp.length; iteramtcp++) {
                if (tcp[iteramtcp].Open) {
                    const markerarmasm = new mapboxgl.Marker()
                        .setLngLat([parseFloat(tcp[iteramtcp].Latlong[1]), parseFloat(tcp[iteramtcp].Latlong[0])])
                        .setPopup(
                            new mapboxgl.Popup()
                                .setHTML("<h1 class='text-xl text-secondary'>" + tcp[iteramtcp].Name + "</h1><p>Made by <b>" + tcp[iteramtcp].Maker.Username + "</b></p><br><a class='text-xl text-secondary btn-accent btn' href ='http://localhost:5173/organization/" + tcp[iteramtcp].OrgId + "/trade/" + tcp[iteramtcp].Id + "'>View</button>")
                                .addTo(map)
                        )
                        .addTo(map)
                    markerarmasm.getElement().addEventListener('click', (e) => {

                        setTimeout(() => {
                            map.fire('closepopups')
                            console.log("udp")
                        }, 1)


                        // popjwt.addTo(map)
                    })
                }
            }
        }).catch((err) => {
            console.log(err)
        })




    })


</script>

<div class="drawer drawer-end">
    <input id="opensettings" type="checkbox" class="drawer-toggle">
    <div class="drawer-content">
        <div id="map" class="z-40 border-8 border-b-secondary border-l-secondary border-r-secondary rounded-xl"></div>
        <Navbar />



    </div>
    <div class="drawer-side">
        <label for="opensettings" class="drawer-overlay"></label>
        <ul class="menu p-4 w-80 bg-base-100 text-base-content text-accent bg-secondary rounded-l-box border-accent">
            <li class="menu-title hover:text-accent">
                <span class="hover:text-accent">General</span>
            </li>
            <li><button><a href="/">Main Page</a></button></li>
            <li>
                <button on:click={logOut}><a>Logout</a></button>
            </li>
            <li class="divider"></li>
            <li class="menu-title hover:text-accent">
                <span class="hover:text-accent">Util</span>
            </li>
            <li><button><a href="/organization">Organizations</a></button></li>
        </ul></div>
</div>
