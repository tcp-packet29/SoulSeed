<script>
    import Navbar from "../lib/navbar.svelte";
    let zipc = "";
    import { browser } from "$app/environment";
    import axios from "axios";
    import Svg from "../lib/svg.svelte";
    function a() {
        console.log("hello");
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

    function createUser() {
        if (browser) {
            axios.get('http://localhost:8080/access/tradesin/' + zipc
            ).then(function (response) {
                console.log(Object.keys(response.data.data.data));
                if (Object.keys(response.data.data.data).length == 0) {
                    document.getElementById("udp").innerHTML = "No Items Found";
                    return;
                }
                let x = Object.keys(response.data.data.data);
                let dothis = ""
                for (let i = 0; i < x.length; i++) {
                    if (x[i] in kv) {
                       dothis += kv[x[i]] + " ";
                    }
                    dothis += x[i] + ", ";
                }
                document.getElementById("udp").innerHTML = dothis;
            }).catch(function (error) {
                document.getElementById("udp").innerHTML = "No Items Found";
                console.log(error);
            });
        }
    }

</script>

<body class="bg-primary"></body>
<Navbar/>
<div class="hero min-h-screen" style="background-image: url(https://media.istockphoto.com/id/612242240/photo/ripe-apples-in-orchard-ready-for-harvesting.jpg?s=612x612&w=0&k=20&c=10kwdSAOJeC-pbr0t9QmmGg6psA9dzGQc6rDcnXLLqo=);">
    <div class="hero-overlay bg-opacity-50"></div>
    <div class="hero-content text-center text-secondary-content align-bottombg-secondary bg-opacity-30 rounded-xl">
        <div class="max-w-md bg-secondary bg-opacity-40 rounded-xl p-5">
            <p class="mb-5 text-7xl font-bold">
                🌱
            </p>
            <h1 class="mb-5 text-5xl font-bold text-secondary">
                Welcome to Soulseed
            </h1>
            <p class="mb-5 text-2xl font-bold text-secondary">
                An app for vendors and customers; with emphasis on flexibility.
            </p>
            <p class="mb-5 text-xl font-bold text-secondary">Create and join organizations from your neighbourhood, and preserve resources for your community!</p>

            <hr>
            <br>
            <button class="btn glass btn-accent text-md text-secondary">
                <a href="/register">Get Started</a>
            </button>
            <br>

            <br>
            <a href="/login" class="link-secondary link">Already Have an Account?</a>

        </div>

<!--        <div data-theme="mycodecompiled" class="card w-96 text-neutral-content glass">-->
<!--            <div class="card-body items-center text-center">-->
<!--                <h2 class="card-title text-secondary text-2xl">Enter Your Zipcode!</h2>-->
<!--                <a class="content-center link-secondary text-sm link" href="/organizationFinder">Want to Find Organizations Instead?</a>-->
<!--                <p class="text-secondary text-md">Enter your Zipcode to find items offered in your area!</p>-->

<!--                <input type="text" placeholder="" class="input input-bordered glass w-full max-w-xs text-secondary " bind:value={zipc}/>-->
<!--                <p class="text-secondary text-xl" id="udp"></p>-->
<!--                <div class="card-actions">-->

<!--                    <button class="btn btn-outline glass" on:click={createUser}>Find</button>-->

<!--                </div>-->
<!--            </div>-->
<!--        </div>-->

    </div>

</div>


