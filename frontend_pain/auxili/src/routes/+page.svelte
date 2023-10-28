<script>
    import Navbar from "../lib/navbar.svelte";
    function a() {
        console.log("hello");
    }
</script>

<body class="bg-primary"></body>
<Navbar/>
<div class="hero min-h-screen" style="background-image: url(https://media.istockphoto.com/id/612242240/photo/ripe-apples-in-orchard-ready-for-harvesting.jpg?s=612x612&w=0&k=20&c=10kwdSAOJeC-pbr0t9QmmGg6psA9dzGQc6rDcnXLLqo=);">
    <div class="hero-overlay bg-opacity-50"></div>
    <div class="hero-content text-center text-secondary-content">
        <div class="max-w-md bg-secondary bg-opacity-20 rounded-xl p-5">
            <p class="mb-5 text-7xl font-bold">
                🌱
            </p>
            <h1 class="mb-5 text-5xl font-bold text-secondary">
                Welcome to Soulseed
            </h1>
            <p class="mb-5 text-2xl font-bold text-secondary">
                An app for vendors and customerstcp; with emphasis on flexibility.
            </p>
            <button class="btn glass btn-accent text-md text-secondary">
                <a href="/register">Get Started</a>
            </button>
        </div>
    </div>
</div>
<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary text-neutral-content">
        <div class="card-body items-center text-center">
                <h2 class="card-title text-neutral">Welcome to Auxili</h2>
            <br>
            <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility. This app uses golang for its backend web servers utilizing gin gonic, as well as rabbitmq for message broking.</p>
            <br>
            <br>
            <div class="card-actions">
                <a class="btn btn-outline btn-accent" href= "/login">Login</a>
                <a class="btn btn-outline btn-accent" href= "/register">Register</a>
                <a class="btn btn-outline btn-accent" href="/about>">About</a>
            </div>
        </div>
    </div>
    </div>
