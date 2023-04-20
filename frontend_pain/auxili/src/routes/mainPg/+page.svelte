<link href='https://api.mapbox.com/mapbox-gl-js/v2.8.1/mapbox-gl.css' rel='stylesheet' />
<style>
    #map {
        width: 100%;
        position: center;
        height: 80%;




    }
</style>
<script>
    import axios from 'axios';
    import { onMount } from 'svelte';
    import mapboxgl from 'mapbox-gl';

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
            axios.get('http://localhost:8080/access/geo/' + zipc, {
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

                    window.location.href = "http://localhost:5173/tradeCreate/" + e.lngLat.wrap().lat + "~amqp~" + e.lngLat.wrap().lng


                }).catch(function (error) {
                    console.log(error)
                    alert("That is not a nearby area!")

            })
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

    onMount(async () => {
        mapboxgl.accessToken = "pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA" //should probably env this

        const map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [-74.5, 40],
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
                .setHTML("<a class='text-xl text-accent' href=''>Create</a>")
                .addTo(map)

            pop.getElement().addEventListener('click', () => {
                setLatLong(e)
            })
        })
    })


</script>

<div class="drawer drawer-end">
    <input id="opensettings" type="checkbox" class="drawer-toggle">
    <div class="drawer-content">
        <div id="map" class="z-40 border-8 border-b-secondary border-l-secondary border-r-secondary rounded-xl"></div>
        <div class="navbar bg-base-200 rounded-b-box z-50 absolute inset-x-0 top-0">
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
            <li class="menu-title hover:text-accent">
                <span class="hover:text-accent">General</span>
            </li>
            <li><button><a>Profile</a></button></li>
            <li>
                <button on:click={logOut}><a>Logout</a></button>
            </li>
            <li class="divider"></li>
            <li class="menu-title hover:text-accent">
                <span class="hover:text-accent">Util</span>
            </li>
            <li><button><a>Organizations</a></button></li>
            <li><button><a>Trades</a></button></li>
            <li><button><a>Deals</a></button></li>
        </ul></div>
</div>
